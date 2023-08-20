package task

import "github.com/robfig/cron/v3"

type Task interface {
	GetName() string
	GetSchedule() string
	cron.Job
}
