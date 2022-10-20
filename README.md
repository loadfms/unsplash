# Unsplash wallpaper change
Every run a new wallpaper for you!

## How to use?

1. Create a [Unsplash](https://unsplash.com/join) account and get an [API Token](https://unsplash.com/oauth/applicationshttps://unsplash.com/developers)
2. Create a **config.yaml** file and follow the configuration sample bellow
3. Execute the unsplash file with `go run main.go`

## config.yaml sample
```
config:
  feh_path: "/usr/bin/feh"
  feh_command: "--bg-fill"
  wallpaper_path: "./wallpapers/unsplash.png"
  unsplash_query: "dark"
  unsplash_api_client_id: "YOUR-CLIENT-ACCESS-KEY-HERE"
```

## Requirements

- [go](https://golang.google.cn/)
- [feh](https://feh.finalrewind.org/)
- Unix-Based System