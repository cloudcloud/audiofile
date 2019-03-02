package audiofile

// Album gives structure around a single album.
type Album struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Artists     []Artist `json:"artists"`
	Length      int      `json:"length"`
	ReleaseYear int      `json:"release_year"`
	Songs       []Song   `json:"songs"`
	TrackCount  int      `json:"track_count"`
	URLText     string   `json:"url_text"`
}
