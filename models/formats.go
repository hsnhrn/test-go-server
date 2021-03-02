package models

type Data struct {
	Name  string `json:"Name"`
	Id    int    `json:"Id"`
	Value int    `json:"Value"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
	Type  int    `json:"Type"`
	Unit  int    `json:"unit"`
}
type Format struct {
	FRev   int    `json:"FRev"`
	Format []Data `json:"Data"`
}
