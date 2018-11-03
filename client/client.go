package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	initiateAuthorization()
}

func initiateAuthorization() *string {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc := cognitoidentityprovider.New(sess)

	initAuth, err := svc.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String("61lmh73b4k5fdogsbhebmmld33"),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String("oscar@mphclub.com"),
			"PASSWORD": aws.String("@Test123"),
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
