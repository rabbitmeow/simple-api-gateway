package controllers

import (
	"github.com/spf13/viper"
)

// ServiceController is
type ServiceController struct{}

// Match fungsi ini digunakan untuk mencocokan service request dgn existing
func (w *ServiceController) Match(service string) string {
	serviceRequest := service
	target := ""
	serviceList := viper.Get("service").([]interface{})
	for _, item := range serviceList {
		if item.(map[string]interface{})["name"] == serviceRequest {
			target = item.(map[string]interface{})["host"].(string) + ":" + item.(map[string]interface{})["port"].(string)
		}
	}
	return target
}
