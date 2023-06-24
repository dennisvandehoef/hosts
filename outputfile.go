package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type OutputFile struct {
	Filename string
	Hosts    []string
}

func (o *OutputFile) DropDuplicateHosts() {
	uniqueMap := make(map[string]bool)

	// Slice to store the unique elements
	uniqueHosts := []string{}

	// Iterate over the original slice
	for _, element := range o.Hosts {
		// Check if the element is already present in the map
		if !uniqueMap[element] {
			// Add the element to the map and slice
			uniqueMap[element] = true
			uniqueHosts = append(uniqueHosts, element)
		}
	}

	sort.Strings(uniqueHosts)

	o.Hosts = uniqueHosts
}

func (o *OutputFile) WriteToFile() {
	o.PrintPiholeFormat()
	o.PrintHostsFormat()

	fmt.Printf("The file %s contains %d uniq hosts.\n", o.Filename, len(o.Hosts))
}

func (o *OutputFile) PrintPiholeFormat() {
	err := os.MkdirAll(filepath.Join("output", o.Filename), os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
	}

	file, err := os.Create(filepath.Join("output", o.Filename, "pihole"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, item := range o.Hosts {
		_, err := writer.WriteString(item + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func (o *OutputFile) PrintHostsFormat() {
	ipAddr := "0.0.0.0"

	err := os.MkdirAll(filepath.Join("output", o.Filename), os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
	}

	file, err := os.Create(filepath.Join("output", o.Filename, "hosts"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, item := range o.Hosts {
		_, err := writer.WriteString(fmt.Sprintf("%s	%s\n", ipAddr, item))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
