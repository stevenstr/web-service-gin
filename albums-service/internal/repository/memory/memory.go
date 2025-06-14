package memory

import (
	model "github.com/stevenstr/web-service-gin/albums-service/pkg"
)

func New() []*model.Album {
	// albums slice to seed record album data.
	return []*model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

}
