package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Breed string `json:"breed"`
	Path  string `json:"path"`
}

var cats = []cat{
	{ID: "0", Name: "Cookie", Color: "Calico", Breed: "Calico", Path: "s3://gpi-cats/Cookie.jpg"},
	{ID: "1", Name: "Pucky", Color: "White", Breed: "White", Path: "s3://gpi-cats/Pucky.jpg"},
	{ID: "2", Name: "Pixie", Color: "Tabby", Breed: "Tabby", Path: "s3://gpi-cats/Pixie.jpg"},
	{ID: "3", Name: "Morris", Color: "Orange", Breed: "Orange", Path: "s3://gpi-cats/Morris.jpg"},
	{ID: "4", Name: "Misty", Color: "Gray", Breed: "Tuxedo", Path: "s3://gpi-cats/Misty.jpg"},
	{ID: "5", Name: "Cassie", Color: "Gray", Breed: "Big and Fluffy", Path: "s3://gpi-cats/Cassie.jpg"},
}

// getCats returns list of all Cats as JSON.
func getCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

// Add a cat to the "database" of cats.
func postCats(c *gin.Context) {
	var newCat cat

	// Call BindJSON to bind the received JSON to newCat
	if err := c.BindJSON(&newCat); err != nil {
		return
	}

	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

// Get a cat by ID.
func getCatByID(c *gin.Context) {
	id := c.Param("id")

	for _, myCat := range cats {
		if myCat.ID == id {
			c.IndentedJSON(http.StatusOK, myCat)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cat not found"})
}

func main() {
	r := gin.Default()
	var getCatEndPoint func(*gin.Context)
	getCatEndPoint = getCats
	r.GET("/cats", getCatEndPoint)

	var postCatEndpoint func(*gin.Context)
	postCatEndpoint = postCats
	r.POST("/cats", postCatEndpoint)

	var getCatByIDEndpoint func(*gin.Context)
	getCatByIDEndpoint = getCatByID
	r.GET("/cat/:id", getCatByIDEndpoint)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
