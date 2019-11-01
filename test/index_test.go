package test

import (
	"GinHello/initRouter"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}


func TestInitGetRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInitPostRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserSave(t *testing.T)  {
	userName := "tyf"
	router := initRouter.SetupRouter();
	w := httptest.NewRecorder();
	req, _ := http.NewRequest(http.MethodGet, "/user/"+ userName, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户" + userName + "已保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	userName := "tyf"
	userAge := "28"
	router := initRouter.SetupRouter();
	w := httptest.NewRecorder();
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+ userName + "&age=" + userAge, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:" + userName +",年龄:" + userAge + "已保存", w.Body.String())
}

func TestUserPostFormEmailErrorAndPasswordError(t *testing.T) {
	router := initRouter.SetupRouter();
	values := url.Values{}
	values.Add("email","qwefs")
	values.Add("password","123")
	values.Add("password-again","234")
	w := httptest.NewRecorder();
	req, _ := http.NewRequest(http.MethodPost,"/user/register", bytes.NewBufferString(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}