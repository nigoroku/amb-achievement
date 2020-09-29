module github.com/kzpolicy/user/middleware

go 1.14

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/gin-gonic/gin v1.6.3
	go.uber.org/zap v1.15.0
	local.packages/models v0.0.0-00010101000000-000000000000
	local.packages/service v0.0.0-00010101000000-000000000000
)

replace local.packages/models => ./../models

replace local.packages/service => ./../service
