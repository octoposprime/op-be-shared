package tserialize

import "encoding/json"

type Serializer interface {
	ToJson() string
}

type Deserializer interface {
	FormJson(src string)
	ToJsonPretty() string
}

type Serialize struct {
	v any
}

func NewSerializer(v any) Serialize {
	return Serialize{v}
}

func (s Serialize) ToJson() string {
	j, _ := json.Marshal(s.v)
	return string(j)
}

func (s Serialize) ToJsonPretty() string {
	b, _ := json.MarshalIndent(s.v, "", "  ")
	return string(b)
}

func (s Serialize) FormJson(src string) any {
	return json.Unmarshal([]byte(src), s.v)
}
