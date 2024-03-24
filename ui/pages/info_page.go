package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/db"
)

func CreateInfoTree(db *db.DbConnection) *tview.TreeView {
	tree := tview.NewTreeView()
	tree.SetBackgroundColor(tcell.ColorBlue)
	tree.SetRoot(tview.NewTreeNode("Info"))

	settings := db.GetDbSettings()
	connection := db.GetDbCurrentConnection()
	locations := db.GetFileLocations()

	dbconnectionMode := tview.NewTreeNode("DatabaseConnection")
	tree.GetRoot().AddChild(dbconnectionMode)
	dbconnectionMode.AddChild(tview.NewTreeNode("Host: " + "localhost"))
	dbconnectionMode.AddChild(tview.NewTreeNode("Port: " + "5432"))
	dbconnectionMode.AddChild(tview.NewTreeNode("User: " + connection.User))
	dbconnectionMode.AddChild(tview.NewTreeNode("Database: " + connection.Database))

	dbNode := tview.NewTreeNode("Database")
	tree.GetRoot().AddChild(dbNode)
	dbNode.AddChild(tview.NewTreeNode("version: " + connection.Version))
	dbNode.AddChild(tview.NewTreeNode("Size: " + connection.Size))

	sNode := tview.NewTreeNode("Settings")
	tree.GetRoot().AddChild(sNode)
	sNode.AddChild(tview.NewTreeNode("SharedBufferSize: " + settings.SharedBufferSize))
	sNode.AddChild(tview.NewTreeNode("MaintenanceWorkMem: " + settings.MaintenanceWorkMem))
	sNode.AddChild(tview.NewTreeNode("TempBuffers: " + settings.TempBuffers))

	sNode.AddChild(tview.NewTreeNode("WorkingMem: " + settings.WorkingMem))
	sNode.AddChild(tview.NewTreeNode("EffectiveCacheSize: " + settings.EffectiveCacheSize))
	sNode.AddChild(tview.NewTreeNode("ListenAddresses: " + settings.ListenAddresses))
	sNode.AddChild(tview.NewTreeNode("MaxConnections: " + settings.MaxConnections))
	sNode.AddChild(tview.NewTreeNode("AutovacuumEnabled: " + settings.AutovacuumEnabled))
	sNode.AddChild(tview.NewTreeNode("Fillfactor: " + settings.Fillfactor))

	fNode := tview.NewTreeNode("File Locations")
	tree.GetRoot().AddChild(fNode)
	fNode.AddChild(tview.NewTreeNode("Data_directory: " + locations.Data_directory))
	fNode.AddChild(tview.NewTreeNode("Config_file: " + locations.Config_file))
	fNode.AddChild(tview.NewTreeNode("Hba_file: " + locations.Hba_file))
	fNode.AddChild(tview.NewTreeNode("Ident_file: " + locations.Ident_file))
	fNode.AddChild(tview.NewTreeNode("External_pid_file: " + locations.External_pid_file))

	return tree
}
