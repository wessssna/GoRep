package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pastry struct {
	ID    string  `json:"id"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var pastries = []pastry{
	{ID: "1", Type: "Cupcake", Name: "Chocolate Delight", Price: 2.99},
	{ID: "2", Type: "Muffin", Name: "Blueberry Bliss", Price: 3.49},
	{ID: "3", Type: "Croissant", Name: "Butterflake Heaven", Price: 4.99},
	{ID: "4", Type: "Roll", Name: "Cinnamon Swirl", Price: 2.79},
}

func main() {
	router := gin.Default()
	router.GET("/pastries", getPastries)
	router.GET("/pastries/:id", getPastryByID)
	router.POST("/pastries", postPastry)
	router.DELETE("/pastries/:id", deletePastryByID)
	router.PUT("/pastries/:id", updatePastryByID)
	router.Run("localhost:8080")
}

func getPastries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pastries)
}

func postPastry(c *gin.Context) {
	var newPastry pastry

	if err := c.BindJSON(&newPastry); err != nil {
		return
	}

	pastries = append(pastries, newPastry)
	c.IndentedJSON(http.StatusCreated, newPastry)
}

func getPastryByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range pastries {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pastry not found"})
}

func deletePastryByID(c *gin.Context) {
	id := c.Param("id")
	for i, p := range pastries {
		if p.ID == id {
			pastries = append(pastries[:i], pastries[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "pastry deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pastry not found"})
}

func updatePastryByID(c *gin.Context) {
	id := c.Param("id")
	for i, p := range pastries {
		if p.ID == id {
			var updatedPastry pastry
			if err := c.BindJSON(&updatedPastry); err != nil {
				return
			}
			pastries[i] = updatedPastry
			c.IndentedJSON(http.StatusOK, updatedPastry)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pastry not found"})
}
