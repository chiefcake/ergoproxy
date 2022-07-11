package storage

import "testing"

func TestMapInsert(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	tests := map[string]struct {
		args args
	}{
		"OK": {
			args: args{
				key:   "<KEY>",
				value: "<VALUE>",
			},
		},
	}

	m := NewMap[string, string]()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m.Insert(test.args.key, test.args.value)
		})
	}
}
