package repository

import (
	"errors"

	models "prism/proxy/models"

	"gorm.io/gorm"
)

type ProxyRepository interface {
	GetAll() []models.Proxy
	Get(id uint64) (*models.Proxy, error)
	Save(newProxy *models.Proxy) error
	Delete(proxyId string) error
}
type proxyRepository struct {
	Db *gorm.DB
}

func NewProxyRepository(Db *gorm.DB) ProxyRepository {
	return &proxyRepository{Db}
}

func (p *proxyRepository) GetAll() []models.Proxy {
	var proxies []models.Proxy
	p.Db.Find(&proxies)

	return proxies
}

func (p *proxyRepository) Get(id uint64) (*models.Proxy, error) {
	var proxy models.Proxy
	proxy.ID = uint(id)
	result := p.Db.First(&proxy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proxy, nil
}

func (p *proxyRepository) Delete(proxyId string) error {
	result := p.Db.Delete(&models.Proxy{}, proxyId)

	if result.RowsAffected == 0 {
		return errors.New(`Error deleting proxy with id ` + proxyId)
	}

	return nil
}

func (p *proxyRepository) Save(newProxy *models.Proxy) error {
	result := p.Db.Save(&newProxy)

	return result.Error
}
