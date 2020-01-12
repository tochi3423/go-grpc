package main

import (
    "context"
    "fmt"
    "../nogizaka_profile/lib"
    "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := nogizaka.NewNogizakaProfileClient(conn)
	req := &nogizaka.MemberName{"白石 麻衣"}
	profile, err := client.FetchProfile(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profile.Name, profile.Birthday, profile.Constellation, profile.Height)
}