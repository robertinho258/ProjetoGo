package config

import (
	"os"

	"github.com/robertinho258/ProjetoGo/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitiliazeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db "

	//Checando para ver se existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("data file not found,creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Criando Database e Conectnado
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
