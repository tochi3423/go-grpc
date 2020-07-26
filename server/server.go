package main

import (
	"context"
	"log"
	"net"
	"./proto"
	"google.golang.org/grpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db, err = gorm.Open("mysql", "souta:souta3423@tcp(localhost:3306)/nogizaka")
)

func main() {
	listener, err := net.Listen("tcp", ":5200")
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
	type Profile struct {
		Id            int64
		Name          string
		Birthday      string
		Blood         string
		Constellation string
		Height        string
		Status        string
		Description   string
		URL           string
 }

	var profile = Profile{}

	db.New().
	Table("members").
	Select("*").
	Where("name = ?", req.Name).
	Find(&profile)

	return &nogizaka.Profile{
		Name: profile.Name,
		Birthday: profile.Birthday,
		Constellation: profile.Constellation,
		Height: profile.Height,
		Status: profile.Status,
		Description: profile.Description,
	}, nil
}