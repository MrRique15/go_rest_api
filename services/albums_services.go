package Services

import (
	"errors"
	Repositores "go_rest_api/repositories"
	Structs "go_rest_api/structs"
)

func ListAlbums() []Structs.Album {
	albumsList := Repositores.ListAlbums()
	return albumsList
}

func AddAlbum(album Structs.Album) (Structs.Album, error) {
	existingAlbum , err := Repositores.FindAlbumById(album.ID)

	if err == nil {
		return existingAlbum, errors.New("album already registered with inserted id")
	}

	newAlbum, err := Repositores.AddAlbum(album)

	return newAlbum, err
}

func GetAlbumById(id string) (Structs.Album, error) {
	foundAlbum, err := Repositores.FindAlbumById(id)
	return foundAlbum, err
}
