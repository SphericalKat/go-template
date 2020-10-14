package crud

import (
	"github.com/SphericalKat/go-template/db"
	"github.com/SphericalKat/go-template/models"
	"github.com/SphericalKat/go-template/schemas"
	"github.com/google/uuid"
)

// CreateUser creates a user in the database or returns an error
func CreateUser(objIn *schemas.UserCreate) (*models.User, error) {
	db := db.GetDB()
	user := &models.User{
		Username: objIn.Username,
		IsActive: objIn.IsActive,
	}

	_, err := db.Model(user).
		OnConflict("DO NOTHING").
		Returning("*").
		Insert()
	return user, err
}

// UpdateUser updates an existing user in the database or returns an error
func UpdateUser(objIn *schemas.UserUpdate) (*models.User, error) {
	db := db.GetDB()

	// convert string in json body to UUID
	id, err := uuid.Parse(objIn.ID)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Username: objIn.Username,
		IsActive: objIn.IsActive,
	}

	_, err = db.Model(user).Returning("*").WherePK().UpdateNotZero()

	return user, err
}

// GetUser fetches an existing user or returns an error
func GetUser(userID string) (*models.User, error) {
	db := db.GetDB()

	// convert string in body to UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user := &models.User{ID: id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes an existing user or returns an error
func DeleteUser(userID string) (*models.User, error) {
	db := db.GetDB()

	// convert string in body to UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user := &models.User{ID: id}
	_, err = db.Model(user).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	return user, nil
}
