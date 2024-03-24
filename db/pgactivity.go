package db

import "time"

type PgStatActivity struct {
	Datid           uint      //OID of the database this backend is connected to
	Datname         string    //Name of the database this backend is connected to
	Pid             int       //Process ID of this backend
	LeaderPid       int       //Process ID of the parallel group leader if this process is a parallel query worker, or process ID of the leader apply worker if this process is a parallel apply worker. NULL indicates that this process is a parallel group leader or leader apply worker, or does not participate in any parallel operation.
	Usesysid        uint      //OID of the user logged into this backend
	Usename         string    //Name of the user logged into this backend
	ApplicationName string    //Name of the application that is connected to this backend
	ClientAddr      string    //IP address of the client connected to this backend. If this field is null, it indicates either that the client is connected via a Unix socket on the server machine or that this is an internal process such as autovacuum.
	ClientHostname  string    //Host name of the connected client, as reported by a reverse DNS lookup of client_addr. This field will only be non-null for IP connections, and only when log_hostname is enabled.
	ClientPort      int       //TCP port number that the client is using for communication with this backend, or -1 if a Unix socket is used. If this field is null, it indicates that this is an internal server process.
	BackendStart    time.Time //Time when this process was started. For client backends, this is the time the client connected to the server.
	XactStart       time.Time //Time when this process' current transaction was started, or null if no transaction is active. If the current query is the first of its transaction, this column is equal to the query_start column.
	QueryStart      time.Time //Time when the currently active query was started, or if state is not active, when the last query was started
	StateChange     time.Time //Time when the state was last changed
	WaitEventType   string    //The type of event for which the backend is waiting, if any; otherwise NULL. See Table 28.4.
	WaitEvent       string    //Wait event name if backend is currently waiting, otherwise NULL. See Table 28.5 through Table 28.13.
	State           string    //state
	BackendXid      int       //Top-level transaction identifier of this backend, if any; see Section 74.1.
	BackendXmin     int       //The current backend's xmin horizon.
	QueryId         int64     //Identifier of this backend's most recent query. If state is active this field shows the identifier of the currently executing query. In all other states, it shows the identifier of last query that was executed. Query identifiers are not computed by default so this field will be null unless compute_query_id parameter is enabled or a third-party module that computes query identifiers is configured.
	Query           string    //Text of this backend's most recent query. If state is active this field shows the currently executing query. In all other states, it shows the last query that was executed. By default the query text is truncated at 1024 bytes; this value can be changed via the parameter track_activity_query_size.
	BackendType     string    //bet
}

const (
	GetActivityQuery = "SELECT * FROM pg_catalog.pg_stat_activity;"
)
