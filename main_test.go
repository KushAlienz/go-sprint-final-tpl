package main

import "testing"

// ─── Тесты generateRandomElements ─────────────────────────────────────

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -10, 0},
		{"один элемент", 1, 1},
		{"маленький слайс", 5, 5},
		{"средний слайс", 1000, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("длина = %d, хотим %d", len(got), tt.want)
			}
			for i, v := range got {
				if v < 0 {
					t.Errorf("элемент %d отрицательный: %d", i, v)
				}
			}
		})
	}
}

// ─── Тесты maximum ────────────────────────────────────────────────────

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"все одинаковые", []int{7, 7, 7, 7}, 7},
		{"возрастающие", []int{1, 2, 3, 4, 5}, 5},
		{"убывающие", []int{9, 8, 7, 6, 5}, 9},
		{"максимум в середине", []int{1, 100, 3, 50, 7}, 100},
		{"максимум в начале", []int{1000, 1, 2, 3}, 1000},
		{"максимум в конце", []int{1, 2, 3, 1000}, 1000},
		{"два максимума", []int{10, 5, 10, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, хотим %d", tt.data, got, tt.want)
			}
		})
	}
}

// ─── Тесты maxChunks ──────────────────────────────────────────────────

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"короче CHUNKS", []int{1, 5, 3}, 5},
		{"ровно CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"кратно CHUNKS", makeRange(1, 64), 64},
		{"не кратно CHUNKS", makeRange(1, 100), 100},
		{"макс в первом чанке", []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 100},
		{"макс в последнем чанке", []int{1, 2, 3, 4, 5, 6, 7, 8, 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, хотим %d", tt.data, got, tt.want)
			}
		})
	}
}

// makeRange возвращает слайс [from, from+1, ..., from+count-1].
func makeRange(from, count int) []int {
	out := make([]int, count)
	for i := 0; i < count; i++ {
		out[i] = from + i
	}
	return out
}

// ─── Сравнение результата однопоточного и параллельного поиска ────────

func TestMaxChunksMatchesMaximum(t *testing.T) {
	sizes := []int{0, 1, 5, 8, 9, 100, 1000, 100_000}
	for _, size := range sizes {
		data := generateRandomElements(size)
		one := maximum(data)
		multi := maxChunks(data)
		if one != multi {
			t.Errorf("размер %d: maximum = %d, maxChunks = %d", size, one, multi)
		}
	}
}
