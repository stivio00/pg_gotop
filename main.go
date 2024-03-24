package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/term"

	_ "github.com/lib/pq"
)

const currentVersion = "1.0.0"

var options = PgGoTopOptions{}

func init() {
	showHelp := flag.Bool(ShowHelpFlag, ShowHelpDefault, "show help")
	showVersion := flag.Bool(ShowVersionFlag, ShowVersionDefault, "show version")
	user := flag.String(UserNameFlag, UserNameDefault, "database user")
	database := flag.String(DatabaseFlag, DatabaseNameDefault, "database name")
	port := flag.Int(PortFlag, PortDefault, "database connection port")
	host := flag.String(HostFlag, HostDefault, "database hostname")

	flag.Parse()
	options.Build(*showHelp, *showVersion, *user, *database, *port, *host)
}

func main() {
	if options.ShowVersion {
		fmt.Printf("pg_gotop version %s\n", currentVersion)
		os.Exit(0)
	}

	if options.ShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	password := ReadPassword("Enter Password: ")
	connStr := CreateConnectionString(options, password)
	db := Connect(connStr)
	postgresVersion := getVersion(db)
	fmt.Printf("Postgresql version : %s", postgresVersion)

	pgActivity := getActivity(db)

	// TUI
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(true)

	table.SetCell(0, 0, tview.NewTableCell("database").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("username").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 2, tview.NewTableCell("app").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 3, tview.NewTableCell("pid").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 4, tview.NewTableCell("state").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
	table.SetCell(0, 5, tview.NewTableCell("type").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))

	rows := len(pgActivity)
	for r := 0; r < rows; r++ {
		color := tcell.ColorWhite

		table.SetCell(r+1, 0, tview.NewTableCell(pgActivity[r].Datname).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 1, tview.NewTableCell(pgActivity[r].Usename).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 2, tview.NewTableCell(pgActivity[r].ApplicationName).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 3, tview.NewTableCell(strconv.Itoa(pgActivity[r].Pid)).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r+1, 4, tview.NewTableCell(pgActivity[r].State).SetTextColor(color).SetAlign(tview.AlignCenter))
		table.SetCell(r, 5, tview.NewTableCell(pgActivity[r].BackendType).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}

func getActivity(db *sql.DB) []PgStatActivity {
	rows, err := db.Query(GetActivityQuery)
	if err != nil {
		log.Fatal(err)
	}
	var activity []PgStatActivity = make([]PgStatActivity, 0)
	for rows.Next() {
		row := PgStatActivity{}
		rows.Scan(&row.Datid,
			&row.Datname,
			&row.Pid,
			&row.LeaderPid,
			&row.Usesysid,
			&row.Usename,
			&row.ApplicationName,
			&row.ClientAddr,
			&row.ClientHostname,
			&row.ClientPort,
			&row.BackendStart,
			&row.XactStart,
			&row.QueryStart,
			&row.StateChange,
			&row.WaitEventType,
			&row.WaitEvent,
			&row.State,
			&row.BackendXid,
			&row.BackendXmin,
			&row.QueryId,
			&row.Query,
			&row.BackendType)
		activity = append(activity, row)
	}
	return activity
}

func getVersion(db *sql.DB) string {
	row, err := db.Query("SELECT version();")
	if err != nil {
		log.Fatal(err)
	}
	var dbVersion string
	row.Next()
	row.Scan(&dbVersion)
	return dbVersion
}

func Connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateConnectionString(options PgGoTopOptions, password string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		options.UserName,
		password,
		options.Host,
		options.Port,
		options.DatabaseName)
}

func ReadPassword(promt string) string {
	fmt.Print(promt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	return string(bytePassword)
}
