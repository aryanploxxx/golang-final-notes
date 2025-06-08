package main

import (
	"context"
)

func main() {
	exampleWithValues()
}

func exampleWithValues() {
	type key int
	const UserKey key = 0

	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, UserKey, "123")

	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		println("this is the user id:", userID)
	} else {
		println("this is a protected route - user id not found")
	}
}
