package main

import (
	"context"
	"fmt"
	"../server/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5200", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := nogizaka.NewNogizakaProfileClient(conn)
	req := &nogizaka.MemberName{"白石麻衣"}
	profile, err := client.FetchProfile(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profile.Name)
	fmt.Println(profile.Birthday)
	fmt.Println(profile.Constellation)
	fmt.Println(profile.Height)
	fmt.Println(profile.Description)
}