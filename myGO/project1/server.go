package main

import (
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
  "github.com/labstack/gommon/log"
  
  "project1"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${method}: -${uri}\n",
	  }))

	//connect to database
	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler.Handler{DB: db}

	//Routes
	e.POST("/signup",h.Signup)
	e.POST("/login", h.Login)

	e.Logger.Fatal(e.Start(":1323"))

}
