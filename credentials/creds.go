package credentials

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	dcred "github.com/docker/docker-credential-helpers/credentials"
)

// SetCredentials will set and return given credentials
func SetCredentials() (string, string) {
	// TODO: Update this with correct URL when we swap repo
	email, password := getCreds()

	url := "https://github.com/BenAndGarys/msconsole-go"
	set("msconsole", url, email, password)

	return email, password
}

// GetCredentials will get creds from keychain
func GetCredentials() (string, string) {
	url := "https://github.com/BenAndGarys/msconsole-go"
	email, password, err := get("msconsole", url)

	if err != nil {
		SetCredentials()
		email, password = GetCredentials()
	}

	return email, password
}

func DeleteCredentials() {
	url := "https://github.com/BenAndGarys/msconsole-go"
	err := del("msconsole", url)

	if err != nil {
		log.Fatal(err)
	}
}

// set adds an item to the keychain
func set(lbl, url, email, secret string) error {
	cr := &dcred.Credentials{
		ServerURL: url,
		Username:  email,
		Secret:    secret,
	}

	dcred.SetCredsLabel(lbl)
	return ns.Add(cr)
}

// get an entry in the keychain
func get(lbl, url string) (string, string, error) {
	dcred.SetCredsLabel(lbl)
	return ns.Get(url)
}

func del(lbl, url string) error {
	dcred.SetCredsLabel(lbl)
	return ns.Delete(url)
}

// getCreds will prompt the user for credentials, in a pretty way!
func getCreds() (string, string) {
	qs := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "What is your email?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Please type enter your password."},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Email    string
		Password string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatal(err)
	}

	return answers.Email, answers.Password
}
