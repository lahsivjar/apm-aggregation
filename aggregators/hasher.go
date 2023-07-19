// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package aggregators

import (
	"hash"

	"github.com/cespare/xxhash/v2"
)

// Hashable represents the hash function interface implemented by aggregation models.
type Hashable interface {
	Hash(xxhash.Digest) xxhash.Digest
}

// HashablePB represents the hash function interface generated for proto definitions.
type HashablePB interface {
	HashPB(hash.Hash, map[string]struct{})
}

// Hasher contains a safe to copy digest.
type Hasher struct {
	digest xxhash.Digest // xxhash.Digest does not contain pointers and is safe to copy
}

// Chain allows chaining hash functions for Hashable interfaces.
func (h Hasher) Chain(hashable Hashable) Hasher {
	return Hasher{digest: hashable.Hash(h.digest)}
}

// ChainPB allows chaining of hash functions for proto models.
func (h Hasher) ChainPB(hashablePB HashablePB) Hasher {
	newH := Hasher{digest: h.digest}
	hashablePB.HashPB(&newH.digest, nil)
	return newH
}

// Sum returns the hash for all the chained interfaces.
func (h Hasher) Sum() uint64 {
	return h.digest.Sum64()
}
