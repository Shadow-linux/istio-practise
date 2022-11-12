package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"mypro/gsrc/pbfiles"
)

const (
	serverName = "grpc.shadow.com"
	//target     = "grpc.shadow.com:31136"
	target     = "grpc.shadow.com:30743"
	clientCert = "certs/out/clientgrpc.crt"
	clientKey  = "certs/out/clientgrpc.key"
	serverCert = "certs/out/grpc.shadow.com.crt"
	serverKey  = "certs/out/grpc.shadow.com.key"
	caCert     = "certs/out/ShadowCA.crt"
)

func main() {

	//单向认证
	//creds, err := credentials.NewClientTLSFromFile(serverCert, serverName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//client, err := grpc.DialContext(context.Background(),
	//	target,
	//	grpc.WithTransportCredentials(creds))

	//双向认证
	cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   serverName,
		RootCAs:      certPool,
	})
	client, err := grpc.DialContext(context.Background(),
		target,
		grpc.WithTransportCredentials(creds))

	// 无认证
	//client, err := grpc.DialContext(context.Background(), target, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	rsp := &pbfiles.ProdResponse{}
	err = client.Invoke(context.Background(),
		"/ProdService/GetProd",
		&pbfiles.ProdRequest{ProdId: 123}, rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Result)

}
