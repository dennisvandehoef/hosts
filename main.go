package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
}

func main() {
	sources := readSourcesFromFile()

	for _, s := range sources {
		s.UpdateDownloadedFiles()
		s.ReadDownloadedFiles()

		fmt.Println()
	}

	outputs := make(map[string]OutputFile)

	for _, s := range sources {
		if s.Active {
			for _, f := range s.Files {
				if o, ok := outputs[f.Output]; ok {
					o.Hosts = append(o.Hosts, f.Hostsfile.Hosts...)
					outputs[f.Output] = o
				} else {
					outputs[f.Output] = OutputFile{Filename: f.Output, Hosts: f.Hostsfile.Hosts}
				}
			}
		}
	}

	for _, of := range outputs {
		of.DropDuplicateHosts()
		of.WriteToFile()
	}
}

func readSourcesFromFile() []Source {
	sources := []Source{}

	err := filepath.Walk("source_data", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			s, err := NewSourceFromFile(path, "update.json")
			if err == nil {
				sources = append(sources, s)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	return sources
}
