//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/version"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func TestQueryTxExecute(t *testing.T) {
	if version.Lt(os.Getenv("YDB_VERSION"), "24.1") {
		t.Skip("query service not allowed in YDB version '" + os.Getenv("YDB_VERSION") + "'")
	}

	scope := newScope(t)

	var (
		columnNames []string
		columnTypes []string
	)
	t.Run("SerializableReadWrite", func(t *testing.T) {
		err := scope.DriverWithLogs().Query().DoTx(scope.Ctx, func(ctx context.Context, tx query.TxActor) (err error) {
			res, err := tx.Query(ctx, "SELECT 1 AS col1")
			if err != nil {
				return err
			}
			rs, err := res.NextResultSet(ctx)
			if err != nil {
				return err
			}
			columnNames = rs.Columns()
			for _, t := range rs.ColumnTypes() {
				columnTypes = append(columnTypes, t.Yql())
			}
			row, err := rs.NextRow(ctx)
			if err != nil {
				return err
			}
			var col1 int
			err = row.ScanNamed(query.Named("col1", &col1))
			if err != nil {
				return err
			}
			return nil
		}, query.WithIdempotent(), query.WithTxSettings(query.TxSettings(query.WithSerializableReadWrite())))
		require.NoError(t, err)
		require.Equal(t, []string{"col1"}, columnNames)
		require.Equal(t, []string{"Int32"}, columnTypes)
	})
	t.Run("SnapshotReadOnly", func(t *testing.T) {
		err := scope.DriverWithLogs().Query().DoTx(scope.Ctx, func(ctx context.Context, tx query.TxActor) (err error) {
			res, err := tx.Query(ctx, "SELECT 1 AS col1")
			if err != nil {
				return err
			}
			rs, err := res.NextResultSet(ctx)
			if err != nil {
				return err
			}
			columnNames = rs.Columns()
			for _, t := range rs.ColumnTypes() {
				columnTypes = append(columnTypes, t.Yql())
			}
			row, err := rs.NextRow(ctx)
			if err != nil {
				return err
			}
			var col1 int
			err = row.ScanNamed(query.Named("col1", &col1))
			if err != nil {
				return err
			}
			return nil
		}, query.WithIdempotent(), query.WithTxSettings(query.TxSettings(query.WithSnapshotReadOnly())))
		require.NoError(t, err)
		require.Equal(t, []string{"col1"}, columnNames)
		require.Equal(t, []string{"Int32"}, columnTypes)
	})
	t.Run("OnlineReadOnly", func(t *testing.T) {
		err := scope.DriverWithLogs().Query().DoTx(scope.Ctx, func(ctx context.Context, tx query.TxActor) (err error) {
			res, err := tx.Query(ctx, "SELECT 1 AS col1")
			if err != nil {
				return err
			}
			rs, err := res.NextResultSet(ctx)
			if err != nil {
				return err
			}
			columnNames = rs.Columns()
			for _, t := range rs.ColumnTypes() {
				columnTypes = append(columnTypes, t.Yql())
			}
			row, err := rs.NextRow(ctx)
			if err != nil {
				return err
			}
			var col1 int
			err = row.ScanNamed(query.Named("col1", &col1))
			if err != nil {
				return err
			}
			return nil
		}, query.WithIdempotent(), query.WithTxSettings(query.TxSettings(query.WithOnlineReadOnly())))
		require.Error(t, err)
		require.ErrorIs(t, err, query.ErrIncompatibleTxSettings)
	})
}

func TestQueryWithCommitTxFlag(t *testing.T) {
	if version.Lt(os.Getenv("YDB_VERSION"), "24.1") {
		t.Skip("query service not allowed in YDB version '" + os.Getenv("YDB_VERSION") + "'")
	}

	scope := newScope(t)
	var count uint64
	err := scope.DriverWithLogs().Query().Do(scope.Ctx, func(ctx context.Context, s query.Session) error {
		tableName := scope.TablePath()
		tx, err := s.Begin(ctx, query.TxSettings(query.WithDefaultTxMode()))
		if err != nil {
			return fmt.Errorf("failed start transaction: %w", err)
		}
		q := fmt.Sprintf("UPSERT INTO `%v` (id, val) VALUES(1, \"2\")", tableName)
		err = tx.Exec(ctx, q, query.WithCommit())
		if err != nil {
			return fmt.Errorf("failed execute insert: %w", err)
		}

		// read row within other (implicit) transaction
		q2 := fmt.Sprintf("SELECT COUNT(*) FROM `%v`", tableName)
		r, err := s.Query(ctx, q2)
		if err != nil {
			return fmt.Errorf("failed query: %w", err)
		}

		rs, err := r.NextResultSet(ctx)
		if err != nil {
			return fmt.Errorf("failed iterate to next result set: %w", err)
		}

		row, err := rs.NextRow(ctx)
		if err != nil {
			return fmt.Errorf("failed iterate to next row: %w", err)
		}

		if err = row.Scan(&count); err != nil {
			return fmt.Errorf("failed scan row: %w", err)
		}
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1), count)
}
