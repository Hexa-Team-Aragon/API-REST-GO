package modelo

type Cancion struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Album     string `json:"album"`
	Artist    string `json:"Artist"`
	Genre     string `json:"Genre"`
	Year      int64  `json:"Year"`
	Url_image string `json:"url_image"`
}
