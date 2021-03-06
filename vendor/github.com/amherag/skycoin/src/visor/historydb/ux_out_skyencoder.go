// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package historydb

import (
	"errors"
	"math"

	"github.com/amherag/skycoin/src/cipher/encoder"
)

// encodeSizeUxOut computes the size of an encoded object of type UxOut
func encodeSizeUxOut(obj *UxOut) uint64 {
	i0 := uint64(0)

	// obj.Out.Head.Time
	i0 += 8

	// obj.Out.Head.BkSeq
	i0 += 8

	// obj.Out.Body.SrcTransaction
	i0 += 32

	// obj.Out.Body.Address.Version
	i0++

	// obj.Out.Body.Address.Key
	i0 += 20

	// obj.Out.Body.Coins
	i0 += 8

	// obj.Out.Body.Hours
	i0 += 8

	// obj.Out.Body.ProgramState
	i0 += 4 + uint64(len(obj.Out.Body.ProgramState))

	// obj.SpentTxnID
	i0 += 32

	// obj.SpentBlockSeq
	i0 += 8

	return i0
}

// encodeUxOut encodes an object of type UxOut to a buffer allocated to the exact size
// required to encode the object.
func encodeUxOut(obj *UxOut) ([]byte, error) {
	n := encodeSizeUxOut(obj)
	buf := make([]byte, n)

	if err := encodeUxOutToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeUxOutToBuffer encodes an object of type UxOut to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeUxOutToBuffer(buf []byte, obj *UxOut) error {
	if uint64(len(buf)) < encodeSizeUxOut(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Out.Head.Time
	e.Uint64(obj.Out.Head.Time)

	// obj.Out.Head.BkSeq
	e.Uint64(obj.Out.Head.BkSeq)

	// obj.Out.Body.SrcTransaction
	e.CopyBytes(obj.Out.Body.SrcTransaction[:])

	// obj.Out.Body.Address.Version
	e.Uint8(obj.Out.Body.Address.Version)

	// obj.Out.Body.Address.Key
	e.CopyBytes(obj.Out.Body.Address.Key[:])

	// obj.Out.Body.Coins
	e.Uint64(obj.Out.Body.Coins)

	// obj.Out.Body.Hours
	e.Uint64(obj.Out.Body.Hours)

	// obj.Out.Body.ProgramState length check
	if uint64(len(obj.Out.Body.ProgramState)) > math.MaxUint32 {
		return errors.New("obj.Out.Body.ProgramState length exceeds math.MaxUint32")
	}

	// obj.Out.Body.ProgramState length
	e.Uint32(uint32(len(obj.Out.Body.ProgramState)))

	// obj.Out.Body.ProgramState copy
	e.CopyBytes(obj.Out.Body.ProgramState)

	// obj.SpentTxnID
	e.CopyBytes(obj.SpentTxnID[:])

	// obj.SpentBlockSeq
	e.Uint64(obj.SpentBlockSeq)

	return nil
}

// decodeUxOut decodes an object of type UxOut from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeUxOut(buf []byte, obj *UxOut) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Out.Head.Time
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Out.Head.Time = i
	}

	{
		// obj.Out.Head.BkSeq
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Out.Head.BkSeq = i
	}

	{
		// obj.Out.Body.SrcTransaction
		if len(d.Buffer) < len(obj.Out.Body.SrcTransaction) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Out.Body.SrcTransaction[:], d.Buffer[:len(obj.Out.Body.SrcTransaction)])
		d.Buffer = d.Buffer[len(obj.Out.Body.SrcTransaction):]
	}

	{
		// obj.Out.Body.Address.Version
		i, err := d.Uint8()
		if err != nil {
			return 0, err
		}
		obj.Out.Body.Address.Version = i
	}

	{
		// obj.Out.Body.Address.Key
		if len(d.Buffer) < len(obj.Out.Body.Address.Key) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Out.Body.Address.Key[:], d.Buffer[:len(obj.Out.Body.Address.Key)])
		d.Buffer = d.Buffer[len(obj.Out.Body.Address.Key):]
	}

	{
		// obj.Out.Body.Coins
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Out.Body.Coins = i
	}

	{
		// obj.Out.Body.Hours
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Out.Body.Hours = i
	}

	{
		// obj.Out.Body.ProgramState

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Out.Body.ProgramState = make([]byte, length)

			copy(obj.Out.Body.ProgramState[:], d.Buffer[:length])
			d.Buffer = d.Buffer[length:]
		}
	}

	{
		// obj.SpentTxnID
		if len(d.Buffer) < len(obj.SpentTxnID) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.SpentTxnID[:], d.Buffer[:len(obj.SpentTxnID)])
		d.Buffer = d.Buffer[len(obj.SpentTxnID):]
	}

	{
		// obj.SpentBlockSeq
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.SpentBlockSeq = i
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeUxOutExact decodes an object of type UxOut from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeUxOutExact(buf []byte, obj *UxOut) error {
	if n, err := decodeUxOut(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
