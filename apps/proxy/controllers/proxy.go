package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	config "prism/proxy/config"
	models "prism/proxy/models"
	repository "prism/proxy/repository"
	utils "prism/proxy/utils"

	"github.com/gin-gonic/gin"
)

type ProxyController interface {
	Delete(context *gin.Context)
	Get(context *gin.Context)
	GetAll(context *gin.Context)
	Save(context *gin.Context)
	InitProxies(context *gin.Engine)

	doCall(method string, url string, headers http.Header, body io.Reader) (*http.Response, error, error)
}

type proxyController struct {
	proxyRepository    repository.ProxyRepository
	appConfig          config.AppConfig
	appNetworkSettings config.AppNetworkSettings
}

func NewProxyController(proxyRepository repository.ProxyRepository, appConfig config.AppConfig, appNetworkSettings config.AppNetworkSettings) ProxyController {
	return &proxyController{
		proxyRepository,
		appConfig,
		appNetworkSettings,
	}
}

func (p *proxyController) Delete(context *gin.Context) {
	var proxyId = context.Param("id")
	err := p.proxyRepository.Delete(proxyId)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": false})
		return
	}

	context.Status(http.StatusNoContent)
}

func (p *proxyController) Get(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": false})
		return
	}
	proxy, err := p.proxyRepository.Get(id)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "status": false})
		return
	}
	context.JSON(http.StatusOK, &proxy)
}

func (p *proxyController) GetAll(context *gin.Context) {
	proxies := p.proxyRepository.GetAll()

	context.JSON(http.StatusOK, proxies)
}

func (p *proxyController) Save(context *gin.Context) {
	var newProxy models.Proxy
	if err := context.ShouldBindJSON(&newProxy); err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": false})
		return
	}

	err := p.proxyRepository.Save(&newProxy)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": false})
		return
	}

	context.JSON(http.StatusCreated, newProxy)
}

func (p *proxyController) InitProxies(router *gin.Engine) {
	proxies := p.proxyRepository.GetAll()

	log.Println("Registering proxies :")
	for _, proxy := range proxies {
		log.Println("[" + fmt.Sprint(models.HttpMethod(proxy.Method)) + "] " + proxy.Path)
	}
	log.Println("Registering " + fmt.Sprint(len(proxies)) + " proxies :")
	router.Any("/proxy/*snail", func(context *gin.Context) {
		log.Println("incoming call : " + context.Request.Method + " " + context.Request.URL.Path)
		proxyIndex := slices.IndexFunc(proxies, func(p models.Proxy) bool {
			pathUtils := utils.NewPathUtils("/proxy/" + strings.TrimLeft(p.Path, "/"))
			log.Printf("%s : %t => %s == %s\n", context.Request.URL.Path, pathUtils.UrlCorrespondsToPath(context.Request.URL.Path), string(p.Method), string(context.Request.Method))
			return pathUtils.UrlCorrespondsToPath(context.Request.URL.Path) && string(p.Method) == string(context.Request.Method)
		})
		if proxyIndex < 0 {
			log.Println("404 proxy not found")
			context.Status(http.StatusNotFound)
		} else {
			proxy := proxies[proxyIndex]
			log.Println(proxy.Name + ": [" + proxy.Method.String() + "] " + proxy.Path + " serving " + proxy.Source)
			switch proxy.Method {
			case models.HTTP_GET:
				p.proxyGETCall(context, proxy)
			case models.HTTP_POST:
				p.proxyPOSTCall(context, proxy)
			case models.HTTP_DELETE:
				p.proxyDELETECall(context, proxy)
			case models.HTTP_PUT:
				p.proxyPUTCall(context, proxy)
			}
		}
	})
}

