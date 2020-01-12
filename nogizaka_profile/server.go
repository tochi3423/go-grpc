package main

import (
	"context"
	"log"
	"net"
	"./lib"
	"google.golang.org/grpc"
	"github.com/coopernurse/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db, err = sql.Open("mysql", "souta:souta3423@tcp(localhost:3306)/nogizaka")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	nogizaka.RegisterNogizakaProfileServer(grpcServer, &server{})
	log.Printf("nogizaka profile server is running!")
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) FetchProfile(ctx context.Context, req *nogizaka.MemberName) (*nogizaka.Profile, error) {
	var profile = []nogizaka.Profile{}
	dbmap.Select(&profile, "SELECT * FROM members WHERE name =" + "\"" + req.Name + "\"" + ";")
	return &nogizaka.Profile{
		Name: profile[0].Name,
		Birthday: profile[0].Birthday,
		Constellation: profile[0].Constellation,
		Height: profile[0].Height,
		Status: profile[0].Status,
	}, nil
}