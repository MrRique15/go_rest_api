package Services

import (
	Repositores "example/web-service-gin/repositories"
	Structs "example/web-service-gin/structs"
)

func ListAlbums() (albumsList []Structs.Album){
	albumsList = Repositores.ListAlbums()
	return
}

func AddAlbum(album Structs.Album) (newAlbum Structs.Album){
	var sameAlbum = Repositores.FindAlbumById(album.ID)

	if(sameAlbum.ID != ""){
		return
	}

	newAlbum = Repositores.AddAlbum(album)
    return
}

func GetAlbumById(id string) (foundAlbum Structs.Album){
	foundAlbum = Repositores.FindAlbumById(id)
	return
}