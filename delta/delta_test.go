package delta

import (
	"reflect"
	"testing"

	"github.com/halimi/rolling-hash/signature"
)

func TestDelta_Diff(t *testing.T) {
	tests := []struct {
		name         string
		chunkSize    int
		originalData []byte
		updatedData  []byte
		expected     map[int][]byte
	}{
		{
			name:         "no differences",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("test"),
			expected:     map[int][]byte{},
		},
		{
			name:         "longer with one",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("testx"),
			expected:     map[int][]byte{4: []byte("x")},
		},
		{
			name:         "shorter with one",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("tes"),
			expected:     map[int][]byte{3: {}},
		},
		{
			name:         "one difference",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("txst"),
			expected:     map[int][]byte{1: []byte("x")},
		},
		{
			name:         "two differences",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("txsx"),
			expected:     map[int][]byte{1: []byte("x"), 3: []byte("x")},
		},
		{
			name:         "shifted to the right",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("txest"),
			expected:     map[int][]byte{1: []byte("x"), 2: []byte("e"), 3: []byte("s"), 4: []byte("t")},
		},
		{
			name:         "shifted to the left",
			chunkSize:    1,
			originalData: []byte("test"),
			updatedData:  []byte("esxt"),
			expected:     map[int][]byte{0: []byte("e"), 1: []byte("s"), 2: []byte("x")},
		},
		{
			name:         "chunk size is half of the data length",
			chunkSize:    2,
			originalData: []byte("test"),
			updatedData:  []byte("etts"),
			expected:     map[int][]byte{0: []byte("et"), 2: []byte("ts")},
		},
		{
			name:         "chunk size is equal with the data length",
			chunkSize:    4,
			originalData: []byte("test"),
			updatedData:  []byte("tets"),
			expected:     map[int][]byte{0: []byte("tets")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDelta(tt.chunkSize)
			s := signature.NewSignature(tt.chunkSize)
			delta, err := d.Diff(s.Signatures(tt.originalData), tt.updatedData)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			if !reflect.DeepEqual(delta, tt.expected) {
				t.Errorf("got = %v, want %v", delta, tt.expected)
			}
		})
	}
}
