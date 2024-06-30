package data

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// Creates a hash (sha1) for the given data and type
// Returns: The hash id
func HashObject(data []byte, objectType string) (string, error) {

	obj := []byte(objectType + "\x00")
	obj = append(obj, data...)

	oid := sha1.New()
	_, err := oid.Write(obj)

	strOid := hex.EncodeToString(oid.Sum(nil))

	return strOid, err

}

// Saves the object data (type + data) using the given hasd ID
func SaveHashObj(obj []byte, h string) error {

	fp := fmt.Sprintf(".ugit/objects/%v", h)

	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(obj)

	return err

}

// Retrieves and verifies the object with the given hash ID and expected type (optional)
// Returns: data of object
func GetObject(oid string, expectedType string) (string, error) {

	fp := fmt.Sprintf(".ugit/objects/%v", oid)

	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer f.Close()

	obj, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	parts := strings.SplitN(string(obj), "\x00", 2)
	if len(parts) < 2 {
		println(parts)
		return "", fmt.Errorf("malformed object")
	}

	objectType := parts[0]
	content := parts[1]
	if expectedType != "" && objectType != expectedType {
		return "", fmt.Errorf("expected: %s, got: %s", expectedType, objectType)
	}

	return string(content), err

}
