package phone

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Recording struct {
	Time         time.Time
	AbsolutePath string
}

func (r *Recording) String() string {
	return fmt.Sprintf("{Time=%s, AbsolutePath=%s}", r.Time, r.AbsolutePath)
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

		phoneDataTime, err := ParsePhoneDataTime(dir.Name())
		if err != nil {
			return nil, err
		}

		recording := &Recording{
			Time:         phoneDataTime,
			AbsolutePath: filepath.Join(path, dir.Name()),
		}

		recordings = append(recordings, recording)
	}

	return recordings, nil
}

const PHONE_TIME = "2006-01-02_15_04_05"

// ParsePhoneDataTime
// Example Time: 2023-02-16_11_47_01
func ParsePhoneDataTime(timeString string) (time.Time, error) {
	return time.Parse(PHONE_TIME, timeString)
}
