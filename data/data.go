package data

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func Data_init() (string, error) {

	var res string
	curr_path, _ := os.Getwd()

	err := os.Mkdir(".ugit/", os.ModePerm)
	if err != nil {

		if err.Error() == "mkdir .ugit/: file exists" {

			res = fmt.Sprintf("ugo-git:: ugit repository already exists in: %s/.ugit", curr_path)
			return res, nil

		} else {
			return "", err
		}
	}

	err = os.Mkdir(".ugit/objects/", os.ModePerm)
	if err != nil {
		return "", err
	}

	res = fmt.Sprintf("ugo-git:: Initialized empty ugit repository in: %s/.ugit", curr_path)
	return res, nil

}

func HashObject(data []byte) (hash.Hash, error) {

	oid := sha1.New()

	_, err := io.WriteString(oid, string(data))

	return oid, err

}

func SaveHashObj(d []byte, h hash.Hash) error {

	hb := hex.EncodeToString(h.Sum(nil))

	fp := fmt.Sprintf(".ugit/objects/%v", hb)

	f, err := os.Create(fp)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(d))

	return err

}
