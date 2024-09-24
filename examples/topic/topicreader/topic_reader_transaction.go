package topicreaderexamples

import (
	"context"

	"github.com/UgnineSirdis/ydb-go-sdk/v3"
	"github.com/UgnineSirdis/ydb-go-sdk/v3/query"
	"github.com/UgnineSirdis/ydb-go-sdk/v3/topic/topicreader"
)

func CommitMessagesToTransaction(ctx context.Context, db *ydb.Driver, reader *topicreader.Reader) error {
	for { // loop
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := db.Query().DoTx(ctx, func(ctx context.Context, tx query.TxActor) error {
			batch, err := reader.PopMessagesBatchTx(ctx, tx) // the batch will be committed with commit the tx
			if err != nil {
				return err
			}
			id := batch.Messages[0].MessageGroupID //nolint:staticcheck

			batchResult, err := processBatch(batch.Context(), batch)
			if err != nil {
				return err
			}

			err = tx.Exec(ctx, `
$last = SELECT MAX(val) FROM table WHERE id=$id;
UPSERT INTO t (id, val) VALUES($id, COALESCE($last, 0) + $value)
`, query.WithParameters(
				ydb.ParamsBuilder().
					Param("$id").Text(id).
					Param("$value").Int64(int64(batchResult)).
					Build(),
			))

			return err
		})
		if err != nil {
			return err
		}
	}
}

func PopWithTransaction(ctx context.Context, db *ydb.Driver, reader *topicreader.Reader) error {
	for { // loop
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := db.Query().DoTx(ctx, func(ctx context.Context, tx query.TxActor) error {
			batch, err := reader.PopMessagesBatchTx(ctx, tx)
			if err != nil {
				return err
			}
			id := batch.Messages[0].MessageGroupID //nolint:staticcheck

			batchResult, err := processBatch(batch.Context(), batch)
			if err != nil {
				return err
			}

			err = tx.Exec(ctx, `
$last = SELECT MAX(val) FROM table WHERE id=$id;
UPSERT INTO t (id, val) VALUES($id, COALESCE($last, 0) + $value)
`, query.WithParameters(
				ydb.ParamsBuilder().
					Param("$id").Text(id).
					Param("$value").Int64(int64(batchResult)).
					Build(),
			))
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}
}

func PopWithTransactionRecreateReader(
	ctx context.Context,
	db *ydb.Driver,
	readerFabric func(ctx context.Context, db *ydb.Driver) (reader *topicreader.Reader, err error),
) error {
	// second loop - for retries
	err := db.Query().Do(ctx, func(ctx context.Context, s query.Session) error {
		reader, err := readerFabric(ctx, db)
		if err != nil {
			return err
		}
		defer reader.Close(ctx)

		for { // loop
			tx, err := s.Begin(ctx, nil)
			if err != nil {
				return err
			}

			batch, err := reader.PopMessagesBatchTx(ctx, tx)
			if err != nil {
				return err
			}
			id := batch.Messages[0].MessageGroupID //nolint:staticcheck

			batchResult, err := processBatch(batch.Context(), batch)
			if err != nil {
				return err
			}

			err = tx.Exec(ctx, `
$last = SELECT MAX(val) FROM table WHERE id=$id;
UPSERT INTO t (id, val) VALUES($id, COALESCE($last, 0) + $value)
`,
				query.WithParameters(
					ydb.ParamsBuilder().
						Param("$id").Text(id).
						Param("$value").Int64(int64(batchResult)).
						Build(),
				))
			if err != nil {
				return err
			}

			if err = tx.CommitTx(ctx); err != nil {
				return err
			}
		}
	})

	return err
}
