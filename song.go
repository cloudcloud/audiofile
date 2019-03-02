package audiofile

// Song structures data for an individual song.
type Song struct {
	Artists     []Artist `json:"artists"`
	Album       Album    `json:"album"`
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	TrackNumber int      `json:"track_number"`
}
