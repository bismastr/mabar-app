package session

import (
	"reflect"

	"cloud.google.com/go/firestore"
)

func addUpdate(updates *[]firestore.Update, path string, value interface{}) {
	if reflect.ValueOf(value).IsZero() {
		return
	}

	*updates = append(*updates, firestore.Update{Path: path, Value: value})
}
