package container

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var C *Container

type Container struct {
	DB  *DB
	Api *Api
}

type DB struct {
	*gorm.DB
}

type Api struct {
	*http.Client
}

func InitContainer() {
	C = NewDefaultContainer()
}

func NewDefaultContainer() *Container {
	db, err := gorm.Open(sqlite.Open("btc.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return &Container{
		DB: &DB{
			db,
		},
		Api: &Api{
			&http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
				},
			},
		},
	}
}
