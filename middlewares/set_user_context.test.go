package middlewares

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/necmettindev/currency-conversion/models/user"
	"github.com/necmettindev/currency-conversion/services/authservice"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetUserContextMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	alice := &user.User{
		Username:  "necmettin",
		FirstName: "",
		LastName:  "",
	}

	svc := authservice.NewAuthService("secret")

	t.Run("Set context values with valid auth", func(t *testing.T) {
		token, _ := svc.IssueToken(*alice)
		bearerToken := "Bearer " + token

		router.GET("/test1", SetUserContext("secret"), func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			assert.EqualValues(t, 0, userID)
			assert.EqualValues(t, 0, c.Request.Context().Value("user_id"))

			username, _ := c.Get("user_username")
			assert.EqualValues(t, "necmettin", username)
			assert.EqualValues(t, "necmettin", c.Request.Context().Value("user_username"))

			c.Status(http.StatusOK)
		})

		request, _ := http.NewRequest("GET", "/test1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		request.Header.Add("Authorization", bearerToken)
		c.Request = request
		router.ServeHTTP(w, request)

		assert.EqualValues(t, http.StatusOK, w.Code)
	})

	t.Run("Set no context value without valid auth", func(t *testing.T) {
		router.GET("/test2", SetUserContext("secret"), func(c *gin.Context) {
			username, _ := c.Get("user_username")
			assert.Nil(t, username)
			assert.Nil(t, c.Request.Context().Value("user_username"))

			userID, _ := c.Get("user_id")
			assert.Nil(t, userID)
			assert.Nil(t, c.Request.Context().Value("user_id"))

			c.Status(http.StatusOK)
		})

		request, _ := http.NewRequest("GET", "/test2", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = request
		router.ServeHTTP(w, request)

		assert.EqualValues(t, http.StatusOK, w.Code)
	})
}
