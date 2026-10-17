package todo

import "time"

type Task struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func NewTask(name string, description string) *Task {
	return &Task{
		Name:        name,
		Description: description,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
}

func (t *Task) Complete() {
	CompletedTime := time.Now()

	t.Completed = true
	t.CompletedAt = &CompletedTime
}

func (t *Task) UnComplete() {
	t.Completed = false
	t.CompletedAt = nil
}
