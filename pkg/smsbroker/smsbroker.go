package smsbroker

import "time"

type SMSBroker struct{}

func (broker *SMSBroker) SendSMS(tel string) (bool, error) {
	time.Sleep(time.Second * 4)
	return true, nil
}
