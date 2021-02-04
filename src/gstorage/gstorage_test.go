package gstorage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/piyuo/libsrv/src/gaccount"
	"github.com/stretchr/testify/assert"
)

func TestNewGstorage(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	cred, err := gaccount.GlobalCredential(ctx)
	assert.Nil(err)
	storage, err := NewGstorage(ctx, cred)
	assert.Nil(err)
	assert.NotNil(storage)
}

func TestGstorageBucket(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	cred, err := gaccount.GlobalCredential(ctx)
	assert.Nil(err)
	storage, err := NewGstorage(ctx, cred)

	bucketName := "gstorage.piyuo.com"

	err = storage.RemoveBucket(ctx, bucketName)
	assert.Nil(err)

	exist, err := storage.IsBucketExists(ctx, bucketName)
	assert.Nil(err)
	assert.False(exist)

	err = storage.AddBucket(ctx, bucketName, "US")
	assert.Nil(err)

	exist, err = storage.IsBucketExists(ctx, bucketName)
	assert.Nil(err)
	assert.True(exist)

	err = storage.RemoveBucket(ctx, bucketName)
	assert.Nil(err)

	exist, err = storage.IsBucketExists(ctx, bucketName)
	assert.Nil(err)
	assert.False(exist)
}

func TestGstorageReadWriteDelete(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	cred, err := gaccount.GlobalCredential(ctx)
	assert.Nil(err)
	storage, err := NewGstorage(ctx, cred)
	bucketName := "gstorage.piyuo.com"
	path := "TestReadWriteDelete.txt"

	err = storage.AddBucket(ctx, bucketName, "US")
	assert.Nil(err)

	found, err := storage.IsFileExists(ctx, bucketName, "", path)
	assert.Nil(err)
	assert.False(found)

	err = storage.WriteText(ctx, bucketName, path, "hi")
	assert.Nil(err)

	found, err = storage.IsFileExists(ctx, bucketName, "", path)
	assert.Nil(err)
	assert.True(found)

	txt, err := storage.ReadText(ctx, bucketName, path)
	assert.Nil(err)
	assert.Equal("hi", txt)

	err = storage.DeleteFile(ctx, bucketName, path)
	assert.Nil(err)

	err = storage.RemoveBucket(ctx, bucketName)
	assert.Nil(err)
}

func TestGstorageCleanBucket(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	cred, err := gaccount.GlobalCredential(ctx)
	assert.Nil(err)
	storage, err := NewGstorage(ctx, cred)
	bucketName := "gstorage.piyuo.com"
	path := "TestCleanBucket.txt"

	err = storage.AddBucket(ctx, bucketName, "US")
	assert.Nil(err)

	for i := 0; i < 1; i++ {
		err = storage.WriteText(ctx, bucketName, fmt.Sprintf("%v%v", path, i), fmt.Sprintf("hi %v", i))
		//fmt.Printf("add object:%v\n", i)
	}
	err = storage.CleanBucket(ctx, bucketName, 25*time.Second)
	assert.Nil(err)
	err = storage.RemoveBucket(ctx, bucketName)
	assert.Nil(err)
}
