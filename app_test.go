package main

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	BrokerSendSMS = func(tel string) (bool, error) {
		return false, nil
	}
	status, err := SendSMS("1234567890")
	fmt.Println(status)
	fmt.Println(err)

}
