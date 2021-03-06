package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank"
	"github.com/zeeraw/riksbank/cli/flags"
	"github.com/zeeraw/riksbank/date"
)

const (
	ratesName  = "rates"
	ratesUsage = "Lists interest or exchange rates between two dates"
)

func (t *Tool) cmdRates() cli.Command {
	return cli.Command{
		Name:   ratesName,
		Usage:  ratesUsage,
		Action: t.actionRates,
		Flags: []cli.Flag{
			flags.From,
			flags.To,
			flags.Aggregate,
			flags.Analysis,
			flags.Series,
		},
	}
}

func (t *Tool) actionRates(c *cli.Context) error {
	ctx := context.Background()
	series := c.StringSlice("series")
	if len(series) < 1 {
		return fmt.Errorf("need to have at least one series")
	}
	from, err := date.Parse(c.String("from"))
	if err != nil {
		return err
	}
	to, err := date.Parse(c.String("to"))
	if err != nil {
		return err
	}
	aggregate, err := riksbank.ParseAggregate(c.String("aggregate"))
	if err != nil {
		return err
	}
	analysis, err := riksbank.ParseAggregateAnalysis(aggregate, c.String("analysis"))
	if err != nil {
		return err
	}

	req := &riksbank.RatesRequest{
		Series:          series,
		From:            from,
		To:              to,
		AggregateMethod: aggregate,
		AnalysisMethod:  analysis,
	}

	res, err := t.Riksbank.Rates(ctx, req)
	if err != nil {
		return err
	}

	return t.renderRates(req, res)
}

func (t *Tool) renderRates(req *riksbank.RatesRequest, res *riksbank.RatesResponse) error {
	const (
		rowFmt = "%s\t %s\t %s\t %s\n"
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(os.Stdout, "Ranging from %s to %s\n", formatDate(req.From), formatDate(req.To))
	if req.AggregateMethod != riksbank.Daily {
		fmt.Fprintf(os.Stdout, "Aggregating %s %s value\n\n", req.AggregateMethod.Name(), req.AnalysisMethod)
	}
	fmt.Fprintf(os.Stdout, "Series %s\n", strings.Join(req.Series, ", "))
	fmt.Fprint(os.Stdout, "\n")
	var valueLabel string
	switch req.AnalysisMethod {
	case riksbank.Mean:
		valueLabel = "Mean value"
	case riksbank.Min:
		valueLabel = "Min value"
	case riksbank.Max:
		valueLabel = "Max value"
	case riksbank.Ultimo:
		valueLabel = "Ultimo value"
	default:
		valueLabel = "Value"
	}
	fmt.Fprintf(w, rowFmt, "Period", "Series ID", "Series Name", valueLabel)
	for _, rate := range res.Rates {
		fmt.Fprintf(w, rowFmt, rate.Period, rate.Series.ID, rate.Series.Name, formatFloat(rate.Value))
	}
	return nil
}
