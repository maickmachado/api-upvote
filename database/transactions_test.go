package database

import (
	"github.com/maickmachado/upvote-api/models"
	"reflect"
	"testing"
)

func TestCheckIfExist(t *testing.T) {
	type args struct {
		responseObject       models.Response
		detailResponseObject models.CryptoData
	}
	tests := []struct {
		name  string
		args  args
		want  *models.CryptoData
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckIfExist(tt.args.responseObject, tt.args.detailResponseObject)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckIfExist() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckIfExist() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFilterTasks(t *testing.T) {
	type args struct {
		filter interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.CryptoDataBase
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FilterTasks(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.CryptoDataBase
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCryptoInfoDataBase(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCryptoInfoDataBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCryptoInfoDataBase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderByVotes(t *testing.T) {
	tests := []struct {
		name    string
		want    []*models.CryptoDataBase
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByVotes()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByVotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByVotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpvote(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Upvote(tt.args.text)
		})
	}
}

func TestUpvote1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "yurifeio",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Upvote(tt.args.text)
		})
	}
}
