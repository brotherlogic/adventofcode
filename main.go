package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/brotherlogic/adventofcode/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/brotherlogic/adventofcode/proto"
)

var (
	port        = flag.Int("port", 8080, "The server port.")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
)

func main() {
	flag.Parse()

	s := &server.Server{}

	// Load the server certificate and its key
	serverCert, err := tls.LoadX509KeyPair("sp/server.pem", "sk/server.key")
	if err != nil {
		log.Fatalf("failed to load server certificate and key. %s.", err)
	}
	log.Printf("GOT: %v", serverCert.Certificate)
	data, err := ioutil.ReadFile("sp/server.pem")
	log.Printf("READ: %v, %v", string(data), err)

	// Load the CA certificate
	trustedCert, err := ioutil.ReadFile("cp/cacert.pem")
	if err != nil {
		log.Fatalf("failed to load trusted certificate. %s.", err)
	}

	// Put the CA certificate to certificate pool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		log.Fatal("failed to append trusted certificate from pem to certificate pool")
	}

	// Create the TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)
	log.Printf("made cred: %v -> %v", string(trustedCert), cred)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", *port, err)
	}
	gs := grpc.NewServer(grpc.Creds(cred))
	//gs := grpc.NewServer()
	pb.RegisterAdventServerServiceServer(gs, s)
	log.Printf("secure server listening at %v", lis.Addr())

	// Setup prometheus export
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	}()

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
