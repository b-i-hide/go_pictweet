package controllers

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

func UserValidator(name, email, password string, c *gin.Context) map[string]string {
	validate_messages := make(map[string]string)
	if len(name) == 0 {
		validate_messages["name"] = "Name can't be blank"
	}
	
	if len(email) == 0 {
		validate_messages["email_blank"] = "Email can't be blank"
	}

	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		validate_messages["email_format"] = "Email format is invalid"
	}

	if len(password) < 6 {
		validate_messages["short_password"] = "Password is too short.(more than 6 characters.)"
	} else if len(password) > 15 {
		validate_messages["long_password"] = "Password is too long.(less than 15 characters.)"
	}
	return validate_messages
}
