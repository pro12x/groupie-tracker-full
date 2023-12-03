package models

// Search : Search struct
type Search struct {
	Artists       []SearchArtist       `json:"searchArtist"`
	Members       []SearchMember       `json:"searchMember"`
	Locations     []SearchLocation     `json:"searchLocation"`
	FirstAlbums   []SearcFirstAlbum    `json:"searchFirstAlbums"`
	CreationDates []SearchCreationDate `json:"searchCreationDates"`
}

type SearchArtist struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type SearchMember struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type SearchLocation struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type SearcFirstAlbum struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type SearchCreationDate struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}
