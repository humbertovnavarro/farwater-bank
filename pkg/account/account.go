package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	MinecraftUUID string
	Password      string
}

func Register(username string, password string, db *gorm.DB) (*Account, error) {
	uuid, err := FetchMinecraftUUID(username)
	if err != nil {
		return nil, err
	}
	account := &Account{
		MinecraftUUID: uuid,
		Password:      password,
	}
	if err := db.FirstOrCreate(account, "minecraft_uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	if err != nil {
		logrus.Error(err)
	}
	return account, nil
}

type MinecraftUUIDResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func FetchMinecraftUUID(username string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	uuid := &MinecraftUUIDResponse{}
	err = json.Unmarshal(respBytes, uuid)
	if err != nil {
		return "", err
	}
	return uuid.ID, nil
}
