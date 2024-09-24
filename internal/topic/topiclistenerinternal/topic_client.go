package topiclistenerinternal

import (
	"context"

	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/grpcwrapper/rawtopic/rawtopicreader"
)

type TopicClient interface {
	StreamRead(connectionCtx context.Context) (rawtopicreader.StreamReader, error)
}
