package colfer

import (
	"io/ioutil"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/pascaldekloe/colfer/testdata/bench"
)

//go:generate go run ./cmd/colf/main.go -p testdata go testdata/bench/scheme.colf
//go:generate protoc --gogofaster_out=. -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/gogo/protobuf/protobuf testdata/bench/scheme.proto
//go:generate flatc -o testdata -g testdata/bench/scheme.fbs

func newTestData(tb testing.TB) []*bench.Colfer {
	return []*bench.Colfer{
		{Key: 1234567890, Host: "db003lz12", Port: 389, Size: 452, Hash: 0x488b5c2428488918, Ratio: 0.99, Route: true},
		{Key: 1234567891, Host: "localhost", Port: 22, Size: 4096, Hash: 0x243048899c24c824, Ratio: 0.20, Route: false},
		{Key: 1234567892, Host: "kdc.local", Port: 88, Size: 1984, Hash: 0x000048891c24485c, Ratio: 0.06, Route: false},
		{Key: 1234567893, Host: "vhost8.dmz.example.com", Port: 27017, Size: 59741, Hash: 0x5c2408488b9c2489, Ratio: 0.0, Route: true},
	}
}

func newProtoBufData(tb testing.TB) []*bench.ProtoBuf {
	testData := newTestData(tb)
	protoBufData := make([]*bench.ProtoBuf, len(testData))
	for i, o := range testData {
		protoBufData[i] = &bench.ProtoBuf{
			Key:   o.Key,
			Host:  o.Host,
			Port:  o.Port,
			Size_: o.Size,
			Hash:  o.Hash,
			Ratio: o.Ratio,
			Route: o.Route,
		}
	}
	return protoBufData
}

// prevent compiler optimization
var (
	holdSerials      [][]byte
	holdData         []*bench.Colfer
	holdProtoBufData []*bench.ProtoBuf
)

