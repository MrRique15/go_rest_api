package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Services "example/web-service-gin/services"
	Structs "example/web-service-gin/structs"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	var albumsList = Services.ListAlbums()

    c.IndentedJSON(http.StatusOK, albumsList)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum Structs.Album

	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

	newAlbum = Services.AddAlbum(newAlbum)

	if(newAlbum.ID == ""){
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "An album already exists with given ID"})
		return
	}

    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

	if(id == "" || id == " "){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "A valid ID must be inserted"})
	}

	var album = Services.GetAlbumById(id)

	fmt.Println(album)

	if(album.ID == ""){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusFound, album)
}