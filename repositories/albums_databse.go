package Repositores

import (
	Structs "example/web-service-gin/structs"
)

// albums slice to seed record album data.
var Albums = []Structs.Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func ListAlbums() (albumsList []Structs.Album){
	albumsList = Albums
	return
}

func AddAlbum(album Structs.Album) (newAlbum Structs.Album){
    // Add the new album to the slice.
    Albums = append(Albums, album)

	newAlbum = Albums[len(Albums)-1]
	return
}

func FindAlbumById(id string)(album Structs.Album){
	// Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range Albums {
        if a.ID == id {
            // c.IndentedJSON(http.StatusOK, a)
			album = a
            return
        }
    }

	return
}