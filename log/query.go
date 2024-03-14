package log

import (
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

// Query makes trace.Query with logging events from details
func Query(l Logger, d trace.Detailer, opts ...Option) (t trace.Query) {
	return internalQuery(wrapLogger(l, opts...), d)
}

//nolint:gocyclo
func internalQuery(
	l *wrapper, //nolint:interfacer
	d trace.Detailer,
) (t trace.Query) {
	t.OnNew = func(info trace.QueryNewStartInfo) func(info trace.QueryNewDoneInfo) {
		if d.Details()&trace.QueryEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "new")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryNewDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := FATAL
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnClose = func(info trace.QueryCloseStartInfo) func(info trace.QueryCloseDoneInfo) {
		if d.Details()&trace.QueryEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "close")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryCloseDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolNew = func(info trace.QueryPoolNewStartInfo) func(trace.QueryPoolNewDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "new")
		l.Log(ctx, "start",
			Int("MinSize", info.MinSize),
			Int("MaxSize", info.MaxSize),
			Int("ProducersCount", info.ProducersCount),
		)
		start := time.Now()

		return func(info trace.QueryPoolNewDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
					Int("MinSize", info.MinSize),
					Int("MaxSize", info.MaxSize),
					Int("ProducersCount", info.ProducersCount),
				)
			} else {
				lvl := FATAL
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					Int("MinSize", info.MinSize),
					Int("MaxSize", info.MaxSize),
					Int("ProducersCount", info.ProducersCount),
					versionField(),
				)
			}
		}
	}
	t.OnPoolClose = func(info trace.QueryPoolCloseStartInfo) func(trace.QueryPoolCloseDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "close")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolCloseDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolProduce = func(info trace.QueryPoolProduceStartInfo) func(trace.QueryPoolProduceDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "produce")
		l.Log(ctx, "start",
			Int("Concurrency", info.Concurrency),
		)
		start := time.Now()

		return func(info trace.QueryPoolProduceDoneInfo) {
			l.Log(ctx, "done",
				latencyField(start),
			)
		}
	}
	t.OnPoolPut = func(info trace.QueryPoolPutStartInfo) func(trace.QueryPoolPutDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "put")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolPutDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolGet = func(info trace.QueryPoolGetStartInfo) func(trace.QueryPoolGetDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "get")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolGetDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolWant = func(info trace.QueryPoolWantStartInfo) func(trace.QueryPoolWantDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "want")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolWantDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolSpawn = func(info trace.QueryPoolSpawnStartInfo) func(trace.QueryPoolSpawnDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "spawn")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolSpawnDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolTry = func(info trace.QueryPoolTryStartInfo) func(trace.QueryPoolTryDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "pool", "try")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolTryDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnPoolWith = func(info trace.QueryPoolWithStartInfo) func(trace.QueryPoolWithDoneInfo) {
		if d.Details()&trace.QueryPoolEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, DEBUG, "ydb", "query", "pool", "with")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryPoolWithDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
					Int("Attempts", info.Attempts),
				)
			} else {
				lvl := ERROR
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					Int("Attempts", info.Attempts),
					versionField(),
				)
			}
		}
	}
	t.OnDo = func(
		info trace.QueryDoStartInfo,
	) func(
		info trace.QueryDoIntermediateInfo,
	) func(
		trace.QueryDoDoneInfo,
	) {
		if d.Details()&trace.QueryEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "do")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryDoIntermediateInfo) func(trace.QueryDoDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}

			return func(info trace.QueryDoDoneInfo) {
				if info.Error == nil {
					l.Log(ctx, "done",
						latencyField(start),
						Int("attempts", info.Attempts),
					)
				} else {
					lvl := ERROR
					if !xerrors.IsYdb(info.Error) {
						lvl = DEBUG
					}
					l.Log(WithLevel(ctx, lvl), "failed",
						latencyField(start),
						Error(info.Error),
						Int("attempts", info.Attempts),
						versionField(),
					)
				}
			}
		}
	}
	t.OnDoTx = func(
		info trace.QueryDoTxStartInfo,
	) func(
		info trace.QueryDoTxIntermediateInfo,
	) func(
		trace.QueryDoTxDoneInfo,
	) {
		if d.Details()&trace.QueryEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "do", "tx")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QueryDoTxIntermediateInfo) func(trace.QueryDoTxDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "done",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}

			return func(info trace.QueryDoTxDoneInfo) {
				if info.Error == nil {
					l.Log(ctx, "done",
						latencyField(start),
						Int("attempts", info.Attempts),
					)
				} else {
					lvl := ERROR
					if !xerrors.IsYdb(info.Error) {
						lvl = DEBUG
					}
					l.Log(WithLevel(ctx, lvl), "failed",
						latencyField(start),
						Error(info.Error),
						Int("attempts", info.Attempts),
						versionField(),
					)
				}
			}
		}
	}
	t.OnSessionCreate = func(info trace.QuerySessionCreateStartInfo) func(info trace.QuerySessionCreateDoneInfo) {
		if d.Details()&trace.QuerySessionEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "session", "create")
		l.Log(ctx, "start")
		start := time.Now()

		return func(info trace.QuerySessionCreateDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
					String("session_id", info.Session.ID()),
					String("session_status", info.Session.Status()),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "done",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnSessionDelete = func(info trace.QuerySessionDeleteStartInfo) func(info trace.QuerySessionDeleteDoneInfo) {
		if d.Details()&trace.QuerySessionEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "session", "delete")
		l.Log(ctx, "start",
			String("session_id", info.Session.ID()),
			String("session_status", info.Session.Status()),
		)
		start := time.Now()

		return func(info trace.QuerySessionDeleteDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}
	t.OnSessionAttach = func(info trace.QuerySessionAttachStartInfo) func(info trace.QuerySessionAttachDoneInfo) {
		if d.Details()&trace.QuerySessionEvents == 0 {
			return nil
		}
		ctx := with(*info.Context, TRACE, "ydb", "query", "session", "attach")
		l.Log(ctx, "start",
			String("session_id", info.Session.ID()),
			String("session_status", info.Session.Status()),
		)
		start := time.Now()

		return func(info trace.QuerySessionAttachDoneInfo) {
			if info.Error == nil {
				l.Log(ctx, "done",
					latencyField(start),
				)
			} else {
				lvl := WARN
				if !xerrors.IsYdb(info.Error) {
					lvl = DEBUG
				}
				l.Log(WithLevel(ctx, lvl), "failed",
					latencyField(start),
					Error(info.Error),
					versionField(),
				)
			}
		}
	}

	return t
}
