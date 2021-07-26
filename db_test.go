package corekv

import (
	"github.com/hardcore-os/corekv/codec"
	"github.com/hardcore-os/corekv/iterator"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	opt := NewDefaultOptions()
	db := Open(opt)
	defer func() { _ = db.Close() }()
	// 写入
	e := codec.NewEntry([]byte("hello"), []byte("coreKV")).WithTTL(1 * time.Second)
	if err := db.Set(e); err != nil {
		t.Fatal(err)
	}
	// 查询
	if entry, err := db.Get([]byte("hello")); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("db.Get key=%s, value=%s, expiresAt=%d", entry.Key, entry.Value, entry.ExpiresAt)
	}
	// 迭代器
	iter := iterator.NewIterator(&iterator.Options{
		Prefix: []byte("he"),
		IsAsc:  false,
	})
	defer func() { _ = iter.Close() }()
	for iter.Rewind(); iter.Valid(); iter.Next() {
		it := iter.Item()
		t.Logf("db.NewIterator key=%s, value=%s", it.Key(), it.Value())
	}
	t.Logf("db.Stats=%+v", db.Info())
	// 删除
	if err := db.Del([]byte("hello")); err != nil {
		t.Fatal(err)
	}
}