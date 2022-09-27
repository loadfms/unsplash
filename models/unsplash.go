package models

import "time"

type UnsplashResponse struct {
	Results []struct {
		ID             string      `json:"id"`
		CreatedAt      time.Time   `json:"created_at"`
		UpdatedAt      time.Time   `json:"updated_at"`
		PromotedAt     interface{} `json:"promoted_at"`
		Width          int         `json:"width"`
		Height         int         `json:"height"`
		Color          string      `json:"color"`
		BlurHash       string      `json:"blur_hash"`
		Description    interface{} `json:"description"`
		AltDescription interface{} `json:"alt_description"`
		Urls           struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
			SmallS3 string `json:"small_s3"`
		} `json:"urls"`
		Links struct {
			Self             string `json:"self"`
			HTML             string `json:"html"`
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
		} `json:"links"`
		Likes                  int           `json:"likes"`
		LikedByUser            bool          `json:"liked_by_user"`
		CurrentUserCollections []interface{} `json:"current_user_collections"`
		Sponsorship            struct {
			ImpressionUrls []string `json:"impression_urls"`
			Tagline        string   `json:"tagline"`
			TaglineURL     string   `json:"tagline_url"`
			Sponsor        struct {
				ID              string      `json:"id"`
				UpdatedAt       time.Time   `json:"updated_at"`
				Username        string      `json:"username"`
				Name            string      `json:"name"`
				FirstName       string      `json:"first_name"`
				LastName        interface{} `json:"last_name"`
				TwitterUsername interface{} `json:"twitter_username"`
				PortfolioURL    string      `json:"portfolio_url"`
				Bio             string      `json:"bio"`
				Location        interface{} `json:"location"`
				Links           struct {
					Self      string `json:"self"`
					HTML      string `json:"html"`
					Photos    string `json:"photos"`
					Likes     string `json:"likes"`
					Portfolio string `json:"portfolio"`
					Following string `json:"following"`
					Followers string `json:"followers"`
				} `json:"links"`
				ProfileImage struct {
					Small  string `json:"small"`
					Medium string `json:"medium"`
					Large  string `json:"large"`
				} `json:"profile_image"`
				InstagramUsername string `json:"instagram_username"`
				TotalCollections  int    `json:"total_collections"`
				TotalLikes        int    `json:"total_likes"`
				TotalPhotos       int    `json:"total_photos"`
				AcceptedTos       bool   `json:"accepted_tos"`
				ForHire           bool   `json:"for_hire"`
				Social            struct {
					InstagramUsername string      `json:"instagram_username"`
					PortfolioURL      string      `json:"portfolio_url"`
					TwitterUsername   interface{} `json:"twitter_username"`
					PaypalEmail       interface{} `json:"paypal_email"`
				} `json:"social"`
			} `json:"sponsor"`
		} `json:"sponsorship"`
		TopicSubmissions struct {
		} `json:"topic_submissions,omitempty"`
		User struct {
			ID              string      `json:"id"`
			UpdatedAt       time.Time   `json:"updated_at"`
			Username        string      `json:"username"`
			Name            string      `json:"name"`
			FirstName       string      `json:"first_name"`
			LastName        interface{} `json:"last_name"`
			TwitterUsername interface{} `json:"twitter_username"`
			PortfolioURL    string      `json:"portfolio_url"`
			Bio             string      `json:"bio"`
			Location        interface{} `json:"location"`
			Links           struct {
				Self      string `json:"self"`
				HTML      string `json:"html"`
				Photos    string `json:"photos"`
				Likes     string `json:"likes"`
				Portfolio string `json:"portfolio"`
				Following string `json:"following"`
				Followers string `json:"followers"`
			} `json:"links"`
			ProfileImage struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Large  string `json:"large"`
			} `json:"profile_image"`
			InstagramUsername string `json:"instagram_username"`
			TotalCollections  int    `json:"total_collections"`
			TotalLikes        int    `json:"total_likes"`
			TotalPhotos       int    `json:"total_photos"`
			AcceptedTos       bool   `json:"accepted_tos"`
			ForHire           bool   `json:"for_hire"`
			Social            struct {
				InstagramUsername string      `json:"instagram_username"`
				PortfolioURL      string      `json:"portfolio_url"`
				TwitterUsername   interface{} `json:"twitter_username"`
				PaypalEmail       interface{} `json:"paypal_email"`
			} `json:"social"`
		} `json:"user"`
		TopicSubmissions0 struct {
			Wallpapers struct {
				Status string `json:"status"`
			} `json:"wallpapers"`
			Film struct {
				Status string `json:"status"`
			} `json:"film"`
		} `json:"topic_submissions,omitempty"`
		TopicSubmissions1 struct {
			Fashion struct {
				Status string `json:"status"`
			} `json:"fashion"`
			People struct {
				Status string `json:"status"`
			} `json:"people"`
			Film struct {
				Status string `json:"status"`
			} `json:"film"`
		} `json:"topic_submissions,omitempty"`
	} `json:"results"`
}
