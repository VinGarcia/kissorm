package kissorm

import "context"

var _ ORMProvider = MockORMProvider{}

// MockORMProvider ...
type MockORMProvider struct {
	InsertFn func(ctx context.Context, records ...interface{}) error
	DeleteFn func(ctx context.Context, ids ...interface{}) error
	UpdateFn func(ctx context.Context, records ...interface{}) error

	QueryFn       func(ctx context.Context, records interface{}, query string, params ...interface{}) error
	QueryOneFn    func(ctx context.Context, record interface{}, query string, params ...interface{}) error
	QueryChunksFn func(ctx context.Context, parser ChunkParser) error

	ExecFn        func(ctx context.Context, query string, params ...interface{}) error
	TransactionFn func(ctx context.Context, fn func(db ORMProvider) error) error
}

// Insert ...
func (m MockORMProvider) Insert(ctx context.Context, records ...interface{}) error {
	return m.InsertFn(ctx, records...)
}

// Delete ...
func (m MockORMProvider) Delete(ctx context.Context, ids ...interface{}) error {
	return m.DeleteFn(ctx, ids...)
}

// Update ...
func (m MockORMProvider) Update(ctx context.Context, records ...interface{}) error {
	return m.UpdateFn(ctx, records...)
}

// Query ...
func (m MockORMProvider) Query(ctx context.Context, records interface{}, query string, params ...interface{}) error {
	return m.QueryFn(ctx, records, query, params...)
}

// QueryOne ...
func (m MockORMProvider) QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error {
	return m.QueryOneFn(ctx, record, query, params...)
}

// QueryChunks ...
func (m MockORMProvider) QueryChunks(ctx context.Context, parser ChunkParser) error {
	return m.QueryChunksFn(ctx, parser)
}

// Exec ...
func (m MockORMProvider) Exec(ctx context.Context, query string, params ...interface{}) error {
	return m.ExecFn(ctx, query, params...)
}

// Transaction ...
func (m MockORMProvider) Transaction(ctx context.Context, fn func(db ORMProvider) error) error {
	return m.TransactionFn(ctx, fn)
}
