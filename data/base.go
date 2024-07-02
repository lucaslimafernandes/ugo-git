package data

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteTree(directory string) (string, error) {

	var slEntries [][]string
	var fType string
	var oid string

	entries, err := os.ReadDir(directory)
	if err != nil {
		return "", err
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
			oid, _ = HashObject(fb, fType)

			// fmt.Printf("%s\t%s\n", oid, path)

		} else if entry.IsDir() {
			fType = "tree"
			oid, _ = WriteTree(path)
		}

		sl := []string{entry.Name(), oid, fType}
		slEntries = append(slEntries, sl)

	}

	var sb strings.Builder
	for _, entry := range slEntries {
		sb.WriteString(fmt.Sprintf("%s %s %s\n", entry[2], entry[1], entry[0]))
	}
	tree := sb.String()

	return tree, nil

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

func iterTreeEntries(oid string) [][]string {

	var entries [][]string

	if oid == "" {
		return entries
	}

	tree, _ := GetObject(oid, "tree")
	scanner := bufio.NewScanner(bytes.NewReader([]byte(tree)))

	for scanner.Scan() {
		entry := scanner.Text()
		parts := strings.SplitN(entry, " ", 3)

		if len(parts) == 3 {
			entries = append(entries, parts)
		}
	}
	return entries

}

func getTree(oid string, basePath string) map[string]string {

	result := make(map[string]string)

	for _, entry := range iterTreeEntries(oid) {

		type_ := entry[0]
		oid := entry[1]
		name := entry[2]

		if strings.Contains(name, "/") || name == ".." || name == "." {
			continue
		}

		path := basePath + name

		if type_ == "blob" {
			result[path] = oid
		} else if type_ == "tree" {
			for k, v := range getTree(oid, path+"/") {
				result[k] = v
			}
		} else {
			fmt.Printf("Unknown tree entry %s\n", type_)
		}
	}

	return result

}

func ReadTree(treeOid string) {

	for path, oid := range getTree(treeOid, "./") {

		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		data, _ := GetObject(oid, "blob")

		err := os.WriteFile(path, []byte(data), os.ModePerm)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", path, err)
		}
	}
}
