package graphql

import (
	"log"

	"github.com/imroc/req"
)

// Info represents the json struct that we get back from Make School.
type Info struct {
	Data struct {
		CurrentUser struct {
			Name         string `json:"name"`
			StudentEmail string `json:"studentEmail"`
		} `json:"currentUser"`
	} `json:"data"`
}

// UserInfo will grab the users name and email from the makeschool GraphQL API
func UserInfo(session *req.Req) (string, string) {

	url := "https://www.makeschool.com/graphql"
	query := "{ currentUser {name studentEmail} }"
	request := map[string]string{"query": query}

	resp, err := session.Post(url, req.BodyJSON(request))
	if err != nil {
		log.Fatal(err)
	}

	var data Info
	resp.ToJSON(&data)

	return data.Data.CurrentUser.Name, data.Data.CurrentUser.StudentEmail
}
