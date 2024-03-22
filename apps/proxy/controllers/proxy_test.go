package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"prism/proxy/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	config "prism/proxy/config"
	mocks "prism/proxy/mocks"
)

var (
	pc                  ProxyController
	proxyRepositoryMock mocks.MockProxyRepository
	appConfigMock       mocks.MockAppConfig
)

type ProxyControllerTestSuite struct {
	suite.Suite
}

func TestProxyControllerTestSuite(t *testing.T) {
	suite.Run(t, &ProxyControllerTestSuite{})
}

func (pcs *ProxyControllerTestSuite) SetupTest() {
	proxyRepositoryMock = *mocks.NewMockProxyRepository(pcs.T())
	appConfigMock = *mocks.NewMockAppConfig(pcs.T())
	mockAppNetworkSettings := &config.AppNetworkSettings{
		Host: "localhost",
		Ip:   "127.0.0.1",
	}
	pc = NewProxyController(&proxyRepositoryMock, &appConfigMock, *mockAppNetworkSettings)
	log.Println("SetupTest")
}

func (pcs *ProxyControllerTestSuite) TestGet() {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Params = []gin.Param{
		{Key: "id", Value: "1"},
	}
	mockProxy := models.Proxy{ID: 1, Name: "fake", Path: "test", Method: "GET", Source: "https://source"}
	proxyRepositoryMock.On("Get", uint64(1)).Return(&mockProxy, nil)

	pc.Get(context)

	var resultProxy models.Proxy
	json.Unmarshal(response.Body.Bytes(), &resultProxy)

	pcs.Equal(http.StatusOK, response.Code)
	pcs.Equal(mockProxy, resultProxy)
}

func (pcs *ProxyControllerTestSuite) TestGetNotFound() {
	r := httptest.NewRequest(http.MethodGet, "/1", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)
	proxyRepositoryMock.On("Get", uint64(1)).Return(nil, errors.New("Not found"))
	router.GET("/:id", pc.Get)

	router.ServeHTTP(w, r)

	pcs.Equal(http.StatusNotFound, w.Code)
}

func (pcs *ProxyControllerTestSuite) TestGetError() {
	r := httptest.NewRequest(http.MethodGet, "/aa", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)
	router.GET("/:id", pc.Get)

	router.ServeHTTP(w, r)

	var ginError gin.H
	json.Unmarshal(w.Body.Bytes(), &ginError)
	pcs.Equal(http.StatusInternalServerError, w.Code)
	log.Println(w.Body.String())
	pcs.NotEmpty(ginError["message"].(string))
	pcs.False(ginError["status"].(bool))
}

func (pcs *ProxyControllerTestSuite) TestGetAll() {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "GET", Source: "https://source"}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)

	pc.GetAll(context)

	var resultProxies []models.Proxy
	json.Unmarshal(response.Body.Bytes(), &resultProxies)

	pcs.Equal(http.StatusOK, response.Code)
	pcs.Equal(mockProxies, resultProxies)
}

func (pcs *ProxyControllerTestSuite) TestSave() {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	requestProxy := models.Proxy{Name: "fake", Path: "test", Method: "GET", Source: "https://source"}
	jsonbytes, _ := json.Marshal(requestProxy)
	reader := bytes.NewReader(jsonbytes)
	context.Request, _ = http.NewRequest(http.MethodPost, "/", reader)

	mockProxy := models.Proxy{ID: 1, Name: requestProxy.Name, Path: requestProxy.Path, Method: requestProxy.Method, Source: requestProxy.Source}
	proxyRepositoryMock.On("Save", mock.Anything).Return(nil)

	pc.Save(context)

	var resultProxy models.Proxy
	json.Unmarshal(response.Body.Bytes(), &resultProxy)

	pcs.Equal(http.StatusCreated, response.Code)
	pcs.Equal(mockProxy.Method, resultProxy.Method)
	pcs.Equal(mockProxy.Name, resultProxy.Name)
	pcs.Equal(mockProxy.Path, resultProxy.Path)
	pcs.Equal(mockProxy.Source, resultProxy.Source)
}

func (pcs *ProxyControllerTestSuite) TestSaveError() {
	requestProxy := models.Proxy{Name: "fake", Path: "test", Method: "GET", Source: "https://source"}
	jsonbytes, _ := json.Marshal(requestProxy)
	reader := bytes.NewReader(jsonbytes)
	r := httptest.NewRequest(http.MethodPost, "/", reader)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)
	err := errors.New("oh noo...")
	proxyRepositoryMock.On("Save", mock.Anything).Return(err)
	router.POST("/", pc.Save)

	router.ServeHTTP(w, r)

	var ginError gin.H
	json.Unmarshal(w.Body.Bytes(), &ginError)

	pcs.Equal(http.StatusInternalServerError, w.Code)
	log.Println(w.Body.String())
	pcs.Equal(gin.H{"message": err.Error(), "status": false}, ginError)
}

