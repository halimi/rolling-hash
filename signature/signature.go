package signature

import (
	"crypto/sha256"
	"encoding/hex"
)

type Signature struct {
	chunkSize  int
	signatures map[string][]int
}

func NewSignature(chunkSize int) *Signature {
	return &Signature{
		chunkSize:  chunkSize,
		signatures: make(map[string][]int),
	}
}

func (s *Signature) Hash(data []byte) string {
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

func (s *Signature) Signatures(data []byte) map[string][]int {
	for i := 0; i < len(data); i += s.chunkSize {
		end := i + s.chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		hash := s.Hash(chunk)
		if chunkList, ok := s.signatures[hash]; ok {
			chunkList = append(chunkList, i)
			s.signatures[hash] = chunkList
		} else {
			s.signatures[hash] = []int{i}
		}
	}

	return s.signatures
}
