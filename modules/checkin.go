package modules

import (
	"fmt"
	"github.com/imroc/req"
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

	// Create a new session
	session := req.New()

	// loginURL is just the url we are getting/posting to to log the user in
	loginURL := "https://www.makeschool.com/login"

	// Get the login page and check for errors
	_, err := session.Get(loginURL)
	if err != nil {
		log.Fatal(err)
	}

	param := req.Param{
		"user[email]": "email",
		"user[password]": "password",
	}

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent": "MSConsole - https://github.comn/BenAndGarys/msconsole-go",
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

	if bannerMessage != "Signed in successfully." {
		log.Fatal(bannerMessage)
	}
	fmt.Println(bannerMessage)

	// Try to log the user in
	resp, err = session.Get(fmt.Sprintf("http://make.sc/attend/%s", args[0]), &grequests.RequestOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the new banner message.
	fmt.Println(getBannerMessage(resp.String()))
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