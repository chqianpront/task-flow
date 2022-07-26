package api

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func Login(c *gin.Context) {
	account := c.Param("account")
	passwd := c.Param("passwd")
	log.Printf("login account: %s\n", account)
	pmd5 := md5.Sum([]byte(passwd))
	hexp := fmt.Sprintf("%x", pmd5)
	log.Printf("show hex pass word: %s\n", hexp)
}
