package audiofile

// File provides structure to encompass all entity attributes related
// to a specific file from the file system.
type File struct {
	Filename string

	Artists []Artist
	Album   Album
	Song    Song
}
