package auth

import (
	"fmt"
	"time"
	"net/http"

	"github.com/labstack/echo/v4"

	// "github.com/msalbrain/authSphere/internals/database"
	service "github.com/msalbrain/authSphere/internals/service"
)

func AuthRoute(e *echo.Echo, u service.UserService, mail service.MailService, jwtoption service.JwtOptionService) {

	g := e.Group("/auth")

	// TODO
	g.POST("/simple-register", simpleRegister(u))
	g.POST("/login", login(u))
	g.POST("/logout", login(u)) // delete refresh token
	g.POST("/refresh-token", login(u))
	g.POST("/forgot-password", login(u))
	g.POST("/reset-password", login(u))
	g.POST("/change-pasword", login(u))
	g.POST("/verify-email", login(u))
	g.POST("/resend-verification", login(u))
	g.POST("/social-login", login(u))
	g.POST("/validate-token", login(u))
	g.POST("/token-expiration", login(u))
	g.POST("/revoke-token", login(u))
	g.POST("/verify-email", login(u))
	g.POST("/deactivate", login(u))
}

type HttpError struct {
	Status        int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

type SimpleSignupUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Bio      string `json:"bio" `
}

type UserReturn struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
}

func simpleRegister(userService service.UserService) func(echo.Context) error {
	return func(c echo.Context) error {
		u := new(SimpleSignupUser)
		err := c.Bind(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, HttpError{
				Status:        http.StatusBadRequest,
				Message:     "Wrong input",
				Description: "",
			})
		}
		if err = c.Validate(u); err != nil {
			return c.JSON(http.StatusConflict, HttpError{
				Status:        http.StatusConflict,
				Message:     "Wrong input",
				Description: "",
			})
		}

		user, _ := userService.FindUserByEmail(u.Email)

		if (user.ID != 0) {
			return c.JSON(http.StatusConflict, HttpError{
				Status:        http.StatusConflict,
				Message:     "User already exists",
				Description: "The provided username or email is already registered. Please log in or use a different username/email.",
			})
		}

		credUser, err := userService.CreateNewEmailUser(u.Name, u.Email, u.Bio, u.Password)
		if err != nil {
			// fmt.Print("This is where the error is comming from")
			fmt.Print(err)
		}



		return c.JSON(http.StatusOK,
			map[string]interface{}{
				"code":    http.StatusOK,
				"message": "New user created",
				"user": UserReturn{
					ID:    credUser.ID,
					Name:  credUser.Name,
					Email: credUser.Email.String,
					Bio:   credUser.Bio.String,
				},
			},
		)
	}
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func login(userService service.UserService) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := new(LoginUser)
		err := c.Bind(u)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, HttpError{
				Status:        http.StatusConflict,
				Message:     "Wrong input",
				Description: "",
			})
		}
		if err = c.Validate(u); err != nil {
			return c.JSON(http.StatusConflict, HttpError{
				Status:        http.StatusConflict,
				Message:     "Wrong input",
				Description: "",
			})
		}

		user, err := userService.FindUserByEmail(u.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized,
				map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "user doesn't exist",
				},
			)
		}

		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized,
				map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "user doesn't exist",
				},
			)
		}

		verifypass := userService.ValidatePassword(user.HashedPassword.String, u.Password)

		if !verifypass {
			return c.JSON(http.StatusUnauthorized,
				map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "invalid email or password",
				},
			)
		}

		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.Value = user.AuthToken
		cookie.Expires = time.Now().Add(24 * time.Hour)

		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "New user created",
			"authtoken": user.AuthToken,
			"user": UserReturn{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email.String,
				Bio:   user.Bio.String},
			})
	}
}
