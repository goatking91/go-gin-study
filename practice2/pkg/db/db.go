package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"time"

	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
	"github.com/goatking91/go-gin-study/practice2/pkg/util"
)

type TestResult struct {
	Date time.Time
}

var (
	DB *gorm.DB
)

func InitDb() (ok bool) {
	ok = true
	env := &util.Env{EnvSource: &util.EnvGetter{}}
	logger.S.Info("Connecting Database Server ....")

	dbUser := env.GetString("DB_USER")
	dbPassword := env.GetString("DB_PASSWORD")
	dbName := env.GetString("DB_NAME")
	dbHost := env.GetString("DB_HOST")
	dbPort := env.GetInt("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	dsnLog := fmt.Sprintf("Database connection string: %s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, "********", dbHost, dbPort, dbName)

	logger.S.Info(dsnLog)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 glogger.Default.LogMode(glogger.Info),
	})

	if err != nil {
		logger.S.Errorf("Database connect fail. %v", err)
		ok = false
		return
	} else {
		logger.S.Info("Database connected")
	}

	// Set connection pool
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(env.GetInt("DB_MAX_IDLE_CONNECTS"))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(env.GetInt("DB_MAX_OPEN_CONNECTS"))

	// SetConnMaxLifetime sets the maximum amount of time config connection may be reused.

	sqlDB.SetConnMaxLifetime(env.GetDuration("DATABASE_MAX_OPEN_CONNECTS"))

	DB = db

	// test database
	testDatabase()
	return
}

func testDatabase() {
	logger.S.Info("Testing Database Server ...")

	var result TestResult
	err := DB.Raw("SELECT now(6) as date").Scan(&result).Error
	if err != nil {
		logger.S.Errorf("Database Query failed, %v", err)
	} else {
		logger.S.Debugf("Database Query datetime(now(6)) result=(%v)", result.Date)
	}

}

// GetDatabaseDateTime Database 의 현재시간을 쿼리해서 리턴
// SQL Builder https://gorm.io/docs/sql_builder.html
func GetDatabaseDateTime() (time.Time, error) {
	var result TestResult
	err := DB.Raw("SELECT now(6) as date").Scan(&result).Error
	if err != nil {
		logger.S.Errorf("Database Query failed. ", err)
	} else {
		logger.S.Debugf("Database Query datetime(now(6)) result=(%v)", result.Date)
	}
	return result.Date, err
}
