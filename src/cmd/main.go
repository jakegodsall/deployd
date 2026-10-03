package main

import (
	"database/sql"
	"fmt"
	"jakegodsall/deployd/src/repository"
	"jakegodsall/deployd/src/repository/sqlite"

	_ "modernc.org/sqlite"
)

var repo repository.DeploymentRepository

// In-memory
// func main() {
// 	repo = memory.NewDeploymentRepository()

// 	deployment, err := repo.FindByApplicationVersion("my-app", 2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(deployment.String())
// }

// SQLite
func main() {
	db, err := sql.Open("sqlite", "deployd.sqlite")
	if err != nil {
		panic(err)
	}

	repo = sqlite.NewDeploymentRepository(db)

	deployments, err := repo.FindAll()

	if err != nil {
		panic(err)
	}

	if len(deployments) == 0 {
		fmt.Println("No deployments stored")
	}

	for _, d := range deployments {
		fmt.Println(d.String())
	}
}
