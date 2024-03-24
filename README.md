# PG Go Top
Another postgresql monitoring tool.


# Intro
pg_gotp is a `top`/`htop` alike utility to monitor postgresql database.
It opens a database conection and lists diferent resource like: 
current backend processes, long running transaction, and IO and memory resources.

It uses interanally the pg_catalog system tables like `pg_stats_activity` and uses system administration funcions like `pg_terminate_backend` to administrates a monitor a postgresql database

# Usage

To show the options 
```bash
$ pg_gotop --help
$ pg_gotop -h
```

To show the current version
```bash
$ pg_gotop --version
$ pg_gotop -v
```

Running pg_gotop will prompt the password if needed to connect to database. 

The default values: 
* **User**: postgres
* **Database Name**: postgres
* **Host**: localhost
* **Port**: 5432 

To use a specific database connection use :
```bash
$ pg_gotop -d $db -U $user -h $host -p $port
password:  
```

# Dependencies
* **tcel**: https://github.com/gdamore/tcell : terminal UI lib
* **pq**: https://github.com/lib/pq : postgres driver


# Build and test

Build project:
```bash
$ go build -o bin/
```

run
```bash
$ bin/pg_gotop
```
