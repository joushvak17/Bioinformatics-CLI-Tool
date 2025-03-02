package parse

import (
	"fmt"
	"github.com/joushvak17/Bioinformatics-CLI-Tool/pkg/parse"
	"github.com/joushvak17/Bioinformatics-CLI-Tool/pkg/sequence"
	"github.com/spf13/cobra"
)

// NewParseCmd creates and returns the `parse` command.
func NewParseCmd() *cobra.Command {
	var (
		// Define flags for the command
		sequenceLength bool
		gcContent      bool

		// TODO: Add additional flags for analyzing the sequences
		// Primarily the GC content, reverse complement, and nucleotide frequency
		// reverseComp    bool
		// nucleotideFreq bool
	)

	parseCmd := &cobra.Command{
		Use:   "parse <file>",
		Short: "Parse and analyze a FASTA file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			records, err := parse.Parse(filename)
			if err != nil {
				fmt.Printf("Error parsing FASTA file: %v\n", err)
				return
			}

			for _, record := range records {
				// fmt.Printf(">%s %s\n%s\n", record.ID, record.Description, record.Sequence)
				fmt.Printf("ID: %s\n", record.ID)
				fmt.Printf("Description: %s\n", record.Description)
				fmt.Printf("Sequence: %s\n\n", record.Sequence)

				if sequenceLength {
					length := len(record.Sequence)
					fmt.Printf("Sequence Length: %d\n", length)
				}

				if gcContent {
					gc := sequence.GCContent(record.Sequence)
					fmt.Printf("GC Content: %.2f%%\n", gc)
				}

				// TODO: Add additional analysis for the sequences
				// if reverseComp {
				// 	// Calculate reverse complement
				// 	reverse := ReverseComplement(record.Sequence)
				// 	fmt.Printf("Reverse Complement: %s\n", reverse)
				// }
				// if nucleotideFreq {
				// 	// Calculate nucleotide frequency
				// 	freq := NucleotideFrequency(record.Sequence)
				// 	fmt.Printf("Nucleotide Frequency: %v\n", freq)
				// }

				fmt.Println() // Add an empty line between records
			}
		},
	}

	// Add flags to the command
	parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")
	parseCmd.Flags().BoolVarP(&gcContent, "gc", "g", false, "Calculate GC content")

	// TODO: Add additional flags for analyzing the sequences
	// Primarily the GC content and the reverse complement
	// parseCmd.Flags().BoolVarP(&reverseComp, "reverse", "r", false, "Calculate reverse complement")

	return parseCmd
}
