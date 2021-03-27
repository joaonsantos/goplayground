package math

import (
  "testing"
)

type cases struct {
  values []float64
  result float64
}



func TestAverage(t *testing.T) {
  tests := []cases{
    { []float64{1,2}, 1.5 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, 0 },
  }

  for _, test := range tests {
    v := Average(test.values)

    if v != test.result {
      t.Error("For", test.values, "Expected", test.result, "got:", v)
    }
  }
}

func TestMax(t *testing.T) {
  tests := []cases{
    { []float64{1,2}, 2 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, 1 },
  }

  for _, test := range tests {
    v := Max(test.values)
    if v != test.result {
      t.Error("For", test.values, "Expected", test.result, "got:", v)
    }
  }
}

func TestMin(t *testing.T) {
  tests := []cases{
    { []float64{1,2}, 1 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, -1 },
  }

  for _, test := range tests {
    v := Min(test.values)
    if v != test.result {
      t.Error("For", test.values, "Expected", test.result, "got:", v)
    }
  }
}
