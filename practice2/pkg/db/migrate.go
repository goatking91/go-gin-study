package db

import "github.com/goatking91/go-gin-study/practice2/pkg/logger"

func init() {
	currentDataBase := DB.Migrator().CurrentDatabase()

	logger.S.Infof("Running auto migrate database name:%s", currentDataBase)

	err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()

	if err != nil {
		logger.S.Errorf("Fail auto migrate database name:%s. %v", currentDataBase, err)
	} else {
		logger.S.Infof("Finished auto migrate database name:%s", currentDataBase)
	}
}
