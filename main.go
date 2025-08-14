package main

import (
	"log"
	"os"

	"github.com/KillWolfVlad/actions-runner-compose/configs"
	"github.com/KillWolfVlad/actions-runner-compose/docker"
	"github.com/KillWolfVlad/actions-runner-compose/server"
)

func main() {
	log.SetOutput(os.Stdout)

	config := configs.LoadConfig()

	dockerClient := docker.InitDockerClient()
	defer dockerClient.Close()

	server.RunServer(config, dockerClient)
}
