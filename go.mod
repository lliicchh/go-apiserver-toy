module github.com/lliicchh/apiserver

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/lexkong/log v0.0.0-20180607165131-972f9cd951fc
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v2.19.12+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf // indirect
	github.com/willf/pad v0.0.0-20190207183901-eccfe5d84172
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	gopkg.in/go-playground/validator.v9 v9.29.1
)

replace github.com/lliicchh/apiserver => /apiserver/handler
