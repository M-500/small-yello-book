package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type NoteTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func TestNote(t *testing.T) {
	suite.Run(t, &NoteTestSuite{})
}
