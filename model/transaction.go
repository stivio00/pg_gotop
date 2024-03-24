package model

import "time"

type Transaction struct {
	Xactid   string
	Sql      string
	Started  time.Time
	Duration time.Duration
	Pid      int
}
