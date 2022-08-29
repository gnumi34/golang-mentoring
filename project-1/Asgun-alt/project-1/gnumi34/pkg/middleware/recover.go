package middleware

import (
	"fmt"
	"log"
	"runtime"

	"github.com/labstack/echo/v4"
)

func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					stack := make([]byte, 32<<10)
					length := runtime.Stack(stack, true)
					stack = stack[:length]
					msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					log.Printf(msg)
					ctx.Error(err)
				}
			}()
			return next(ctx)
		}
	}
}
