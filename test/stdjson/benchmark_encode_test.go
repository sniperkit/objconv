package json

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/segmentio/objconv/test"
)

var (
	benchmarkEncoder = json.NewEncoder(ioutil.Discard)
)

func BenchmarkEncodeNil(b *testing.B) { test.BenchmarkEncodeNil(b, benchmarkEncoder) }

func BenchmarkEncodeBoolTrue(b *testing.B)  { test.BenchmarkEncodeBoolTrue(b, benchmarkEncoder) }
func BenchmarkEncodeBoolFalse(b *testing.B) { test.BenchmarkEncodeBoolFalse(b, benchmarkEncoder) }

func BenchmarkEncodeIntZero(b *testing.B)  { test.BenchmarkEncodeIntZero(b, benchmarkEncoder) }
func BenchmarkEncodeIntShort(b *testing.B) { test.BenchmarkEncodeIntShort(b, benchmarkEncoder) }
func BenchmarkEncodeIntLong(b *testing.B)  { test.BenchmarkEncodeIntLong(b, benchmarkEncoder) }

func BenchmarkEncodeUintZero(b *testing.B)  { test.BenchmarkEncodeUintZero(b, benchmarkEncoder) }
func BenchmarkEncodeUintShort(b *testing.B) { test.BenchmarkEncodeUintShort(b, benchmarkEncoder) }
func BenchmarkEncodeUintLong(b *testing.B)  { test.BenchmarkEncodeUintLong(b, benchmarkEncoder) }

func BenchmarkEncodeFloatZero(b *testing.B)  { test.BenchmarkEncodeFloatZero(b, benchmarkEncoder) }
func BenchmarkEncodeFloatShort(b *testing.B) { test.BenchmarkEncodeFloatShort(b, benchmarkEncoder) }
func BenchmarkEncodeFloatLong(b *testing.B)  { test.BenchmarkEncodeFloatLong(b, benchmarkEncoder) }

func BenchmarkEncodeStringEmpty(b *testing.B) { test.BenchmarkEncodeStringEmpty(b, benchmarkEncoder) }
func BenchmarkEncodeStringShort(b *testing.B) { test.BenchmarkEncodeStringShort(b, benchmarkEncoder) }
func BenchmarkEncodeStringLong(b *testing.B)  { test.BenchmarkEncodeStringLong(b, benchmarkEncoder) }

func BenchmarkEncodeBytesEmpty(b *testing.B) { test.BenchmarkEncodeBytesEmpty(b, benchmarkEncoder) }
func BenchmarkEncodeBytesShort(b *testing.B) { test.BenchmarkEncodeBytesShort(b, benchmarkEncoder) }
func BenchmarkEncodeBytesLong(b *testing.B)  { test.BenchmarkEncodeBytesLong(b, benchmarkEncoder) }

func BenchmarkEncodeSliceInterfaceEmpty(b *testing.B) {
	test.BenchmarkEncodeSliceInterfaceEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceInterfaceShort(b *testing.B) {
	test.BenchmarkEncodeSliceInterfaceShort(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceInterfaceLong(b *testing.B) {
	test.BenchmarkEncodeSliceInterfaceLong(b, benchmarkEncoder)
}

func BenchmarkEncodeSliceStringEmpty(b *testing.B) {
	test.BenchmarkEncodeSliceStringEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceStringShort(b *testing.B) {
	test.BenchmarkEncodeSliceStringShort(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceStringLong(b *testing.B) {
	test.BenchmarkEncodeSliceStringLong(b, benchmarkEncoder)
}

func BenchmarkEncodeSliceBytesEmpty(b *testing.B) {
	test.BenchmarkEncodeSliceBytesEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceBytesShort(b *testing.B) {
	test.BenchmarkEncodeSliceBytesShort(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceBytesLong(b *testing.B) {
	test.BenchmarkEncodeSliceBytesLong(b, benchmarkEncoder)
}

func BenchmarkEncodeSliceStructEmpty(b *testing.B) {
	test.BenchmarkEncodeSliceStructEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceStructShort(b *testing.B) {
	test.BenchmarkEncodeSliceStructShort(b, benchmarkEncoder)
}
func BenchmarkEncodeSliceStructLong(b *testing.B) {
	test.BenchmarkEncodeSliceStructLong(b, benchmarkEncoder)
}

func BenchmarkEncodeMapStringStringEmpty(b *testing.B) {
	test.BenchmarkEncodeMapStringStringEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringStringShort(b *testing.B) {
	test.BenchmarkEncodeMapStringStringShort(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringStringLong(b *testing.B) {
	test.BenchmarkEncodeMapStringStringLong(b, benchmarkEncoder)
}

func BenchmarkEncodeMapStringInterfaceEmpty(b *testing.B) {
	test.BenchmarkEncodeMapStringInterfaceEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringInterfaceShort(b *testing.B) {
	test.BenchmarkEncodeMapStringInterfaceShort(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringInterfaceLong(b *testing.B) {
	test.BenchmarkEncodeMapStringInterfaceLong(b, benchmarkEncoder)
}

func BenchmarkEncodeMapStringStructEmpty(b *testing.B) {
	test.BenchmarkEncodeMapStringStructEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringStructShort(b *testing.B) {
	test.BenchmarkEncodeMapStringStructShort(b, benchmarkEncoder)
}
func BenchmarkEncodeMapStringStructLong(b *testing.B) {
	test.BenchmarkEncodeMapStringStructLong(b, benchmarkEncoder)
}

func BenchmarkEncodeMapSliceEmpty(b *testing.B) {
	test.BenchmarkEncodeMapSliceEmpty(b, benchmarkEncoder)
}
func BenchmarkEncodeMapSliceShort(b *testing.B) {
	test.BenchmarkEncodeMapSliceShort(b, benchmarkEncoder)
}
func BenchmarkEncodeMapSliceLong(b *testing.B) {
	test.BenchmarkEncodeMapSliceLong(b, benchmarkEncoder)
}

func BenchmarkEncodeStructEmpty(b *testing.B) { test.BenchmarkEncodeStructEmpty(b, benchmarkEncoder) }
func BenchmarkEncodeStructShort(b *testing.B) { test.BenchmarkEncodeStructShort(b, benchmarkEncoder) }
func BenchmarkEncodeStructLong(b *testing.B)  { test.BenchmarkEncodeStructLong(b, benchmarkEncoder) }