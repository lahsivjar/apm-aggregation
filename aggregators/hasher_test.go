// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package aggregators

import (
	"hash"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cespare/xxhash/v2"
)

type testHashable func(xxhash.Digest) xxhash.Digest

func (f testHashable) Hash(h xxhash.Digest) xxhash.Digest {
	return f(h)
}

type testHashablePB func(hash.Hash)

func (f testHashablePB) HashPB(h hash.Hash, _ map[string]struct{}) {
	f(h)
}

func TestHasher(t *testing.T) {
	a := Hasher{}
	b := a.Chain(testHashable(func(h xxhash.Digest) xxhash.Digest {
		h.WriteString("1")
		return h
	}))
	c := a.Chain(testHashable(func(h xxhash.Digest) xxhash.Digest {
		h.WriteString("1")
		return h
	}))
	assert.NotEqual(t, a, b)
	assert.Equal(t, b, c)

	// Ensure the struct does not change after calling Sum
	c.Sum()
	assert.Equal(t, b, c)
}

func TestHasherPB(t *testing.T) {
	base := Hasher{}
	h1 := base.ChainPB(testHashablePB(func(h hash.Hash) {
		h.Write([]byte("1"))
	}))
	h2 := h1.ChainPB(testHashablePB(func(h hash.Hash) {
		h.Write([]byte("2"))
	}))
	h1Copy := Hasher{}.ChainPB(testHashablePB(func(h hash.Hash) {
		h.Write([]byte("1"))
	}))

	assert.Equal(t, h1, h1Copy)
	assert.Equal(t, h1.Sum(), h1Copy.Sum())
	assert.NotEqual(t, h1, h2)
	assert.NotEqual(t, h1.Sum(), h2.Sum())
}

func BenchmarkHasher(b *testing.B) {
	keys := getTestHashables()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h := Hasher{}
		for _, k := range keys {
			h.Chain(k)
		}
		h.Sum()
	}
}

func BenchmarkHasherPB(b *testing.B) {
	keys := getTestHashablePBs()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h := Hasher{}
		for _, k := range keys {
			h.ChainPB(k)
		}
		h.Sum()
	}
}

func getTestHashablePBs() []HashablePB {
	key1, key2 := getTestKeys()
	return []HashablePB{key1.ToProto(), key2.ToProto()}
}

func getTestHashables() []Hashable {
	key1, key2 := getTestKeys()
	return []Hashable{key1, key2}
}

func getTestKeys() (ServiceAggregationKey, TransactionAggregationKey) {
	key1 := ServiceAggregationKey{
		Timestamp:           time.Now(),
		ServiceName:         "test",
		ServiceEnvironment:  "testing",
		ServiceLanguageName: "go",
		AgentName:           "go-agent",
	}
	key2 := TransactionAggregationKey{
		TraceRoot:              false,
		ContainerID:            "test-container",
		KubernetesPodName:      "test-pod",
		ServiceVersion:         "v0",
		ServiceNodeName:        "test-node",
		ServiceRuntimeName:     "test-runtime",
		ServiceRuntimeVersion:  "v0",
		ServiceLanguageVersion: "v0",
		HostHostname:           "test-host",
		HostName:               "test-hostname",
		HostOSPlatform:         "test-hostos",
		EventOutcome:           "success",
		TransactionName:        "testtxn",
		TransactionType:        "testtype",
		TransactionResult:      "success",
		FAASColdstart:          123,
		FAASID:                 "test-faasid",
		FAASName:               "test-fassname",
		FAASVersion:            "v0",
		FAASTriggerType:        "test-trigger",
		CloudProvider:          "test-cloud",
		CloudRegion:            "test-region",
		CloudAvailabilityZone:  "test-zone",
		CloudServiceName:       "test-svc",
		CloudAccountID:         "test-acc",
		CloudAccountName:       "test-accname",
		CloudMachineType:       "test-machine",
		CloudProjectID:         "test-prj",
		CloudProjectName:       "test-prjname",
	}
	return key1, key2
}
