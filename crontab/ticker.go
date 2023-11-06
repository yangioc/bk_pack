package crontab

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/rs/xid"
)

type Task struct {
	Name     string
	EntryID  cron.EntryID
	Function func()
	Spec     string
}

type SchedulerTicker struct {
	Cron *cron.Cron
	Jobs *tickerBroadcast
}

func New() *SchedulerTicker {
	return &SchedulerTicker{
		Cron: cron.New(cron.WithSeconds()),
		Jobs: new(tickerBroadcast),
	}
}

func (s *SchedulerTicker) Run() {
	s.Cron.Run()
}

// 每幾秒，無限迴圈執行,
// secondTime: 每幾秒可以執行,
// function: 要執行的 function
func (s *SchedulerTicker) AddTask(name, spec string, function func()) error {
	if name == "" {
		name = xid.New().String()
	}

	if _, exist := s.Jobs.Load(name); !exist {
		entryId, err := s.Cron.AddFunc(spec, function)
		if err != nil {
			return err
		}

		task := &Task{
			Name:     name,
			EntryID:  entryId,
			Function: function,
			Spec:     spec,
		}
		s.Jobs.Store(task.Name, task)
	} else {
		return fmt.Errorf("this taskName: %v have been used", name)
	}

	return nil
}

// 中斷任務,
// eventName: 要中斷的任務 id
func (s *SchedulerTicker) DeleteTask(taskName string) {
	if task, exist := s.Jobs.Load(taskName); exist {
		s.Cron.Remove(task.EntryID)
		s.Jobs.Delete(task.Name)
	}
}

func (s *SchedulerTicker) DeleteAllTask() {
	s.Jobs.Range(func(key string, t *Task) bool {
		s.Cron.Remove(t.EntryID)
		s.Jobs.Delete(t.Name)
		return true
	})
}

func (s *SchedulerTicker) GetTask(taskName string) (t *Task, exist bool) {
	return s.Jobs.Load(taskName)
}

func (s *SchedulerTicker) GetAllTasks() []*Task {
	tList := make([]*Task, 0)
	s.Jobs.Range(func(key string, t *Task) bool {
		tList = append(tList, t)
		return true
	})
	return tList
}
