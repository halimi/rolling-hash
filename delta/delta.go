package delta

import (
	"errors"

	"github.com/halimi/rolling-hash/signature"
)

type Delta struct {
	chunkSize int
	chunks    map[int][]byte
}

func NewDelta(chunkSize int) *Delta {
	return &Delta{
		chunkSize: chunkSize,
		chunks:    make(map[int][]byte),
	}
}

func (d *Delta) Diff(signatures map[string][]int, data []byte) (map[int][]byte, error) {
	if len(signatures) == 0 || len(data) == 0 {
		return d.chunks, errors.New("empty data")
	}

	sig := signature.NewSignature(d.chunkSize)

	for i := 0; i < len(data); i += d.chunkSize {
		end := i + d.chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		hash := sig.Hash(chunk)

		if chunkList, ok := signatures[hash]; ok {
			// looking for existing elements
			found := false
			for k, v := range chunkList {
				if v == i {
					found = true
					// remove the found element from the list
					chunkList = append(chunkList[:k], chunkList[k+1:]...)
				}
			}

			// not found, new element
			if !found {
				d.chunks[i] = chunk
			}

			// remove the already found elements
			if len(chunkList) > 0 {
				signatures[hash] = chunkList
			} else {
				delete(signatures, hash)
			}

		} else {
			d.chunks[i] = chunk
		}
	}

	// add removed elements to the delta list
	if len(signatures) > 0 {
		for _, leftovers := range signatures {
			for _, v := range leftovers {
				if _, ok := d.chunks[v]; !ok {
					d.chunks[v] = []byte{}
				}
			}
		}
	}

	return d.chunks, nil
}
