package random

// type Links struct {
// 	Self      string `json:"self"`
// 	HTML      string `json:"html"`
// 	Photos    string `json:"photos"`
// 	Likes     string `json:"likes"`
// 	Portfolio string `json:"portfolio"`
// }
// type User struct {
// 	Id                string `json:"id"`
// 	UpdatedAt         string `json:"updated_at"`
// 	Username          string `json:"username"`
// 	Name              string `json:"name"`
// 	PortfolioUrl      string `json:"portfolio_url"`
// 	Bio               string `json:"bio"`
// 	Location          string `json:"location"`
// 	TotalLikes        string `json:"total_likes"`
// 	TotalPhotos       string `json:"total_photos"`
// 	TotalCollections  string `json:"total_collections"`
// 	InstagramUsername string `json:"instagram_username"`
// 	TwitterUsername   string `json:"twitter_username"`
// 	Links             Links
// }

//	type Collection struct {
//		ID              string `json:"id"`
//		Title           string `json:"title"`
//		PublishedAt     string `json:"published_at"`
//		LastCollectedAt string `json:"last_collected_at"`
//		UpdatedAt       string `json:"updated_at"`
//		CoverPhoto      Photo `json:"cover_photo"`
//		User            User
//	}
type Links struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type Photo struct {
	ID    string `json:"id"`
	Links Links  `json:"links"`
}