func (p *proxyController) proxyGETCall(context *gin.Context, proxy models.Proxy) {
	timer := setTimer(p.appConfig)
	path := strings.Replace(context.Request.RequestURI, "/proxy/", "/", 1)
	sourceUrl := proxy.Source + path
	resp, requestErr, responseErr := p.doCall("GET", sourceUrl, context.Request.Header, context.Request.Body)
	if requestErr != nil {
		fmt.Println("Error creating PUT request:", requestErr)
		return
	}
	if responseErr != nil {
		log.Println(responseErr)
		context.Error(responseErr)
		context.String(http.StatusInternalServerError, responseErr.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	<-timer.C

	context.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (p *proxyController) proxyPOSTCall(context *gin.Context, proxy models.Proxy) {
	timer := setTimer(p.appConfig)
	resp, requestErr, responseErr := p.doCall("POST", proxy.Source, context.Request.Header, context.Request.Body)
	if requestErr != nil {
		fmt.Println("Error creating PUT request:", requestErr)
		return
	}
	if responseErr != nil {
		log.Println(responseErr)
		context.Error(responseErr)
		context.String(http.StatusInternalServerError, responseErr.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	<-timer.C

	context.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (p *proxyController) proxyPUTCall(context *gin.Context, proxy models.Proxy) {
	timer := setTimer(p.appConfig)

	resp, requestErr, responseErr := p.doCall("PUT", proxy.Source, context.Request.Header, context.Request.Body)
	if requestErr != nil {
		fmt.Println("Error creating PUT request:", requestErr)
		return
	}
	if responseErr != nil {
		log.Println(responseErr)
		context.Error(responseErr)
		context.String(http.StatusInternalServerError, responseErr.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	<-timer.C

	context.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (p *proxyController) proxyDELETECall(context *gin.Context, proxy models.Proxy) {
	timer := setTimer(p.appConfig)

	resp, requestErr, responseErr := p.doCall("DELETE", proxy.Source, context.Request.Header, nil)
	if requestErr != nil {
		fmt.Println("Error creating PUT request:", requestErr)
		return
	}
	if responseErr != nil {
		log.Println(responseErr)
		context.Error(responseErr)
		context.String(http.StatusInternalServerError, responseErr.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	<-timer.C

	context.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (p *proxyController) doCall(method string, url string, headers http.Header, body io.Reader) (*http.Response, error, error) {
	fmt.Printf("calling [%s] %s\n", method, url)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("Error creating %s request: %s", method, err)
		return nil, err, nil
	}

	p.addDefaultHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return resp, nil, nil
}

func (p *proxyController) addDefaultHeaders(req *http.Request) {
	var remoteAddr = req.RemoteAddr
	localIp := p.appNetworkSettings.Ip

	forwardedHeader := fmt.Sprintf("for=%s;proto=%s;by=%s", remoteAddr, req.Proto, localIp)
	req.Header.Add("Forwarded", forwardedHeader)
	var xForwardedForHeader = req.Header.Get("X-Forwarded-For")
	if xForwardedForHeader == "" {
		xForwardedForHeader = remoteAddr
	}
	xForwardedForHeader = fmt.Sprintf("%s, %s", xForwardedForHeader, localIp)
	req.Header.Add("X-Forwarded-For", xForwardedForHeader)
	req.Header.Add("X-Forwarded-Proto", req.Proto)
	req.Header.Add("X-Forwarded-Host", p.appNetworkSettings.Host)
	req.Header.Add("Via", "snail-proxy")
}

func getNbSecondsToWait(appConfig config.AppConfig) int {
	conf := appConfig.Get()
	return conf.Delay
}

func setTimer(appConfig config.AppConfig) *time.Timer {
	nbSecondsToWait := getNbSecondsToWait(appConfig)
	timer := time.NewTimer(time.Duration(nbSecondsToWait) * time.Second)
	fmt.Printf("timer set for %d seconds\n", nbSecondsToWait)
	return timer
}

/*
func getBody(resp *http.Response) string {
	encoding := resp.Header.Get("Content-Encoding")

	body, err := decompressBody(resp.Body, encoding)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return body
}

func decompressBody(r io.Reader, encoding string) (string, error) {
	var reader io.Reader
	var err error

	switch encoding {
	case "gzip":
		reader, err = gzip.NewReader(r)
		if err != nil {
			return "", err
		}
	case "deflate":
		reader = flate.NewReader(r)
	case "br":
		reader = brotli.NewReader(r)
	default:
		return "", fmt.Errorf("unsupported encoding: %s", encoding)
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
*/
