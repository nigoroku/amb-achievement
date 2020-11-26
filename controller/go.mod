module github.com/nigoroku/amb-achievement/controller

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	local.packages/models v0.0.0-00010101000000-000000000000
	local.packages/generated v0.0.0-00010101000000-000000000000
	local.packages/service v0.0.0-00010101000000-000000000000
)

replace local.packages/models => ./../models

replace local.packages/generated => ./../models/generated

replace local.packages/service => ./../service
