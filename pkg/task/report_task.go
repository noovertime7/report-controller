package task

import (
	"fmt"
	defaultv1 "github.com/noovertime7/report-controller/api/v1"
)

type reportTask struct {
	Name     string
	Schedule string
	report   *defaultv1.Report
}

func NewReportTask(report *defaultv1.Report) Task {
	return &reportTask{
		Name:     report.Name,
		Schedule: report.Spec.Schedule,
		report:   report,
	}
}

func (r *reportTask) GetName() string {
	return r.Name
}

func (r *reportTask) GetSchedule() string {
	return r.Schedule
}

func (r *reportTask) Run() {
	fmt.Println("start", r.report.Name)
	//fmt.Println("report", r.report)
}
