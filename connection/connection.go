package connection

import (
	"fmt"
	config "goSqlRaw/configs"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connDBOnce sync.Once
	dbInstance *DB
)

type DB struct {
	*gorm.DB
}

func connectDB(config *config.DBConfig) error {
	// need to configure this
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.User, config.Password, config.DBName, config.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
		return err
	}
	fmt.Println("Database connected successfully.")
	dbInstance = &DB{conn}
	return nil
}

func ConnectDB(config *config.DBConfig) *DB {
	connDBOnce.Do(func() {
		err := connectDB(config)
		if err != nil {
			panic("failed to connect DB: " + err.Error())
		}
	})
	return dbInstance
}

func GetDBInstance() *DB {
	return dbInstance
}

// Migration at first deletes all the tables if exists and then create new tables
func (db *DB) Migration() error {

	// to avoid duplicate data entry during every run
	db.deleteTables()

	err := db.createDepartmentTable()
	if err != nil {
		log.Println("error creating department table", "error: ", err)
		return err
	}

	err = db.createTeachersTable()
	if err != nil {
		log.Println("error creating teacher table", "error: ", err)
		return err
	}

	err = db.createStudentsTable()
	if err != nil {
		log.Println("error creating student table", "error: ", err)
		return err
	}
	return nil
}

// InsertInitialDataIntoTable inserts intial given data into the tables
func (db *DB) InsertInitialDataIntoTable() error {
	err := db.initialInsertIntoDepartmentsTable()
	if err != nil {
		log.Println("error inserting batch data into departments table", "error: ", err)
		return err
	}

	err = db.initialInsertIntoTeachersTable()
	if err != nil {
		log.Println("error inserting batch data into teachers table", "error: ", err)
		return err
	}

	err = db.initialInsertIntoStudentsTable()
	if err != nil {
		log.Println("error inserting batch data into departments table", "error: ", err)
		return err
	}
	return nil
}

// deletes table if exists
func (db *DB) deleteTables() {
	db.Exec(`DROP TABLE STUDENTS`)
	db.Exec(`DROP TABLE TEACHERS`)
	db.Exec(`DROP TABLE DEPARTMENTS`)
}
