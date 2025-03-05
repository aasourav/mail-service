package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aasourav/mail-service/utils"
	mailsvc "github.com/aasourav/proto/mail-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	mailsvc.UnimplementedEmailServiceServer
}

type verification struct {
	VerificationLink string
	Title            string
	Name             string
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	mailsvc.RegisterEmailServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("listening on port: 9000")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) SendVerificationEmail(c context.Context, req *mailsvc.MailServiceRequest) (*mailsvc.MailServiceResponse, error) {
	log.Print("received request from client: ", req.Name)
	emailData := verification{
		VerificationLink: req.VerificationLink,
		Title:            "AES Meal",
		Name:             req.Name,
	}
	err := utils.SendEmail(req.Email, "verify admin account", "templates/emailtemplate.html", emailData)
	if err != nil {
		return &mailsvc.MailServiceResponse{Response: "send email failed"}, err
	}

	return &mailsvc.MailServiceResponse{Response: "email set successfully"}, nil
}
