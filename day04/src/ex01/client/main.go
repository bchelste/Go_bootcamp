package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var newOrder Order
	var cntr int
	var moneyAmount int

	flag.StringVar(&newOrder.CandyType, "k", "", "accepts two-letter abbreviation for the candy type")
	flag.IntVar(&cntr, "c", 0, "count of candy to buy")
	flag.IntVar(&moneyAmount, "m", 0, "amount of money you \"gave to machine\"")
	flag.Parse()

	newOrder.CandyCount = int32(cntr)
	newOrder.Money = int32(moneyAmount)

	currentOrder, err := json.MarshalIndent(newOrder, "", "    ")
	if (err != nil) {
		log.Fatalln("MarshalIndent error with your order params")
		return
	}
	orderReader := bytes.NewReader(currentOrder)

	client := getClient()
	resp, err := client.Post("https://candy.tld:25565/buy_candy", "application/json", orderReader)
	must(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	must(err)

	// fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))

	switch (resp.StatusCode) {
	case 400:
		fmt.Println("---400")
		var answer InlineResponse400
		err := json.Unmarshal(body, &answer)
		if (err != nil) {
			log.Println("something wrong with response 400")
			return
		}
		fmt.Printf("Failure: %d - \"%s\"\n", resp.StatusCode, answer.Error_)
		return
	case 402:
		fmt.Println("---402")
		var answer InlineResponse402
		err := json.Unmarshal(body, &answer)
		if (err != nil) {
			log.Println("something wrong with response 402")
			return
		}
		fmt.Printf("Failure: %d - \"%s\"\n", resp.StatusCode, answer.Error_)
		return
	case 201:
		fmt.Println("---201")
		var answer InlineResponse201
		err := json.Unmarshal(body, &answer)
		if (err != nil) {
			log.Println("something wrong with response 201")
			return
		}
		fmt.Printf("%s Your change is %d\n", answer.Thanks, answer.Change)
		return
	default:
		fmt.Println("Unknown type of answer, try again.")
	}
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../minica.pem")
	cp.AppendCertsFromPEM(data)

	// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

	config := &tls.Config{
		// Certificates: []tls.Certificate{c},
		RootCAs:              	cp,
		GetClientCertificate:	ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate:	CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}