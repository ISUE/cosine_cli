package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	outputPath string
	mergeCmd   = &cobra.Command{
		Use:              "merge",
		Short:            "Merges the vicon and phone trial data",
		Run:              ExecuteMergeCmd,
		TraverseChildren: true,
	}
)

func ExecuteMergeCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("outputPath=%s\n", outputPath)
	fmt.Printf("viconTrialPath=%s\n", viconTrialPath)
	fmt.Printf("phoneTrialPaths=%s\n", phoneTrialPaths)

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "",
		"The path the merge output")

	rootCmd.AddCommand(mergeCmd)
}
