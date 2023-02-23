# Golang based API Authenticator 

## Introduction

Available methods:
- JWT token based
- Cookie based

## Usage
```golang
package main

import (
	"log"
	"os"

	"github.com/gauravsarma1992/goapiauth"
)

func main() {
	var (
		auth        *goapiauth.Authenticator
		tokenString string
		parsedId    string
		userId      string
		err         error
	)
	userId = "Stephen Hawking"
	if auth, err = goapiauth.New(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	if tokenString, err = auth.GenerateToken(userId); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	if parsedId, err = auth.GetUserFromToken(tokenString); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(tokenString, parsedId)
}
```
