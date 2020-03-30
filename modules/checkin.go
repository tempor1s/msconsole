package modules

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tempor1s/msconsole/creds"
	"github.com/tempor1s/msconsole/graphql"

	"github.com/imroc/req"

	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
)

// CheckinModule is the source code that allows a user to checkin to a class using a code.
func CheckinModule(cmdCtx *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter a a checkin code to this command. Example: `ms checkin dog`")
		return
	}

	// Create a new session
	session := req.New()

	// Log the user in and set cookies :)
	loginUser(session, false)

	// Get the logged in users name and email from Graph QL
	name, email := graphql.UserInfo(session)

	fmt.Printf("\nName: %s\nMS Email: %s\n\n", name, email)

	// Try to log the user in
	resp, err := session.Get(fmt.Sprintf("http://make.sc/attend/%s", args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Log usage of the command :)
	http.Post("https://msconsole.dev.benlafferty.me/log", "application/txt", nil)

	// Print the new banner message.
	bannerMessage := getBannerMessage(resp.String())
	fmt.Print(colorBannerMessage(bannerMessage))
}

func loginUser(session *req.Req, retry bool) {
	// loginURL is just the url we are getting/posting to to log the user in
	loginURL := "https://www.makeschool.com/login"

	// Get the login page and check for errors
	_, err := session.Get(loginURL)
	if err != nil {
		log.Fatal(err)
	}

	// If we are on a retry, we want to delete the old creds
	if retry {
		creds.DeleteCredentials()
	}

	// Get username and password
	email, password := creds.GetCredentials()

	if email == "" || password == "" {
		loginUser(session, true)
	}

	param := req.Param{
		"user[email]":    email,
		"user[password]": password,
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "MSConsole - https://github.comn/tempor1s/msconsole",
	}

	resp, err := session.Post(loginURL, param, header)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we get a 200 status code from our request
	if resp.Response().StatusCode != 200 {
		log.Fatalf("Got non-ok status code from response... Code: %d", resp.Response().StatusCode)
	}

	// Get the banner message from the response body
	bannerMessage := getBannerMessage(resp.String())

	// If banner message is not successful login, then we retry login
	if bannerMessage == "Invalid Email or password." {
		fmt.Println("\n\x1b[1;31mInvalid username or password. Please try again!\x1b[0m")
		loginUser(session, true)
	} else if bannerMessage != "Signed in successfully." {
		log.Fatal(bannerMessage)
	} else {
		fmt.Print(colorBannerMessage(bannerMessage))
	}
}

func getBannerMessage(page string) string {
	htmlData, err := htmlquery.Parse(strings.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}

	// nodes := htmlquery.Find(htmlData, "//*[@id='js-header']/div[3]/div/text()")
	// nodes := htmlquery.Find(htmlData, "/html/body/main/header/div[4]")
	nodes := htmlquery.Find(htmlData, "/html/body/main/header/div[4]/div/text()")

	return strings.TrimSpace(nodes[0].Data)
}

func colorBannerMessage(message string) string {
	switch message {
	// red
	case "You are not registered for this class.":
		fallthrough
	case "You need to be connected to Make School Wi-Fi to check-in.":
		return fmt.Sprintf("\x1b[1;31m%s\x1b[0m\n", message)
		// green
	case "You have already checked in as for this class.":
		fallthrough
	case "You have checked in as present for this class.":
		fallthrough
	case "Signed in successfully.":
		fallthrough
	case "You have checked in tardy for this class.":
		return fmt.Sprintf("\x1b[1;32m%s\x1b[0m\n", message)
		// yellow
	case "You code is not related to any class.":
		fallthrough
	case "You cannot check-in after a class is already over":
		fallthrough
	default:
		return fmt.Sprintf("\033[93m%s\x1b[0m\n", message)
	}
}
