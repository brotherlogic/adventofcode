package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"
	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, cancel := utils.ManualContext("adventofcode", time.Minute)
	defer cancel()

	clientCert, err := tls.LoadX509KeyPair("/home/simon/keys/client.pem", "/home/simon/keys/client.key")
	if err != nil {
		log.Fatalf("Unable to load: %v", err)
	}

	trustedCert, err := ioutil.ReadFile("/home/simon/keys/cacert.pem")

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		log.Fatalf("Unable to append cert: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewAdventServerServiceClient(conn)
	res, err := client.Solve(ctx, &pb.SolveRequest{Year: 2017, Day: 12, Part: 2})

	fmt.Printf("%v -> %v\n", res, err)
}
