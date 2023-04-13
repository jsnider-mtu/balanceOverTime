package main

import (
    "os"
    "reflect"
    "testing"
    "time"
)

func init() {
    exp = map[string]float64{"1": 1600.00, "2": 171.91, "5": 32.98, "12": 500.00, "18": 25.00, "25": 251.65, "29": 490.60}
    bf = 1000.00
}

//func check(e error) {
//    if e != nil {
//        panic(e)
//    }
//}

type PaydaysTestCase struct {
    startDate string
    defaultExpected [5]time.Time
    twoWeeksExpected [5]time.Time
}

type DelayPaymentsTestCase struct {
    d0, d1 int
    d2 float64
    strExpected map[string]string
    floatExpected map[string]float64
}

type AddPaymentTestCase struct {
    e0 int
    e1 float64
    expected map[string]float64
}

type ExcludePaymentTestCase struct {
    e0 int
    e1 float64
    expected map[string]float64
}

type GetPaidTestCase struct {
    amount float64
    expected float64
}

type UpdateSubLnTestCase struct {
    d string
    payday bool
    expected string
}

// Test each function with different dates

func TestPaydays(t *testing.T) {
    tc0_0_default, err := time.Parse("2006-01-02", "2023-01-15")
    check(err)
    tc0_1_default, err := time.Parse("2006-01-02", "2023-01-30")
    check(err)
    tc0_2_default, err := time.Parse("2006-01-02", "2023-02-15")
    check(err)
    tc0_3_default, err := time.Parse("2006-01-02", "2023-02-28")
    check(err)
    tc0_4_default, err := time.Parse("2006-01-02", "2023-03-15")
    check(err)

    tc0_0_tw, err := time.Parse("2006-01-02", "2023-01-16")
    check(err)
    tc0_1_tw, err := time.Parse("2006-01-02", "2023-01-30")
    check(err)
    tc0_2_tw, err := time.Parse("2006-01-02", "2023-02-13")
    check(err)
    tc0_3_tw, err := time.Parse("2006-01-02", "2023-02-27")
    check(err)
    tc0_4_tw, err := time.Parse("2006-01-02", "2023-03-13")
    check(err)

    tc1_0_default, err := time.Parse("2006-01-02", "2023-01-30")
    check(err)
    tc1_1_default, err := time.Parse("2006-01-02", "2023-02-15")
    check(err)
    tc1_2_default, err := time.Parse("2006-01-02", "2023-02-28")
    check(err)
    tc1_3_default, err := time.Parse("2006-01-02", "2023-03-15")
    check(err)
    tc1_4_default, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)

    tc1_0_tw, err := time.Parse("2006-01-02", "2023-01-31")
    check(err)
    tc1_1_tw, err := time.Parse("2006-01-02", "2023-02-14")
    check(err)
    tc1_2_tw, err := time.Parse("2006-01-02", "2023-02-28")
    check(err)
    tc1_3_tw, err := time.Parse("2006-01-02", "2023-03-14")
    check(err)
    tc1_4_tw, err := time.Parse("2006-01-02", "2023-03-28")
    check(err)

    tc2_0_default, err := time.Parse("2006-01-02", "2023-02-15")
    check(err)
    tc2_1_default, err := time.Parse("2006-01-02", "2023-02-28")
    check(err)
    tc2_2_default, err := time.Parse("2006-01-02", "2023-03-15")
    check(err)
    tc2_3_default, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc2_4_default, err := time.Parse("2006-01-02", "2023-04-15")
    check(err)

    tc2_0_tw, err := time.Parse("2006-01-02", "2023-02-16")
    check(err)
    tc2_1_tw, err := time.Parse("2006-01-02", "2023-03-02")
    check(err)
    tc2_2_tw, err := time.Parse("2006-01-02", "2023-03-16")
    check(err)
    tc2_3_tw, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc2_4_tw, err := time.Parse("2006-01-02", "2023-04-13")
    check(err)

    tc3_0_default, err := time.Parse("2006-01-02", "2023-02-28")
    check(err)
    tc3_1_default, err := time.Parse("2006-01-02", "2023-03-15")
    check(err)
    tc3_2_default, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc3_3_default, err := time.Parse("2006-01-02", "2023-04-15")
    check(err)
    tc3_4_default, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)

    tc3_0_tw, err := time.Parse("2006-01-02", "2023-03-03")
    check(err)
    tc3_1_tw, err := time.Parse("2006-01-02", "2023-03-17")
    check(err)
    tc3_2_tw, err := time.Parse("2006-01-02", "2023-03-31")
    check(err)
    tc3_3_tw, err := time.Parse("2006-01-02", "2023-04-14")
    check(err)
    tc3_4_tw, err := time.Parse("2006-01-02", "2023-04-28")
    check(err)

    tc4_0_default, err := time.Parse("2006-01-02", "2023-03-15")
    check(err)
    tc4_1_default, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc4_2_default, err := time.Parse("2006-01-02", "2023-04-15")
    check(err)
    tc4_3_default, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)
    tc4_4_default, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)

    tc4_0_tw, err := time.Parse("2006-01-02", "2023-03-16")
    check(err)
    tc4_1_tw, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc4_2_tw, err := time.Parse("2006-01-02", "2023-04-13")
    check(err)
    tc4_3_tw, err := time.Parse("2006-01-02", "2023-04-27")
    check(err)
    tc4_4_tw, err := time.Parse("2006-01-02", "2023-05-11")
    check(err)

    tc5_0_default, err := time.Parse("2006-01-02", "2023-03-30")
    check(err)
    tc5_1_default, err := time.Parse("2006-01-02", "2023-04-15")
    check(err)
    tc5_2_default, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)
    tc5_3_default, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)
    tc5_4_default, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)

    tc5_0_tw, err := time.Parse("2006-01-02", "2023-03-31")
    check(err)
    tc5_1_tw, err := time.Parse("2006-01-02", "2023-04-14")
    check(err)
    tc5_2_tw, err := time.Parse("2006-01-02", "2023-04-28")
    check(err)
    tc5_3_tw, err := time.Parse("2006-01-02", "2023-05-12")
    check(err)
    tc5_4_tw, err := time.Parse("2006-01-02", "2023-05-26")
    check(err)

    tc6_0_default, err := time.Parse("2006-01-02", "2023-04-15")
    check(err)
    tc6_1_default, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)
    tc6_2_default, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)
    tc6_3_default, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)
    tc6_4_default, err := time.Parse("2006-01-02", "2023-06-15")
    check(err)

    tc6_0_tw, err := time.Parse("2006-01-02", "2023-04-16")
    check(err)
    tc6_1_tw, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)
    tc6_2_tw, err := time.Parse("2006-01-02", "2023-05-14")
    check(err)
    tc6_3_tw, err := time.Parse("2006-01-02", "2023-05-28")
    check(err)
    tc6_4_tw, err := time.Parse("2006-01-02", "2023-06-11")
    check(err)

    tc7_0_default, err := time.Parse("2006-01-02", "2023-04-30")
    check(err)
    tc7_1_default, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)
    tc7_2_default, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)
    tc7_3_default, err := time.Parse("2006-01-02", "2023-06-15")
    check(err)
    tc7_4_default, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)

    tc7_0_tw, err := time.Parse("2006-01-02", "2023-05-01")
    check(err)
    tc7_1_tw, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)
    tc7_2_tw, err := time.Parse("2006-01-02", "2023-05-29")
    check(err)
    tc7_3_tw, err := time.Parse("2006-01-02", "2023-06-12")
    check(err)
    tc7_4_tw, err := time.Parse("2006-01-02", "2023-06-26")
    check(err)

    tc8_0_default, err := time.Parse("2006-01-02", "2023-05-15")
    check(err)
    tc8_1_default, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)
    tc8_2_default, err := time.Parse("2006-01-02", "2023-06-15")
    check(err)
    tc8_3_default, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)
    tc8_4_default, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)

    tc8_0_tw, err := time.Parse("2006-01-02", "2023-05-16")
    check(err)
    tc8_1_tw, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)
    tc8_2_tw, err := time.Parse("2006-01-02", "2023-06-13")
    check(err)
    tc8_3_tw, err := time.Parse("2006-01-02", "2023-06-27")
    check(err)
    tc8_4_tw, err := time.Parse("2006-01-02", "2023-07-11")
    check(err)

    tc9_0_default, err := time.Parse("2006-01-02", "2023-05-30")
    check(err)
    tc9_1_default, err := time.Parse("2006-01-02", "2023-06-15")
    check(err)
    tc9_2_default, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)
    tc9_3_default, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)
    tc9_4_default, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)

    tc9_0_tw, err := time.Parse("2006-01-02", "2023-05-31")
    check(err)
    tc9_1_tw, err := time.Parse("2006-01-02", "2023-06-14")
    check(err)
    tc9_2_tw, err := time.Parse("2006-01-02", "2023-06-28")
    check(err)
    tc9_3_tw, err := time.Parse("2006-01-02", "2023-07-12")
    check(err)
    tc9_4_tw, err := time.Parse("2006-01-02", "2023-07-26")
    check(err)

    tc10_0_default, err := time.Parse("2006-01-02", "2023-06-15")
    check(err)
    tc10_1_default, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)
    tc10_2_default, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)
    tc10_3_default, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)
    tc10_4_default, err := time.Parse("2006-01-02", "2023-08-15")
    check(err)

    tc10_0_tw, err := time.Parse("2006-01-02", "2023-06-16")
    check(err)
    tc10_1_tw, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)
    tc10_2_tw, err := time.Parse("2006-01-02", "2023-07-14")
    check(err)
    tc10_3_tw, err := time.Parse("2006-01-02", "2023-07-28")
    check(err)
    tc10_4_tw, err := time.Parse("2006-01-02", "2023-08-11")
    check(err)

    tc11_0_default, err := time.Parse("2006-01-02", "2023-06-30")
    check(err)
    tc11_1_default, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)
    tc11_2_default, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)
    tc11_3_default, err := time.Parse("2006-01-02", "2023-08-15")
    check(err)
    tc11_4_default, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)

    tc11_0_tw, err := time.Parse("2006-01-02", "2023-07-01")
    check(err)
    tc11_1_tw, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)
    tc11_2_tw, err := time.Parse("2006-01-02", "2023-07-29")
    check(err)
    tc11_3_tw, err := time.Parse("2006-01-02", "2023-08-12")
    check(err)
    tc11_4_tw, err := time.Parse("2006-01-02", "2023-08-26")
    check(err)

    tc12_0_default, err := time.Parse("2006-01-02", "2023-07-15")
    check(err)
    tc12_1_default, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)
    tc12_2_default, err := time.Parse("2006-01-02", "2023-08-15")
    check(err)
    tc12_3_default, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)
    tc12_4_default, err := time.Parse("2006-01-02", "2023-09-15")
    check(err)

    tc12_0_tw, err := time.Parse("2006-01-02", "2023-07-16")
    check(err)
    tc12_1_tw, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)
    tc12_2_tw, err := time.Parse("2006-01-02", "2023-08-13")
    check(err)
    tc12_3_tw, err := time.Parse("2006-01-02", "2023-08-27")
    check(err)
    tc12_4_tw, err := time.Parse("2006-01-02", "2023-09-10")
    check(err)

    tc13_0_default, err := time.Parse("2006-01-02", "2023-07-30")
    check(err)
    tc13_1_default, err := time.Parse("2006-01-02", "2023-08-15")
    check(err)
    tc13_2_default, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)
    tc13_3_default, err := time.Parse("2006-01-02", "2023-09-15")
    check(err)
    tc13_4_default, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)

    tc13_0_tw, err := time.Parse("2006-01-02", "2023-07-31")
    check(err)
    tc13_1_tw, err := time.Parse("2006-01-02", "2023-08-14")
    check(err)
    tc13_2_tw, err := time.Parse("2006-01-02", "2023-08-28")
    check(err)
    tc13_3_tw, err := time.Parse("2006-01-02", "2023-09-11")
    check(err)
    tc13_4_tw, err := time.Parse("2006-01-02", "2023-09-25")
    check(err)

    tc14_0_default, err := time.Parse("2006-01-02", "2023-08-15")
    check(err)
    tc14_1_default, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)
    tc14_2_default, err := time.Parse("2006-01-02", "2023-09-15")
    check(err)
    tc14_3_default, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)
    tc14_4_default, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)

    tc14_0_tw, err := time.Parse("2006-01-02", "2023-08-16")
    check(err)
    tc14_1_tw, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)
    tc14_2_tw, err := time.Parse("2006-01-02", "2023-09-13")
    check(err)
    tc14_3_tw, err := time.Parse("2006-01-02", "2023-09-27")
    check(err)
    tc14_4_tw, err := time.Parse("2006-01-02", "2023-10-11")
    check(err)

    tc15_0_default, err := time.Parse("2006-01-02", "2023-08-30")
    check(err)
    tc15_1_default, err := time.Parse("2006-01-02", "2023-09-15")
    check(err)
    tc15_2_default, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)
    tc15_3_default, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)
    tc15_4_default, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)

    tc15_0_tw, err := time.Parse("2006-01-02", "2023-08-31")
    check(err)
    tc15_1_tw, err := time.Parse("2006-01-02", "2023-09-14")
    check(err)
    tc15_2_tw, err := time.Parse("2006-01-02", "2023-09-28")
    check(err)
    tc15_3_tw, err := time.Parse("2006-01-02", "2023-10-12")
    check(err)
    tc15_4_tw, err := time.Parse("2006-01-02", "2023-10-26")
    check(err)

    tc16_0_default, err := time.Parse("2006-01-02", "2023-09-15")
    check(err)
    tc16_1_default, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)
    tc16_2_default, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)
    tc16_3_default, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)
    tc16_4_default, err := time.Parse("2006-01-02", "2023-11-15")
    check(err)

    tc16_0_tw, err := time.Parse("2006-01-02", "2023-09-16")
    check(err)
    tc16_1_tw, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)
    tc16_2_tw, err := time.Parse("2006-01-02", "2023-10-14")
    check(err)
    tc16_3_tw, err := time.Parse("2006-01-02", "2023-10-28")
    check(err)
    tc16_4_tw, err := time.Parse("2006-01-02", "2023-11-11")
    check(err)

    tc17_0_default, err := time.Parse("2006-01-02", "2023-09-30")
    check(err)
    tc17_1_default, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)
    tc17_2_default, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)
    tc17_3_default, err := time.Parse("2006-01-02", "2023-11-15")
    check(err)
    tc17_4_default, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)

    tc17_0_tw, err := time.Parse("2006-01-02", "2023-10-01")
    check(err)
    tc17_1_tw, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)
    tc17_2_tw, err := time.Parse("2006-01-02", "2023-10-29")
    check(err)
    tc17_3_tw, err := time.Parse("2006-01-02", "2023-11-12")
    check(err)
    tc17_4_tw, err := time.Parse("2006-01-02", "2023-11-26")
    check(err)

    tc18_0_default, err := time.Parse("2006-01-02", "2023-10-15")
    check(err)
    tc18_1_default, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)
    tc18_2_default, err := time.Parse("2006-01-02", "2023-11-15")
    check(err)
    tc18_3_default, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)
    tc18_4_default, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)

    tc18_0_tw, err := time.Parse("2006-01-02", "2023-10-16")
    check(err)
    tc18_1_tw, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)
    tc18_2_tw, err := time.Parse("2006-01-02", "2023-11-13")
    check(err)
    tc18_3_tw, err := time.Parse("2006-01-02", "2023-11-27")
    check(err)
    tc18_4_tw, err := time.Parse("2006-01-02", "2023-12-11")
    check(err)

    tc19_0_default, err := time.Parse("2006-01-02", "2023-10-30")
    check(err)
    tc19_1_default, err := time.Parse("2006-01-02", "2023-11-15")
    check(err)
    tc19_2_default, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)
    tc19_3_default, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)
    tc19_4_default, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)

    tc19_0_tw, err := time.Parse("2006-01-02", "2023-10-31")
    check(err)
    tc19_1_tw, err := time.Parse("2006-01-02", "2023-11-14")
    check(err)
    tc19_2_tw, err := time.Parse("2006-01-02", "2023-11-28")
    check(err)
    tc19_3_tw, err := time.Parse("2006-01-02", "2023-12-12")
    check(err)
    tc19_4_tw, err := time.Parse("2006-01-02", "2023-12-26")
    check(err)

    tc20_0_default, err := time.Parse("2006-01-02", "2023-11-15")
    check(err)
    tc20_1_default, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)
    tc20_2_default, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)
    tc20_3_default, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)
    tc20_4_default, err := time.Parse("2006-01-02", "2024-01-15")
    check(err)

    tc20_0_tw, err := time.Parse("2006-01-02", "2023-11-16")
    check(err)
    tc20_1_tw, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)
    tc20_2_tw, err := time.Parse("2006-01-02", "2023-12-14")
    check(err)
    tc20_3_tw, err := time.Parse("2006-01-02", "2023-12-28")
    check(err)
    tc20_4_tw, err := time.Parse("2006-01-02", "2024-01-11")
    check(err)

    tc21_0_default, err := time.Parse("2006-01-02", "2023-11-30")
    check(err)
    tc21_1_default, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)
    tc21_2_default, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)
    tc21_3_default, err := time.Parse("2006-01-02", "2024-01-15")
    check(err)
    tc21_4_default, err := time.Parse("2006-01-02", "2024-01-30")
    check(err)

    tc21_0_tw, err := time.Parse("2006-01-02", "2023-12-01")
    check(err)
    tc21_1_tw, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)
    tc21_2_tw, err := time.Parse("2006-01-02", "2023-12-29")
    check(err)
    tc21_3_tw, err := time.Parse("2006-01-02", "2024-01-12")
    check(err)
    tc21_4_tw, err := time.Parse("2006-01-02", "2024-01-26")
    check(err)

    tc22_0_default, err := time.Parse("2006-01-02", "2023-12-15")
    check(err)
    tc22_1_default, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)
    tc22_2_default, err := time.Parse("2006-01-02", "2024-01-15")
    check(err)
    tc22_3_default, err := time.Parse("2006-01-02", "2024-01-30")
    check(err)
    tc22_4_default, err := time.Parse("2006-01-02", "2024-02-15")
    check(err)

    tc22_0_tw, err := time.Parse("2006-01-02", "2023-12-16")
    check(err)
    tc22_1_tw, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)
    tc22_2_tw, err := time.Parse("2006-01-02", "2024-01-13")
    check(err)
    tc22_3_tw, err := time.Parse("2006-01-02", "2024-01-27")
    check(err)
    tc22_4_tw, err := time.Parse("2006-01-02", "2024-02-10")
    check(err)

    tc23_0_default, err := time.Parse("2006-01-02", "2023-12-30")
    check(err)
    tc23_1_default, err := time.Parse("2006-01-02", "2024-01-15")
    check(err)
    tc23_2_default, err := time.Parse("2006-01-02", "2024-01-30")
    check(err)
    tc23_3_default, err := time.Parse("2006-01-02", "2024-02-15")
    check(err)
    tc23_4_default, err := time.Parse("2006-01-02", "2024-02-29")
    check(err)

    tc23_0_tw, err := time.Parse("2006-01-02", "2023-12-31")
    check(err)
    tc23_1_tw, err := time.Parse("2006-01-02", "2024-01-14")
    check(err)
    tc23_2_tw, err := time.Parse("2006-01-02", "2024-01-28")
    check(err)
    tc23_3_tw, err := time.Parse("2006-01-02", "2024-02-11")
    check(err)
    tc23_4_tw, err := time.Parse("2006-01-02", "2024-02-25")
    check(err)

    paydaysTestCases := []PaydaysTestCase{
        PaydaysTestCase{
            startDate: "2023-01-02",
            defaultExpected: [5]time.Time{tc0_0_default, tc0_1_default, tc0_2_default, tc0_3_default, tc0_4_default},
            twoWeeksExpected: [5]time.Time{tc0_0_tw, tc0_1_tw, tc0_2_tw, tc0_3_tw, tc0_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-01-17",
            defaultExpected: [5]time.Time{tc1_0_default, tc1_1_default, tc1_2_default, tc1_3_default, tc1_4_default},
            twoWeeksExpected: [5]time.Time{tc1_0_tw, tc1_1_tw, tc1_2_tw, tc1_3_tw, tc1_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-02-02",
            defaultExpected: [5]time.Time{tc2_0_default, tc2_1_default, tc2_2_default, tc2_3_default, tc2_4_default},
            twoWeeksExpected: [5]time.Time{tc2_0_tw, tc2_1_tw, tc2_2_tw, tc2_3_tw, tc2_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-02-17",
            defaultExpected: [5]time.Time{tc3_0_default, tc3_1_default, tc3_2_default, tc3_3_default, tc3_4_default},
            twoWeeksExpected: [5]time.Time{tc3_0_tw, tc3_1_tw, tc3_2_tw, tc3_3_tw, tc3_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-03-02",
            defaultExpected: [5]time.Time{tc4_0_default, tc4_1_default, tc4_2_default, tc4_3_default, tc4_4_default},
            twoWeeksExpected: [5]time.Time{tc4_0_tw, tc4_1_tw, tc4_2_tw, tc4_3_tw, tc4_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-03-17",
            defaultExpected: [5]time.Time{tc5_0_default, tc5_1_default, tc5_2_default, tc5_3_default, tc5_4_default},
            twoWeeksExpected: [5]time.Time{tc5_0_tw, tc5_1_tw, tc5_2_tw, tc5_3_tw, tc5_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-04-02",
            defaultExpected: [5]time.Time{tc6_0_default, tc6_1_default, tc6_2_default, tc6_3_default, tc6_4_default},
            twoWeeksExpected: [5]time.Time{tc6_0_tw, tc6_1_tw, tc6_2_tw, tc6_3_tw, tc6_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-04-17",
            defaultExpected: [5]time.Time{tc7_0_default, tc7_1_default, tc7_2_default, tc7_3_default, tc7_4_default},
            twoWeeksExpected: [5]time.Time{tc7_0_tw, tc7_1_tw, tc7_2_tw, tc7_3_tw, tc7_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-05-02",
            defaultExpected: [5]time.Time{tc8_0_default, tc8_1_default, tc8_2_default, tc8_3_default, tc8_4_default},
            twoWeeksExpected: [5]time.Time{tc8_0_tw, tc8_1_tw, tc8_2_tw, tc8_3_tw, tc8_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-05-17",
            defaultExpected: [5]time.Time{tc9_0_default, tc9_1_default, tc9_2_default, tc9_3_default, tc9_4_default},
            twoWeeksExpected: [5]time.Time{tc9_0_tw, tc9_1_tw, tc9_2_tw, tc9_3_tw, tc9_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-06-02",
            defaultExpected: [5]time.Time{tc10_0_default, tc10_1_default, tc10_2_default, tc10_3_default, tc10_4_default},
            twoWeeksExpected: [5]time.Time{tc10_0_tw, tc10_1_tw, tc10_2_tw, tc10_3_tw, tc10_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-06-17",
            defaultExpected: [5]time.Time{tc11_0_default, tc11_1_default, tc11_2_default, tc11_3_default, tc11_4_default},
            twoWeeksExpected: [5]time.Time{tc11_0_tw, tc11_1_tw, tc11_2_tw, tc11_3_tw, tc11_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-07-02",
            defaultExpected: [5]time.Time{tc12_0_default, tc12_1_default, tc12_2_default, tc12_3_default, tc12_4_default},
            twoWeeksExpected: [5]time.Time{tc12_0_tw, tc12_1_tw, tc12_2_tw, tc12_3_tw, tc12_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-07-17",
            defaultExpected: [5]time.Time{tc13_0_default, tc13_1_default, tc13_2_default, tc13_3_default, tc13_4_default},
            twoWeeksExpected: [5]time.Time{tc13_0_tw, tc13_1_tw, tc13_2_tw, tc13_3_tw, tc13_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-08-02",
            defaultExpected: [5]time.Time{tc14_0_default, tc14_1_default, tc14_2_default, tc14_3_default, tc14_4_default},
            twoWeeksExpected: [5]time.Time{tc14_0_tw, tc14_1_tw, tc14_2_tw, tc14_3_tw, tc14_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-08-17",
            defaultExpected: [5]time.Time{tc15_0_default, tc15_1_default, tc15_2_default, tc15_3_default, tc15_4_default},
            twoWeeksExpected: [5]time.Time{tc15_0_tw, tc15_1_tw, tc15_2_tw, tc15_3_tw, tc15_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-09-02",
            defaultExpected: [5]time.Time{tc16_0_default, tc16_1_default, tc16_2_default, tc16_3_default, tc16_4_default},
            twoWeeksExpected: [5]time.Time{tc16_0_tw, tc16_1_tw, tc16_2_tw, tc16_3_tw, tc16_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-09-17",
            defaultExpected: [5]time.Time{tc17_0_default, tc17_1_default, tc17_2_default, tc17_3_default, tc17_4_default},
            twoWeeksExpected: [5]time.Time{tc17_0_tw, tc17_1_tw, tc17_2_tw, tc17_3_tw, tc17_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-10-02",
            defaultExpected: [5]time.Time{tc18_0_default, tc18_1_default, tc18_2_default, tc18_3_default, tc18_4_default},
            twoWeeksExpected: [5]time.Time{tc18_0_tw, tc18_1_tw, tc18_2_tw, tc18_3_tw, tc18_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-10-17",
            defaultExpected: [5]time.Time{tc19_0_default, tc19_1_default, tc19_2_default, tc19_3_default, tc19_4_default},
            twoWeeksExpected: [5]time.Time{tc19_0_tw, tc19_1_tw, tc19_2_tw, tc19_3_tw, tc19_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-11-02",
            defaultExpected: [5]time.Time{tc20_0_default, tc20_1_default, tc20_2_default, tc20_3_default, tc20_4_default},
            twoWeeksExpected: [5]time.Time{tc20_0_tw, tc20_1_tw, tc20_2_tw, tc20_3_tw, tc20_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-11-17",
            defaultExpected: [5]time.Time{tc21_0_default, tc21_1_default, tc21_2_default, tc21_3_default, tc21_4_default},
            twoWeeksExpected: [5]time.Time{tc21_0_tw, tc21_1_tw, tc21_2_tw, tc21_3_tw, tc21_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-12-02",
            defaultExpected: [5]time.Time{tc22_0_default, tc22_1_default, tc22_2_default, tc22_3_default, tc22_4_default},
            twoWeeksExpected: [5]time.Time{tc22_0_tw, tc22_1_tw, tc22_2_tw, tc22_3_tw, tc22_4_tw},
        },
        PaydaysTestCase{
            startDate: "2023-12-17",
            defaultExpected: [5]time.Time{tc23_0_default, tc23_1_default, tc23_2_default, tc23_3_default, tc23_4_default},
            twoWeeksExpected: [5]time.Time{tc23_0_tw, tc23_1_tw, tc23_2_tw, tc23_3_tw, tc23_4_tw},
        },
    }

    for cind, c := range paydaysTestCases {
        tn, err = time.Parse("2006-01-02", c.startDate)
        check(err)
        default_actual := Paydays(c.startDate, false)
        if default_actual != c.defaultExpected {
            t.Errorf("[Test #%d] Expected %v, got %v", cind, c.defaultExpected, default_actual)
        }

        tw_actual := Paydays(c.startDate, true)
        if tw_actual != c.twoWeeksExpected {
            t.Errorf("[Test #%d] Expected %v, got %v", cind, c.twoWeeksExpected, tw_actual)
        }
    }
}

func TestDelayPayments(t *testing.T) {
    delayPaymentsTestCases := []DelayPaymentsTestCase{
        DelayPaymentsTestCase{
            d0: 1,
            d1: 15,
            d2: 1300.00,
            strExpected: map[string]string{"1": "15"},
            floatExpected: map[string]float64{"1": 1300.00},
        },
        DelayPaymentsTestCase{
            d0: 2,
            d1: 15,
            d2: 0.0,
            strExpected: map[string]string{"2": "15"},
            floatExpected: map[string]float64{"2": 171.91},
        },
        DelayPaymentsTestCase{
            d0: 3,
            d1: 15,
            d2: 500.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
        DelayPaymentsTestCase{
            d0: 0,
            d1: 15,
            d2: 500.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
        DelayPaymentsTestCase{
            d0: 7,
            d1: 32,
            d2: 500.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
        DelayPaymentsTestCase{
            d0: 15,
            d1: 15,
            d2: 500.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
        DelayPaymentsTestCase{
            d0: 1,
            d1: 15,
            d2: -500.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
        DelayPaymentsTestCase{
            d0: 1,
            d1: 15,
            d2: 1700.00,
            strExpected: map[string]string{},
            floatExpected: map[string]float64{},
        },
    }

    for _, c := range delayPaymentsTestCases {
        dexp = map[string]string{}
        vdexp = map[string]float64{}
        str_actual, float_actual := DelayPayments(c.d0, c.d1, c.d2)
        if !reflect.DeepEqual(str_actual, c.strExpected) {
            t.Errorf("Expected %v, got %v", c.strExpected, str_actual)
        }
        if !reflect.DeepEqual(float_actual, c.floatExpected) {
            t.Errorf("Expected %v, got %v", c.floatExpected, float_actual)
        }
    }
}

func TestAddPayment(t *testing.T) {
    addPaymentTestCases := []AddPaymentTestCase{
        AddPaymentTestCase{
            e0: 3,
            e1: 550.00,
            expected: map[string]float64{"3": 550.00},
        },
        AddPaymentTestCase{
            e0: 0,
            e1: 550.00,
            expected: map[string]float64{},
        },
        AddPaymentTestCase{
            e0: 32,
            e1: 550.00,
            expected: map[string]float64{},
        },
        AddPaymentTestCase{
            e0: 3,
            e1: -550.00,
            expected: map[string]float64{},
        },
    }

    for _, c := range addPaymentTestCases {
        actual := AddPayment(c.e0, c.e1)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Expected %v, got %v", c.expected, actual)
        }
    }
}

func TestExcludePayment(t *testing.T) {
    excludePaymentTestCases := []ExcludePaymentTestCase{
        ExcludePaymentTestCase{
            e0: 1,
            e1: 550.00,
            expected: map[string]float64{"1": 550.00},
        },
        ExcludePaymentTestCase{
            e0: 0,
            e1: 550.00,
            expected: map[string]float64{},
        },
        ExcludePaymentTestCase{
            e0: 32,
            e1: 550.00,
            expected: map[string]float64{},
        },
        ExcludePaymentTestCase{
            e0: 1,
            e1: -550.00,
            expected: map[string]float64{},
        },
        ExcludePaymentTestCase{
            e0: 3,
            e1: 550.00,
            expected: map[string]float64{},
        },
        ExcludePaymentTestCase{
            e0: 2,
            e1: 550.00,
            expected: map[string]float64{},
        },
    }

    for _, c := range excludePaymentTestCases {
        actual := ExcludePayment(c.e0, c.e1)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Expected %v, got %v", c.expected, actual)
        }
    }
}

func TestGetPaid(t *testing.T) {
    getPaidTestCases := []GetPaidTestCase{
        GetPaidTestCase{
            amount: 1500.00,
            expected: 2500.00,
        },
        GetPaidTestCase{
            amount: -1500.00,
            expected: 1000.00,
        },
    }

    for _, c := range getPaidTestCases {
        bf = 1000.00
        actual := GetPaid(c.amount)
        if actual != c.expected {
            t.Errorf("Expected %f, got %f", c.expected, actual)
        }
    }
}

func TestAddMDVIP(t *testing.T) {
    expected := map[string]float64{"1": 2050.00, "2": 171.91, "5": 32.98, "12": 500.00, "18": 25.00, "25": 251.65, "29": 490.60}
    actual := AddMDVIP()
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}

func TestSubMDVIP(t *testing.T) {
    exp = map[string]float64{"1": 2050.00, "2": 171.91, "5": 32.98, "12": 500.00, "18": 25.00, "25": 251.65, "29": 490.60}
    expected := map[string]float64{"1": 1600.00, "2": 171.91, "5": 32.98, "12": 500.00, "18": 25.00, "25": 251.65, "29": 490.60}
    actual := SubMDVIP()
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}

func TestUpdateSubLn(t *testing.T) {
    updateSubLnTestCases := []UpdateSubLnTestCase{
        UpdateSubLnTestCase{
            d: "3",
            payday: true,
            expected: "",
        },
        UpdateSubLnTestCase{
            d: "12",
            payday: true,
            expected: " - 500",
        },
        UpdateSubLnTestCase{
            d: "3",
            payday: false,
            expected: "",
        },
        UpdateSubLnTestCase{
            d: "4",
            payday: false,
            expected: " - 90.36",
        },
        UpdateSubLnTestCase{
            d: "6",
            payday: false,
            expected: " - 190.36",
        },
        UpdateSubLnTestCase{
            d: "18",
            payday: false,
            expected: " - 25 + 20",
        },
        UpdateSubLnTestCase{
            d: "0",
            payday: false,
            expected: "",
        },
        UpdateSubLnTestCase{
            d: "32",
            payday: false,
            expected: "",
        },
    }

    for cind, c := range updateSubLnTestCases {
        exp = map[string]float64{"1": 1600.00, "2": 171.91, "5": 32.98, "12": 500.00, "18": 25.00, "25": 251.65, "29": 490.60}
        dexp = map[string]string{"2": "4"}
        vdexp = map[string]float64{"2": 90.36}
        ddexp = map[string]float64{"4": 90.36}
        exexp = map[string]float64{"6": 190.36}
        nexp = map[string]float64{"18": 20.00}
        subln = ""

        actual := UpdateSubLn(c.d, c.payday)
        if actual != c.expected {
            t.Errorf("[Test #%d] Expected \"%s\", got \"%s\"", cind, c.expected, actual)
        }
    }
}

// Test version option

func TestCurrentVersion(t *testing.T) {
    actual := CurrentVersion()
    ver := os.Getenv("APP_VERSION")
    if ver == "" {
        ver = "DEVELOPMENT"
    }
    if actual != "Version " + ver {
        t.Errorf("Expected \"Version %s\", got \"%s\"", ver, actual)
    }
}
