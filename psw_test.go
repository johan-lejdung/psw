package psw_test

import (
	"context"
	"testing"

	psw "github.com/johan-lejdung/psw"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func Test_CreateSubscription_NotExist(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the server without using TLS.
	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	assert.NoError(t, err)

	defer conn.Close()
	ctx := context.Background()
	// Use the connection when creating a pubsub client.
	c, err := pubsub.NewClient(ctx, "testproj", option.WithGRPCConn(conn))
	assert.NoError(t, err)
	defer c.Close()

	topic, err := psw.CreateTopicIfNotExists(c, "testtopic")
	assert.NoError(t, err)

	_, err = psw.CreateSubscriptionIfNotExists(c, "testsub", topic)
	assert.NoError(t, err)
}

func Test_CreateSubscription_Exist(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the server without using TLS.
	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	assert.NoError(t, err)

	defer conn.Close()
	ctx := context.Background()
	// Use the connection when creating a pubsub client.
	c, err := pubsub.NewClient(ctx, "testproj", option.WithGRPCConn(conn))
	assert.NoError(t, err)
	defer c.Close()

	topic, err := psw.CreateTopicIfNotExists(c, "testtopic")
	assert.NoError(t, err)

	_, err = psw.CreateSubscriptionIfNotExists(c, "testsub", topic)
	assert.NoError(t, err)

	_, err = psw.CreateSubscriptionIfNotExists(c, "testsub", topic)
	assert.NoError(t, err)
}

func Test_CreateTopic_NotExist(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the server without using TLS.
	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	assert.NoError(t, err)

	defer conn.Close()
	ctx := context.Background()
	// Use the connection when creating a pubsub client.
	c, err := pubsub.NewClient(ctx, "testproj", option.WithGRPCConn(conn))
	assert.NoError(t, err)
	defer c.Close()

	_, err = psw.CreateTopicIfNotExists(c, "testtopic")
	assert.NoError(t, err)
}

func Test_CreateTopic_Exist(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the server without using TLS.
	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	assert.NoError(t, err)

	defer conn.Close()
	ctx := context.Background()
	// Use the connection when creating a pubsub client.
	c, err := pubsub.NewClient(ctx, "testproj", option.WithGRPCConn(conn))
	assert.NoError(t, err)
	defer c.Close()

	_, err = psw.CreateTopicIfNotExists(c, "testtopic")
	assert.NoError(t, err)

	_, err = psw.CreateTopicIfNotExists(c, "testtopic")
	assert.NoError(t, err)
}
