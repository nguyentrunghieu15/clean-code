package main

import "log"

type Adminstrator struct {
	uuid     string
	username string
}

var admin *Adminstrator

func getInstanceAdmin() *Adminstrator {
	if admin == nil {
		admin = &Adminstrator{}
		log.Println("Create admin!: ", &admin)
		return admin
	}
	log.Println("Get admin: ", &admin)
	return admin
}

func main() {
	for i := 0; i < 5; i++ {
		getInstanceAdmin()
	}
}
