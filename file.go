package main

import (
	"fmt"
	"path"
	"strings"
)

type File struct {
	Hostsfile Hostsfile

	Output string `json:"output"`
	URL    string `json:"url"`
}

func (f *File) FileName() string {
	_, filename := path.Split(f.URL)

	parts := []string{
		f.Output,
		fmt.Sprintf("%d", len(f.URL)),
		filename,
	}

	return strings.Join(parts, "-")
}
