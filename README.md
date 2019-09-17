Golang OAuth@Mail.ru (It's very simple implementation)
--------------------
Setup Application https://o2.mail.ru/app/

```sh
go get github.com/rubaxa/oauth-mailru.go
```

### Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/rubaxa/oauth-mailru.go"
	"github.com/rubaxa/oauth-mailru.go/button"
)

func main() {
	mailru := oauth.Client{
		ClientID:     oauth.GetEnvVar("O2_CLIENT_ID"),
		ClientSecret: oauth.GetEnvVar("O2_CLIENT_SECRET"),
		RedirectURI:  oauth.GetEnvVar("O2_REDIRECT_URI", "http://localhost:7000/mailru-oauth"),
	}

	mailruButtonHtml := button.Render(mailru, button.Options{
		Text: "Continue with Mail.ru",
		Size: "50px",
	})

	// Render button
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; encoding=utf-8;")
		w.Write([]byte(mailruButtonHtml))
	})

	// Exchange `code` to `access_token`
	http.HandleFunc("/mailru-oauth", func(w http.ResponseWriter, r *http.Request) {
		auth := mailru.ParseResponse(r)
		token, err := mailru.Exchange(auth)

		w.Header().Add("Content-Type", "text/plain; encoding=utf-8;")
		if err != nil {
			w.Write([]byte("Error: " + err.String()))
			return;
		}

		w.Write([]byte("AccessToken: " + token.AccessToken))
	})

	// Start http server
	fmt.Println("OAuth@Mail.ru example running at http://localhost:7000/")
	http.ListenAndServe(":7000", nil)
}
```
