package modules

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
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
		"user[email]": "benlaugherty@gmail.com", // Make school dashboard login here
		"user[password]": "{RfV.4G8{Kv>tU82zYbR", // Make school password here
	}

	// Parse string into html node so we can parse
	htmlData, err := htmlquery.Parse(strings.NewReader(loginPage.String()))
	if err != nil {
		log.Fatal(err)
	}

	// Find all hidden inputs
	nodes := htmlquery.Find(htmlData, "//form//input[@type='hidden']")

	// Get authenticity token
	for _, node := range nodes {
		for i, item := range node.Attr {
			if item.Val == "authenticity_token" {
				formData[item.Val] = node.Attr[i + 1].Val
			}
		}
	}

	fmt.Println(formData)

	// Post that data to the login page, and get the response
	resp, err := s.Session.Post(loginURL, &grequests.RequestOptions{Data:formData, RedirectLimit:3})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	resp, err = s.Session.Get(fmt.Sprintf("http://make.sc/attend/%s", args[0]), &grequests.RequestOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getBannerMessage(resp.String()))
}

// Session represents stuff to do with making requests and keeping a session to do so
type Session struct {
	Session *grequests.Session
	DefaultRO *grequests.RequestOptions
}

func NewSession() *Session {
	ro := &grequests.RequestOptions{RedirectLimit:100}
	session := grequests.NewSession(ro)

	return &Session{Session:session, DefaultRO:ro}
}

func getBannerMessage(page string) string {
	htmlData, err := htmlquery.Parse(strings.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}

	nodes := htmlquery.Find(htmlData, "//*[@id='js-header']/div[3]/div/text()")
	// TODO: Decide if we wanna trim or not
	return strings.TrimSpace(nodes[0].Data)
}