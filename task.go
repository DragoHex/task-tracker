package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"text/tabwriter"
	"time"
)

type TaskStatus int

func (e TaskStatus) String() string {
	return [...]string{"todo", "in-progress", "done"}[e-1]
}

func (e TaskStatus) EnumIndex() int {
	return int(e)
}

const (
	_ TaskStatus = iota
	ToDo
	InProgress
	Done
)

type Task struct {
	Id          int        `json:"id,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty"`
	CreateAt    time.Time  `json:"create_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

func NewTask() *Task {
	return &Task{
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
		Status:    ToDo,
	}
}

func (t *Task) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "id:\t%d\t\n", t.Id)
	fmt.Fprintf(w, "description:\t%s\t\n", t.Description)
	fmt.Fprintf(w, "status:\t%s\t\n", t.Status)
	fmt.Fprintf(w, "id:\t%s\t\n", t.CreateAt.Format("02 Jan 06 15:04 IST"))
	fmt.Fprintf(w, "id:\t%s\t\n", t.UpdatedAt.Format("02 Jan 06 15:04 IST"))
	w.Flush()
}

type TaskData struct {
	CurrentID int    `json:"current_id,omitempty"`
	Count     int    `json:"count,omitempty"`
	Tasks     []Task `json:"tasks,omitempty"`
}

func NewTaskData() *TaskData {
	return &TaskData{
		CurrentID: 0,
		Count:     0,
	}
}

func (t *TaskData) Add(task *Task) {
	t.CurrentID = t.CurrentID + 1
	t.Count = t.Count + 1
	task.Id = t.CurrentID
	t.Tasks = append(t.Tasks, *task)
}

func (t *TaskData) Save() error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("data.json", data, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskData) Delete(id int) error {
	t.Count = t.Count - 1
	for i, task := range t.Tasks {
		if task.Id == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			break
		}
	}
	return fmt.Errorf("no task found with %d ID", id)
}
