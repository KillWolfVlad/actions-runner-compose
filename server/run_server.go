package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KillWolfVlad/actions-runner-compose/configs"
	"github.com/KillWolfVlad/actions-runner-compose/runners"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/moby/moby/client"
)

func RunServer(config configs.Config, dockerClient *client.Client) {
	hook, err := github.New(github.Options.Secret(config.WebhookSecret))

	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc(config.WebhookPath, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PingEvent, github.WorkflowJobEvent)

		if err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)

			return
		}

		switch payload.(type) {

		case github.WorkflowJobPayload:
			workflowJob := payload.(github.WorkflowJobPayload)

			err = runners.QueueRunner(config, dockerClient, workflowJob)

			if err != nil {
				log.Println(err)

				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, err)

				return
			}
		}

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "ok")
	})

	addr := fmt.Sprintf(":%d", config.Port)

	log.Printf("server run and listen on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}
