package main

import (
	"crypto/tls"
	"fmt"
    "log"
	"net/http"

    "golang.org/x/crypto/acme/autocert"
)

func indexHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
    fmt.Fprintf(w, html)
}

func main() {
	log.Printf("starting")
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("shankarvellal.com", "www.shankarvellal.com"),
        Cache:      autocert.DirCache("certs"), //folder for storing certificates
    }

    server := &http.Server{
        Addr: ":443",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }

	http.HandleFunc("/", indexHandler)
	if err := server.ListenAndServeTLS("", ""); err != nil {
        log.Fatalf(err.Error())
    } //key and cert are comming from Let's Encrypt
}
