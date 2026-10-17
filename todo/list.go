package todo


type List struct{
	tasks map[string]Task

}

func NewList() *List{
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List)AddTask(task Task) error{
	if _, ok := l.tasks[task.Name]; ok {
		return ErrTaskAleradyExists
	}

	l.tasks[task.Name] = task
	return nil
}

func (l *List)DeleteTask(name string) error{
	if _, ok := l.tasks[name]; !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, name)

	return nil
}

func (l *List) CompleteTask(name string) error {
	task, ok := l.tasks[name]
	if !ok{
		return ErrTaskNotFound
	}

	task.Complete()
	l.tasks[name] = task

	return nil
}

func (l *List) UnCompleteTask(name string) error{
	task, ok := l.tasks[name]
	
	if !ok{
		return ErrTaskNotFound
	}

	task.UnComplete()
	l.tasks[name] = task

	return nil
}


func (l *List) GetAllTasks() map[string]Task{
	tmp := make(map[string]Task)

	for name, task := range l.tasks{
		tmp[name] = task
	}

	return tmp
}

func (l *List) GetTask(name string) (Task, error) {
	task, ok := l.tasks[name]
	if !ok{
		return Task{}, ErrTaskNotFound
	}

	
	return task, nil
}

func (l *List) GetUncTasks() map[string]Task {
	unc := make(map[string]Task)

	for name, task := range l.tasks{
		if !task.Completed{
			unc[name] = task
		}
	}
	
	return unc
}

func (l *List) UpdateTaskDescription(name string, NewDescription string) error {
	tmpTask, ok := l.tasks[name]

	if !ok{
		return ErrTaskNotFound
	}

	tmpTask.Description = NewDescription
	l.tasks[name] = tmpTask

	return nil
}

func (l *List) UpdateTaskName(name string, NewName string) error {
	tmpTask, ok := l.tasks[name]

	if !ok{
		return ErrTaskNotFound
	}

	if NewName != name {
		if _, exists := l.tasks[NewName]; exists {
			return ErrTaskAleradyExists
		}
	}

	tmpTask.Name = NewName
	delete(l.tasks, name)
	l.tasks[NewName] = tmpTask

	return nil
}

func (l *List) UpdateFullTask(name string, NewName string, NewDescription string) error {
	tmpTask,ok := l.tasks[name]

	if !ok {
		return ErrTaskNotFound
	}

	if NewName != name {
		if _, exists := l.tasks[NewName]; exists {
			return ErrTaskAleradyExists
		}
	}

	tmpTask.Name = NewName
	tmpTask.Description = NewDescription

	delete(l.tasks, name)
	l.tasks[NewName] = tmpTask

	return nil
}
