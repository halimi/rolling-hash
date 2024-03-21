package signature

import (
	"reflect"
	"testing"
)

func TestSignature_Hash(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{
			name:     "empty data",
			data:     []byte{},
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "test data",
			data:     []byte("test"),
			expected: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := NewSignature(1)
			got := sig.Hash(tt.data)
			if got != tt.expected {
				t.Errorf("got: %v, want: %v", got, tt.expected)
			}
		})
	}
}

func TestSignature_Signatures(t *testing.T) {
	tests := []struct {
		name      string
		chunkSize int
		data      []byte
		expected  map[string][]int
	}{
		{
			name:      "chunk size is equal with the data length",
			chunkSize: 4,
			data:      []byte("test"),
			expected:  map[string][]int{"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08": {0}},
		},
		{
			name:      "chunk size is half of the data length",
			chunkSize: 2,
			data:      []byte("test"),
			expected: map[string][]int{
				"2d6c9a90dd38f6852515274cde41a8cd8e7e1a7a053835334ec7e29f61b918dd": {0},
				"56af4bde70a47ae7d0f1ebb30e45ed336165d5c9ec00ba9a92311e33a4256d74": {2},
			},
		},
		{
			name:      "chunk size is one",
			chunkSize: 1,
			data:      []byte("test"),
			expected: map[string][]int{
				"e3b98a4da31a127d4bde6e43033f66ba274cab0eb7eb1c70ec41402bf6273dd8": {0, 3},
				"3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea": {1},
				"043a718774c572bd8a25adbeb1bfcd5c0256ae11cecf9f9c3f925d0e52beaf89": {2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := NewSignature(tt.chunkSize)
			got := sig.Signatures(tt.data)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, want: %v", got, tt.expected)
			}
		})
	}
}
