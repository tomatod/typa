package typa

import (
  "reflect"
)

var (
  int8Type  reflect.Type = reflect.TypeOf((int8)(0))
  int16Type              = reflect.TypeOf((int16)(0))
  int32Type              = reflect.TypeOf((int32)(0))
  int64Type              = reflect.TypeOf((int64)(0))
  intType                = reflect.TypeOf((int)(0))
  uint8Type              = reflect.TypeOf((uint8)(0))
  uint16Type             = reflect.TypeOf((uint16)(0))
  uint32Type             = reflect.TypeOf((uint32)(0))
  uint64Type             = reflect.TypeOf((uint64)(0))
  uintType               = reflect.TypeOf((uint)(0))
  float32Type            = reflect.TypeOf((float32)(0))
  float64Type            = reflect.TypeOf((float64)(0))
  stringType             = reflect.TypeOf((string)(""))
  boolType               = reflect.TypeOf((bool)(false))
)

func Int8(i interface{}) (int8, bool) {
  v, ok := convert(i, int64Type)
  if !ok {
    return 0, false
  }
  return int8(v.Int()), true
}

func Int16(i interface{}) (int16, bool) {
  v, ok := convert(i, int32Type)
  if !ok {
    return 0, false
  }
  return int16(v.Int()), true
}

func Int32(i interface{}) (int32, bool) {
  v, ok := convert(i, int64Type)
  if !ok {
    return 0, false
  }
  return int32(v.Int()), true
}

func Int64(i interface{}) (int64, bool) {
  v, ok := convert(i, int64Type)
  if !ok {
    return 0, false
  }
  return int64(v.Int()), true
}

func Int(i interface{}) (int, bool) {
  v, ok := convert(i, intType)
  if !ok {
    return 0, false
  }
  return int(v.Int()), true
}

func Float32(i interface{}) (float32, bool) {
  v, ok := convert(i, float32Type)
  if !ok {
    return 0, false
  }
  return float32(v.Float()), true
}

func Float64(i interface{}) (float64, bool) {
  v, ok := convert(i, float64Type)
  if !ok {
    return 0, false
  }
  return v.Float(), true
}

func Uint8(i interface{}) (uint8, bool) {
  v, ok := convert(i, uint8Type)
  if !ok {
    return 0, false
  }
  return uint8(v.Uint()), true
}

func Uint16(i interface{}) (uint16, bool) {
  v, ok := convert(i, uint16Type)
  if !ok {
    return 0, false
  }
  return uint16(v.Uint()), true
}

func Uint32(i interface{}) (uint32, bool) {
  v, ok := convert(i, uint32Type)
  if !ok {
    return 0, false
  }
  return uint32(v.Uint()), true
}

func Uint64(i interface{}) (uint64, bool) {
  v, ok := convert(i, uint64Type)
  if !ok {
    return 0, false
  }
  return v.Uint(), true
}

func Uint(i interface{}) (uint, bool) {
  v, ok := convert(i, uintType)
  if !ok {
    return 0, false
  }
  return uint(v.Uint()), true
}

func String(i interface{}) (string, bool) {
  v, ok := convert(i, stringType)
  if !ok {
    return "", false
  }
  return v.String(), true
}

func Bool(i interface{}) (bool, bool) {
  v, ok := convert(i, boolType)
  if !ok {
    return false, false
  }
  return v.Bool(), true
}

func InternalTypeConvert(i interface{}, t reflect.Type) (interface{}, bool) {
  v, ok := convert(i, t)
  if !ok {
    return nil, false
  }
  return v.Interface(), true
}

func convert(i interface{}, t reflect.Type) (*reflect.Value, bool) {
  v := reflect.ValueOf(i)
  if !v.CanConvert(t) {
    return nil, false
  }
  v = v.Convert(t)
  return &v, true
}
