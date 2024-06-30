package data

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteTree(directory string) (string, error) {

	var slEntries [][]string
	var fType string

	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {

		path := fmt.Sprintf("%s/%s", directory, entry.Name())

		if isIgnored(path) {
			continue
		}

		// if is File
		if !entry.IsDir() {

			fType = "blob"

			fl, _ := os.Open(path)
			fb, _ := io.ReadAll(fl)
			oid, _ := HashObject(fb, "blob")

			// fmt.Printf("%s\t%s\n", oid, path)

		} else if entry.IsDir() {
			fType = "tree"
			oid, _ := WriteTree(path)
		}

		// slEntries = append(slEntries, {entry.})

	}

	return "", nil

}

func isIgnored(path string) bool {

	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part == ".ugit" || part == ".git" {
			return true
		}
	}

	return false
}
