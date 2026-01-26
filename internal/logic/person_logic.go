package logic

import (
	"crud/domain"
	"log"

	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PersonRepositoryInterface interface {
	PostPersonQuery(p domain.Person) error
	GetPersonQuery() ([]domain.Person, error)
	GetPersonIdQuery(id int) (domain.Person, error)
	PutPersonIdQuery(id int, p domain.Person) error
	DelPersonIdQuery(id int) error
}

type PersonHandler struct {
	Repo PersonRepositoryInterface
}

func (ph *PersonHandler) PostPersonHandler(c echo.Context) error {

	var p domain.Person

	err := c.Bind(&p)
	if err != nil {
		log.Println("PostpersonHandler error:", err)
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	err = ph.Repo.PostPersonQuery(p)
	if err != nil {
		log.Println("PostpersonHandler error:", err)
		return c.String(http.StatusInternalServerError, "DB error")
	}

	return c.String(http.StatusOK, "Done!")
}

func (ph *PersonHandler) GetPersonHandler(c echo.Context) error {

	personlist, err := ph.Repo.GetPersonQuery()
	if err != nil {
		log.Println("GetpersonHandler error:", err)
		return c.String(http.StatusInternalServerError, "DB error")
	}

	return c.JSON(http.StatusOK, personlist)
}

func (ph *PersonHandler) GetPersonIdHandler(c echo.Context) error {

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println("GetpersonidHandler error:", err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	p, err := ph.Repo.GetPersonIdQuery(id)
	if err != nil {
		log.Println("GetpersonidHandler error:", err)
		return c.String(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, p)
}

func (ph *PersonHandler) PutPersonIdHandler(c echo.Context) error {

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println("PutpersonidHandler error:", err)
		return c.String(http.StatusNotFound, "id not found")
	}

	var p domain.Person

	err = c.Bind(&p)
	if err != nil {
		fmt.Println("error:", err)
		log.Println("PutpersonidHandler error:", err)
		return c.String(http.StatusBadRequest, "Invalid JSON")
	}

	err = ph.Repo.PutPersonIdQuery(id, p)
	if err != nil {
		log.Println("PutpersonidHandler error:", err)
		return c.String(http.StatusInternalServerError, "DB error")
	}

	return c.String(http.StatusOK, "Done!")
}

func (ph *PersonHandler) DelPersonIdHandler(c echo.Context) error {

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Println("DelpersonidHandler error:", err)
		return c.String(http.StatusNotFound, "id not found")
	}

	err = ph.Repo.DelPersonIdQuery(id)
	if err != nil {
		log.Println("DelpersonidHandler error:", err)
		return c.String(http.StatusInternalServerError, "DB error")
	}

	return c.String(http.StatusOK, "Done!")
}
