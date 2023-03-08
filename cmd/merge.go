package cmd

import (
	"cosine_cli/phone"
	"cosine_cli/vicon"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

var (
	outputPath string
	mergeCmd   = &cobra.Command{
		Use:              "merge",
		Short:            "Merges the vicon and phone recording data",
		Run:              ExecuteMergeCmd,
		TraverseChildren: true,
	}
)

func ExecuteMergeCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("outputPath=%s\n", outputPath)
	fmt.Printf("viconRecordingPath=%s\n", viconRecordingPath)
	fmt.Printf("phoneRecordingPaths=%s\n", phoneRecordingPaths)

	phoneRecordings := make([]*phone.Recording, 0)

	for _, phoneRecordingPath := range phoneRecordingPaths {
		recordings, err := phone.GetRecordings(phoneRecordingPath)
		cobra.CheckErr(err)
		phoneRecordings = append(phoneRecordings, recordings...)
	}

	fmt.Printf("phoneRecordings=%v\n", phoneRecordings)

	data, err := vicon.NewViconData(viconRecordingPath)
	cobra.CheckErr(err)

	recordings, err := data.Parse()
	cobra.CheckErr(err)

	sort.SliceStable(recordings, func(i, j int) bool {
		return recordings[i].StartTime.Before(recordings[j].StartTime)
	})

	fmt.Printf("recordings=%v\n", recordings)

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "",
		"The path the merge output")

	rootCmd.AddCommand(mergeCmd)
}
