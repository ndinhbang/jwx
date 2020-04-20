// this file was auto-generated by internal/cmd/gentypes/main.go: DO NOT EDIT

package jwa_test

import (
	"testing"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/stretchr/testify/assert"
)

func TestContentEncryptionAlgorithm(t *testing.T) {
	t.Run(`accept jwa constant A128CBC_HS256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A128CBC_HS256), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128CBC_HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A128CBC-HS256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A128CBC-HS256"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128CBC_HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A128CBC-HS256`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A128CBC-HS256"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128CBC_HS256, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant A128GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A128GCM), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A128GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A128GCM"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A128GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A128GCM"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A128GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant A192CBC_HS384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A192CBC_HS384), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192CBC_HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A192CBC-HS384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A192CBC-HS384"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192CBC_HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A192CBC-HS384`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A192CBC-HS384"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192CBC_HS384, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant A192GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A192GCM), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A192GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A192GCM"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A192GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A192GCM"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A192GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant A256CBC_HS512`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A256CBC_HS512), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256CBC_HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A256CBC-HS512`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A256CBC-HS512"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256CBC_HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A256CBC-HS512`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A256CBC-HS512"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256CBC_HS512, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept jwa constant A256GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(jwa.A256GCM), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept the string A256GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept("A256GCM"), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
	t.Run(`accept fmt.Stringer for A256GCM`, func(t *testing.T) {
		t.Parallel()
		var dst jwa.ContentEncryptionAlgorithm
		if !assert.NoError(t, dst.Accept(stringer{src: "A256GCM"}), `accept is successful`) {
			return
		}
		if !assert.Equal(t, jwa.A256GCM, dst, `accepted value should be equal to constant`) {
			return
		}
	})
}
