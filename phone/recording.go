package phone

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

type Recording struct {
	StartTime    time.Time
	AbsolutePath string
	EndTime      time.Time
}

func (r *Recording) String() string {
	return fmt.Sprintf(
		"{StartTime=%s, EndTime=%s, AbsolutePath=%s}",
		r.StartTime,
		r.EndTime,
		r.AbsolutePath,
	)
}

func GetRecordings(path string) ([]*Recording, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	recordings := make([]*Recording, 0)

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		fmt.Printf("Parsing phone data for %s\n", dir.Name())

		phoneDataTime, err := ParsePhoneDataTime(dir.Name())
		if err != nil {
			return nil, err
		}

		fmt.Printf("Getting video length for %s\n", dir.Name())
		duration, err := GetVideoLength(filepath.Join(path, dir.Name(), "recording.mp4"))
		if err != nil {
			if strings.Contains(err.Error(), "recording.mp4: no such file or directory") {
				fmt.Printf("Skipping %s because there is a missing recording.mp4 file\n", dir.Name())
				continue
			}
			return nil, err
		}

		recording := &Recording{
			StartTime:    phoneDataTime,
			EndTime:      phoneDataTime.Add(*duration),
			AbsolutePath: filepath.Join(path, dir.Name()),
		}

		fmt.Printf("Finished processing for %s\n", dir.Name())
		recordings = append(recordings, recording)
	}

	return recordings, nil
}

func GetVideoLength(path string) (*time.Duration, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ffprobe.ProbeReader(context.Background(), fileReader)
	if err != nil {
		return nil, err
	}

	duration := data.Format.Duration()
	return &duration, nil
}

const PHONE_TIME = "2006-01-02_15_04_05"

// ParsePhoneDataTime
// Example Time: 2023-02-16_11_47_01
func ParsePhoneDataTime(timeString string) (time.Time, error) {
	return time.Parse(PHONE_TIME, timeString)
}
