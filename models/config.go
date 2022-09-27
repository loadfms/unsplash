package models

type Configs struct {
	Config struct {
		FehPath             string `yaml:"feh_path"`
		FehCommand          string `yaml:"feh_command"`
		WallpaperPath       string `yaml:"wallpaper_path"`
		UnsplashQuery       string `yaml:"unsplash_query"`
		UnsplashAPIClientID string `yaml:"unsplash_api_client_id"`
	} `yaml:"config"`
}
