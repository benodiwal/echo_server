package main

import (
	"os"
	"net/http"
	"github.com/labstack/echo/v4"
	"io"
)

func main() {

	// Initislzie echo
	e := echo.New()

	// Get route
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.POST("/save", save)
	e.POST("/save2", save2)

	e.Logger.Fatal(e.Start(":1234"))

}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello there!!\n")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: " + team + " " + "member: " + member + "\n")
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name: " + name + "email: " + email + "\n")
}

func save2(c echo.Context) error {

	// Get name
	name := c.FormValue("name")

	// Get avate
	avatar, err := c.FormFile("avatar")
	if (err != nil) {
		return err
	}

	// Source
	src, err := avatar.Open()
	if (err != nil) {
		return err
	}
	defer src.Close()


	// Destination
	dst, err := os.Create(avatar.Filename)
	if (err != nil) {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank You " + name + "</b>")

}
