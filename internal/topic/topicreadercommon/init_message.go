package topicreadercommon

import "github.com/UgnineSirdis/ydb-go-sdk/v3/internal/grpcwrapper/rawtopic/rawtopicreader"

func CreateInitMessage(consumer string, selectors []*PublicReadSelector) *rawtopicreader.InitRequest {
	res := &rawtopicreader.InitRequest{
		Consumer: consumer,
	}

	res.TopicsReadSettings = make([]rawtopicreader.TopicReadSettings, len(selectors))
	for i, selector := range selectors {
		settings := &res.TopicsReadSettings[i]
		settings.Path = selector.Path
		settings.PartitionsID = selector.Partitions
		if !selector.ReadFrom.IsZero() {
			settings.ReadFrom.HasValue = true
			settings.ReadFrom.Value = selector.ReadFrom
		}
		if selector.MaxTimeLag != 0 {
			settings.MaxLag.HasValue = true
			settings.MaxLag.Value = selector.MaxTimeLag
		}
	}

	return res
}
