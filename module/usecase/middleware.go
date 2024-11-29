package usecase


import "github.com/labstack/echo/v4"

func (m *Methods) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		
		return next(c)
	}
}


