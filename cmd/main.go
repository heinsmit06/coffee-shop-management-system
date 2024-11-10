package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"hot-coffee/internal/handler"
)

var (
	dirPath       string
	port          string
	nameBlackList = []string{"objects.csv", "buckets.csv"}
)

func main() {
	// err := os.MkdirAll("./test/another", 0750)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	Run()
}

func Run() {
	dirPath, port = InitFlags()

	// check if the directory already exist
	// _, err := os.Stat(dirPath)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 	}
	// } else {
	// 	err = fmt.Errorf("Directory %s already exists", dirPath)
	// 	log.Fatal(err)
	// }
	err := os.MkdirAll(dirPath, 0750)
	if err != nil {
		log.Fatal("Failed to make directory: ", err)
	}

	mux := handler.SetupRoutes()

	fmt.Println("Server started on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func InitFlags() (string, string) {
	var dirPath string
	var port string
	var help bool

	flag.BoolVar(&help, "help", false, "-help for usage information")
	flag.StringVar(
		&dirPath,
		"dir",
		"",
		"path to the directory where the files will be stored as arguments",
	)

	flag.StringVar(&port, "port", "8080", "port number API gonna listening to")
	flag.Parse()

	if help {
		fmt.Println(`Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
		os.Exit(0)
	}

	if dirPath == "" {
		fmt.Fprintf(
			os.Stderr,
			"Please determine path by -dir flag\n Use -help flag for more information\n",
		)
		os.Exit(1)
	}

	dirPath += "/"

	return dirPath, port
}
