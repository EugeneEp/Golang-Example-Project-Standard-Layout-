package user

import (
	"projectname/internal/project/domain/basic"
	"projectname/pkg/crypt"
	"strconv"
	"time"
)

const salt = `salt`

type (
	Entity struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		CreatedAt   int64  `json:"created_at"`
	}

	Get struct {
		basic.Request
		ID string `json:"id"`
	}

	GetResult struct {
		Entity
	}

	Create struct {
		ID          string
		DisplayName string `query:"display_name"`
		CreatedAt   int64  `json:"created_at"`
	}

	DeleteOverdue struct {
		TimeRange int64
	}

	CreateResult struct {
		Entity
	}
)

// GenerateID Метод для генерации поля ID структуры Create
func (c *Create) GenerateID() {
	v := strconv.Itoa(int(time.Now().Unix())) + salt
	c.ID = crypt.GetMD5Hash(v)
}

// SetCreatedAt Метод для того, чтобы установить дату создания пользователя
// В случае, если тип хранения данных - windows registry
func (c *Create) SetCreatedAt() {
	c.CreatedAt = time.Now().Unix()
}
