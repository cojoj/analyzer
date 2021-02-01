package models

import "net/url"

type Report struct {
	*url.URL
	DTD               string
	Title             string
	Headings          []Heading
	ContainsLoginForm bool
	Links             []Link
}

type Heading struct {
	Level  string
	Amount int
}

type Link struct {
	URL      *url.URL
	Internal bool
	Count    int
	Status   *Status
}

type Status struct {
	Reachable  bool
	StatusCode int
}
