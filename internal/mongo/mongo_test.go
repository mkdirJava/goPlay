package mongo

import (
	"fmt"
	"os"
	"testing"
)

type test_type struct {
	Updates `bson:"updates"`
	Name    string
}

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
			args: args{"test", []interface{}{&test_type{Name: "bob"}}},
			want: []interface{}{},
		},
	}

	os.Setenv("MONGO_URI", "mongodb://root:example@localhost:27018")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savedIds := Save(tt.args.collectionName, tt.args.documents)
			initialRetrieved := Get[test_type](tt.args.collectionName, savedIds)
			if len(initialRetrieved) == 0 {
				panic("Saved should of inserted one document")
			}
			Delete[test_type](tt.args.collectionName, savedIds)
			deletedRetrieved := Get[test_type](tt.args.collectionName, savedIds)
			if deletedRetrieved[0].Name != "bob" {
				panic(fmt.Errorf("Deleted retrieve should be bob"))
			}
			if deletedRetrieved[0].Deleted_at.IsZero() {
				panic(fmt.Errorf("deleted_at should not be a zero value"))
			}

		})
	}
}
