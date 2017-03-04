// Package sakura provides a Sakura hash tree encoder as described in the paper
// "Sakura: a flexible coding for tree hashing". The paper is available here:
// http://keccak.noekeon.org/Sakura.pdf
package sakura

import (
	"errors"
	"hash"
	"io"
)

// Hasher provides a source of hash.Hash implementations.
type Hasher func() hash.Hash

// Filter returns a reader and writer pair that are capable of manipulating
// data before it is hashed.
type Filter func() (io.Reader, io.Writer)

// BlockSize represents a block size as a mantissa and exponent in the formula:
//
//   Pow(2, Exponent) * (2 * Mantissa + 1)
type BlockSize struct {
	Mantissa uint8
	Exponent uint8
}

// Value returns the block size as a total number of bytes.
func (bs BlockSize) Value() int {
	return 1 << bs.Exponent * (2*int(bs.Mantissa) + 1)
}

// HashingMode is a Sakura tree mode that describes how the tree is encoded.
type HashingMode struct {
	Hash       Hasher    // Source of hash.Hash implementations.
	Kangaroo   bool      // Does the mode apply Kangaroo hopping, wherein the first node is nested in its parent?
	Alignment  uint8     // The number of bytes that nodes will be aligned to.
	Interleave BlockSize // Block size for interleaving values.
}

// Hop is a hop in a hop tree.
//
// A hop must also implement either ChainingHop or MessageHop, but not both.
type Hop interface {
	// ChainingValue returns the already computed chaining value of the hop, which
	// is the output of a hash function.
	//
	// If the value has not yet been calculated, or if the implementation does
	// not cache chaining values, nil will be returned.
	ChainingValue() (hash []byte)

	// SetChainingValue provides the hop implementation with the calculated
	// chaining value.
	//
	// The implementation may retain this value and return it in future calls to
	// ChainingValue(), or the value may be discarded.
	SetChainingValue(hash []byte)
}

// ChainingHop is a source of chaining values.
type ChainingHop interface {
	// Child returns the child hop at index i.
	Child(i int) Hop
	// Degree returns the the number of children.
	Degree() int
}

// MessageHop is a source of message bits.
type MessageHop interface {
	io.Reader
}

// Encoder is a Sakura tree encoder.
type Encoder struct {
	mode HashingMode
	//pool bithash.Pool
}

// New returns a new encoder with the given hashing mode.
func New(mode HashingMode) *Encoder {
	return &Encoder{
		mode: mode,
	}
}

// Final encodes the given hop as a final node and returns the hash.
func (e *Encoder) Final(hop Hop) (hash []byte, err error) {
	return nil, errors.New("not implemented")
}

// Inner encodes the given hop as an inner node and returns the hash.
func (e *Encoder) Inner(hop Hop) (hash []byte, err error) {
	return nil, errors.New("not implemented")
}
