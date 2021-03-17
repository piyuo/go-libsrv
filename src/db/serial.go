package db

import "context"

// Serial can generate serial Sequence in low frequency. 1 per second, use it with high frequency will cause error
//
type Serial interface {

	// NumberRX return sequence number, number is unique and serial, please be aware serial can only generate one sequence per second, use it with high frequency will cause error and  must used it in transaction with NumberWX()
	//
	//	err := Transaction(ctx, func(ctx context.Context,tx db.Transaction) error {
	//		num, err:= serial.NumberRX(ctx,tx)
	//		err := serial.NumberWX(ctx,tx)
	//	})
	//
	NumberRX(ctx context.Context, transaction Transaction) (int64, error)

	// NumberWX commit NumberRX
	//
	//	err := Transaction(ctx, func(ctx context.Context,tx db.Transaction) error {
	//		num, err:= serial.NumberRX(ctx,tx)
	//		err := serial.NumberWX(ctx,tx)
	//	})
	//
	NumberWX(ctx context.Context, transaction Transaction) error

	// Clear all shards
	//
	//	err = serial.Clear(ctx,100)
	//
	Clear(ctx context.Context, max int) (bool, error)

	// ShardsCount returns shards count
	//
	//	count, err = serial.ShardsCount(ctx)
	//
	ShardsCount(ctx context.Context) (int, error)
}
