package comments_test

import (
	"net/url"
	"testing"

	"github.com/matryer/infoshare/comments"
	"github.com/stretchr/testify/assert"
)

func TestSimplifyURL(t *testing.T) {

	tests := []struct {
		URL    string
		Simple string
	}{
		{
			URL:    "http://infoshare.pl/",
			Simple: "infoshare.pl/",
		}, {
			URL:    "http://infoshare.pl",
			Simple: "infoshare.pl/",
		},
		{
			URL:    "http://infoshare.pl/path/to/SOMETHING",
			Simple: "infoshare.pl/path/to/something",
		},
		{
			URL:    "https://infoshare.pl/path/to/SOMETHING",
			Simple: "infoshare.pl/path/to/something",
		},
	}

	for _, test := range tests {
		u, err := url.Parse(test.URL)
		assert.NoError(t, err)
		actual := comments.SimplifyURL(u)
		assert.Equal(t, actual, test.Simple)
	}

}
