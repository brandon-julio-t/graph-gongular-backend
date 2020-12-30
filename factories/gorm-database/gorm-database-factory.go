package gorm_database

import (
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Factory struct{}

func (*Factory) Create() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	migrateDatabase(db)
	//seedDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	if err := db.AutoMigrate(
		new(model.UserRole),
		new(model.User),
		new(model.FileUpload),
	); err != nil {
		log.Fatal(err)
	}
}

func seedDatabase(db *gorm.DB) {
	adminRole := &model.UserRole{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Name: "Admin",
	}

	userRole := &model.UserRole{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Name: "User",
	}

	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	adminDOB, _ := time.Parse("2006-01-02", "1970-01-01")

	userHash, _ := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	userDOB, _ := time.Parse("2006-01-02", "1970-01-01")

	adminUser := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Admin",
		Email:       "admin@admin.com",
		Password:    string(adminHash),
		DateOfBirth: adminDOB,
		Gender:      "Male",
		Address:     "Admin Address",
		UserRole:    adminRole,
	}

	regularUser := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "User",
		Email:       "user@user.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "User Address",
		UserRole:    userRole,
	}

	db.Create(adminUser)
	db.Create(regularUser)
}