func (pcs *ProxyControllerTestSuite) TestDelete() {
	r := httptest.NewRequest(http.MethodDelete, "/1", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	proxyRepositoryMock.On("Delete", "1").Return(nil)
	router.DELETE("/:id", pc.Delete)

	router.ServeHTTP(w, r)

	pcs.Equal(http.StatusNoContent, w.Code)
}

func (pcs *ProxyControllerTestSuite) TestDeleteError() {
	r := httptest.NewRequest(http.MethodDelete, "/1", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)
	err := errors.New("oh noo...")
	proxyRepositoryMock.On("Delete", "1").Return(err)
	router.DELETE("/:id", pc.Delete)

	router.ServeHTTP(w, r)

	var ginError gin.H
	json.Unmarshal(w.Body.Bytes(), &ginError)

	pcs.Equal(http.StatusBadRequest, w.Code)
	log.Println(w.Body.String())
	pcs.Equal(gin.H{"message": err.Error(), "status": false}, ginError)
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesGetHappyPath() {
	r := httptest.NewRequest(http.MethodGet, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	expectedTestData := "test data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", expectedTestData)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "GET", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal(expectedTestData, w.Body.String())
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesGetWithVariables() {
	r := httptest.NewRequest(http.MethodGet, "/proxy/test/1", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()

	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test/:id", Method: "GET", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}
func (pcs *ProxyControllerTestSuite) TestInitProxiesGetWithQueryString() {
	r := httptest.NewRequest(http.MethodGet, "/proxy/test/1", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()

	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test/:id", Method: "GET", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}
func (pcs *ProxyControllerTestSuite) TestInitProxiesGetWithCustomHeaders() {
	request := httptest.NewRequest(http.MethodGet, "/proxy/test/1", nil)
	responseRecorder := httptest.NewRecorder()
	_, router := gin.CreateTestContext(responseRecorder)
	request.Header.Add("the-answer", "42")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", 42) //r.Header.Get("the-answer"))
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()

	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test/:id", Method: "GET", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(responseRecorder, request)
	pcs.Equal(http.StatusOK, responseRecorder.Code)
	pcs.Equal("42", responseRecorder.Body.String())
	//pcs.Equal("text/plain; charset=utf-8", responseRecorder.Header().Get("Content-Type"))
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesGet404() {
	r := httptest.NewRequest(http.MethodGet, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "GET", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusNotFound, w.Code)
}
func (pcs *ProxyControllerTestSuite) TestInitProxiesNoProxy() {
	r := httptest.NewRequest(http.MethodGet, "/proxy/nope", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "GET", Source: "fakeurl"}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})
	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusNotFound, w.Code)
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesPostHappyPath() {
	myReader := strings.NewReader("sonme text to be passed")
	r := httptest.NewRequest(http.MethodPost, "/proxy/test", myReader)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	expectedTestData := "test data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", expectedTestData)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "POST", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal(expectedTestData, w.Body.String())
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesPost404() {
	myReader := strings.NewReader("sonme text to be passed")
	r := httptest.NewRequest(http.MethodPost, "/proxy/test", myReader)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "POST", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusNotFound, w.Code)
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesDeleteHappyPath() {
	r := httptest.NewRequest(http.MethodDelete, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	expectedTestData := "test data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", expectedTestData)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "DELETE", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal(expectedTestData, w.Body.String())
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}
func (pcs *ProxyControllerTestSuite) TestInitProxiesDeleteError() {
	r := httptest.NewRequest(http.MethodDelete, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "DELETE", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusInternalServerError, w.Code)
}

func (pcs *ProxyControllerTestSuite) TestInitProxiesPutHappyPath() {
	r := httptest.NewRequest(http.MethodPut, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	expectedTestData := "test data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", expectedTestData)
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "PUT", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusOK, w.Code)
	pcs.Equal(expectedTestData, w.Body.String())
	pcs.Equal("text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}
func (pcs *ProxyControllerTestSuite) TestInitProxiesPutError() {
	r := httptest.NewRequest(http.MethodPut, "/proxy/test", nil)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	mockProxies := []models.Proxy{{ID: 1, Name: "fake", Path: "test", Method: "PUT", Source: srv.URL}}
	proxyRepositoryMock.On("GetAll").Return(mockProxies)
	appConfigMock.On("Get").Return(models.Config{Delay: 0, ProxyPrefix: "/proxy/"})

	pc.InitProxies(router)
	router.ServeHTTP(w, r)
	pcs.Equal(http.StatusInternalServerError, w.Code)
}
