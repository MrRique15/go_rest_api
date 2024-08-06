package Services

import (
	Repositores "go_rest_api/repositories"
	Structs "go_rest_api/structs"
)

func ListAlbums() (albumsList []Structs.Album) {
	albumsList = Repositores.ListAlbums()
	return
}

func AddAlbum(album Structs.Album) (newAlbum Structs.Album) {
	var sameAlbum = Repositores.FindAlbumById(album.ID)

	if sameAlbum.ID != "" {
		return
	}

	newAlbum = Repositores.AddAlbum(album)
	return
}

func GetAlbumById(id string) (foundAlbum Structs.Album) {
	foundAlbum = Repositores.FindAlbumById(id)
	return
}
