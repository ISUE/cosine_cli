package vicon

import (
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	Path string
}

func (d *Data) Parse() ([]*Trial, error) {
	dirs, err := os.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	trials := make([]*Trial, 0)

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		if !strings.HasSuffix(dir.Name(), ".xcp") {
			continue
		}

		trialPath := filepath.Join(d.Path, dir.Name())
		trial, err := NewTrial(trialPath)
		if err != nil {
			return nil, err
		}
		trials = append(trials, trial)
	}

	return trials, nil
}

func NewViconData(path string) (*Data, error) {

	return &Data{Path: path}, nil
}
