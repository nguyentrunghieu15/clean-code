package main

import "fmt"

type NotifierInterface interface {
	Send(string)
}

type BaseDecoraterNotifier struct {
	Wapper NotifierInterface
}

func (s *BaseDecoraterNotifier) Send(msg string) {
	if s.Wapper != nil {
		s.Wapper.Send(msg)
	}
}

type SMSNotifierDerorater struct {
	BaseDecoraterNotifier
}

func (s *SMSNotifierDerorater) Send(msg string) {
	fmt.Println("Send message cross SMS :", msg)
	s.BaseDecoraterNotifier.Send(msg)
}

type FacebookNotifierDerorater struct {
	BaseDecoraterNotifier
}

func (s *FacebookNotifierDerorater) Send(msg string) {
	fmt.Println("Send message cross Face :", msg)
	s.BaseDecoraterNotifier.Send(msg)
}

type ZaloNotifierDerorater struct {
	BaseDecoraterNotifier
}

func (s *ZaloNotifierDerorater) Send(msg string) {
	fmt.Println("Send message cross Zalo :", msg)
	s.BaseDecoraterNotifier.Send(msg)
}

func main() {
	isEnableFacebook := true
	isEnableSMS := true
	isEnableZalo := true
	var stack NotifierInterface
	if isEnableFacebook {
		stack = &FacebookNotifierDerorater{}
	}
	if isEnableSMS {
		stack = &SMSNotifierDerorater{BaseDecoraterNotifier{Wapper: stack}}
	}
	if isEnableZalo {
		stack = &ZaloNotifierDerorater{BaseDecoraterNotifier{Wapper: stack}}
	}

	stack.Send("Hi")
}
