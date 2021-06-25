package field

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// Custom type to sort keys in resulting JSON
type OrderedMap map[int]Field

func (om OrderedMap) MarshalJSON() ([]byte, error) {
	keys := make([]int, 0, len(om))
	for k := range om {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for _, i := range keys {
		b, err := json.Marshal(om[i])
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("\"%d\":", i))
		buf.Write(b)

		// don't add "," if it's the last item
		if i == keys[len(keys)-1] {
			break
		}

		buf.Write([]byte{','})
	}
	buf.Write([]byte{'}'})

	return buf.Bytes(), nil
}
