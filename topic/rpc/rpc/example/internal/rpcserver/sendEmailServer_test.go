package rpcserver

import (
	"testing"

	"context"
	"fmt"

	"example/internal/protodatasvr"
)

func TestSendEmailServer_SendSms(t *testing.T) {

	sms := &SendEmailServer{}

	toAddr := []string{"2483777000@qq.com"}

	rsp, err := sms.SendEmail(context.Background(), &protodatasvr.SendEmailRequest{
		Address: toAddr,
		Subject: "MySubject",
		Body:    "This is body.",
	})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(rsp)
}
