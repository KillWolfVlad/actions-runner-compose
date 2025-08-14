package runners

import (
	"errors"
	"sync"

	"github.com/KillWolfVlad/actions-runner-compose/configs"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/moby/moby/client"
)

var (
	queuedWorkflowJobs []github.WorkflowJobPayload
	activeRunners      = 0
	queueRunnerMutex   sync.Mutex
)

func QueueRunner(config configs.Config, dockerClient *client.Client, workflowJob github.WorkflowJobPayload) error {
	runnersForStart := 0
	var workflowJobsForStart []github.WorkflowJobPayload

	queueRunnerMutex.Lock()

	switch workflowJob.Action {

	case "queued":
		queuedWorkflowJobs = append(queuedWorkflowJobs, workflowJob)

	case "completed":
		if activeRunners > 0 {
			activeRunners--
		}
	}

	if config.MaxRunners <= 0 {
		runnersForStart = len(queuedWorkflowJobs)
	} else {
		avialiableRunners := config.MaxRunners - activeRunners

		runnersForStart = min(avialiableRunners, len(queuedWorkflowJobs))
	}

	activeRunners += runnersForStart

	for i := 0; i < runnersForStart; i++ {
		workflowJobsForStart = append(workflowJobsForStart, queuedWorkflowJobs[0])
		queuedWorkflowJobs = queuedWorkflowJobs[1:]
	}

	queueRunnerMutex.Unlock()

	if runnersForStart > 0 {
		var wg sync.WaitGroup
		wg.Add(runnersForStart)

		errCh := make(chan error, runnersForStart)

		for i := 0; i < runnersForStart; i++ {
			go func(workflowJobForStart github.WorkflowJobPayload) {
				defer wg.Done()

				if err := startRunner(config, dockerClient, workflowJobForStart); err != nil {
					queueRunnerMutex.Lock()

					activeRunners--

					queueRunnerMutex.Unlock()

					errCh <- err
				}
			}(workflowJobsForStart[i])
		}

		wg.Wait()
		close(errCh)

		var errs error

		for err := range errCh {
			errs = errors.Join(errs, err)
		}

		return errs
	}

	return nil
}
