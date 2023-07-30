package model

import "time"

type Request struct {
	Currency    string
	Average     float64
	StartDate   time.Time
	EndDate     time.Time
	RequestDate time.Time
}
