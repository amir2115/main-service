package middleware

import (
	"MainService/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthenticationResponse struct {
	MessageCode int    `json:"message_code"`
	Message     string `json:"message"`
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			c.Abort()
		}
		auth := getAuthenticationResult(token)
		if auth.MessageCode != 404 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			c.Abort()
		}
		c.Next()
	}
}

func getAuthenticationResult(token string) AuthenticationResponse {
	requestBody, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		log.Fatalln(err)
	}
	response, err := http.Post("", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data AuthenticationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}
