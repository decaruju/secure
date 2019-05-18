package repository

import (
	"github.com/satori/go.uuid"
	"secure/model"
)

func CreateApikey(user *model.User) (*model.Apikey, error) {
	db, err := db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = DeleteAllApikeys(user)
	if err != nil {
		return nil, err
	}

	key, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	apiKey := &model.Apikey{
		Key:    key.String(),
		UserID: user.ID,
	}

	db.Create(apiKey)
	return apiKey, nil
}

func DeleteAllApikeys(user *model.User) error {
	db, err := db()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	print(user.ID)
	err = db.Where("user_id = ?", user.ID).Delete(model.Apikey{}).Error
	return err
}
