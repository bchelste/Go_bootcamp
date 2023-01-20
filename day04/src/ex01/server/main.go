package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
)

var candyStorage []Item = initStorage()

func main() {
	log.Printf("Candy server was started")
	server := getServer()
	http.HandleFunc("/buy_candy", BuyCandy)
	must(server.ListenAndServeTLS("", ""))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")
	w.Write([]byte("Hey GopherCon!"))
}

func getServer() *http.Server {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../minica.pem")
	cp.AppendCertsFromPEM(data)

	// c, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")

	tls := &tls.Config{
		// Certificates:          []tls.Certificate{c},
		ClientCAs:             cp,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		GetCertificate:        CertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: CertificateChains,
	}

	server := &http.Server{
		Addr:      ":25565",
		TLSConfig: tls,
	}
	return server
}

func must(err error) {
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}