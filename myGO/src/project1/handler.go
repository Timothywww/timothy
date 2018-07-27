package handler

import (
	"github.com/globalsign/mgo"
)

type (
	Handler struct {
		DB *mgo.Session
	}
)
