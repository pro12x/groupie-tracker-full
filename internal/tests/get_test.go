package tests

import (
	"bim/internal/handlers"
	"bim/internal/models"
	"bim/internal/pkg"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var (
	ID      = pkg.RandomID(1, 52)
	art, _  = handlers.GetArtist(ID)
	arts, _ = handlers.GetArtists()
	json, _ = handlers.GetJSON("https://groupietrackers.herokuapp.com/api")
	search  = handlers.GetSearch("usa", arts)
)

func TestGetArtist(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Artist
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Artist " + strconv.Itoa(ID),
			args: args{
				id: ID,
			},
			want:    art,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetArtist(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("❌ GetArtist() error = %v, wantErr %v", err, tt.wantErr)
				// fmt.Println("❌ FAIL:\n\tError:", err, "\n\tWant:", tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArtist() got = %v, want %v", got, tt.want)
				// fmt.Println("❌ FAIL:\n\tError:", got, "\n\tWant:", tt.want)
			} else {
				fmt.Println("✔ PASSED")
			}
		})
	}
}

func TestGetArtists(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Artist
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Passed Test",
			want:    arts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetArtists()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArtists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArtists() got = %v, want %v", got, tt.want)
			} else {
				fmt.Println("✔ PASSED")
			}
		})
	}
}

func TestGetDate(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Date
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetDate(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDates(t *testing.T) {
	tests := []struct {
		name    string
		want    models.Dates
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetDates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetLocation(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocations(t *testing.T) {
	tests := []struct {
		name    string
		want    models.Locations
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetLocations()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandom(t *testing.T) {
	type args struct {
		nb int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Artist
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetRandom(tt.args.nb)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRelation(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Relation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetRelation(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRelation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRelation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRelations(t *testing.T) {
	tests := []struct {
		name    string
		want    models.Relations
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetRelations()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRelations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRelations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSearch(t *testing.T) {
	type args struct {
		query   string
		artists []models.Artist
	}
	tests := []struct {
		name string
		args args
		want models.Search
	}{
		// TODO: Add test cases.
		{
			name: "Location",
			args: args{
				query:   "usa",
				artists: arts,
			},
			want: search,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handlers.GetSearch(tt.args.query, tt.args.artists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSearch() = %v, want %v", got, tt.want)
			} else {
				// fmt.Println("✔ PASSED")
				fmt.Println("✔ PASSED", search, tt.want)
			}
		})
	}
}

func Test_getJSON(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Fetch Data",
			args: args{
				url: "https://groupietrackers.herokuapp.com/api",
			},
			want:    json,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.GetJSON(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getJSON() got = %v, want %v", got, tt.want)
			} else {
				fmt.Println("✔ PASSED")
			}
		})
	}
}
