module github.com/ahmetaniltokatli/golang_rest_api

go 1.17

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/robfig/cron v1.2.0
)

require (
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
)

replace github.com/ahmetaniltokatli/golang_rest_api/pkg/db/redis => ../redis
