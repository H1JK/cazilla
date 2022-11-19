package main

import (
	"flag"
	"log"
	"os"

	"github.com/h1jk/cazilla"
	"github.com/h1jk/cazilla/embed"
	"github.com/h1jk/cazilla/fetch"
)

var (
	out = flag.String("out", "mozilla_included.pem", "root CA list output")
)

func main() {
	flag.Parse()

	log.Println("㊙ Loading embed CA list")
	if !cazilla.CA.AppendCertsFromPEM(embed.MozillaIncludedCAPEM) {
		log.Println("☠ Failed to load embed CA list")
	}
	cazilla.ConfigureDefault()

	log.Println("✉ Downloading to", *out)

	pem, err := fetch.DownloadPEM(nil)
	if err != nil {
		log.Println("☠ Error when downloading:", err)
		return
	}

	err = os.WriteFile(*out, pem, 0o666)
	if err != nil {
		log.Println("☠ Error when writing output:", err)
		return
	}

	log.Println("🎉 Download succeed!", len(pem), "bytes have been written to", *out)
}
