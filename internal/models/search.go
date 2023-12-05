package models

// Search : Search struct
type Search struct {
	Artists       []SearchArtist       `json:"artists"`
	Members       []SearchMember       `json:"members"`
	Locations     []SearchLocation     `json:"locations"`
	FirstAlbums   []SearcFirstAlbum    `json:"firstAlbums"`
	CreationDates []SearchCreationDate `json:"creationDates"`
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
