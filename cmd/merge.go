package cmd

import (
	"cosine_cli/phone"
	"cosine_cli/trial"
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

		newRecordings := make([]*phone.Recording, 0)

		for _, recording := range recordings {
			if recording.EndTime.Sub(recording.StartTime).Seconds() < 60.0*4 {
				fmt.Printf("Excluding %s due to length of recording being under 4 minutes\n", recording.AbsolutePath)
				continue
			}
			newRecordings = append(newRecordings, recording)
		}

		phoneRecordings = append(phoneRecordings, newRecordings...)
	}

	//fmt.Printf("phoneRecordings=%v\n", phoneRecordings)

	data, err := vicon.NewViconData(viconRecordingPath)
	cobra.CheckErr(err)

	viconRecordings, err := data.Parse()
	cobra.CheckErr(err)

	sort.SliceStable(viconRecordings, func(i, j int) bool {
		return viconRecordings[i].StartTime.Before(viconRecordings[j].StartTime)
	})

	//fmt.Printf("viconRecordings=%v\n", viconRecordings)

	trials, err := trial.MatchRecordings(phoneRecordings, viconRecordings)
	cobra.CheckErr(err)

	fmt.Printf("trials=%s\n", trials)
	fmt.Printf("Trial count=%d\n", len(trials))

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "",
		"The path the merge output")

	rootCmd.AddCommand(mergeCmd)
}
