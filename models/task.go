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

var (
	tasks  []Task
	nextID = 1
)

func GetAllTasks() []Task {
	return tasks
}

func AddTask(req TaskCreation) int {
	task := Task{
		ID:          nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	tasks = append(tasks, task)
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
