package ulids

import (
	"database/sql/driver"
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func TestULIDNew(t *testing.T) {
	uid := New()
	require.NotEmpty(t, uid.String())
	require.NotEqual(t, ulid.ULID{}.String(), uid.String())
}

func TestULIDValuer(t *testing.T) {
	var _ driver.Valuer = &ULID{}

	uid := New()

	dr, err := uid.Value()
	require.NoError(t, err)

	st := dr.(string)
	require.NotEmpty(t, st)
	require.Equal(t, uid.String(), st)
}

func TestULIDParse(t *testing.T) {
	uid := New()
	id := uid.String()

	uid2, err := Parse(id)
	require.NoError(t, err)
	require.Equal(t, uid, uid2)
}

func TestMustParse(t *testing.T) {
	t.Run("should panic on fail to parse", func(t *testing.T) {
		defer func() {
			rec := recover()
			require.Equal(t, "ulid: bad data size when unmarshaling", rec.(error).Error())
		}()

		id := MustParse("")
		require.Empty(t, id.String())
	})

	t.Run("should OK", func(t *testing.T) {
		defer func() {
			rec := recover()
			require.Empty(t, rec)
		}()

		id := MustParse(New().String())
		require.NotEmpty(t, id.String())
	})
}
