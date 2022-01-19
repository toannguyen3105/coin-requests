package client

import (
	"crypto/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// FetchCrypto is exported to show message
func FetchCrypto(fiat string, crypto string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=74b3dc72b1330faa6789ac71d3c1511ba542862a&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var cResp model.Cryptoresponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("Oh shit! An error occurred, please try again")
	}

	return cResp.TextOutput(), nil
}
