package value

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/xtest"
)

func ptr[T any]() interface{} {
	var zeroValue T

	return &zeroValue
}

func value2ptr[T any](v T) *T {
	return &v
}

func unwrapPtr(v interface{}) interface{} {
	return reflect.ValueOf(v).Elem().Interface()
}

func loadLocation(T *testing.T, name string) *time.Location {
	loc, err := time.LoadLocation(name)
	require.NoError(T, err)

	return loc
}

func TestCastTo(t *testing.T) {
	testsCases := []struct {
		name  string
		value Value
		dst   interface{}
		exp   interface{}
		err   error
	}{
		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   (interface{})(nil),
			err:   errNilDestination,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   ptr[Value](),
			exp:   TextValue("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(TextValue("test")),
			dst:   ptr[Value](),
			exp:   OptionalValue(TextValue("test")),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   ptr[string](),
			exp:   "test",
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(TextValue("test")),
			dst:   ptr[*string](),
			exp:   value2ptr("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   ptr[[]byte](),
			exp:   []byte("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(TextValue("test")),
			dst:   ptr[*[]byte](),
			exp:   value2ptr([]byte("test")),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   ptr[[]byte](),
			exp:   []byte("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TextValue("test"),
			dst:   ptr[int](),
			err:   ErrCannotCast,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: BytesValue([]byte("test")),
			dst:   ptr[Value](),
			exp:   BytesValue([]byte("test")),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: BytesValue([]byte("test")),
			dst:   ptr[string](),
			exp:   "test",
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: BytesValue([]byte("test")),
			dst:   ptr[[]byte](),
			exp:   []byte("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: BytesValue([]byte("test")),
			dst:   ptr[[]byte](),
			exp:   []byte("test"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: BytesValue([]byte("test")),
			dst:   ptr[int](),
			err:   ErrCannotCast,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: JSONDocumentValue(`{"test": "text"}"`),
			dst:   ptr[string](),
			exp:   `{"test": "text"}"`,
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: JSONDocumentValue(`{"test":"text"}"`),
			dst:   ptr[Value](),
			exp:   JSONDocumentValue(`{"test":"text"}"`),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(JSONDocumentValue(`{"test": "text"}"`)),
			dst:   ptr[*[]byte](),
			exp:   value2ptr([]byte(`{"test": "text"}"`)),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: JSONDocumentValue(`{"test":"text"}"`),
			dst:   ptr[[]byte](),
			exp:   []byte(`{"test":"text"}"`),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: BoolValue(true),
			dst:   ptr[bool](),
			exp:   true,
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: BoolValue(true),
			dst:   ptr[Value](),
			exp:   BoolValue(true),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(BoolValue(true)),
			dst:   ptr[*bool](),
			exp:   value2ptr(true),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[int32](),
			exp:   int32(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[Value](),
			exp:   Int32Value(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[int64](),
			exp:   int64(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[float32](),
			exp:   float32(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[float64](),
			exp:   float64(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(Int32Value(123)),
			dst:   ptr[*int32](),
			exp:   value2ptr(int32(123)),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[string](),
			exp:   "123",
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int32Value(123),
			dst:   ptr[[]byte](),
			exp:   []byte("123"),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: Int64Value(123),
			dst:   ptr[int64](),
			exp:   int64(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int64Value(123),
			dst:   ptr[Value](),
			exp:   Int64Value(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(Int64Value(123)),
			dst:   ptr[*int64](),
			exp:   value2ptr(int64(123)),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int64Value(123),
			dst:   ptr[float64](),
			exp:   float64(123),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int64Value(123),
			dst:   ptr[string](),
			exp:   "123",
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: Int64Value(123),
			dst:   ptr[[]byte](),
			exp:   []byte("123"),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: DoubleValue(1.23),
			dst:   ptr[float64](),
			exp:   1.23,
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: DoubleValue(1.23),
			dst:   ptr[Value](),
			exp:   DoubleValue(1.23),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: OptionalValue(DoubleValue(1.23)),
			dst:   ptr[*float64](),
			exp:   value2ptr(1.23),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: IntervalValueFromDuration(time.Second),
			dst:   ptr[time.Duration](),
			exp:   time.Second,
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: IntervalValueFromDuration(time.Second),
			dst:   ptr[Value](),
			exp:   IntervalValueFromDuration(time.Second),
			err:   nil,
		},

		// nanoseconds are ignored in YDB timestamps
		{
			name:  xtest.CurrentFileLine(),
			value: TimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local)),
			dst:   ptr[time.Time](),
			exp:   time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local)),
			dst:   ptr[Value](),
			exp:   TimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local)),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TimestampValue(uint64(time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local).Unix())),
			dst:   ptr[uint64](),
			exp:   uint64(time.Date(2024, 1, 2, 3, 4, 5, 0, time.Local).Unix()),
			err:   nil,
		},

		// nanoseconds are ignored in YDB timestamps
		{
			name:  xtest.CurrentFileLine(),
			value: TzTimestampValue("2024-01-02T03:04:05.000000,Europe/Moscow"),
			dst:   ptr[time.Time](),
			exp:   time.Date(2024, 1, 2, 3, 4, 5, 0, loadLocation(t, "Europe/Moscow")),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TzTimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)),
			dst:   ptr[string](),
			exp:   "2024-01-02T03:04:05.000000Z",
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TzTimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)),
			dst:   ptr[[]byte](),
			exp:   []byte("2024-01-02T03:04:05.000000Z"),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: TzTimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)),
			dst:   ptr[Value](),
			exp:   TzTimestampValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)),
			err:   nil,
		},

		{
			name:  xtest.CurrentFileLine(),
			value: DateValue(uint32(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC).Unix())),
			dst:   ptr[int32](),
			exp:   int32(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC).Unix()),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: DateValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)),
			dst:   ptr[uint64](),
			exp:   uint64(DateValueFromTime(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC))),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: DateValueFromTime(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)),
			dst:   ptr[int64](),
			exp:   int64(DateValueFromTime(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC))),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: DateValueFromTime(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)),
			dst:   ptr[time.Time](),
			exp:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			err:   nil,
		},
		{
			name:  xtest.CurrentFileLine(),
			value: DateValueFromTime(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)),
			dst:   ptr[Value](),
			exp:   DateValueFromTime(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)),
			err:   nil,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				require.NoError(t, CastTo(tt.value, tt.dst))
				require.Equal(t, tt.exp, unwrapPtr(tt.dst))
			} else {
				require.ErrorIs(t, CastTo(tt.value, tt.dst), tt.err)
			}
		})
	}
}
