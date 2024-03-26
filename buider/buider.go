package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Country  string
	Province string
	Street   string
	Ward     string
}

// AtStreet implements AddressBuilder.
func (s *Address) AtStreet(street string) *Address {
	s.Street = street
	return s
}

// AtWard implements AddressBuilder.
func (s *Address) AtWard(ward string) *Address {
	s.Ward = ward
	return s
}

// InCountry implements AddressBuilder.
func (s *Address) InCountry(country string) *Address {
	s.Country = country
	return s
}

// InProvince implements AddressBuilder.
func (s *Address) InProvince(province string) *Address {
	s.Province = province
	return s
}

func (s *Address) Build() *Address {
	return s
}

type Student struct {
	ID          string
	Fname       string
	Lname       string
	PhoneNumber string
	Age         int16
	Address     Address
}

// At implements StudentBuilder.
func (s *Student) At(address *Address) *Student {
	s.Address = *address
	return s
}

func (s *Student) AtStreet(street string) *Student {
	s.Address.AtStreet(street)
	return s
}

// AtWard implements AddressBuilder.
func (s *Student) AtWard(ward string) *Student {
	s.Address.AtWard(ward)
	return s
}

// InCountry implements AddressBuilder.
func (s *Student) InCountry(country string) *Student {
	s.Address.InCountry(country)
	return s
}

// InProvince implements AddressBuilder.
func (s *Student) InProvince(province string) *Student {
	s.Address.InProvince(province)
	return s
}

// HasPhoneNumber implements StudentBuilder.
func (s *Student) HasPhoneNumber(phoneNumber string) *Student {
	s.PhoneNumber = phoneNumber
	return s
}

// WithAge implements StudentBuilder.
func (s *Student) WithAge(age int16) *Student {
	s.Age = age
	return s
}

// WithFName implements StudentBuilder.
func (s *Student) WithFName(fName string) *Student {
	s.Fname = fName
	return s
}

// WithId implements StudentBuilder.
func (s *Student) WithId(id string) *Student {
	s.ID = id
	return s
}

// WithLName implements StudentBuilder.
func (s *Student) WithLName(lName string) *Student {
	s.Lname = lName
	return s
}

// Build implements StudentBuilder.
func (s *Student) Build() *Student {
	return s
}

type AddressBuilder interface {
	InCountry(string) *Address
	InProvince(string) *Address
	AtStreet(string) *Address
	AtWard(string) *Address
	Build() *Address
}

type StudentBuilder interface {
	WithId(string) *Student
	WithFName(string) *Student
	WithLName(string) *Student
	HasPhoneNumber(string) *Student
	WithAge(int16) *Student
	InCountry(string) *Student
	At(*Address) *Student
	InProvince(string) *Student
	AtStreet(string) *Student
	AtWard(string) *Student
	Build() *Student
}

func NewStudentBuider() StudentBuilder {
	return &Student{}
}

func NewAddressBuilder() AddressBuilder {
	return &Address{}
}

func main() {
	student1 := NewStudentBuider().
		WithId("877074").
		HasPhoneNumber("04762382").
		WithAge(22).
		WithFName("Nguyen").
		WithLName("Hiro").
		At(
			NewAddressBuilder().
				AtStreet("Le Quan").
				AtWard("Thach Khoi").
				InProvince("Hai Duong").
				InCountry("Viet Nam").
				Build(),
		).
		Build()
	if v, e := json.Marshal(student1); e == nil {
		fmt.Println(string(v))
	}
}
