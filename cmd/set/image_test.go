package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageType(t *testing.T) {
	testCases := []struct {
		name             string
		withAbsolutePath string
		expectType       string
	}{
		{
			name:             "Should work for a realistic path with a PNG extension",
			withAbsolutePath: "/home/user/Pictures/image.png",
			expectType:       "png",
		},
		{
			name:             "Should work for a realistic path with a JPG extension",
			withAbsolutePath: "/home/user/Pictures/image.jpg",
			expectType:       "jpg",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			img := image{AbsolutePath: tc.withAbsolutePath}

			assert.Equal(t, tc.expectType, img.Type())
		})
	}

}
