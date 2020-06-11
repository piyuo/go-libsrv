package data

import "context"

// Query is query interface
type Query interface {
	Where(path, op string, value interface{}) Query
	OrderBy(path string) Query
	OrderByDesc(path string) Query
	Limit(n int) Query
	//Offset(n int) IQuery //in firestore will bill extra mony on offset
	Run(callback func(o Object)) error
}

// AbstractQuery is query object need to implement
type AbstractQuery struct {
	Query
	ctx       context.Context
	newObject func() Object
	limit     int
}
