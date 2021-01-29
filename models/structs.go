package models

// Data that we take from json file
type Data struct {
	Title       string
	Link        string
	Description string
	Content     string
}

// DataCont ...
type DataCont struct {
	Name string
	Data []*Data
}
