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

	catFileCmd := &cobra.Command{
		Use:   "cat-file",
		Short: "It can print an object by its OID.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			optCatFile(args[0])
		},
	}

	writeTree := &cobra.Command{
		Use:   "write-tree",
		Short: "The write-tree is for storing a whole directory.",
		Run: func(cmd *cobra.Command, args []string) {
			optWriteTree()
		},
	}

	readTree := &cobra.Command{
		Use:   "read-tree",
		Short: "This command will take an OID of a tree and extract it to the working directory.",
		Run: func(cmd *cobra.Command, args []string) {
			optReadTree(args[0])
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(hashCmd)
	rootCmd.AddCommand(catFileCmd)
	rootCmd.AddCommand(writeTree)
	rootCmd.AddCommand(readTree)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

// Initialize a new git repository
func optInit() {

	res, err := data.Data_init()
	if err != nil {
		log.Fatalln("ugo-git:: Error:", err)
	}

	fmt.Println(res)

}

// Hash the content of the file using SHA-1.
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

	res, err := data.HashObject(fb, "blob")
	if err != nil {
		log.Fatalln(err)
	}

	obj := append([]byte("blob\x00"), fb...)
	err = data.SaveHashObj(obj, res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

}

// It can print an object by its OID.
func optCatFile(fl string) {

	res, err := data.GetObject(fl, "blob")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

}

func optWriteTree() {

	res, err := data.WriteTree(".")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)

}

func optReadTree(read string) {

	data.ReadTree(read)

}
