package main

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
