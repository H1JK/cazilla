package main

import (
	"log"
	"net/http"

	"github.com/h1jk/cazilla"
	"github.com/h1jk/cazilla/embed"
)

func main() {
	// this loads pre-downloaded CA list from cazilla.
	// note that the CA list may change after a while,
	// so keep a frequent update if you are using this.
	cazilla.CA.AppendCertsFromPEM(embed.MozillaIncludedCAPEM)

	// apply cazilla shared CA pool to http.DefaultTransport,
	// which is used by http.DefaultClient.
	// if you are using custom http client, use cazilla.ConfigureHTTPTransport or configure it by yourself.
	cazilla.ConfigureDefault()

	req, _ := http.NewRequest(http.MethodGet, "https://gstatic.com/generate_204", nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("error when requesting", err)
	}
	response.Body.Close()

	if response.StatusCode == 204 {
		log.Println("successfully requested google! response code", response.StatusCode)
	} else {
		log.Panic("request failed! response code", response.StatusCode)
	}
}
