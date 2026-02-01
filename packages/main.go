package main

import (
	"fmt"
    "github.com/fatih/color"
"github.com/annuvrat/GO_bs/auth"
	"github.com/annuvrat/GO_bs/user"
)




func main(){

auth.LoginWithCreds("annu","mypassword")



    session := auth.GetSession()

	fmt.Println("session", session)


	user := user.User{
		Id:    1,
		Name:  "Annu",
		Email: "annu@example.com",
	}

	fmt.Println("user", user,user.Name)

    color.Cyan(user.Email)


}