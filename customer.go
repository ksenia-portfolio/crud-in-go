package main

import (
	"time"
)

type Customer struct{
	Id uint16
	DOB time.Time
	FirstName,
	LastName,
	Gender,
	Email,
	Address   string
}

func (c Customer) FormatDOB() string{
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	dob := c.DOB  // example: 1999-12-31 00:00:00 +0000 UTC
	return dob.Format(layoutISO) // example: December 31, 1999

}
