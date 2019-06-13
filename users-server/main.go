package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	grpc_health_v1 "github.com/maelvls/users-grpc/schema/health/v1"
	"github.com/maelvls/users-grpc/schema/user"
	"github.com/maelvls/users-grpc/users-server/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Set during build, e.g.: go build  -ldflags"-X main.version=$(git
// describe)". Some vars are commented out so that 'golangci-lint run' us
// happy, but you can un-comment them and use them.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// I'll try to use the go-idiomatic 'happy-path' way.
func main() {

	// Set the port according to PORT.
	port := 8000
	if str := os.Getenv("PORT"); str != "" {
		var err error
		port, err = strconv.Atoi(str)
		if err != nil {
			log.Errorf("PORT not valid: %v", str)
			os.Exit(1)
		}
	}

	// Set the log format according to LOG_FORMAT.
	switch os.Getenv("LOG_FORMAT") {
	case "", "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.Errorf(`expected LOG_FORMAT: 'json', 'text'`)
		os.Exit(1)
	}

	// Set to verbose.
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.TraceLevel)
	}

	log.Printf("serving on port %v, version %s (git %s, built on %s)", port, version, commit, date)

	if err := Run(port); err != nil {
		log.Errorf("launching server: %v", err)
		os.Exit(1)
	}
}

// Run starts the server
func Run(port int) error {
	svc := service.NewUserImpl()
	err := svc.LoadSampleUsers()
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	user.RegisterUserServiceServer(srv, svc)
	grpc_health_v1.RegisterHealthServer(srv, &service.HealthImpl{})

	// Maybe we should let the user choose which address he wants to bind
	// to; in our case, when the host is unspecified (:80 is equivalent to
	// 0.0.0.0:80) then the local system. See: https://godoc.org/net#Dial
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		os.Exit(1)
	}
	return srv.Serve(lis)
}
