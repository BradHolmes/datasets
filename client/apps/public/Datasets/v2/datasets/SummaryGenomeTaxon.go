package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

func createSummaryGenomeTaxonCmd(sgf SummaryGenomeFlag, assemblyRequestFlag AssemblyRequestFlag) *cobra.Command {
	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)"
	tem := cmdflags.NewTaxExactMatchFlag()
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse)
	flagSets := []cmdflags.FlagInterface{tem, iff}

	cmd := &cobra.Command{
		Use:   "taxon",
		Short: fmt.Sprintf("Print a data report containing genome metadata by %s", inputDescription),
		Long: fmt.Sprintf(`
Print a data report containing genome metadata by %s. The data report is returned in JSON format.`, inputDescription),
		Example: `  datasets summary genome taxon human
  datasets summary genome taxon "mus musculus"
  datasets summary genome taxon 10116`,

		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var taxons []string
			for _, arg := range iff.InputIDArgs {
				_, taxError := RetrieveTaxIdForTaxon(
					arg,
					true,
					openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENOME,
					"genome",
				)
				if taxError != nil {
					return taxError
				}
				taxons = append(taxons, arg)
			}

			return getGenomeSummary(NewDefaultRequestIterator(GetGenomeReportsTaxonRequest(taxons, tem.IsTaxExactMatch())), sgf, assemblyRequestFlag)
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())

	return cmd
}
