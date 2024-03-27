// Cơ chế để theo dõi và nhận thông báo đối tượng khác
package main

import "fmt"

type Message string

type Subcriber interface {
	Update(Message)
}

type Publisher interface {
	Notify(Message)
	Subscribe(Subcriber)
	UnSubscribe(Subcriber)
}

type Forecast struct {
	ListListener []Subcriber
}

// Notify implements Publisher.
func (s *Forecast) Notify(msg Message) {
	for _, v := range s.ListListener {
		v.Update(msg)
	}
}

// Subscribe implements Publisher.
func (s *Forecast) Subscribe(subcriber Subcriber) {
	s.ListListener = append(s.ListListener, subcriber)
}

// UnSubscribe implements Publisher.
func (s *Forecast) UnSubscribe(subcriber Subcriber) {
	// find index of Subscriber
	var idx int
	var isExists bool
	for i, v := range s.ListListener {
		if v == subcriber {
			idx = i
			break
		}
	}
	if isExists {
		s.ListListener = append(s.ListListener[:idx], s.ListListener[idx+1:]...)
	}
}

type Person struct {
	Name string
}

func (p *Person) Update(msg Message) {
	fmt.Printf("%v got a message :%v \n", p.Name, msg)
}

func main() {
	var p Publisher = &Forecast{}
	var person1 = &Person{Name: "Thanh"}
	var person2 = &Person{Name: "Le"}
	p.Subscribe(person1)
	p.Subscribe(person2)
	p.Notify("Rain")
}
