module github.com/kzpolicy/user

go 1.14

require (
	github.com/appleboy/gin-jwt v2.5.0+incompatible
	github.com/appleboy/gin-jwt/v2 v2.6.4 // indirect
	github.com/clipperhouse/fsnotify v1.1.0 // indirect
	github.com/clipperhouse/gen v4.1.1+incompatible // indirect
	github.com/clipperhouse/slice v0.0.0-20200107170738-a74fc3888fd9 // indirect
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/kzpolicy/amb-todo/models v0.0.0-20200802145439-572500e27288 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/fasttemplate v1.1.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.7.1+incompatible // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/dgrijalva/jwt-go.v3 v3.2.0 // indirect
	local.packages/controller v0.0.0-00010101000000-000000000000
	local.packages/middleware v0.0.0-00010101000000-000000000000
	local.packages/models v0.0.0-00010101000000-000000000000
	local.packages/service v0.0.0-00010101000000-000000000000 // indirect
)

replace local.packages/models => ./models

replace local.packages/middleware => ./middleware

replace local.packages/controller => ./controller

replace local.packages/service => ./service
