package main

import (
	"bufio"
	"bytes"
	"fmt"
	"path"
	"strings"

	"html/template"
)

// User -- user
type User struct {
	FirstName string `json:"firstName"  bson:"FirstName"`
	LastName  string `json:"lastName"  bson:"LastName"`
	UserName  string `json:"userName"  bson:"UserName"`
	Email     string `json:"email"  bson:"Email"`
}

func main() {
	templatePath := "/home/chadit/Projects/src/github.com/chadit/QikTrackerApi/main/templates"
	u := User{FirstName: "John", LastName: "Smith", UserName: "", Email: "test@test.com"}
	if r, err := generateWelcomeEmail(u, templatePath, false); err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println(r)
		fmt.Println("")
	}
}

func generateWelcomeEmail(user User, templatePath string, useEmbeddedTemaplate bool) (string, error) {
	var b bytes.Buffer
	var err error
	buffer := bufio.NewWriter(&b)
	if useEmbeddedTemaplate {
		welcomeEmail := `"
		<!DOCTYPE html>
		<html>
		  <head>
		  </head>
		  <body>
				<h1>Welcome {{ .FirstName}} {{ .LastName}} to QikTracker</h1>
				</body>
			</html>
		"`

		tmpl := template.New("page")
		if tmpl, err = tmpl.Parse(welcomeEmail); err != nil {
			return "", err
		}
		tmpl.Execute(buffer, user)
		buffer.Flush()

		r := strings.NewReplacer("\n", "", "\t", "")
		return r.Replace(b.String()), nil
	}
	return generateWelcomeEmail2(user, templatePath)
}

func generateWelcomeEmail2(user User, templatePath string) (string, error) {
	var b bytes.Buffer
	buffer := bufio.NewWriter(&b)
	t := template.New("welcome")
	var err error
	if t, err = template.ParseFiles(path.Join(templatePath, "welcome.html")); err != nil {
		return "", err
	}
	t.Execute(buffer, user)
	buffer.Flush()

	return b.String(), nil
}
