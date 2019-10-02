package main

import (
	"fmt"
	"os"
	"sort"
)

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// check if file exists
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// delete file
func deleteFile(file string) {
	if exists(file) {
		var err = os.Remove(file)
		if err != nil {
			return
		}
	}
}

func linksToString(links map[int64]int) (output [][]string) {
	keys := make([]int, len(links))
	i := 0
	total := 0
	for k, v := range links {
		keys[i] = int(k)
		i++
		total += v
	}
	sort.Ints(keys)
	for _, key := range keys {
		record := []string{
			fmt.Sprintf("%v", key),
			fmt.Sprintf("%v", float64(links[int64(key)])/float64(total)),
		}
		output = append(output, record)
	}
	return output
}
