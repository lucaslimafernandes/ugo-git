package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lucaslimafernandes/ugo-git/data"
	"github.com/spf13/cobra"
)

func main() {
	parseArgs()
}

func parseArgs() {
	rootCmd := &cobra.Command{
		Use:   "ugit",
		Short: "uGit is a version control system",
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new git repository",
		Run: func(cmd *cobra.Command, args []string) {
			optInit()
		},
	}

	hashCmd := &cobra.Command{
		Use:   "hash-object",
		Short: "Hash the content of the file using SHA-1.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			optHashObj(args[0])
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(hashCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func optInit() {

	res, err := data.Data_init()
	if err != nil {
		log.Fatalln("ugo-git:: Error:", err)
	}

	fmt.Println(res)

}

func optHashObj(fl string) {

	f, err := os.Open(fl)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fb, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := data.HashObject(fb)
	if err != nil {
		log.Fatalln(err)
	}

	data.SaveHashObj(fb, res)

}
