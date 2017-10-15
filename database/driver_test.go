package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"portchecker/db_models"
)

var (
	email = "email@localhost"
)

func TestInsert(t *testing.T) {

	db := NewDBDriver()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&db_models.CheckAgent{})

	// Create
	db.Create(&db_models.CheckAgent{Hostname: "local"})

	// Read
	var agent db_models.CheckAgent
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
