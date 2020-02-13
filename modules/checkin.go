package modules

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
	"github.com/spf13/cobra"
)

// CheckinModule is the source code that allows a user to checkin to a class using a code.
func CheckinModule(cmdCtx *cobra.Command, args[]string) {
	if len(args) == 0 {
		fmt.Println("Please enter a a checkin code to this command. Example: `ms checkin dog`")
		return
	}

	// s represents our simple session wrapper that sets up request options and such
	s := NewSession()
	// loginURL is just the url we are getting/posting to to log the user in
	loginURL := "https://www.makeschool.com/login"

	// Get the login page and check for errors
	loginPage, err := s.Session.Get(loginURL, s.DefaultRO)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we get a 200 status code from our request
	if loginPage.Ok == false {
		log.Fatalf("Got non-ok status code from response... Code: %d", loginPage.StatusCode)
	}

	// Create a map with the info we want to login with
	formData := map[string]string{
		"user[email]": "", // Make school dashboard login here
		"user[password]": "", // Make school password here
	}

	// Create new request options with that data
	ro := &grequests.RequestOptions{Data: formData}

	// Post that data to the login page, and get the response
	resp, err := s.Session.Post(loginURL, ro)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())
	fmt.Println(resp.StatusCode)

	fmt.Printf("Your checkin code is %s\n", args[0])
}

// Session represents stuff to do with making requests and keeping a session to do so
type Session struct {
	Session *grequests.Session
	DefaultRO *grequests.RequestOptions
}

func NewSession() *Session {
	ro := &grequests.RequestOptions{}
	session := grequests.NewSession(ro)

	return &Session{Session:session, DefaultRO:ro}
}

func login() {}

func checkin() {}
