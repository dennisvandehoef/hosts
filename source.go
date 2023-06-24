package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Source struct {
	Directory string

	Name        string `json:"name"`
	Description string `json:"description"`
	HomeURL     string `json:"homeurl"`
	Issues      string `json:"issues"`
	Files       []File `json:"files"`
	License     string `json:"license"`
	Active      bool   `json:"active"`
}

func NewSourceFromFile(dirname, jsonFile string) (Source, error) {
	full_path := filepath.Join(dirname, jsonFile)
	_, err := os.Stat(full_path)
	if err != nil {
		return Source{}, err
	}

	// Read the contents of the "update.json" file
	data, err := ioutil.ReadFile(full_path)
	if err != nil {
		return Source{}, err
	}

	// Unmarshal the JSON data into the struct
	source := Source{Directory: dirname}
	err = json.Unmarshal(data, &source)
	if err != nil {
		return Source{}, err
	}

	return source, nil
}

func (s *Source) UpdateDownloadedFiles() {
	for _, f := range s.Files {
		// Send an HTTP GET request to the file URL
		response, err := http.Get(f.URL)
		if err != nil {
			fmt.Println("Error downloading:", err)
			return
		}
		defer response.Body.Close()

		// Create the output file
		file, err := os.Create(filepath.Join(s.Directory, f.FileName()))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// Copy the response body to the output file
		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Printf("Downloaded %s successfully!\n", f.URL)
	}
}

func (s *Source) ReadDownloadedFiles() {
	for i, f := range s.Files {
		file, err := NewHostsfileFromHostsFile(s.Directory, f.FileName())
		if err == nil {
			s.Files[i].Hostsfile = file
		}
		fmt.Printf("The file %s/%s contains %d hosts\n", s.Directory, f.FileName(), len(file.Hosts))
	}
}
