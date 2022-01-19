package main

import (
	"crypto/client"
	"flag"
	"fmt"
	"log"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "BTC", "The name of the fiat currency you would like to know the price of your crypto in",
	)

	nameOfCrypto := flag.String(
		"crypto", "ELF", "Input the name of the CryptoCurrency you would like to know the price",
	)
	flag.Parse()

	crypto, err := client.FetchCrypto(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
