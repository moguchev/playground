package bench

import (
	"testing"

	"github.com/lib/pq"
)

var a = []int64{25955, 44616, 44875, 45004, 45297, 45720, 45721, 45739, 47737, 48135, 51396, 51398, 51404, 51405, 51406, 51407, 51409, 51415, 51416, 51419, 51421, 51422, 51430, 51431, 51435, 51440, 51449, 51451, 51452, 51453, 51454, 51461, 51462, 51908, 51910, 51912, 51915, 51916, 51920, 51924, 51926, 51928, 51933, 51950, 51954, 51960, 51961, 51962, 51963, 51974,
	25955, 44616, 44875, 45004, 45297, 45720, 45721, 45739, 47737, 48135, 51396, 51398, 51404, 51405, 51406, 51407, 51409, 51415, 51416, 51419, 51421, 51422, 51430, 51431, 51435, 51440, 51449, 51451, 51452, 51453, 51454, 51461, 51462, 51908, 51910, 51912, 51915, 51916, 51920, 51924, 51926, 51928, 51933, 51950, 51954, 51960, 51961, 51962, 51963, 51974}

func BenchmarkPqArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := pq.Array(a).Value(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCustomArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ArrayInt64(a); err != nil {
			b.Fatal(err)
		}
	}
}
