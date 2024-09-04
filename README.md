# rss

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/BurntSushi/toml
go get -u gopkg.in/natefinch/lumberjack.v2
go get -u github.com/sirupsen/logrus    
go get -u github.com/gin-contrib/cors
go get -u github.com/go-ozzo/ozzo-validation/v4
go get -u gorm.io/gorm@latest
go get -u github.com/cosmtrek/air@latest
go get -u github.com/gin-gonic/gin@latest
go get -u github.com/go-gormigrate/gormigrate/v2@latest
go get -u github.com/go-ozzo/ozzo-validation/v4@latest
go get -u github.com/golang-jwt/jwt/v4@latest
go get -u github.com/google/wire@latest
go get -u gorm.io/driver/postgres@latest
go get -u github.com/BurntSushi/toml@latest
go get -u github.com/gin-contrib/cors@latest
go get -u gorm.io/driver/mysql



❯ wire ./pkg/wire/wire.go
❯ air -c ./pkg/air/.air.toml
