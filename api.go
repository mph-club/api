package main

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/server"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	database.CreateSchema()

	server.CreateAndListen()
}

func initiateAuthorization() *string {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc := cognitoidentityprovider.New(sess)

	initAuth, err := svc.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String("2895spmk9o6vuqctbedc0v9f38"),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String("oscar@mphclub.com"),
			"PASSWORD": aws.String("Hunter2!!"),
		},
	})

	if err != nil {
		log.Println("fail!")
		log.Println(err)
	} else {
		log.Println("success!")
		log.Println(initAuth)

		token := initAuth.AuthenticationResult.AccessToken

		return token
	}

	return aws.String("attempt to auth failed")
}
