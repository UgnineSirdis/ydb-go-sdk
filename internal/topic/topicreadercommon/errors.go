package topicreadercommon

import (
	"errors"

	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/xerrors"
)

var PublicErrCommitSessionToExpiredSession = xerrors.Wrap(errors.New("ydb: commit to expired session"))
