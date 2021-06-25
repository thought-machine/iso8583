package field

import (
	"testing"

	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/stretchr/testify/require"
)

func TestNumericField(t *testing.T) {
	spec := &Spec{
		Length:      10,
		Description: "Field",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
		Pad:         padding.Left(' '),
	}
	numeric := NewNumeric(spec)

	numeric.SetBytes([]byte("100"))
	require.Equal(t, 100, numeric.Value)

	packed, err := numeric.Pack()
	require.NoError(t, err)
	require.Equal(t, "       100", string(packed))

	length, err := numeric.Unpack([]byte("      9876"))
	require.NoError(t, err)
	require.Equal(t, 10, length)

	b, err := numeric.Bytes()
	require.NoError(t, err)
	require.Equal(t, "9876", string(b))

	require.Equal(t, 9876, numeric.Value)

	numeric = NewNumeric(spec)
	numeric.SetData(NewNumericValue(9876))
	packed, err = numeric.Pack()
	require.NoError(t, err)
	require.Equal(t, "      9876", string(packed))

	numeric = NewNumeric(spec)
	data := NewNumericValue(0)
	numeric.SetData(data)
	length, err = numeric.Unpack([]byte("      9876"))
	require.NoError(t, err)
	require.Equal(t, 10, length)
	require.Equal(t, 9876, data.Value)
}

func TestNumericFieldWithNotANumber(t *testing.T) {
	numeric := NewNumeric(&Spec{
		Length:      10,
		Description: "Field",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
		Pad:         padding.Left(' '),
	})

	numeric.SetBytes([]byte("hello"))
	require.Equal(t, 0, numeric.Value)

	packed, err := numeric.Pack()

	require.NoError(t, err)
	require.Equal(t, "         0", string(packed))

	_, err = numeric.Unpack([]byte("hhhhhhhhhh"))

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to convert into number")
}