func BenchmarkMarshal(b *testing.B) {
	testData := newTestData(b)
	holdSerials = make([][]byte, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdSerials {
		var err error
		holdSerials[i], err = testData[i%len(testData)].MarshalBinary()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalProtoBuf(b *testing.B) {
	testData := newProtoBufData(b)
	holdSerials = make([][]byte, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdSerials {
		var err error
		holdSerials[i], err = testData[i%len(testData)].Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalFlatBuf(b *testing.B) {
	testData := newTestData(b)
	holdSerials = make([][]byte, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdSerials {
		o := testData[i%len(testData)]

		builder := flatbuffers.NewBuilder(0)
		host := builder.CreateString(o.Host)
		bench.FlatBuffersStart(builder)
		bench.FlatBuffersAddKey(builder, o.Key)
		bench.FlatBuffersAddHost(builder, host)
		bench.FlatBuffersAddPort(builder, o.Port)
		bench.FlatBuffersAddSize(builder, o.Size)
		bench.FlatBuffersAddHash(builder, o.Hash)
		bench.FlatBuffersAddRatio(builder, o.Ratio)
		if o.Route {
			bench.FlatBuffersAddRoute(builder, 1)
		} else {
			bench.FlatBuffersAddRoute(builder, 0)
		}
		builder.Finish(bench.FlatBuffersEnd(builder))

		holdSerials[i] = builder.Bytes[builder.Head():]
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	testData := newTestData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		serials[i], _ = o.MarshalBinary()
	}

	holdData = make([]*bench.Colfer, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdData {
		o := new(bench.Colfer)
		holdData[i] = o

		err := o.UnmarshalBinary(serials[i%len(serials)])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalProtoBuf(b *testing.B) {
	testData := newProtoBufData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		var err error
		serials[i], err = o.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}

	holdProtoBufData = make([]*bench.ProtoBuf, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdProtoBufData {
		o := new(bench.ProtoBuf)
		holdProtoBufData[i] = o

		err := o.Unmarshal(serials[i%len(serials)])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalFlatBuf(b *testing.B) {
	testData := newTestData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		builder := flatbuffers.NewBuilder(0)
		host := builder.CreateString(o.Host)
		bench.FlatBuffersStart(builder)
		bench.FlatBuffersAddKey(builder, o.Key)
		bench.FlatBuffersAddHost(builder, host)
		bench.FlatBuffersAddPort(builder, o.Port)
		bench.FlatBuffersAddSize(builder, o.Size)
		bench.FlatBuffersAddHash(builder, o.Hash)
		bench.FlatBuffersAddRatio(builder, o.Ratio)
		if o.Route {
			bench.FlatBuffersAddRoute(builder, 1)
		} else {
			bench.FlatBuffersAddRoute(builder, 0)
		}
		builder.Finish(bench.FlatBuffersEnd(builder))
		serials[i] = builder.FinishedBytes()
	}

	holdData = make([]*bench.Colfer, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := range holdData {
		o := new(bench.Colfer)
		holdData[i] = o

		bytes := serials[i%len(serials)]
		buf := new(bench.FlatBuffers)
		buf.Init(bytes, flatbuffers.GetUOffsetT(bytes))
		o.Key = buf.Key()
		o.Host = string(buf.Host())
		o.Port = buf.Port()
		o.Size = buf.Size()
		o.Hash = buf.Hash()
		o.Ratio = buf.Ratio()
		o.Route = buf.Route() == 1
	}
}

func BenchmarkMarshalReuse(b *testing.B) {
	testData := newTestData(b)
	var buf []byte

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		o := testData[i%len(testData)]

		l, err := o.MarshalLen()
		if err != nil {
			b.Fatal(err)
		}
		if l > len(buf) {
			buf = make([]byte, l+100)
		}

		o.MarshalTo(buf)
		ioutil.Discard.Write(buf[:l])
	}
}

func BenchmarkMarshalProtoBufReuse(b *testing.B) {
	testData := newProtoBufData(b)
	var buf []byte

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		o := testData[i%len(testData)]

		l := o.Size()
		if l > len(buf) {
			buf = make([]byte, l+100)
		}

		_, err := o.MarshalTo(buf)
		if err != nil {
			b.Fatal(err)
		}
		ioutil.Discard.Write(buf[:l])
	}
}

func BenchmarkMarshalFlatBufReuse(b *testing.B) {
	testData := newTestData(b)
	builder := flatbuffers.NewBuilder(0)

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		o := testData[i%len(testData)]

		builder.Reset()
		host := builder.CreateString(o.Host)
		bench.FlatBuffersStart(builder)
		bench.FlatBuffersAddKey(builder, o.Key)
		bench.FlatBuffersAddHost(builder, host)
		bench.FlatBuffersAddPort(builder, o.Port)
		bench.FlatBuffersAddSize(builder, o.Size)
		bench.FlatBuffersAddHash(builder, o.Hash)
		bench.FlatBuffersAddRatio(builder, o.Ratio)
		if o.Route {
			bench.FlatBuffersAddRoute(builder, 1)
		} else {
			bench.FlatBuffersAddRoute(builder, 0)
		}
		builder.Finish(bench.FlatBuffersEnd(builder))

		ioutil.Discard.Write(builder.Bytes[builder.Head():])
	}
}

func BenchmarkUnmarshalReuse(b *testing.B) {
	testData := newTestData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		serials[i], _ = o.MarshalBinary()
	}

	o := new(bench.Colfer)

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		err := o.UnmarshalBinary(serials[i%len(serials)])
		if err != nil {
			b.Fatal(err)
		}
	}

	holdData = []*bench.Colfer{o}
}

func BenchmarkUnmarshalProtoBufReuse(b *testing.B) {
	testData := newProtoBufData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		var err error
		serials[i], err = o.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}

	o := new(bench.ProtoBuf)

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		err := o.Unmarshal(serials[i%len(serials)])
		if err != nil {
			b.Fatal(err)
		}
	}

	holdProtoBufData = []*bench.ProtoBuf{o}
}

func BenchmarkUnmarshalFlatBufReuse(b *testing.B) {
	testData := newTestData(b)
	serials := make([][]byte, len(testData))
	for i, o := range testData {
		builder := flatbuffers.NewBuilder(0)
		host := builder.CreateString(o.Host)
		bench.FlatBuffersStart(builder)
		bench.FlatBuffersAddKey(builder, o.Key)
		bench.FlatBuffersAddHost(builder, host)
		bench.FlatBuffersAddPort(builder, o.Port)
		bench.FlatBuffersAddSize(builder, o.Size)
		bench.FlatBuffersAddHash(builder, o.Hash)
		bench.FlatBuffersAddRatio(builder, o.Ratio)
		if o.Route {
			bench.FlatBuffersAddRoute(builder, 1)
		} else {
			bench.FlatBuffersAddRoute(builder, 0)
		}
		builder.Finish(bench.FlatBuffersEnd(builder))
		serials[i] = builder.FinishedBytes()
	}

	o := new(bench.Colfer)
	buf := new(bench.FlatBuffers)

	b.ReportAllocs()
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		bytes := serials[i%len(serials)]
		buf.Init(bytes, flatbuffers.GetUOffsetT(bytes))
		o.Key = buf.Key()
		o.Host = string(buf.Host())
		o.Port = buf.Port()
		o.Size = buf.Size()
		o.Hash = buf.Hash()
		o.Ratio = buf.Ratio()
		o.Route = buf.Route() == 1
	}
	holdData = []*bench.Colfer{o}
}
