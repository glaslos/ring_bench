package ring

import (
	"os"
	"testing"

	bmharper "github.com/bmharper/ringbuffer"
	cloudflare "github.com/cloudflare/buffer"
	dullgiulio "github.com/dullgiulio/ringbuf"
	ericaro "github.com/ericaro/ringbuffer"
	mediafly "github.com/mediafly/ringbuffer"
	textnode "github.com/textnode/gringo"
	eapache "gopkg.in/eapache/queue.v1"
)

func BenchmarkEricaro(b *testing.B) {
	q := ericaro.New(10)
	for n := 0; n < b.N; n++ {
		q.Add(1)
		q.Remove(1)
	}
}

func BenchmarkEapache(b *testing.B) {
	q := eapache.New()
	for n := 0; n < b.N; n++ {
		q.Add("123")
		q.Remove()
	}
}

func BenchmarkDullgiulio(b *testing.B) {
	r := dullgiulio.NewRingbufBytes(64)
	reader := dullgiulio.NewReaderBytes(r)
	go r.Ringbuf().Run()
	for n := 0; n < b.N; n++ {
		r.Write([]byte("123"))
		buf := make([]byte, 3)
		reader.Read(buf)
	}
}

func BenchmarkGringb(b *testing.B) {
	r := textnode.NewGringo()
	pl := *textnode.NewPayload(1)
	for n := 0; n < b.N; n++ {
		r.Write(pl)
		r.Read()
	}
}

func BenchmarkMediafly(b *testing.B) {
	r := mediafly.NewStringBuffer(10)
	for n := 0; n < b.N; n++ {
		r.Append("123")
	}
}

func BenchmarkBMHarper(b *testing.B) {
	r := bmharper.Ring{}
	for n := 0; n < b.N; n++ {
		r.Write([]byte("123"))
		buf := make([]byte, 3)
		r.Read(buf)
	}
}

func BenchmarkCloudflare(b *testing.B) {
	buf, err := cloudflare.New("123", 80*12)
	if err != nil {
		b.Error(err)
	}
	defer os.Remove("123")
	for n := 0; n < b.N; n++ {
		if err = buf.Insert([]byte("123")); err != nil {
			b.Error(err)
		}
		_, err = buf.Pop()
		if err != nil {
			b.Error(err)
		}
	}
}
