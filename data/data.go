package data

import (
	"context"
)

// NewDB create db from default provider
//
//	db := data.NewDB(ctx)
//	err := db.Put(ctx, &greet)
//	retrive := Greet{}
//	retrive.SetID(greet.ID())
//	err = db.Get(ctx, &retrive)
func NewDB(ctx context.Context) (DB, error) {
	return NewFirestoreDB(ctx)
}