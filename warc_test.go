package warc_test

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/edsu/warc"
)

func TestReader(t *testing.T) {
	f, _ := os.Open("testdata/test.warc.gz")
	reader := warc.NewWarcReader(f)
	count := 0
	for {
		rec, err := reader.ReadRecord()
		if rec == nil {
			break
		}
		if err != nil {
			t.Error("received error from ReadRecord", err)
		}
		count += 1
	}
	assert.Equal(t, count, 10)
}

func TestRecord(t *testing.T) {
	f, _ := os.Open("testdata/test.warc.gz")
	reader := warc.NewWarcReader(f)
	rec, err := reader.ReadRecord()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, rec.Version, "WARC/1.0")
	assert.Equal(t, len(rec.Headers), 7)
	assert.Equal(t, rec.Headers["content-length"], "251")
}
