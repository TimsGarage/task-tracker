package model

import "time"

type Project struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Todos []Todo `json:"todos" gorm:"constraint:OnDelete:CASCADE"`
}

type Todo struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	ProjectID uint       `json:"projectId" gorm:"not null;index"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	DueDate   *time.Time `json:"dueDate"`
	Priority  string     `json:"priority" gorm:"not null;default:medium"`
	Notes     string     `json:"notes"`
	Tags      StringList `json:"tags" gorm:"type:text;not null;default:'[]'"`
	Subtasks  []Subtask  `json:"subtasks" gorm:"constraint:OnDelete:CASCADE"`
}

type Subtask struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	TodoID uint   `json:"todoId" gorm:"not null;index"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

const (
	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)

func ValidPriority(priority string) bool {
	switch priority {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return true
	default:
		return false
	}
}
