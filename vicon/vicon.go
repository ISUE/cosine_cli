package vicon

import (
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
		recording, err := NewRecording(recordingPath)
		if err != nil {
			return nil, err
		}
		recordings = append(recordings, recording)
	}

	return recordings, nil
}

func NewViconData(path string) (*Data, error) {
	return &Data{Path: path}, nil
}
