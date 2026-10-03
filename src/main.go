package main

import (
	"fmt"
	"jakegodsall/deployd/src/repository"
	"jakegodsall/deployd/src/repository/memory"
)

var repo repository.DeploymentRepository

func main() {
	repo = memory.NewDeploymentRepository()

	deployments, err := repo.FindAllByApplication("my-app")
	if err != nil {
		panic(err)
	}

	for _, d := range deployments {
		fmt.Println(d.String())
	}
}
