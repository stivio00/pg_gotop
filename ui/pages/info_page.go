package pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stivio00/pg_gotop/db"
)

func CreateInfoTree(db *db.DbConnection) *tview.TreeView {
	tree := tview.NewTreeView()
	tree.SetGraphics(true)
	tree.SetBackgroundColor(tcell.ColorBlue)

	rootNode := tview.NewTreeNode("Info")
	rootNode.SetSelectable(true)
	tree.SetRoot(rootNode)

	settings := db.GetDbSettings()
	connection := db.GetDbCurrentConnection()
	locations := db.GetFileLocations()

	dbconnectionNode := createDbConnectionNode(connection)
	tree.GetRoot().AddChild(dbconnectionNode)

	dbNode := createDbNode(connection)
	tree.GetRoot().AddChild(dbNode)

	sNode := createSettingsNode(settings)
	tree.GetRoot().AddChild(sNode)

	fNode := createfileNode(locations)
	tree.GetRoot().AddChild(fNode)

	return tree
}

func createfileNode(locations db.FileLocations) *tview.TreeNode {
	fNode := tview.NewTreeNode("File Locations")

	fNode.AddChild(tview.NewTreeNode("Data_directory: " + locations.Data_directory))
	fNode.AddChild(tview.NewTreeNode("Config_file: " + locations.Config_file))
	fNode.AddChild(tview.NewTreeNode("Hba_file: " + locations.Hba_file))
	fNode.AddChild(tview.NewTreeNode("Ident_file: " + locations.Ident_file))
	fNode.AddChild(tview.NewTreeNode("External_pid_file: " + locations.External_pid_file))

	return fNode
}

func createSettingsNode(settings db.DbSettings) *tview.TreeNode {
	sNode := tview.NewTreeNode("Settings")

	sNode.AddChild(tview.NewTreeNode("SharedBufferSize: " + settings.SharedBufferSize))
	sNode.AddChild(tview.NewTreeNode("MaintenanceWorkMem: " + settings.MaintenanceWorkMem))
	sNode.AddChild(tview.NewTreeNode("TempBuffers: " + settings.TempBuffers))

	sNode.AddChild(tview.NewTreeNode("WorkingMem: " + settings.WorkingMem))
	sNode.AddChild(tview.NewTreeNode("EffectiveCacheSize: " + settings.EffectiveCacheSize))
	sNode.AddChild(tview.NewTreeNode("ListenAddresses: " + settings.ListenAddresses))
	sNode.AddChild(tview.NewTreeNode("MaxConnections: " + settings.MaxConnections))
	sNode.AddChild(tview.NewTreeNode("AutovacuumEnabled: " + settings.AutovacuumEnabled))
	sNode.AddChild(tview.NewTreeNode("Fillfactor: " + settings.Fillfactor))

	return sNode
}

func createDbNode(connection db.DbCurrentConnection) *tview.TreeNode {
	dbNode := tview.NewTreeNode("Database")
	dbNode.SetSelectable(true)

	dbNode.AddChild(tview.NewTreeNode("version: " + connection.Version))
	dbNode.AddChild(tview.NewTreeNode("Size: " + connection.Size))

	return dbNode
}

func createDbConnectionNode(connection db.DbCurrentConnection) *tview.TreeNode {
	dbconnectionNode := tview.NewTreeNode("DatabaseConnection")
	dbconnectionNode.SetSelectable(true)

	dbconnectionNode.AddChild(tview.NewTreeNode("Host: " + "localhost"))
	dbconnectionNode.AddChild(tview.NewTreeNode("Port: " + "5432"))
	dbconnectionNode.AddChild(tview.NewTreeNode("User: " + connection.User))
	dbconnectionNode.AddChild(tview.NewTreeNode("Database: " + connection.Database))

	return dbconnectionNode
}
