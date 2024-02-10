package utils

import (
	"api-exemple/app/data/response"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SendErroBindJson(c *gin.Context, arg interface{}) error {
	if err := c.ShouldBindJSON(&arg); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, response.Msg{Message: "Invalid json"})
		} else {
			c.JSON(http.StatusBadRequest, response.Msg{Message: err.Error()})
		}
		return err
	}

	return nil
}

func SendNotFoundOrError(ctx *gin.Context, err error) error {
	if err != nil {
		if err != pgx.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, response.Msg{Message: err.Error()})
			return err
		}
		ctx.JSON(http.StatusNotFound, response.Msg{Message: "Register not found"})
		return err
	}

	return nil
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
