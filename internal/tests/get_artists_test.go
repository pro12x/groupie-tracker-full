package tests

import (
	"bim/internal/handlers"
	"bim/internal/models"
	"fmt"
	"reflect"
	"testing"
)

var arts, _ = handlers.GetArtists()

func TestGetArtists(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Artist
		wantErr bool
	}{
		{
			name:    "Passed Test",
			want:    arts,
			wantErr: false,
		},
		{
			name:    "Failed Test",
			want:    []models.Artist{},
			wantErr: true,
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
				fmt.Println("✔ PASS")
			}
		})
	}
}
