package typa

import (
  "testing"
  "reflect"
  "encoding/json"
)

func TestInt(t *testing.T) {
  var (
    v_origin interface{}  = 127
    v_expect int8         = 127
  )
  v_result, ok := Int8(v_origin)
  if !ok || v_result != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestFloat32(t *testing.T) {
  var (
    v_origin interface{}  = 1.234567
    v_expect float32      = 1.234567
  )
  v_result, ok := Float32(v_origin)
  if !ok || v_result != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestUint16(t *testing.T) {
  var (
    v_origin interface{}  = 12345
    v_expect uint16       = 12345
  )
  v_result, ok := Uint16(v_origin)
  if !ok || v_result != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestString(t *testing.T) {
  var (
    v_origin interface{}  = "This is test message."
    v_expect string       = "This is test message."
  )
  v_result, ok := String(v_origin)
  if !ok || v_result != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestBool(t *testing.T) {
  var (
    v_origin interface{}  = true
    v_expect bool         = true
  )
  v_result, ok := Bool(v_origin)
  if !ok || v_result != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestInternalTypeConvert(t *testing.T) {
  var (
    v_origin interface{}  = 77
    v_expect byte         = 77
  )
  v_result, ok := InternalTypeConvert(v_origin, reflect.TypeOf((byte)(0)))
  if !ok || v_result.(byte) != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestSenarioJsonParse(t *testing.T) {
  data := `
  {
    "string"  : "v1",
    "int"     : 123456,
    "float32" : 7.7777,
    "float64" : 987654,
    "bool"    : true
  }
`
  v_expect := map[string]interface{} {
    "string":  string("v1"),
    "int":     int(123456),
    "float32": float32(7.7777),
    "float64": float64(987654),
    "bool":    bool(true),
  }
  v_result := map[string]interface{} {}

  json.Unmarshal([]byte(data), &v_result)
  v_result["string"], _  = InternalTypeConvert(v_result["string"], stringType)
  v_result["int"], _     = InternalTypeConvert(v_result["int"], intType)
  v_result["float32"], _ = InternalTypeConvert(v_result["float32"], float32Type)
  v_result["float64"], _ = InternalTypeConvert(v_result["float64"], float64Type)
  v_result["bool"], _    = InternalTypeConvert(v_result["bool"], boolType)
  if !reflect.DeepEqual(v_result, v_expect) {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}

func TestConvet(t *testing.T) {
  var (
    v_origin interface{}  = 77777777
    v_expect float64      = 77777777
  )
  v_result, ok := convert(v_origin, float64Type)
  if !ok || v_result.Float() != v_expect {
    t.Fatalf("\nresult: %v\nexpect: %v", v_result, v_expect)
  }
}
