package crontab

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

// 任務資料
type CronJob struct {
	Name        string            // 任務名稱
	Spec        string            // 任務排程
	CallBack    func(interface{}) // 排程啟動任務
	TaskContext interface{}       // 任務內容
}

func (self *CronJob) Run() {
	self.CallBack(self.TaskContext)
}

type CronHandle struct {
	Cron    *cron.Cron
	taskMap map[string]cron.EntryID

	mu sync.RWMutex
}

func New() *CronHandle {
	return &CronHandle{
		Cron:    cron.New(cron.WithSeconds()),
		taskMap: make(map[string]cron.EntryID),
	}
}

func (self *CronHandle) Run() {
	self.Cron.Run()
}

func (self *CronHandle) AddFunc(name, spec string, function func()) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	if _, ok := self.taskMap[name]; !ok {
		return fmt.Errorf("Job %s already exist.", name)
	}

	id, err := self.Cron.AddFunc(spec, function)
	if err != nil {
		return err
	}

	self.taskMap[name] = id
	return nil
}

func (self *CronHandle) AddJob(task *CronJob) error {
	self.mu.Lock()
	defer self.mu.Unlock()

	id, err := self.Cron.AddJob(task.Spec, task)
	if err != nil {
		return err
	}

	self.taskMap[task.Name] = id
	return nil
}

func (self *CronHandle) DeleteTask(taskName string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	id, ok := self.taskMap[taskName]
	if !ok {
		return
	}

	delete(self.taskMap, taskName)
	self.Cron.Remove(id)
}

func (self *CronHandle) DeleteAllTask() {
	self.mu.Lock()
	defer self.mu.Unlock()

	for _, id := range self.taskMap {
		self.Cron.Remove(id)
	}

	self.taskMap = make(map[string]cron.EntryID)
}
