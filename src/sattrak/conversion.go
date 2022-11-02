package main

import (
  "math"
	"strconv"
)

func setBitFlag(flag bool, statusByte byte, position int) byte {
	if flag {
		return statusByte | (1 << position)
	}
	return statusByte & (0xff ^ (1 << position))
}

func intToBCD(num interface{}) (bcd []byte) {
	var num64 int64
	switch num.(type) {
	case int:
		num64 = int64(num.(int))
	case int64:
		num64 = num.(int64)
	default:
	}
	numStr := strconv.FormatInt(num64, 10)
	for _, x := range numStr {
		bcd = append(bcd, byte(x-'0'))
	}
	return
}

func padByteSlice(i []byte, size int) (o []byte) {
	o = i[:]
	for x := len(i); x < size; x++ {
		o = append([]byte{0}, o...)
	}
	return
}

func float64ToDecimalPointAssumed(f float64) (s string) {
  if f < 0 {
    s += "-"
  } else {
    s += " "
  }
  if f == 0 {
    s += "00000+0"
  } else {
    l10 := int(math.Log10(math.Abs(f)))
    s += float64ToFixedString(f/math.Pow(10, float64(l10)), 5)
    if l10 < 0 {
      s += "-"
    } else {
      s += "+"
    }
    s += intToFixedString(l10, 1)
  }
  return
}

func float64ToFixedString(f float64, places int) string {
  _, dec := math.Modf(math.Abs(f))
  decInt := int64(math.Round(math.Abs(dec*math.Pow(10, float64(places)))))
  decIntStr := strconv.FormatInt(decInt, 10)
  for len(decIntStr) < places {
    decIntStr = "0"+decIntStr
  }
  return decIntStr[:places]
}

func intToFixedString(i int, places int) string {
  intStr := strconv.Itoa(int(math.Abs(float64(i))))
  for len(intStr) < places {
    intStr = "0"+intStr
  }
  return intStr[len(intStr)-places:]
}

func int64ToFixedString(i int64, places int) string {
  int64Str := strconv.FormatInt(int64(math.Abs(float64(i))), 10)
  for len(int64Str) < places {
    int64Str = "0"+int64Str
  }
  return int64Str[len(int64Str)-places:]
}

func stringToFixedString(s string, places int) string {
  for len(s) < places {
    s += " "
  }
  return s[:places]
}
