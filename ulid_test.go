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
