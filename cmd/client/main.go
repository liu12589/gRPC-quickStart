package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"useCpu/pb"
	"useCpu/sample"
)

func main() {
	// 获取命令行中的服务端地址
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	//创建一个传送笔记本电脑详细信息的客户端
	laptopClient := pb.NewLaptopServiceClient(conn)

	//创建一个电脑对象，并且获取电脑的详细信息
	laptop := sample.NewLaptop()
	//将电脑信息封装为req
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop:", err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)

}
