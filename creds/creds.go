package creds

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	keychain "github.com/keybase/go-keychain"
)

// SetCredentials will add the credentials to the keychain
func SetCredentials() (string, string) {
	email, password := getCreds()

	item := keychain.NewGenericPassword("MSConsole", email, "MSConsole Credentials", []byte(password), "msconsole")
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		log.Fatal("duplicate item in keychain")
	}

	return email, password
}

// GetCredentials will get creds from keychain
func GetCredentials() (string, string) {
	accounts, err := keychain.GetAccountsForService("MSConsole")
	if err != nil {
		log.Fatal(err)
	}

	if len(accounts) < 1 {
		email, password := SetCredentials()
		return email, password
	}

	email := accounts[0]

	storedAccount, err := keychain.GetGenericPassword("MSConsole", email, "MSConsole Credentials", "msconsole")

	if err != nil {
		log.Fatal(err)
	}

	return email, string(storedAccount)
}

// DeleteCredentials will delete bad username and password from go-keychain
func DeleteCredentials() {
	accounts, err := keychain.GetAccountsForService("MSConsole")
	if err != nil {
		log.Fatal(err)
	}

	if len(accounts) < 1 {
		SetCredentials()
	}

	email := accounts[0]

	err = keychain.DeleteGenericPasswordItem("MSConsole", email)
	if err != nil {
		log.Fatal(err)
	}
}

// getCreds will prompt the user for creds, in a pretty way!
func getCreds() (string, string) {
	qs := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "What is your email?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Please enter your password."},
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
