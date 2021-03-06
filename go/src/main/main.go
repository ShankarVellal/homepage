package main

import (
	"crypto/tls"
    "log"
	"net/http"

    "golang.org/x/crypto/acme/autocert"
)

func main() {
	log.Printf("starting")
    go http.ListenAndServe(":80", http.HandlerFunc(redirect))

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
    fs := http.StripPrefix("/", http.FileServer(http.Dir("public")))
    http.Handle("/", fs)

	if err := server.ListenAndServeTLS("", ""); err != nil {
        log.Fatalf(err.Error())
    } //key and cert are comming from Let's Encrypt
}

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host + req.URL.Path 
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target, http.StatusMovedPermanently)
}