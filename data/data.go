package data

import (
	"fmt"
	"os"
)

// Initialize .ugit directory
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
