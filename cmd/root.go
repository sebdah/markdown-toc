package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sebdah/markdown-toc/toc"
	"github.com/spf13/cobra"
)

var (
	// header is the string injected as a header for the table of contents.
	header string

	// skipHeader is indicating whether or not a header should be injected in
	// the table of contents.
	skipHeader bool

	// replaceToC is indicating whether we should replace the table of contents
	// in the input file. This assumes that there are two tags indicating where
	// the ToC starts and where it ends:
	//
	// Start:   <!-- ToC start -->
	// End:     <!-- ToC end -->
	//
	// If these tags are not found, the table of contents will be injected on
	// top all existing content in the markdown file.
	replaceToC bool

	// inline indicates whether we should do an inline replacement of the ToC or
	// if we should print to stdout. Default is to print to stdout.
	inline bool
)

var RootCmd = &cobra.Command{
	Use:   "markdown-toc <file>",
	Short: "Generate a table of contents for your markdown file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}

		t, err := toc.Build(d, header, !skipHeader)
		if err != nil {
			return err
		}

		if replaceToC {
			t = toc.Replace(d, t)
		}

		if inline {
			f, err := os.Open(args[0])
			if err != nil {
				return err
			}

			i, err := f.Stat()
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(args[0], []byte(strings.Join(t, "\n")), i.Mode())
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s\n", strings.Join(t, "\n"))
		}

		return nil
	},
}

func init() {
	RootCmd.Flags().StringVar(&header, "header", "# Table of Contents", "Text to use for the header for the ToC")
	RootCmd.Flags().BoolVar(&skipHeader, "skip-header", false, "If this is set there will be no header for the ToC")
	RootCmd.Flags().BoolVar(&replaceToC, "replace", false, "If the replace flag is set the full markdown will be returned and any existing ToC replaced")
	RootCmd.Flags().BoolVar(&inline, "inline", false, "Overwrite the input file with the output from this command. Should be used together with --replace")
}
