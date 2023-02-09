module web-service-gin

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.2
	xnw.com/core v0.0.0-00010101000000-000000000000
)

replace xnw.com/core => ../core

replace xnw.com/utils => ../utils
