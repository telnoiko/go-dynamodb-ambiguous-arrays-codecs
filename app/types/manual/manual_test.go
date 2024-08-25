package manual

import (
	"reflect"
	"testing"
)

func Test_ConvertToArray(t *testing.T) {
	type args struct {
		field any
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "nil input",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "string",
			args:    args{"apple"},
			want:    []string{"apple"},
			wantErr: false,
		},
		{
			name:    "[]string",
			args:    args{[]string{"apple", "banana"}},
			want:    []string{"apple", "banana"},
			wantErr: false,
		},
		{
			name:    "[]interface{}",
			args:    args{[]interface{}{"apple", "banana", 1}},
			want:    []string{"apple", "banana"},
			wantErr: false,
		},
		{
			name:    "unsupported type",
			args:    args{[]bool{true}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToArray(tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("tryConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tryConvert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
