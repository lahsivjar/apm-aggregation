// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.2.0
// Source: proto/labels.proto

package aggregationpb

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *GlobalLabels) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		elastic_apm_GlobalLabels_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Label) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		elastic_apm_Label_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *NumericLabel) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		elastic_apm_NumericLabel_hashpb_sum(m, hasher, ignore)
	}
}