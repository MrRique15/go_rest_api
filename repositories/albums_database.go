package Repositores

import (
	"errors"
	Structs "go_rest_api/structs"
)

var Albums = []Structs.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func ListAlbums() []Structs.Album {
	albumsList := Albums
	return albumsList
}

func AddAlbum(album Structs.Album) (Structs.Album, error) {
	Albums = append(Albums, album)

	newAlbum := Albums[len(Albums)-1]

	if(newAlbum.ID == ""){
		return newAlbum, errors.New("error during album insertion")
	}

	return newAlbum, nil
}

func FindAlbumById(id string) (Structs.Album, error) {
	var album Structs.Album

	for _, a := range Albums {
		if a.ID == id {
			album = a
			return album, nil
		}
	}

	return album, errors.New("album not found")
}
