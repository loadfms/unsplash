package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/loadfms/unsplash/models"
)

func Run(configs *models.Configs) (err error) {
	unsplash, err := getPicureUrl(configs)
	if err != nil {
		return err
	}

	err = saveFile(unsplash, configs)
	if err != nil {
		return err
	}

	err = setNewWallpaper(configs)

	return err
}

func saveFile(unsplash models.UnsplashResponse, configs *models.Configs) (err error) {
	randomItem := getRandomNumber(0, 9)
	resp, err := http.Get(unsplash.Results[randomItem].Links.Download)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(configs.Config.WallpaperPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func getPicureUrl(configs *models.Configs) (unsplash models.UnsplashResponse, err error) {
	var unsplashResponse models.UnsplashResponse
	randomPage := getRandomNumber(1, 999)
	url := fmt.Sprintf("https://api.unsplash.com/search/photos/?query=%s&client_id=%s&page=%d", configs.Config.UnsplashQuery, configs.Config.UnsplashAPIClientID, randomPage)
	resp, err := http.Get(url)

	if err != nil {
		return unsplashResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return unsplashResponse, err
	}

	err = json.Unmarshal(body, &unsplashResponse)
	if err != nil {
		return unsplashResponse, err
	}

	return unsplashResponse, err
}

func setNewWallpaper(configs *models.Configs) (err error) {
	cmd := exec.Command(configs.Config.FehPath, configs.Config.FehCommand, configs.Config.WallpaperPath)
	err = cmd.Run()
	return err
}

func getRandomNumber(min int, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}
