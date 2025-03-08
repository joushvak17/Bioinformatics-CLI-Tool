package cli

import (
	"github.com/joushvak17/SeqCraft/internal/parse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "SeqCraft",
	Short: "SeqCraft - Bioinformatics CLI Tool written in Go",
	Long:  "SeqCraft - CLI tool for sequence analysis, alignment, and structure analysis, all accessible through an easy-to-use command line interface.",
}

// Execute runs the root command and returns any error encountered.
// It should be called to start the CLI application.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add the parse command to the root command
	rootCmd.AddCommand(parse.NewParseCmd())
}

// TODO: The following will list out the things that need to be done in the future

// Advanced Features
// - Sequence Alignment: Pairwise sequence alignment and support for MSA
// - Sequence Translation: Translate DNA/RNA sequences to protein sequences 
// - Motif Search: Search for specific motifs in sequences
// - Secondary Structure Prediction: Predict secondary structure of protein sequences
// - Phylogenetic Analysis: Construct phylogenetic trees from sequence data

// Improve Usability
// - Add more detailed help messages for each command
// - Progress bars for long-running tasks
// - Error handling and reporting
// - Input validation and error messages

// Logging
// - Log levels: Use different log levels, like INFO, WARNING, ERROR, DEBUG
// - Log Files: Write logs to a file for debugging and analysis

// Optimization
// - Parallel Processing: Use concurrency to speed up processing
// - Memory Optimization: Reduce memory usage for large datasets

// Add Support for Additional File Formats
// Add a Web Interface