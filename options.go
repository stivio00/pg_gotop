package main

import (
	"flag"
	"fmt"
)

type PgGoTopOptions struct {
	ShowHelp     bool
	ShowVersion  bool
	UserName     string
	DatabaseName string
	Port         int
	Host         string
}

const (
	ShowHelpDefault     = false
	ShowVersionDefault  = false
	UserNameDefault     = "postgres"
	DatabaseNameDefault = "postgres"
	PortDefault         = 5432
	HostDefault         = "localhost"
)

const (
	ShowHelpFlag    = "help"
	ShowVersionFlag = "version"
	UserNameFlag    = "user"
	DatabaseFlag    = "database"
	PortFlag        = "port"
	HostFlag        = "host"
)

const (
	RefreshRate = 100
)

func (o *PgGoTopOptions) Build(help bool, version bool, user string, database string, port int, host string) {
	o.ShowHelp = help
	o.ShowVersion = version
	o.UserName = user
	o.DatabaseName = database
	o.Host = host
	o.Port = port
}

var Options = PgGoTopOptions{}

func init() {
	showHelp := flag.Bool(ShowHelpFlag, ShowHelpDefault, "show help")
	showVersion := flag.Bool(ShowVersionFlag, ShowVersionDefault, "show version")
	user := flag.String(UserNameFlag, UserNameDefault, "database user")
	database := flag.String(DatabaseFlag, DatabaseNameDefault, "database name")
	port := flag.Int(PortFlag, PortDefault, "database connection port")
	host := flag.String(HostFlag, HostDefault, "database hostname")

	flag.Parse()
	Options.Build(*showHelp, *showVersion, *user, *database, *port, *host)
}

func (o PgGoTopOptions) CreateConnectionString(password string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		o.UserName,
		password,
		o.Host,
		o.Port,
		o.DatabaseName)
}
