package main

import (
	"crypto/tls"
    "crypto/x509"
	"fmt"
	"golang.org/x/crypto/acme/autocert"
    "log"
	"net/http"
)

var (
    client *http.Client
    pool   *x509.CertPool
)

func init() {
    log.Println("running init")
    pool = x509.NewCertPool()
    pool.AppendCertsFromPEM(pemCerts)
    client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homePage)
}

func main() {
	log.Printf("starting request")
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("shankarvellal.com"),
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

const homePage = 
`<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
  </head>
  <body>
    <h1>Hello, world!</h1>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
  </body>
</html>
`