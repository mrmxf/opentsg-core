//Package auth searches for the required tokens to access gitlab, github and aws in enviroment variables and user input.
package auth

import (
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var auths map[string]Auth

//AuthInit generates a map of configuration objects for each available connection
//based off the available user enviroments and user input.
//This should be called before any other functions in mmReader
func AuthInit(profile string, keyparams []string) {

	auths = make(map[string]Auth)
	auths["s3_profile"] = s3Auth(profile, keyparams)
	auths["git_auth"] = gitlabAuth(keyparams)
	auths["github_auth"] = gitHubAuth(keyparams)

}

//AuthGet returns the map of authorisation types for a system
func AuthGet() map[string]Auth {
	//prevent nil errors down the line
	if auths == nil {
		AuthInit("", nil) //make the map with no keys if not previously declared
	}
	return auths
}

//Auth contains the tokens for any type of sytem
type Auth struct {
	Token any
}

//S3AuthDetail contains the required information to access S3
type S3AuthDetail struct {
	Creds  *credentials.Credentials
	Region string
}

//make a struct of credentials and region
//make variac based on strings of keys
func s3Auth(profile string, keyParams []string) Auth {
	var a Auth
	var s3_contents S3AuthDetail
	a.Token = &s3_contents
	fmt.Println("searching for aws default profile")
	if profile == "" {
		profile = "default"
	}
	creds := credentials.NewSharedCredentials(config.DefaultSharedCredentialsFilename(), profile)
	if _, err := creds.Get(); err == nil {
		s3_contents.Creds = creds
		return a
	} else {
		fmt.Println(err)
	}

	creds = credentials.NewEnvCredentials()
	fmt.Println("searching for aws enviroment variables")
	if _, err := creds.Get(); err == nil {
		s3_contents.Creds = creds
		return a
	}
	if len(keyParams) < 2 {
		fmt.Println("not enough AWS paramters passed")
		return a
	}

	fmt.Println("Extracting aws values from user input")
	//use these regexp to match the string
	reg := regexp.MustCompile("(us|ap|ca|cn|eu|sa)-(central|(north|south)?(east|west)?)-\\d$")
	sk := regexp.MustCompile("^[A-Za-z0-9/+=]{40}$") //these are taken from https://aws.amazon.com/blogs/security/a-safer-way-to-distribute-aws-credentials-to-ec2/
	akid := regexp.MustCompile("^[A-Z0-9]{20}$")     //without the lookarounds as it is not go compatiable

	var secret string
	var id string
	for _, text := range keyParams {
		switch {
		case reg.MatchString(text):
			s3_contents.Region = text
		case sk.MatchString(text):
			secret = text
		case akid.MatchString(text):
			id = text
		}
	}
	s3_contents.Creds = credentials.NewStaticCredentials(id, secret, "")
	//manually apply items
	return a
}

//make a struct of credentials and region
//make variac based on strings of keys
func gitlabAuth(keyParams []string) Auth {
	var a Auth
	//look for more names/ask bruce
	tokenNames := []string{"GITLAB_PAT"}

	for _, env := range tokenNames {
		if val := os.Getenv(env); val != "" {
			a.Token = val
			return a
		}
	}

	if len(keyParams) > 0 {
		gitlabT := regexp.MustCompile("^glpat-[0-9a-zA-Z\\-\\_]{20}$")
		for _, v := range keyParams {
			if gitlabT.MatchString(v) {
				a.Token = v
				return a
			}
		}
	}
	//else assign an empty token to preven errors
	a.Token = ""
	return a
}

//make a struct of credentials and region
//make variac based on strings of keys
func gitHubAuth(keyParams []string) Auth {
	var a Auth
	//look for more names/ask bruce
	tokenNames := []string{"GITHUB_PAT"}

	for _, env := range tokenNames {
		if val := os.Getenv(env); val != "" {
			a.Token = val
			return a
		}
	}

	if len(keyParams) > 0 {
		githubT := regexp.MustCompile("^ghp_[0-9a-zA-Z]{36}$")
		for _, v := range keyParams {
			if githubT.MatchString(v) {
				a.Token = v
				return a
			}
		}
	}
	//else assign an empty token to preven errors
	a.Token = ""
	return a
}
