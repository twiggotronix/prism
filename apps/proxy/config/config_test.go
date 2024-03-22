package config

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"prism/proxy/models"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	redisClient *redis.Client
)
var (
	redisKeys   = [2]string{"api-proxy-delay", "api-proxy-path-prefix"}
	redisValues = [2]string{"42", "/fake/"}
	//redisErrorValues = [2]string{"delay error", "path error"}
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

func TestConfigGetIfNoData(t *testing.T) {
	appConf := NewAppConfig(redisClient)
	config := appConf.Get()

	assert.Equal(t, 0, config.Delay)
}

func TestConfigGet(t *testing.T) {
	mock := redismock.NewNiceMock(redisClient)
	for i, key := range redisKeys {
		mock.On("Get", key).Return(redis.NewStringResult(redisValues[i], nil))
	}

	appConf := NewAppConfig(mock)
	config := appConf.Get()

	assert.Equal(t, 42, config.Delay)
	assert.Equal(t, "/fake/", config.ProxyPrefix)
}

func TestConfigSet(t *testing.T) {
	mock := redismock.NewNiceMock(redisClient)
	for i, key := range redisKeys {
		if key == "api-proxy-delay" {
			delay, _ := strconv.ParseInt(redisValues[i], 10, 64)
			mock.On("Set", key, int(delay), time.Duration(0)).Return(redis.NewStatusCmd(redisValues[i]))
		} else {
			mock.On("Set", key, redisValues[i], time.Duration(0)).Return(redis.NewStatusCmd(redisValues[i]))
		}
	}

	appConf := NewAppConfig(mock)
	conf := models.Config{
		Delay:       42,
		ProxyPrefix: "/fake/",
	}
	err := appConf.Set(conf)

	assert.Nil(t, err)
}
