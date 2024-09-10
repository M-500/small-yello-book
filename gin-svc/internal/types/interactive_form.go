package types

import "gin-svc/internal/domain"

type InteractiveForm struct {
	ResourceId   string            `json:"resource_id" binding:"required"`
	Action       domain.IntrAction `json:"action" binding:"required"`
	ResourceType string            `json:"resource_type" binding:"required"`
	OwnerId      string            `json:"owner_id" binding:"required"`
}
