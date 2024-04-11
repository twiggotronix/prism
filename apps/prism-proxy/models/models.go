package models

type HttpMethod string

const (
	HTTP_GET    HttpMethod = "GET"
	HTTP_POST              = "POST"
	HTTP_PUT               = "PUT"
	HTTP_DELETE            = "DELETE"
)

func (e HttpMethod) String() string {
	switch e {
	case HTTP_GET:
		return "GET"
	case HTTP_POST:
		return "POST"
	case HTTP_PUT:
		return "POST"
	case HTTP_DELETE:
		return "POST"
	default:
		return e.String()
	}
}

type Proxy struct {
	ID     uint       `gorm:"primaryKey" json:"id,omitempty"`
	Name   string     `json:"name" binding:"required"`
	Path   string     `json:"path" binding:"required"`
	Method HttpMethod `json:"method" gorm:"type:enum('GET', 'POST', 'PUT', 'DELETE')" default:"GET" binding:"required"`
	Source string     `json:"source" binding:"required"`
}

type Config struct {
	Delay       int    `json:"delay"`
	ProxyPrefix string `json:"proxyPrefix"`
}

type DbConfig struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"databaseName"`
}
