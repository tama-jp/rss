package db

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	entities "github.com/tama-jp/rss/internal/domain/entities"
	"github.com/tama-jp/rss/internal/frameworks/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DataBase struct {
	Connection            *gorm.DB
	TransactionConnection *gorm.DB
	Config                *config.Config
	Count                 int
}

func NewDB(config *config.Config) (*DataBase, error) {
	var db *gorm.DB
	var err error

	dbConfig := gorm.Config{}

	if config.DB.Debug == 1 {
		dbConfig.Logger = logger.Default.LogMode(logger.Info)

	}

	fmt.Println("config.DB.DataBase:" + config.DB.DataBase)

	switch config.DB.DataBase {

	case "sqlite":
		fmt.Println("sqlite:")

		db, err = gorm.Open(sqlite.Open(config.DB.FileName), &dbConfig)
		if err != nil {
			panic("failed to connect database")
		}

	case "postgresql":
		fmt.Println("postgresql:")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
			config.DB.Host, config.DB.User, config.DB.Password, config.DB.Name, config.DB.Port)

		fmt.Println("dsn:" + dsn)
		db, err = gorm.Open(postgres.Open(dsn), &dbConfig)

	//case "sqlserver":
	//	// 未検証
	//	fmt.Println("SQL Server:")
	//
	//	dsn := fmt.Sprintf("sqlserver://gorm:LoremIpsum86@%s:%s?database=%s",
	//		config.DB.Host, config.DB.Port, config.DB.Name)
	//
	//	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	case "mysql":
		// 未検証
		fmt.Println("mySQL:")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

		fmt.Println("dsn:" + dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	if err != nil {
		panic(fmt.Errorf("fatal error: %s\n", err))
	}

	var database = &DataBase{
		Connection: db,
		Config:     config,
		Count:      100,
	}

	return database, nil
}

func (db *DataBase) Migration() *gorm.DB {
	// https://github.com/go-gormigrate/gormigrate
	// github.com/go-gormigrate/gormigrate/v2

	fmt.Println("Migration:", "Start")
	m := gormigrate.New(db.Connection, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "0",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.UserRole{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_roles")
			},
		},
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				name := "権限なし"
				roleName := "no_authority"
				bitCode := entities.UserRoleNoAuthority
				model := entities.Model{ID: 1}
				userRole := entities.UserRole{
					Model:    model,
					Name:     name,
					RoleName: roleName,
					BitCode:  bitCode,
				}

				return tx.Create(&userRole).Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				name := "一般ユーザ"
				roleName := "default"
				bitCode := entities.UserRoleDefault
				model := entities.Model{ID: 2}
				userRole := entities.UserRole{
					Model:    model,
					Name:     name,
					RoleName: roleName,
					BitCode:  bitCode,
				}

				return tx.Create(&userRole).Error
			},
		},
		{
			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				name := "管理者"
				roleName := "super_user"
				bitCode := entities.UserRoleSuperUser
				model := entities.Model{ID: 3}

				userRole := entities.UserRole{
					Model:    model,
					Name:     name,
					RoleName: roleName,
					BitCode:  bitCode,
				}

				return tx.Create(&userRole).Error
			},
		},
		{
			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "7",
			Migrate: func(tx *gorm.DB) error {

				userName := db.Config.Admin.UserName
				password := db.Config.Admin.Password
				lastName := db.Config.Admin.LastName
				firstName := db.Config.Admin.FirstName

				user := entities.User{
					UserName:    userName,
					LastName:    lastName,
					FirstName:   firstName,
					Password:    GetSHA512String(password),
					RoleBitCode: entities.UserRoleDefault + entities.UserRoleSuperUser, // スーパーユーザ
				}

				return tx.Create(&user).Error
			},
		},
		{
			ID: "8",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.UserAuth{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("user_auths")
			},
		},
		{
			ID: "9",
			Migrate: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&entities.User{}, "EmployeeNumber") == false {
					return tx.Migrator().AddColumn(&entities.User{}, "EmployeeNumber")
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn(&entities.User{}, "EmployeeNumber")
			},
		},
		{
			ID: "10",
			Migrate: func(tx *gorm.DB) error {
				if tx.Migrator().HasIndex(&entities.User{}, "idx_user_employee_number") == false {
					return tx.Migrator().CreateIndex(&entities.User{}, "idx_user_employee_number")
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropIndex(&entities.User{}, "idx_user_employee_number")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migration did run successfully")

	return db.Connection.Begin()
}

func (db *DataBase) UserRolesSeqIDReset() {
	db.Connection.Exec("select setval('user_roles_id_seq',(select max(id) from user_roles))")
}

// Begin begins a transaction
func (db *DataBase) Begin() {
	db.TransactionConnection = db.Connection.Begin()
}

func (db *DataBase) Rollback() {
	db.TransactionConnection.Rollback()
}

func (db *DataBase) Commit() {
	db.TransactionConnection.Commit()
}

func (db *DataBase) Connect() *gorm.DB {
	return db.Connection
}

func GetSHA512String(s string) string {
	r := sha512.Sum512([]byte(s))
	h := hex.EncodeToString(r[:])

	return h
}
