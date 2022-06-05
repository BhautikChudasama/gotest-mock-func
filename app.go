package main

import (
	"fmt"

	"github.com/bhautikchudasama/golang_mocktest/pkg/smsbroker"
)

func main() {
	status, err := SendSMS("1234567890")
	if err != nil {
		fmt.Println(err)
	}
	if status {
		fmt.Println("SMS send successfully")
	}
}

func InitBroker() *smsbroker.SMSBroker {
	return new(smsbroker.SMSBroker)
}

var BrokerSendSMS func(string) (bool, error) = nil

func SendSMS(tel string) (bool, error) {
	broker := InitBroker()
	if BrokerSendSMS == nil {
		BrokerSendSMS = broker.SendSMS
	}
	return BrokerSendSMS(tel)
}
