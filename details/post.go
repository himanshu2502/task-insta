package details

import "time"

type Post struct {
	ID       string
	Caption  string
	ImageURL string
	PostedTimestamp time.Duration
}
