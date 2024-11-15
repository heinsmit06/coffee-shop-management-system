package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"hot-coffee/internal/handler"
)

var (
	dirPath       string
	port          string
	nameBlackList = []string{"objects.csv", "buckets.csv"}
)

func main() {
	Run()
}

func Run() {
	dirPath, port = InitFlags()

	err := os.MkdirAll(dirPath, 0o750)
	if err != nil {
		log.Fatal("Failed to make directory: ", err)
	}

	mux := handler.SetupServer(dirPath)

	fmt.Println("Server started on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func InitFlags() (string, string) {
	var dirPath string
	var port string
	var help bool

	flag.BoolVar(&help, "help", false, "usage information")
	flag.StringVar(
		&dirPath,
		"dir",
		"",
		"path to the directory where the files will be stored as arguments",
	)

	flag.StringVar(&port, "port", "8000", "port number API gonna listening to")
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

	portNumber, _ := strconv.Atoi(port)
	if portNumber < 1024 || portNumber > 65535 {
		fmt.Fprintf(os.Stderr, "Port number is not allowed: must be in between [1024, 65535]\n")
		os.Exit(1)
	}
	// if dirPath == "" {
	// 	fmt.Fprintf(
	// 		os.Stderr,
	// 		"Please determine path by -dir flag\n Use -help flag for more information\n",
	// 	)
	// 	os.Exit(1)
	// }

	if dirPath != "" {
		dirPath = filepath.Clean(dirPath)
		if strings.HasPrefix(dirPath, "..") {
			fmt.Fprintf(os.Stderr, "Invalid directory path: %s\n", dirPath)
			os.Exit(1)
		}
	} else {
		dirPath = "data" // Set default directory if not provided
	}

	// Convert relative path to absolute
	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %s\n", err)
		os.Exit(1)
	}
	absDirPath += "/"

	fmt.Println(absDirPath)
	return absDirPath, port
}
