package account

import (
	"encoding/json"
	"errors"
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

func GetByUUID(uuid string, db *gorm.DB) (*Account, error) {
	account := &Account{}
	if err := db.First(account, "minecraft_uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func GetByID(id uint, db *gorm.DB) (*Account, error) {
	account := &Account{}
	if err := db.First(account, id).Error; err != nil {
		return nil, err
	}
	return account, nil
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

type MinecraftAccount struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func FetchMinecraftUsername(id string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/user/profile/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	mcAccount := &MinecraftAccount{}
	if err := json.Unmarshal(respBytes, mcAccount); err != nil {
		return "", err
	}
	return mcAccount.Name, nil
}

func FetchMinecraftUUID(username string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("could not fetch minecraft uuid")
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	mcAccount := &MinecraftAccount{}
	err = json.Unmarshal(respBytes, mcAccount)
	if err != nil {
		return "", err
	}
	return mcAccount.ID, nil
}
