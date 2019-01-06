package audiofile

// Artist is a house for all data related to a specific Artist.
type Artist struct {
	Name string `json:"name"`
	//Albums []Album `json:"albums"`
	Status string `json:"status"`
	ID     string `json:"id"`
}
