package app

import (
	"cwa/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerApp struct {
}

func NewServerApp() *ServerApp {
	return &ServerApp{}
}

func (s *ServerApp) Start() {
	r := gin.Default()
	r.GET("/customer/:name", s.HandleGetCustomer)
	r.GET("/meter/:id", s.HandleGetUsage)
	r.Run(":8080")
}

func (s *ServerApp) HandleGetCustomer(c *gin.Context) {
	aquaflow := dto.Meters{
		Meters: []dto.MeterData{
			{Customer: "Aquaflow", Building: "Treatment Plant A", Id: "1111-1111-1111"},
			{Customer: "Aquaflow", Building: "Treatment Plant B", Id: "1111-1111-2222"}}}

	albers := dto.Meters{
		Meters: []dto.MeterData{
			{Customer: "Albers Facilities Management", Building: "Student Halls", Id: "1111-1111-3333"}}}

	name := c.Param("name")

	switch name {
	case "Aquaflow":
		c.IndentedJSON(http.StatusOK, aquaflow)
	case "Albers":
		c.IndentedJSON(http.StatusOK, albers)
	default:
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not found"})
	}
}

func (s *ServerApp) HandleGetUsage(c *gin.Context) {
	meterUsageMap := map[string]int{"1111-1111-1111": 2000, "1111-1111-2222": 3000, "1111-1111-3333": 4000}
	id := c.Param("id")
	if meterUsageMap[id] == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Meter not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, dto.MeterUsage{Id: id, Usage: meterUsageMap[id]})
}
