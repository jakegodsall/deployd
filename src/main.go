package main

import (
	"fmt"
	"jakegodsall/deployd/src/repository"
	"jakegodsall/deployd/src/repository/memory"
)

var repo repository.DeploymentRepository

func main() {
	repo = memory.NewDeploymentRepository()

	deployment, err := repo.FindByApplicationVersion("my-app", 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(deployment.String())
}
