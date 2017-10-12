package database

import (
	"portchecker/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	email = "email@localhost"
)

func TestInsert(t *testing.T) {

	db := NewDBDriver()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.CheckAgent{})

	// Create
	db.Create(&models.CheckAgent{Hostname: "local"})

	// Read
	var agent models.CheckAgent
	res := db.First(&agent, "Hostname = ?", "local")

	assert.NotNil(t, res)

	// Delete
	//db.Delete(&agent)

}

func TestAutoMigrateDB(t *testing.T) {
	db := NewDBDriver()
	defer db.Close()

	AutoMigrateDB(db)
}

// Force to remove test user
func TestTearDown(t *testing.T) {
	// remove user in order to not change initial state
	db := NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	//res := db.Exec("delete from checkagent where email = ?", email).Error
	//assert.Nil(t, res)
}
