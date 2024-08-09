package usecases

import (
	"errors"
	"strings"
	domain "task_manager/Domain"
)

type TaskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(repository domain.TaskRepository)*TaskUseCase{
	return &TaskUseCase{
		taskRepository: repository,
	}
}

func (taskUseCase *TaskUseCase) GetAllTasks(role string, uid string)([]domain.Task, error){
	if role == "" || (uid == "" && strings.ToUpper(role) == "USER"){
		return nil, errors.New("user not logged in")
	}
	allTasks, err := taskUseCase.taskRepository.GetAllTasks(role, uid)
	if err != nil{
		return nil, err
	}
	return allTasks, nil
}
func (taskUseCase *TaskUseCase) GetTaskById(id string, uid string, role string)(domain.Task, error){
	if role == "" || (uid == "" && strings.ToUpper(role) == "USER") || id == ""{
		return domain.Task{}, errors.New("user not logged in")
	}
	task, err := taskUseCase.taskRepository.GetTaskById(id, uid, role)
	if err != nil{
		return domain.Task{}, err
	}
	return task, nil
}
func (taskUseCase *TaskUseCase) DeleteTaskById(id string, uid string, role string)(bool, error){
	if role == "" || (uid == "" && strings.ToUpper(role) == "USER") || id == ""{
		return false, errors.New("user not logged in")
	}
	deleted, err := taskUseCase.taskRepository.DeleteTaskById(id, uid,role)
	if err != nil{
		return false, err
	}
	return deleted, nil
}
func (taskUseCase *TaskUseCase) CreateTask(newTask domain.Task, uid string, role string)(interface{}, error){
	if role == "" || (uid == "" && strings.ToUpper(role) == "USER") {  //|| newTask == domain.Task{}{
		return nil, errors.New("user not logged in")
	}
	createdId, err := taskUseCase.taskRepository.CreateTask(newTask, uid, role)
	if err != nil{
		return false, err
	}
	return createdId, nil
}
func (taskUseCase *TaskUseCase) UpdateTask(id string, uid string, role string, newTask domain.Task)(bool, error){
	if role == "" || (uid == "" && strings.ToUpper(role) == "USER") || id == ""{
		return false, errors.New("user not logged in")
	}
	updated, err := taskUseCase.taskRepository.UpdateTask(id, uid, role, newTask)
	if err != nil{
		return false, err
	}
	return updated, nil
}