package steel

import (
	"context"
	"encoding/json"
)

type CursorPage[T any] struct {
	Items      []T
	cursor     string
	hasMore    bool
	fetch      func(ctx context.Context, cursor string) (*CursorPage[T], error)
	ctx        context.Context
	nextCursor func([]T, string) string
}

func (p *CursorPage[T]) GetNextPage() (*CursorPage[T], error) {
	if !p.hasMore {
		return nil, nil
	}
	cursor := p.nextCursor(p.Items, p.cursor)
	if cursor == "" {
		return nil, nil
	}
	return p.fetch(p.ctx, cursor)
}

func decodeCursorPage[T any](raw map[string]json.RawMessage, itemsField, cursorField, hasMoreField string) ([]T, string, bool, error) {
	var items []T
	if msg, ok := raw[itemsField]; ok {
		if err := json.Unmarshal(msg, &items); err != nil {
			return nil, "", false, err
		}
	}
	var cursor string
	if cursorField != "" {
		if msg, ok := raw[cursorField]; ok {
			_ = json.Unmarshal(msg, &cursor)
		}
	}
	hasMore := len(items) > 0
	if hasMoreField != "" {
		if msg, ok := raw[hasMoreField]; ok {
			_ = json.Unmarshal(msg, &hasMore)
		}
	}
	return items, cursor, hasMore, nil
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
	for a.idx >= len(a.page.Items) {
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
	a.cur = a.page.Items[a.idx]
	a.idx++
	return true
}

func (a *CursorPageAutoPager[T]) Current() T {
	return a.cur
}

func (a *CursorPageAutoPager[T]) Err() error {
	return a.err
}
