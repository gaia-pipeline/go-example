package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func CreateUser() error {
	log.Println("CreateUser has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateUser has been finished!")

	return nil
}

func MigrateDB() error {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("MigrateDB has been finished!")

	return nil
}

func CreateNamespace() error {
	log.Println("CreateNamespace has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateNamespace has been finished!")

	return nil
}

func CreateDeployment() error {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateDeployment has been finished!")

	return nil
}

func CreateService() error {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateService has been finished!")

	return nil
}

func CreateIngress() error {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("CreateIngress has been finished!")

	return nil
}

func Cleanup() error {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	log.Println("Cleanup has been finished!")

	return nil
}

func main() {
	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
