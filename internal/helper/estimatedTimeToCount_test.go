package helper

import (
    "testing"
)

func TestCountEstimatedTimeToFix(t *testing.T) {
    tests := []struct {
        typeOfDamage         string
        hasBeenDamagedBefore bool
        want                 int
    }{
        {"Rusak Sebagian", false, 2},
        {"Rusak Sedang", false, 4},
        {"Rusak Parah", false, 7},
        {"Rusak Sebagian", true, 4},
        {"Rusak Sedang", true, 8},
        {"Rusak Parah", true, 14},
    }

    for _, tt := range tests {
        got := CountEstimatedTimeToFix(tt.typeOfDamage, tt.hasBeenDamagedBefore)
        if got != tt.want {
            t.Errorf("CountEstimatedTimeToFix(%v, %v) = %v, want %v", tt.typeOfDamage, tt.hasBeenDamagedBefore, got, tt.want)
        }
    }
}
