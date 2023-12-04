package main

import "testing"
 
func TestCalibrationCode1(t *testing.T) {
    code := get_calibration_code("1abc2")
    if code != 12 {
        t.Fail()
    }
}

func TestCalibrationCode2(t *testing.T) {
    code := get_calibration_code("pqr3stu8vwx")
    if code != 38 {
        t.Fail()
    }
}

func TestCalibrationCode3(t *testing.T) {
    code := get_calibration_code("a1b2c3d4e5f")
    if code != 15 {
        t.Fail()
    }
}

func TestCalibrationCode4(t *testing.T) {
    code := get_calibration_code("treb7uchet")
    if code != 77 {
        t.Fail()
    }
}
