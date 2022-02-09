package delivery

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"time"
)

type PostHandlers struct {
}

type pagination struct {
	LastCreatedAt string
	LastID        string
}

func parseQuery(c echo.Context) error {
	// проверка, что об поля переданы, либо не переданы
	lastId := c.QueryParam("last_id")
	lastCreatedAt := c.QueryParam("last_created_at")
	// валидация lastId
	if err := validation.Validate(lastId, validation.Required, is.UUIDv4); err != nil {
		return err
	}
	// валидация lastCreatedAt
	if err := validation.Validate(lastCreatedAt, validation.Required); err != nil {
		return err
	}
	// Парсинг даты
	timePattern := "2006-01-02T15:04:05MST"
	if _, err := time.Parse(timePattern, lastCreatedAt); err != nil {
		return fmt.Errorf("date validation error (%s). please, provide correct format: %s", err.Error(), timePattern)
	}
	return nil
}

type errorResp struct {
	Message string `json:"message"`
}

func (e errorResp) str(msg string) string {

}

func AllPosts(c echo.Context) error {
	// Достать пагинацию (дата создания и ID) ИЛИ по дефолту лимит = 10
	//lastId, lastCreatedAt := "", ""
	err := parseQuery(c)
	if err != nil {
		return c.JSON(400, errorResp{Message: err.Error()})
	}
	return c.JSON(200, struct{}{})
}
