package main

import (
	"github.com/bearbin/go-age"
	_ "github.com/bearbin/go-age"
	"regexp"
	"strings"
	"time"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type CustomerForm struct {
	Id uint16
	FirstName,
	LastName,
	Gender,
	DOB,
	Email,
	Address   string
	Errors 	  map[string]string
}

func (cf *CustomerForm) ValidateCreateForm() bool{
	cf.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(cf.Email))



	// first name max length is 100; is required; not empty
	if len(cf.FirstName) > 100{
		cf.Errors["FirstName"] = "Can't exceed 100 characters long."
	}
	if len(cf.FirstName) == 0{
		cf.Errors["FirstName"] = "Is required field."
	}
	if strings.TrimSpace(cf.FirstName) == ""{
		cf.Errors["FirstName"] = "Can't contain spaces only."
	}

	// last name max length is 100; is required; not empty
	if len(cf.LastName) > 100{
		cf.Errors["LastName"] = "Can't exceed 100 characters."
	}
	if len(cf.LastName) == 0{
		cf.Errors["LastName"] = "Is required field."
	}
	if strings.TrimSpace(cf.LastName) == ""{
		cf.Errors["LastName"] = "Can't contain spaces only."
	}

	// gender is required;
	if len(cf.Gender) == 0{
		cf.Errors["Gender"] = "Is required field."
	}

	// dob - age from 18 to 60; required
	if cf.IsValidAge() == false{
		cf.Errors["DOB"] = "Age should be from 18 to 60."
	}

	// email has valid email pattern; required
	if match == false {
		cf.Errors["Email"] = "Please enter a valid email address."
	}
	// address max length is 200; optional
	if len(cf.Address) > 200{
		cf.Errors["Address"] = "Can't exceed 200 characters."
	}
	return len(cf.Errors) == 0
}

func (cf *CustomerForm) ValidateUpdateForm() bool{
	cf.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(cf.Email))



	// first name max length is 100; is required; not empty
	if len(cf.FirstName) > 100{
		cf.Errors["FirstName"] = "Can't exceed 100 characters long."
	}
	if len(cf.FirstName) == 0{
		cf.Errors["FirstName"] = "Is required field."
	}
	if strings.TrimSpace(cf.FirstName) == ""{
		cf.Errors["FirstName"] = "Can't contain spaces only."
	}

	// last name max length is 100; is required; not empty
	if len(cf.LastName) > 100{
		cf.Errors["LastName"] = "Can't exceed 100 characters."
	}
	if len(cf.LastName) == 0{
		cf.Errors["LastName"] = "Is required field."
	}
	if strings.TrimSpace(cf.LastName) == ""{
		cf.Errors["LastName"] = "Can't contain spaces only."
	}

	// dob - age from 18 to 60; required
	if cf.IsValidAge() == false{
		cf.Errors["DOB"] = "Age should be from 18 to 60."
	}

	// email has valid email pattern; required
	if match == false {
		cf.Errors["Email"] = "Please enter a valid email address."
	}
	// address max length is 200; optional
	if len(cf.Address) > 200{
		cf.Errors["Address"] = "Can't exceed 200 characters."
	}
	return len(cf.Errors) == 0
}

func (cf *CustomerForm) IsValidAge() bool {
	cDOB := GetDOB(cf.DOB)
	cAge := age.Age(cDOB)

	return cAge >= 18 && cAge <= 60
}

func GetDOB(dob string) time.Time {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	t, _ := time.Parse(layoutISO, dob)
	return t
}

