package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"k8s.io/klog/v2"
	"sync"
)

type Manager interface {
	AddTaskWithJob(job Task) error
	Len() int
	Start(name string) error
	Stop(name string) error
	StopAll()
}

var Mgr = NewTaskManager()

type manager struct {
	taskMap   map[string]*cron.Cron
	taskMutex *sync.Mutex
}

func NewTaskManager() Manager {
	return &manager{
		taskMap:   make(map[string]*cron.Cron),
		taskMutex: &sync.Mutex{},
	}
}

func (tm *manager) AddTaskWithJob(job Task) error {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	_, ok := tm.getCron(job.GetName())
	if !ok {
		c := tm.addTaskLocked(job.GetName())
		_, err := c.AddJob(job.GetSchedule(), job)
		if err != nil {
			return fmt.Errorf("task add job error: %w", err)
		}
		return nil
	}

	return nil
}

func (tm *manager) addTaskLocked(name string) *cron.Cron {
	klog.Infof("add task %s", name)
	c := cron.New(cron.WithSeconds())
	tm.taskMap[name] = c
	return c
}

func (tm *manager) getCron(name string) (*cron.Cron, bool) {
	c, ok := tm.taskMap[name]
	return c, ok
}

func (tm *manager) Start(name string) error {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	c, ok := tm.getCron(name)
	if !ok {
		return fmt.Errorf("task %s not added", name)
	}
	c.Start()
	return nil
}

func (tm *manager) Stop(name string) error {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	c, ok := tm.getCron(name)
	if !ok {
		return fmt.Errorf("task %s not added", name)
	}
	c.Stop()

	delete(tm.taskMap, name)

	return nil
}

func (tm *manager) Len() int {
	return len(tm.taskMap)
}

func (tm *manager) StopAll() {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	if tm.Len() == 0 {
		return
	}

	for name, c := range tm.taskMap {
		c.Stop()
		delete(tm.taskMap, name)
		klog.Infof("%s task stop", name)
	}
}
