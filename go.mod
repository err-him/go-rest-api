module go-rest-api

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	go.uber.org/zap v1.13.0
	gopkg.in/natefinch/lumberjack.v1 v1.0.0-20140618183000-8ec9c6b748e0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/err-him/go-rest-api/config => /Users/himanshu/go-learning/go-rest-api/config

replace github.com/err-him/go-rest-api/logger => /Users/himanshu/go-learning/go-rest-api/logger

replace github.com/err-him/go-rest-api/src/controller => /Users/himanshu/go-learning/go-rest-api/src/controller

//replace github.com/err-him/go-rest-api/src/interfaces => /Users/himanshu/go-learning/go-rest-api/src/interfaces
replace github.com/err-him/go-rest-api/src/handler => /Users/himanshu/go-learning/go-rest-api/src/handler

replace github.com/err-him/go-rest-api/src/constant => /Users/himanshu/go-learning/go-rest-api/src/constant
