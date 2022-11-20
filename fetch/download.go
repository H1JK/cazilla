package fetch

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"strings"

	"github.com/h1jk/cazilla"
)

// DownloadCSV downloads Mozilla provided Included CA Certificates (CSV with PEM of raw certificate data) and return its original CSV data.
func DownloadCSV(c *http.Client) ([]byte, error) {
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Get(cazilla.SourceCAListURL)
	if err != nil {
		return nil, err
	}
	respBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

// DownloadPEM downloads Mozilla provided Included CA Certificates (CSV with PEM of raw certificate data) and parses it, return parsed, ready-to-import PEM data.
func DownloadPEM(c *http.Client) ([]byte, error) {
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Get(cazilla.SourceCAListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	source := csv.NewReader(resp.Body)
	builder := bytes.Buffer{}
	builder.Grow(256 * 1024)
	defer builder.Reset()
	headers, err := source.Read()
	if err != nil {
		return nil, err
	}
	index := make(map[string]int, 32)
	for i, header := range headers {
		index[header] = i
	}
	for {
		record, err := source.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if country := record[index["Geographic Focus"]]; strings.Contains(country, "China") {
			continue // because they are crazy
		}
		if b := record[index["Trust Bits"]]; !strings.Contains(b, "Websites") {
			continue // usually unused
		}
		builder.WriteByte('\n')
		builder.WriteString(record[index["Common Name or Certificate Name"]])
		builder.WriteByte('\n')
		cert := record[index["PEM Info"]]
		builder.WriteString(cert[1 : len(cert)-1])
		builder.WriteByte('\n')
	}
	return builder.Bytes(), nil
}
