package main

import (
	"testing"
)

func Test_validate(t *testing.T) {
	type args struct {
		buf []byte
	}
	test1 := args{[]byte{84, 104, 105, 115, 32, 105, 115, 32, 110, 111, 116, 32, 105, 109, 97, 103, 101, 32, 99, 111, 110, 116, 101, 110, 116, 10}}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test to check in correct file type throws error",
			test1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.buf); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
