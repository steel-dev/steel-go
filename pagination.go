package steel

import "context"

type CursorPage[T any] struct {
	Sessions   []T    `json:"sessions"`
	NextCursor string `json:"nextCursor"`
	fetch      func(ctx context.Context, cursor string) (*CursorPage[T], error)
	ctx        context.Context
	nextCursor func([]T, string) string
}

func (p *CursorPage[T]) GetNextPage() (*CursorPage[T], error) {
	if !(len(p.Sessions) > 0) {
		return nil, nil
	}
	cursor := p.nextCursor(p.Sessions, p.NextCursor)
	if cursor == "" {
		return nil, nil
	}
	return p.fetch(p.ctx, cursor)
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	err  error
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{page: page, err: err}
}

func (a *CursorPageAutoPager[T]) Next() bool {
	if a.err != nil || a.page == nil {
		return false
	}
	for a.idx >= len(a.page.Sessions) {
		next, err := a.page.GetNextPage()
		if err != nil {
			a.err = err
			return false
		}
		if next == nil {
			return false
		}
		a.page = next
		a.idx = 0
	}
	a.cur = a.page.Sessions[a.idx]
	a.idx++
	return true
}

func (a *CursorPageAutoPager[T]) Current() T {
	return a.cur
}

func (a *CursorPageAutoPager[T]) Err() error {
	return a.err
}
