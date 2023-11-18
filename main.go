package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

// album represents data about a record album.
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album {
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	app := gin.Default()
    app.GET("/albums", getAlbums)
	app.POST("/albums", createAlbum)
	app.GET("/albums/:id", getAlbumById)
	app.DELETE("/albums/:id", removeAlbumById)

    app.Run("localhost:8080")
}

// fetches all albums in the system
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// creates a new album from data submitted by the client
func createAlbum(context *gin.Context) {
	var newAlbum Album

	// gin's BindJSON method binds the request json to appropriate struct
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// If the request JSON is OK. add it to the album collection
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, albums)
}

// fetches an existing album by its id
func getAlbumById(context *gin.Context) {
	id := context.Param("id")
	
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func removeAlbumById (context *gin.Context) {
	id := context.Param("id")

	for index, album := range albums {
		if album.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			fmt.Println("Updated albums", albums)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "album was deleted"})
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
