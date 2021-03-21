package main

type SearchForm struct{
	FirstName,
	LastName string
}

func (sf *SearchForm) ValidateSearchForm() bool{
	// first name or last name are required
	return len(sf.FirstName) > 0 || len(sf.LastName) > 0
}