package tests

import (
	"errors"
	"fmt"
	"github.com/ArmiKGG/in_memory_cache/internal/cache"
	"testing"
)

func TestSet(t *testing.T) {
	memoryCache := cache.New()
	result := memoryCache.Set("unknownKey", map[string]any{"test_1": 1, "test_2": true})

	if !errors.Is(nil, result) {
		t.Errorf("Expected %v, but got %v", nil, result)
	}

	result = memoryCache.Set("unknownKey", map[string]any{"test_42": 12, "test_23": false})

	expected := fmt.Errorf("value related with this <%s> key founded", "unknownKey")
	if !(expected.Error() == result.Error()) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGet(t *testing.T) {
	memoryCache := cache.New()
	_ = memoryCache.Set("unknownKey", "Hey")

	val, err := memoryCache.Get("unknownKey")

	result := "Hey"

	if !errors.Is(nil, err) || (val != result) {
		t.Errorf("Expected %v, but got %v", nil, result)
	}

}

func TestDelete(t *testing.T) {
	memoryCache := cache.New()
	_ = memoryCache.Set("unknownKey", "Hey")
	err := memoryCache.Delete("unknownKey")
	if err != nil {
		t.Errorf("Expected %v, but got %v", nil, err)
	}
}
