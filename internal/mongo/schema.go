package mongo

import swagger "github.com/mkdirJava/goPlay/internal/web/go"

type WorkFlow struct {
	Id    int64        `json:"id" bson:"id"`
	Name  string       `json:"name"`
	Tasks swagger.Task `json:"Tasks"`
}
