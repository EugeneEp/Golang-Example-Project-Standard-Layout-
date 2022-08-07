package access

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"projectname/internal/project/infrastructure/logger"
	"projectname/internal/project/interfaces/web/echo/middleware"
	"strings"
)

const (
	accessToken = `access_token`
	salt        = `salt`
)

// SetInvoker Функция устанавливает значение CurrentUserID
func SetInvoker() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := context.(*middleware.Context)

			var log *zap.Logger

			if err := ctx.Container.Fill(logger.BaseServiceName, &log); err != nil {
				return err
			}

			token := getToken(ctx)
			CurrentUserID := getCurrentUserIDFromToken(token)

			ctx.Set(middleware.KeyUserID, CurrentUserID)
			return handlerFunc(context)
		}
	}
}

// getToken Функция ищет передан ли accessToken и возвращает его, в случае наличия токена
func getToken(ctx *middleware.Context) string {
	header := ctx.Request().Header.Get(echo.HeaderAuthorization)

	if header == `` {
		return ctx.QueryParam(accessToken)
	}

	if strings.Count(header, ` `) != 1 {
		return ctx.QueryParam(accessToken)
	}

	token := header[strings.IndexByte(header, ' ')+1:]

	if token == `` {
		return ctx.QueryParam(accessToken)
	}

	return token
}

// getCurrentUserIDFromToken Функция парсит переданный токен и формирует из него CurrentUserID
func getCurrentUserIDFromToken(token string) string {
	h := md5.New()
	h.Write([]byte(token + salt))
	return hex.EncodeToString(h.Sum(nil))
}
