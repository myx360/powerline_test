package dto

type Meters struct {
	Meters []MeterData
}

type MeterData struct {
	Customer string
	Building string
	Id       string
}

type MeterUsage struct {
	Id    string
	Usage int
}
