package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

<<<<<<< Updated upstream
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
=======
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
	//conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
>>>>>>> Stashed changes
	if err != nil {
		log.Fatalf("Bad dial: %v -> %v", err, cred)
	}

	client := pb.NewAdventOfCodeServiceClient(conn)
	iclient := pb.NewAdventOfCodeInternalServiceClient(conn)

	switch os.Args[2] {
	case "cookie":
		_, err := iclient.SetCookie(ctx, &pb.SetCookieRequest{Cookie: os.Args[3]})
		if err != nil {
			fmt.Printf("Unable to set cookie: %v\n", err)
		}
	case "solve":
		sflags := flag.NewFlagSet("solve", flag.ExitOnError)
		year := sflags.Int("year", -1, "year")
		day := sflags.Int("day", -1, "day")
		part := sflags.Int("part", -1, "part")
		if err := sflags.Parse(os.Args[3:]); err == nil {
			res, err := client.Solve(ctx, &pb.SolveRequest{Year: int32(*year), Day: int32(*day), Part: int32(*part)})
			fmt.Printf("%v -> %v\n", res, err)
		} else {
			fmt.Printf("Parse error: %v", err)
		}
	case "upload":
		sflags := flag.NewFlagSet("solve", flag.ExitOnError)
		year := sflags.Int("year", -1, "year")
		day := sflags.Int("day", -1, "day")
		data := sflags.String("data", "", "data")
		if err := sflags.Parse(os.Args[3:]); err == nil {
			data, err := ioutil.ReadFile(*data)
			if err != nil {
				log.Fatalf("Unable to run upload: %v", err)
			}
			res, err := iclient.Upload(ctx, &pb.UploadRequest{Year: int32(*year), Day: int32(*day), Part: 1, Data: string(data)})
			fmt.Printf("%v -> %v\n", res, err)
		}
	case "solution":
		sflags := flag.NewFlagSet("solve", flag.ExitOnError)
		year := sflags.Int("year", -1, "year")
		day := sflags.Int("day", -1, "day")
		part := sflags.Int("part", -1, "part")
		solution := sflags.Int("solution", 0, "solution")
		bigsolution := sflags.Int("big_solution", 0, "big solution")
		if err := sflags.Parse(os.Args[3:]); err == nil {
			res, err := iclient.AddSolution(ctx, &pb.AddSolutionRequest{Solution: &pb.Solution{Year: int32(*year), Day: int32(*day), Part: int32(*part), Answer: int32(*solution), BigAnswer: int64(*bigsolution)}})
			fmt.Printf("%v -> %v\n", res, err)
		}
	case "gsolution":
		sflags := flag.NewFlagSet("solve", flag.ExitOnError)
		year := sflags.Int("year", -1, "year")
		day := sflags.Int("day", -1, "day")
		part := sflags.Int("part", -1, "part")
		if err := sflags.Parse(os.Args[3:]); err == nil {
			res, err := iclient.GetSolution(ctx, &pb.GetSolutionRequest{Year: int32(*year), Day: int32(*day), Part: int32(*part)})
			fmt.Printf("%v -> %v\n", res, err)
		}
	case "tight":
		cclient := pb.NewSolverServiceClient(conn)
		solution, err := cclient.Solve(ctx, &pb.SolveRequest{Year: 2023, Day: 1, Part: 1})
		fmt.Printf("%v %v\n", solution, err)
	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}
