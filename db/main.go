package database

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "gqlgen-todos/graph/model"
)

type dbConfig struct {
    host     string
    port     int
    user     string
    dbname   string
    password string
}

var config = dbConfig{"localhost", 5432, "postgres", "test", "root"}

func getDatabaseUrl() string {
    return fmt.Sprintf(
        "host=%s port=%d user=%s dbname=%s password=%s",
        config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
    db, err := gorm.Open("postgres","host='localhost' port=5432 user='saurav.sardana' dbname='test' password='root' sslmode=disable")
    return db, err
}

func RunMigrations(db *gorm.DB) {
    if !db.HasTable(&model.Question{}) {
        db.CreateTable(&model.Question{})
    }
    if !db.HasTable(&model.Choice{}) {
        db.CreateTable(&model.Choice{})
        db.Model(&model.Choice{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
    }
}