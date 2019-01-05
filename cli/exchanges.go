package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zeeraw/riksbank/swea"
)

func (r *runner) cmdExchanges() cli.Command {
	return cli.Command{
		Name:      "exchanges",
		Usage:     "Lists series for currency exchange rates to SEK",
		UsageText: "foo",
		Action:    r.actionExchanges,
		Flags: []cli.Flag{
			r.flagLang(),
		},
	}
}

func (r *runner) actionExchanges(c *cli.Context) error {
	ctx := context.Background()
	req := &swea.GetAllCrossNamesRequest{
		Language: swea.Language(r.lang),
	}
	res, err := r.api.GetAllCrossNames(ctx, req)
	if err != nil {
		return err
	}
	r.renderCrossNames(res)
	return nil
}

func (r *runner) renderCrossNames(res *swea.GetAllCrossNamesResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "ID\t Name\t Description\n")
	for _, series := range res.Series {

		fmt.Fprintf(w, "%s\t %s\t %s\n", series.ID, series.Name, series.Description[4:])
	}
}