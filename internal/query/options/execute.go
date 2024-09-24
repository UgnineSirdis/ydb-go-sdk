package options

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Query"
	"google.golang.org/grpc"

	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/params"
	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/query/tx"
	"github.com/UgnineSirdis/ydb-go-sdk/v3/internal/stats"
	"github.com/UgnineSirdis/ydb-go-sdk/v3/retry"
)

var (
	_ Execute = callOptionsOption(nil)
	_ Execute = (*txCommitOption)(nil)
	_ Execute = parametersOption{}
	_ Execute = (*txControlOption)(nil)
	_ Execute = syntaxOption(0)
	_ Execute = statsModeOption{}
	_ Execute = execModeOption(0)
)

type (
	Syntax    Ydb_Query.Syntax
	ExecMode  Ydb_Query.ExecMode
	StatsMode Ydb_Query.StatsMode

	// executeSettings is a holder for execute settings
	executeSettings struct {
		syntax        Syntax
		params        params.Parameters
		execMode      ExecMode
		statsMode     StatsMode
		statsCallback func(queryStats stats.QueryStats)
		callOptions   []grpc.CallOption
		txControl     *tx.Control
		retryOptions  []retry.Option
	}

	// Execute is an interface for execute method options
	Execute interface {
		applyExecuteOption(s *executeSettings)
	}

	ExecuteNoTx interface {
		thisOptionIsNotForExecuteOnTx()
	}

	// execute options
	callOptionsOption []grpc.CallOption
	txCommitOption    struct{}
	parametersOption  params.Parameters
	txControlOption   tx.Control
	syntaxOption      = Syntax
	statsModeOption   struct {
		mode     StatsMode
		callback func(stats.QueryStats)
	}
	execModeOption = ExecMode
)

func (s *executeSettings) RetryOpts() []retry.Option {
	return s.retryOptions
}

func (s *executeSettings) StatsCallback() func(stats.QueryStats) {
	return s.statsCallback
}

func (t txCommitOption) applyExecuteOption(s *executeSettings) {
	s.txControl.Commit = true
}

func (txControl *txControlOption) applyExecuteOption(s *executeSettings) {
	s.txControl = (*tx.Control)(txControl)
}

func (txControl *txControlOption) thisOptionIsNotForExecuteOnTx() {}

func (syntax Syntax) applyExecuteOption(s *executeSettings) {
	s.syntax = syntax
}

const (
	SyntaxYQL        = Syntax(Ydb_Query.Syntax_SYNTAX_YQL_V1)
	SyntaxPostgreSQL = Syntax(Ydb_Query.Syntax_SYNTAX_PG)
)

func (params parametersOption) applyExecuteOption(s *executeSettings) {
	s.params = append(s.params, params...)
}

func (opts callOptionsOption) applyExecuteOption(s *executeSettings) {
	s.callOptions = append(s.callOptions, opts...)
}

func (mode StatsMode) applyExecuteOption(s *executeSettings) {
	s.statsMode = mode
}

func (mode ExecMode) applyExecuteOption(s *executeSettings) {
	s.execMode = mode
}

const (
	ExecModeParse    = ExecMode(Ydb_Query.ExecMode_EXEC_MODE_PARSE)
	ExecModeValidate = ExecMode(Ydb_Query.ExecMode_EXEC_MODE_VALIDATE)
	ExecModeExplain  = ExecMode(Ydb_Query.ExecMode_EXEC_MODE_EXPLAIN)
	ExecModeExecute  = ExecMode(Ydb_Query.ExecMode_EXEC_MODE_EXECUTE)
)

const (
	StatsModeBasic   = StatsMode(Ydb_Query.StatsMode_STATS_MODE_BASIC)
	StatsModeNone    = StatsMode(Ydb_Query.StatsMode_STATS_MODE_NONE)
	StatsModeFull    = StatsMode(Ydb_Query.StatsMode_STATS_MODE_FULL)
	StatsModeProfile = StatsMode(Ydb_Query.StatsMode_STATS_MODE_PROFILE)
)

func defaultExecuteSettings() executeSettings {
	return executeSettings{
		syntax:    SyntaxYQL,
		execMode:  ExecModeExecute,
		statsMode: StatsModeNone,
		txControl: tx.DefaultTxControl(),
	}
}

func ExecuteSettings(opts ...Execute) *executeSettings {
	settings := defaultExecuteSettings()

	for _, opt := range opts {
		if opt != nil {
			opt.applyExecuteOption(&settings)
		}
	}

	return &settings
}

func (s *executeSettings) TxControl() *tx.Control {
	return s.txControl
}

func (s *executeSettings) CallOptions() []grpc.CallOption {
	return s.callOptions
}

func (s *executeSettings) Syntax() Syntax {
	return s.syntax
}

func (s *executeSettings) ExecMode() ExecMode {
	return s.execMode
}

func (s *executeSettings) StatsMode() StatsMode {
	return s.statsMode
}

func (s *executeSettings) Params() *params.Parameters {
	if len(s.params) == 0 {
		return nil
	}

	return &s.params
}

func WithParameters(parameters *params.Parameters) parametersOption {
	return parametersOption(*parameters)
}

var (
	_ Execute = ExecMode(0)
	_ Execute = StatsMode(0)
	_ Execute = ExecMode(0)
	_ Execute = StatsMode(0)
	_ Execute = txCommitOption{}
	_ Execute = (*txControlOption)(nil)
)

func WithCommit() txCommitOption {
	return txCommitOption{}
}

func WithExecMode(mode ExecMode) execModeOption {
	return mode
}

func WithSyntax(syntax Syntax) syntaxOption {
	return syntax
}

func (opt statsModeOption) applyExecuteOption(s *executeSettings) {
	s.statsMode = opt.mode
	s.statsCallback = opt.callback
}

func WithStatsMode(mode StatsMode, callback func(stats.QueryStats)) statsModeOption {
	return statsModeOption{
		mode:     mode,
		callback: callback,
	}
}

func WithCallOptions(opts ...grpc.CallOption) callOptionsOption {
	return opts
}

func WithTxControl(txControl *tx.Control) *txControlOption {
	return (*txControlOption)(txControl)
}
