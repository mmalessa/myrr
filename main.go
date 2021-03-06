package main

import (
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"

	// services (plugins)
	"github.com/spiral/roadrunner/service/env"
	"github.com/spiral/roadrunner/service/gzip"
	"github.com/spiral/roadrunner/service/headers"
	"github.com/spiral/roadrunner/service/health"
	"github.com/spiral/roadrunner/service/http"
	"github.com/spiral/roadrunner/service/limit"
	"github.com/spiral/roadrunner/service/metrics"
	"github.com/spiral/roadrunner/service/reload"
	"github.com/spiral/roadrunner/service/rpc"
	"github.com/spiral/roadrunner/service/static"

	// additional commands and debug handlers
	_ "github.com/spiral/roadrunner/cmd/rr/http"
	_ "github.com/spiral/roadrunner/cmd/rr/limit"

	"github.com/spiral/jobs/v2"
	"github.com/spiral/jobs/v2/broker/amqp"
	"github.com/spiral/jobs/v2/broker/beanstalk"
	"github.com/spiral/jobs/v2/broker/ephemeral"
	"github.com/spiral/jobs/v2/broker/sqs"

	_ "github.com/spiral/jobs/v2/cmd/rr-jobs/jobs"
)

func main() {
	rr.Container.Register(env.ID, &env.Service{})
	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(http.ID, &http.Service{})
	rr.Container.Register(metrics.ID, &metrics.Service{})
	rr.Container.Register(headers.ID, &headers.Service{})
	rr.Container.Register(static.ID, &static.Service{})
	rr.Container.Register(limit.ID, &limit.Service{})
	rr.Container.Register(health.ID, &health.Service{})
	rr.Container.Register(gzip.ID, &gzip.Service{})
	rr.Container.Register(reload.ID, &reload.Service{})

	rr.Container.Register(rpc.ID, &rpc.Service{})
	rr.Container.Register(jobs.ID, &jobs.Service{
		Brokers: map[string]jobs.Broker{
			"amqp":      &amqp.Broker{},
			"ephemeral": &ephemeral.Broker{},
			"beanstalk": &beanstalk.Broker{},
			"sqs":       &sqs.Broker{},
		},
	})

	// you can register additional commands using cmd.CLI
	rr.Execute()
}
