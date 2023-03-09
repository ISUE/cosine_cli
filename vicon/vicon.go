package vicon

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	Path string
}

func (d *Data) Parse() ([]*Recording, error) {
	dirs, err := os.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	recordings := make([]*Recording, 0)

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		if !strings.HasSuffix(dir.Name(), ".xcp") {
			continue
		}

		recordingPath := filepath.Join(d.Path, dir.Name())

		if len(strings.Split(recordingPath, "_")) < 6 {
			log.Printf("Ignoring %s because it does not fit our file format\n", dir.Name())
			continue
		}

		recording, err := NewRecording(recordingPath)
		if err != nil {
			log.Printf("Ignoring %s due to error: %v\n", recordingPath, err)
			continue
		}
		recordings = append(recordings, recording)
	}

	return recordings, nil
}

func NewViconData(path string) (*Data, error) {
	return &Data{Path: path}, nil
}
