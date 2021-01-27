package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ramailh/auth-server/http/controller/oauth"
	"github.com/ramailh/auth-server/model"
	"github.com/ramailh/auth-server/service/user"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"test": "echo"})
}

func LoginPost(c *gin.Context) {
	var param model.User

	if err := c.Bind(&param); err != nil {
		log.Println(err)
	}

	result, err := user.Login(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to login, password and username combination not matched!", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success, welcome " + param.Username, "status": true, "data": result})
}

func Google(c *gin.Context) {
	url := oauth.UrlOauth()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := oauth.Conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	client := oauth.Conf.Client(context.Background(), token)

	responseInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer responseInfo.Body.Close()

	jsonInfo, _ := ioutil.ReadAll(responseInfo.Body)

	var googleParam model.Google
	json.Unmarshal(jsonInfo, &googleParam)

	result, err := user.LoginGoogle(googleParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to login, email not found!", "status": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success, welcome " + googleParam.Name, "status": true, "data": result})
}
