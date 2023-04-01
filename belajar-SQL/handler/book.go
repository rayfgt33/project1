package handler

import (
	"belajar-SQL/model"
	"belajar-SQL/respond"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		respond.InternalServerError(c, err.Error())
		return
	}

	respond.Ok(c, res)
}

func (h HttpServer) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.GetBookByID(int64(idInt))
	if err != nil {
		if err.Error() == respond.ErrNotFound {
			respond.NotFound(c, err.Error())
			return
		}
		respond.InternalServerError(c, err.Error())
		return
	}

	respond.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	in.ID = idInt
	// call service
	res, err := h.app.UpdateBook(in)
	if err != nil {
		respond.InternalServerError(c, err.Error())
		return
	}

	respond.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		respond.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		respond.InternalServerError(c, err.Error())
		return
	}

	respond.OkWithMessage(c, "Book deleted Successfully")
}
