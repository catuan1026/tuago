package tasks

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type TaskInf interface {
	Next() bool          // 是否继续
	EndpointTime() int64 //执行时间
	Exec() error         //执行
	Name() string        //任务名称
}

type TaskCenter struct {
	tasks    chan TaskInf
	taskPool []TaskInf //当任务暂停时，将任务放入任务池
	status   int8      //状态 1:暂停 0:运行
}

func (t *TaskCenter) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-t.tasks:
			if t.status == 1 { //是否暂停
				if t.taskPool == nil {
					t.taskPool = make([]TaskInf, 0)
				}
				t.taskPool = append(t.taskPool, task)
				continue
			}
			//执行任务
			go func() {
				defer func() {
					if err := recover(); err != nil {
						//TODO
						logrus.WithFields(logrus.Fields{
							"tip":       "任务执行异常",
							"task_name": task.Name(),
						}).Error("task exec panic", err)
					}
				}()

				if task.EndpointTime() <= time.Now().Unix() {
					//执行任务
					if err := task.Exec(); err != nil {
						//TODO
						logrus.WithFields(logrus.Fields{
							"tip":       "任务执行异常",
							"task_name": task.Name(),
						}).Error("task exec error", err)
					}
				}
				//判断是否继续
				if task.Next() {
					time.Sleep(time.Second * 1)
					t.tasks <- task
				}
			}()
		}
	}
}

func (t *TaskCenter) AddTask(tasks ...TaskInf) {
	for _, task := range tasks {
		t.tasks <- task
	}
}

// Pause 暂停任务
func (t *TaskCenter) Pause() {
	t.status = 1
}

// Resume 恢复任务
func (t *TaskCenter) Resume() {
	t.status = 0
	for _, task := range t.taskPool {
		t.tasks <- task
	}
	t.taskPool = nil
}
