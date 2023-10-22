package cmd

import (
	"plandex/lib"
	"plandex/types"

	"github.com/spf13/cobra"
)

var (
	recursive bool
	maxDepth  int
	namesOnly bool
	truncate  bool
	maxTokens int
	note      string
)

var contextLoadCmd = &cobra.Command{
	Use:     "load [files-or-urls...]",
	Aliases: []string{"l"},
	Short:   "Load context from various inputs",
	Long:    `Load context from a file path, a directory, a URL, text, or piped data.`,
	Args:    cobra.MinimumNArgs(1),
	Run:     contextLoad,
}

func init() {
	contextLoadCmd.Flags().StringVarP(&note, "note", "n", "", "Add a note to the context")
	contextLoadCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Search directories recursively")
	contextLoadCmd.Flags().IntVarP(&maxDepth, "depth", "d", -1, "Maximum depth for recursive directory search (-1 means no limit)")
	contextLoadCmd.Flags().BoolVar(&namesOnly, "names", false, "Only process file names")
	contextLoadCmd.Flags().BoolVar(&truncate, "truncate", false, "Truncate contents if tokens exceed maximum")
	contextLoadCmd.Flags().IntVar(&maxTokens, "max", -1, "Maximum limit for number of tokens")

	// can be called via plandex load or plandex context load
	RootCmd.AddCommand(contextLoadCmd)
	contextCmd.AddCommand(contextLoadCmd)
}

func contextLoad(cmd *cobra.Command, args []string) {
	lib.LoadContextOrDie(&types.LoadContextParams{
		Note:      note,
		MaxTokens: maxTokens,
		Recursive: recursive,
		MaxDepth:  maxDepth,
		NamesOnly: namesOnly,
		Truncate:  truncate,
		Resources: args,
	})
}