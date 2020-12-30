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

	adam := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Adam",
		Email:       "adam@adam.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Adam Address",
		UserRole:    userRole,
	}

	jensen := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Jensen",
		Email:       "jensen@jensen.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Jensen Address",
		UserRole:    userRole,
	}

	hakurei := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Hakurei",
		Email:       "hakurei@hakurei.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Hakurei Address",
		UserRole:    userRole,
	}

	reimu := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Reimu",
		Email:       "reimu@reimu.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Reimu Address",
		UserRole:    userRole,
	}

	marisa := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Marisa",
		Email:       "marisa@marisa.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Marisa Address",
		UserRole:    userRole,
	}

	kirisame := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Kirisame",
		Email:       "kirisame@kirisame.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Kirisame Address",
		UserRole:    userRole,
	}

	dante := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Dante",
		Email:       "dante@dante.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Dante Address",
		UserRole:    userRole,
	}

	vergil := &model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        "Vergil",
		Email:       "vergil@vergil.com",
		Password:    string(userHash),
		DateOfBirth: userDOB,
		Gender:      "Male",
		Address:     "Vergil Address",
		UserRole:    userRole,
	}

	db.Create(adminUser)
	db.Create(regularUser)
	db.Create(adam)
	db.Create(jensen)
	db.Create(hakurei)
	db.Create(reimu)
	db.Create(marisa)
	db.Create(kirisame)
	db.Create(dante)
	db.Create(vergil)
}
