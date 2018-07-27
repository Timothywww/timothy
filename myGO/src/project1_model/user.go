package model

import (
	"github.com/globalsign/mgo/bson"
)

type (
	User struct {
		ID        bson.ObjectId `json:"id"  bson:"_id,omitempty"`
		Account   string        `json:"account" bson:"account"`
		Password  string        `json:"password,omitempty"  bson:"password"`
		Token     string        `json:"token,omitempty" bson:"-"`
		Group	  int        	`json:"group,omitempty" bson:"group,omitempty"`
	}
)
