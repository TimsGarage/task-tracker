package repository

import (
	"fmt"
	"need-some-todo-overview/backend/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Projects

func (repository *Repository) ListProjects() ([]model.Project, error) {
	var projects []model.Project

	err := repository.db.
		Order("id ASC").
		Preload("Todos", func(query *gorm.DB) *gorm.DB {
			return query.Order("id DESC")
		}).
		Preload("Todos.Subtasks", func(query *gorm.DB) *gorm.DB {
			return query.Order("id ASC")
		}).
		Find(&projects).
		Error

	if err != nil {
		return nil, fmt.Errorf("list projects: %w", err)
	}

	return projects, nil
}

func (repository *Repository) FindProject(
	projectID uint,
) (*model.Project, error) {
	var project model.Project

	err := repository.db.
		First(&project, projectID).
		Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repository *Repository) CreateProject(
	project *model.Project,
) error {
	if err := repository.db.Create(project).Error; err != nil {
		return fmt.Errorf("create project: %w", err)
	}

	return nil
}

func (repository *Repository) UpdateProjectTitle(
	projectID uint,
	title string,
) (*model.Project, error) {
	result := repository.db.
		Model(&model.Project{}).
		Where("id = ?", projectID).
		Update("title", title)

	if result.Error != nil {
		return nil, fmt.Errorf("update project: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return repository.FindProject(projectID)
}

func (repository *Repository) DeleteProject(
	projectID uint,
) error {
	result := repository.db.Delete(
		&model.Project{},
		projectID,
	)

	if result.Error != nil {
		return fmt.Errorf("delete project: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Todos

func (repository *Repository) FindTodo(
	todoID uint,
) (*model.Todo, error) {
	var todo model.Todo

	err := repository.db.
		Preload("Subtasks", func(query *gorm.DB) *gorm.DB {
			return query.Order("id ASC")
		}).
		First(&todo, todoID).
		Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (repository *Repository) CreateTodo(
	todo *model.Todo,
) error {
	if err := repository.db.Create(todo).Error; err != nil {
		return fmt.Errorf("create todo: %w", err)
	}

	return nil
}

func (repository *Repository) UpdateTodo(
	todoID uint,
	updates map[string]any,
) (*model.Todo, error) {
	result := repository.db.
		Model(&model.Todo{}).
		Where("id = ?", todoID).
		Updates(updates)

	if result.Error != nil {
		return nil, fmt.Errorf("update todo: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return repository.FindTodo(todoID)
}

func (repository *Repository) DeleteTodo(
	todoID uint,
) error {
	result := repository.db.Delete(
		&model.Todo{},
		todoID,
	)

	if result.Error != nil {
		return fmt.Errorf("delete todo: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Subtasks

func (repository *Repository) FindSubtask(
	subtaskID uint,
) (*model.Subtask, error) {
	var subtask model.Subtask

	if err := repository.db.
		First(&subtask, subtaskID).
		Error; err != nil {
		return nil, err
	}

	return &subtask, nil
}

func (repository *Repository) CreateSubtask(
	subtask *model.Subtask,
) error {
	if err := repository.db.Create(subtask).Error; err != nil {
		return fmt.Errorf("create subtask: %w", err)
	}

	return nil
}

func (repository *Repository) UpdateSubtask(
	subtaskID uint,
	updates map[string]any,
) (*model.Subtask, error) {
	result := repository.db.
		Model(&model.Subtask{}).
		Where("id = ?", subtaskID).
		Updates(updates)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"update subtask: %w",
			result.Error,
		)
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return repository.FindSubtask(subtaskID)
}

func (repository *Repository) DeleteSubtask(
	subtaskID uint,
) error {
	result := repository.db.Delete(
		&model.Subtask{},
		subtaskID,
	)

	if result.Error != nil {
		return fmt.Errorf(
			"delete subtask: %w",
			result.Error,
		)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
