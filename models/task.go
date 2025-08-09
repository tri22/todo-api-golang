package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskCreation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskUpdate struct {
	Completed bool `json:"completed"`
}

var (
	tasks = []Task{
		{
			ID:          1,
			Title:       "Sample Task",
			Description: "This is a sample task",
			Completed:   false,
		},
	}
	nextID = 2
)

func GetAllTasks() []Task {
	return tasks
}

func AddTask(req TaskCreation) int {
	//mapping request to new task
	task := Task{
		ID:          nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	tasks = append(tasks, task)
	//increase id when create successfully
	nextID++

	return task.ID
}

func UpdateTaskStatus(id int, completed bool) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = completed
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func GetTaskByID(id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}
