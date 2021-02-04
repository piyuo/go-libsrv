package gcp

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/src/key"
	"github.com/piyuo/libsrv/src/region"
	"github.com/stretchr/testify/assert"
)

func TestGcpCredential(t *testing.T) {
	assert := assert.New(t)

	//should create google credential
	bytes, err := key.BytesWithoutCache("gcloud.json")
	assert.Nil(err)
	ctx := context.Background()
	cred, err := makeCredential(ctx, bytes)
	assert.Nil(err)
	assert.NotNil(cred)

	//should keep global credential
	assert.Nil(globalCredential)
	cred, err = GlobalCredential(ctx)
	assert.Nil(err)
	assert.NotNil(cred)
	assert.NotNil(globalCredential)
}

func TestGcpCreateCredential(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	cred, err := CreateCredential(ctx, "gcloud.json")
	assert.Nil(err)
	assert.NotNil(cred)
	cred, err = CreateCredential(ctx, "notExist.json")
	assert.NotNil(err)
	assert.Nil(cred)
}

func TestGcpDataCredentialByRegion(t *testing.T) {
	assert := assert.New(t)
	region.Current = "us"
	ctx := context.Background()
	cred, err := RegionalCredential(ctx)
	assert.Nil(err)
	assert.NotNil(cred)

	region.Current = "jp"
	cred, err = RegionalCredential(ctx)
	assert.Nil(err)
	assert.NotNil(cred)

	region.Current = "be"
	cred, err = RegionalCredential(ctx)
	assert.Nil(err)
	assert.NotNil(cred)
}

func TestGcpCredentialWhenContextCanceled(t *testing.T) {
	assert := assert.New(t)
	deadline := time.Now().Add(time.Duration(1) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	time.Sleep(time.Duration(2) * time.Millisecond)
	_, err := GlobalCredential(ctx)
	assert.NotNil(err)
	_, err = RegionalCredential(ctx)
	assert.NotNil(err)
}
