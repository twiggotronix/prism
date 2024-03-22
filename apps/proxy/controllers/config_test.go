package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"

	config "prism/proxy/config"
	mocks "prism/proxy/mocks"
	models "prism/proxy/models"
)

var (
	cc          ConfigController
	redisClient *redis.Client
)
var (
	redisTypes       = [2]string{"int", "string"}
	redisKeys        = [2]string{"api-proxy-delay", "api-proxy-path-prefix"}
	redisValues      = [2]string{"42", "/fake/"}
	redisErrorValues = [2]string{"delay error", "path error"}
)

type ConfigControllerTestSuite struct {
	suite.Suite
}

func TestConfigControllerTestSuite(t *testing.T) {
	suite.Run(t, &ConfigControllerTestSuite{})
}

func (ccs *ConfigControllerTestSuite) SetupTest() {
	appConfig := mocks.NewMockAppConfig(ccs.T())

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	cc = NewConfigController(appConfig)
}

func TestConfigGet(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	mock := redismock.NewNiceMock(redisClient)
	for i, key := range redisKeys {
		mock.On("Get", key).Return(redis.NewStringResult(redisValues[i], nil))
	}

	conf := config.NewAppConfig(mock)
	configController := NewConfigController(conf)
	configController.Get(context)

	var config models.Config
	json.Unmarshal(response.Body.Bytes(), &config)
	expectedDelay, _ := strconv.ParseInt(redisValues[0], 10, 64)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, int(expectedDelay), config.Delay)
	assert.Equal(t, redisValues[1], config.ProxyPrefix)
}

func TestConfigGetRedisDefaultValues(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	c := models.Config{
		Delay: 10,
	}
	jsonbytes, _ := json.Marshal(c)
	reader := bytes.NewReader(jsonbytes)

	context.Request, _ = http.NewRequest(http.MethodPost, "/", reader)
	mock := redismock.NewNiceMock(redisClient)
	for i, key := range redisKeys {
		mock.On("Get", key).Return(redis.NewStringResult("", errors.New(redisErrorValues[i])))
	}

	conf := config.NewAppConfig(mock)
	configController := NewConfigController(conf)
	configController.Get(context)

	var configResult models.Config
	json.Unmarshal(response.Body.Bytes(), &configResult)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 0, configResult.Delay)
	assert.Equal(t, "/proxy/", configResult.ProxyPrefix)
}

func TestConfigGetRedisDelayError(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	c := models.Config{
		Delay: 10,
	}
	jsonbytes, _ := json.Marshal(c)
	reader := bytes.NewReader(jsonbytes)

	context.Request, _ = http.NewRequest(http.MethodPost, "/", reader)
	mock := redismock.NewNiceMock(redisClient)
	for i, key := range redisKeys {
		if key == "api-proxy-delay" {
			mock.On("Get", key).Return(redis.NewStringResult("no work", nil))
		} else {
			mock.On("Get", key).Return(redis.NewStringResult(redisValues[i], nil))
		}
	}

	conf := config.NewAppConfig(mock)
	configController := NewConfigController(conf)
	configController.Get(context)

	var configResult models.Config
	json.Unmarshal(response.Body.Bytes(), &configResult)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 0, configResult.Delay)
	assert.Equal(t, "/fake/", configResult.ProxyPrefix)
}

func TestConfigSet(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	c := models.Config{
		Delay:       42,
		ProxyPrefix: "/fake/",
	}
	jsonbytes, _ := json.Marshal(c)
	reader := bytes.NewReader(jsonbytes)

	context.Request, _ = http.NewRequest(http.MethodPost, "/", reader)
	mock := redismock.NewNiceMock(redisClient)

	for i, key := range redisKeys {
		if redisTypes[i] == "int" {
			expectedInt, _ := strconv.ParseInt(redisValues[i], 10, 64)
			mock.On("Set", key, int(expectedInt), time.Duration(0)).Return(redis.NewStatusCmd(expectedInt))
		} else {
			mock.On("Set", key, redisValues[i], time.Duration(0)).Return(redis.NewStatusCmd(redisValues[i]))
		}
	}

	conf := config.NewAppConfig(mock)
	configController := NewConfigController(conf)
	configController.Set(context)

	mock.AssertNumberOfCalls(t, "Set", len(redisKeys))
	for _, call := range mock.Calls {
		log.Printf("%v", call.Arguments)
		for i, key := range redisKeys {
			if key == call.Arguments[0] {
				if redisTypes[i] == "int" {
					expectedInt, _ := strconv.ParseInt(redisValues[i], 10, 64)
					assert.IsEqual(call.Arguments[1], expectedInt)
				} else {
					assert.IsEqual(call.Arguments[1], redisValues[i])
				}
				assert.IsEqual(call.Arguments[2], time.Duration(0))
			}
		}
	}

}
