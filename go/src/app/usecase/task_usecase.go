package usecase

import (
	"app/model"
	"app/repository"
	"app/validator"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUseCase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUseCase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUseCase {
	return &taskUseCase{tr, tv}
}

func (tu *taskUseCase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	// リポジトリを通じて全てのタスクを取得する
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}

	// レスポンス用のタスクスライスを作成する
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (tu *taskUseCase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	// リポジトリを通じて特定のIDのタスクを取得する
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	// レスポンス用のタスクを作成する
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タスクのバリデーションを行う
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// リポジトリを通じてタスクを作成する
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	// レスポンス用のタスクを作成する
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクのバリデーションを行う
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// リポジトリを通じてタスクを更新する
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	// レスポンス用のタスクを作成する
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) DeleteTask(userId uint, taskId uint) error {
	// タスクの削除をリポジトリを通じて行う
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}

	return nil
}
