package task

import (
	"fmt"
	"testing"
	"time"
)

type MockJob struct {
}

func (j *MockJob) GetName() string {
	return "task1"
}

func (j *MockJob) Run() {
	fmt.Println("task1 task run")
}

func (j *MockJob) GetSchedule() string {
	return "* * * * * *"
}

func TestTaskManager(t *testing.T) {
	// 创建任务管理器实例
	taskManager := NewTaskManager()

	// 添加任务
	err := taskManager.AddTaskWithJob(&MockJob{})
	if err != nil {
		t.Errorf("Failed to add task: %v", err)
	}

	// 启动任务
	err = taskManager.Start("task1")
	if err != nil {
		t.Errorf("Failed to start task: %v", err)
	}

	// 等待一段时间
	time.Sleep(10 * time.Second)

	// 停止任务
	err = taskManager.Stop("task1")
	if err != nil {
		t.Errorf("Failed to stop task: %v", err)
	}

	// 获取任务管理器长度
	length := taskManager.Len()
	if length != 0 {
		t.Errorf("Unexpected task manager length: expected 1, got %d", length)
	}

	// 停止所有任务
	taskManager.StopAll()

	// 再次获取任务管理器长度
	length = taskManager.Len()
	if length != 0 {
		t.Errorf("Unexpected task manager length after StopAll(): expected 0, got %d", length)
	}
}
