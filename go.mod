module github.com/mchmarny/dapr-http-cron-handler-template

go 1.14

replace github.com/dapr/go-sdk => github.com/mchmarny/go-sdk v0.8.15

require (
	github.com/dapr/go-sdk v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/assert/v2 v2.0.1
	github.com/sirupsen/logrus v1.6.0
)
