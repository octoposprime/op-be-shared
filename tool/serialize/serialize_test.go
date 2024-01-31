package tserialize

import (
	"reflect"
	"testing"
)

func Test_NewSerializer(t *testing.T) {
	type testData struct {
		test string
	}
	type args struct {
		v testData
	}
	tests := []struct {
		name string
		args args
		want Serialize
	}{
		{
			name: "NewSerializer",
			args: args{
				v: testData{
					test: "test",
				},
			},
			want: Serialize{
				v: testData{
					test: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSerializer(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSerializer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerialize_ToJson(t *testing.T) {
	type testData struct {
		test string
	}
	type fields struct {
		v testData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ToJson",
			fields: fields{
				v: testData{
					test: "test",
				},
			},
			want: "{\"test\":\"test\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Serialize{
				v: tt.fields.v,
			}
			if got := s.ToJson(); got != tt.want {
				t.Errorf("Serialize.ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerialize_ToJsonPretty(t *testing.T) {
	type testData struct {
		test string
	}
	type fields struct {
		v testData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ToJsonPretty",
			fields: fields{
				v: testData{
					test: "test",
				},
			},
			want: "{\n  \"test\": \"test\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Serialize{
				v: tt.fields.v,
			}
			if got := s.ToJsonPretty(); got != tt.want {
				t.Errorf("Serialize.ToJsonPretty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerialize_FormJson(t *testing.T) {
	type testData struct {
		test string
	}
	type fields struct {
		v testData
	}
	type args struct {
		src string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			name: "FormJson",
			fields: fields{
				v: testData{
					test: "test",
				},
			},
			args: args{
				src: "{\"test\":\"test\"}",
			},
			want: testData{
				test: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Serialize{
				v: tt.fields.v,
			}
			if got := s.FormJson(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize.FormJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
