package mongo

import (
	"reflect"
	"testing"
)

func TestCRUD(t *testing.T) {
	type args struct {
		collectionName string
		documents      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "application must be abel to CRUD a collection",
			args: args{
				collectionName: "test",
				documents:      []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Save(tt.args.collectionName, tt.args.documents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
