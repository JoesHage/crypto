// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong64 && !purego && gc

package salsa

import "golang.org/x/sys/cpu"

// XORKeyStreamVX is implemented in salsa20_loong64.s.
//
//go:noescape
func XORKeyStreamVX(out, in *byte, n uint64, nonce, key *byte)

// XORKeyStream crypts bytes from in to out using the given key and counters.
// In and out must overlap entirely or not at all. Counter
// contains the raw salsa20 counter bytes (both nonce and block counter).
func XORKeyStream(out, in []byte, counter *[16]byte, key *[32]byte) {
	if len(in) == 0 {
		return
	}
	_ = out[len(in)-1]
	if cpu.Loong64.HasLSX {
		XORKeyStreamVX(&out[0], &in[0], uint64(len(in)), &counter[0], &key[0])
	} else {
		genericXORKeyStream(out, in, counter, key)
	}
}
