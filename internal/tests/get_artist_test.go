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
	ID     = pkg.RandomID(1, 52)
	art, _ = handlers.GetArtist(ID)
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
				fmt.Println("✔ PASS")
			}
		})
	}
}
