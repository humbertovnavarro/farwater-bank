package minecraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type MinecraftAccount struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func FetchUsername(minecraftUUID string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/user/profile/%s", minecraftUUID)
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

func FetchUUID(minecraftUsername string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", minecraftUsername)
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
