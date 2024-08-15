package integration

import (
	"bytes"
	"encoding/json"
	"gin-svc/internal/integration/startup"
	"gin-svc/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 权限相关的集成测试

// RoleTestSuite 是 Role 的集成测试的 Test Suite
type RoleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
	cli    redis.Cmdable
}

// 预期的返回值 之所以用泛型，是防止any在反序列化的时候出问题
type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// 实现 SetupAllSuite 接口
func (s *RoleTestSuite) SetupSuite() {
	// 在所有测试开始之前，做一些事情 比如初始化gin的engine 初始化MySQL，redis等，用于测试
	s.db = startup.InitDB()
	s.cli = startup.InitRedis()

	server := gin.Default()

	s.server = server

}

// TestCreate 测试创建角色
func (s *RoleTestSuite) TestCreate() {
	t := s.T()
	testCase := []struct {
		name string

		before func(t *testing.T) // 集成测试之前准备数据
		after  func(t *testing.T) // 集成测试之后清理数据

		form types.CreateRoleReq // 预期的输入

		wantCode int // 预期的HTTP响应码

		wantRes Result[int64] // 预期的返回值 因为角色创建成功会返回角色的ID

		wantErr error
	}{
		{},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// 分层3个部分，1 构造请求，2. 执行逻辑，3.校验结果
			tc.before(t)
			reqBody, err := json.Marshal(tc.form)
			assert.NoError(t, err)                                                          // 断言请求体的序列化是否成功
			req, err := http.NewRequest(http.MethodPost, "/role", bytes.NewReader(reqBody)) // 模拟构造http的post请求
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json") // 设置请求头 数据格式为json
			resp := httptest.NewRecorder()                     // 模拟http响应
			s.server.ServeHTTP(resp, req)                      // 使用测试套件里的Server对象
			assert.Equal(t, tc.wantCode, resp.Code)            // 断言响应码是否符合预期
			if resp.Code != 200 {
				return
			}
			require.NoError(t, err)
			var webRes Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			assert.Equal(t, tc.wantRes, webRes)
			tc.after(t)
		})
	}
}

func TestRole(t *testing.T) {
	suite.Run(t, &RoleTestSuite{})
}
