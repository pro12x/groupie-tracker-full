package models

type MainData struct {
	ArtistList  []Artist  // All Artist
	ArtistOne   Artist    // One Artist
	Relations   Relations // All Relations
	RelationOne Relation  // One Relation
	Locations   Locations // All Locations
	LocationOne Location  // One Location
	Dates       Dates     // All Dates
	DateOne     Date      // One Date
	AppInfos    App       // App Information
	Error       ErrorData // Error information
}

type App struct {
	AppName   string // Name of the application
	PageTitle string // Title of the page
	Attr      string // Name of the page
}

type ErrorData struct {
	Code    string // Error code
	Message string // Error message
}
