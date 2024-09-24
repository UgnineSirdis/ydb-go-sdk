package topicoptions

import "github.com/UgnineSirdis/ydb-go-sdk/v3/internal/grpcwrapper/rawtopic"

// DropOption type for drop options. Not used now.
type DropOption interface {
	ApplyDropOption(request *rawtopic.DropTopicRequest)
}
