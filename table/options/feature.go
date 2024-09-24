package options

import (
	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/feature"
)

type FeatureFlag = feature.Flag

const (
	FeatureEnabled  = feature.Enabled
	FeatureDisabled = feature.Disabled
)
