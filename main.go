package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/stivio00/pg_gotop/db"
	"github.com/stivio00/pg_gotop/ui"
	"golang.org/x/term"

	_ "github.com/lib/pq"
)

const currentVersion = "1.0.0"

func main() {
	if Options.ShowVersion {
		fmt.Printf("pg_gotop version %s\n", currentVersion)
		os.Exit(0)
	}

	if Options.ShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	password := ReadPassword("Enter Password: ")
	connStr := Options.CreateConnectionString(password)
	db := db.New(connStr)

	app := ui.New(db)
	app.BuildLayout()
	app.Run()
}

func ReadPassword(promt string) string {
	fmt.Print(promt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}
	return string(bytePassword)
}
