package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Hostsfile struct {
	Dirname  string
	Filename string
	Hosts    []string
}

func NewHostsfileFromHostsFile(dirname, filename string) (Hostsfile, error) {
	hostfile := Hostsfile{Dirname: dirname, Filename: filename}

	file, err := os.Open(filepath.Join(dirname, filename))
	if err != nil {
		fmt.Println("Error:", err)
		return hostfile, err
	}
	defer file.Close()

	// Define a regular expression pattern to match the domain name
	pattern := `(?:www\.)?([a-zA-Z0-9.-]+)`
	regex := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore lines starting with "#" as they are a comment
		if strings.HasPrefix(line, "#") {
			continue
		}
		// Ignore blank lines
		if len(line) > 0 {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				// Some source files give us lines with a comment character not on the first place
				if strings.HasPrefix(parts[len(parts)-1], "#") {
					continue
				}

				// Find the domain name in the string
				match := regex.FindStringSubmatch(parts[len(parts)-1])
				if len(match) > 1 {
					hostfile.Hosts = append(hostfile.Hosts, match[1])
				}

			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
		return hostfile, err
	}

	return hostfile, nil
}
