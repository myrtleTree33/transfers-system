// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "gorm.io/driver/postgres"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
	dsn   = flags.String("dsn", "postgresql://local:local@localhost:5432/postgres", "postgres connection string")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	db, err := goose.OpenDBWithDriver("postgres", string(*dsn))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
