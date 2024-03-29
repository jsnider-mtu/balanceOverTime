package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

var (
	err        error
	tn         time.Time
	tim        time.Time
	pda        [5]time.Time
	dat        []byte
	bf         float64
	mdvipadded bool    = false
	mdvipprice float64 = 512.5
	subln      string
)

var exp = make(map[string]float64)
var dexp = make(map[string]string)
var ddexp = make(map[string]float64)
var vdexp = make(map[string]float64)
var exexp = make(map[string]float64)
var nexp = make(map[string]float64)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Paydays(startDate string, twoWeeks bool) [5]time.Time {
	var ref time.Time
	var pdarray [5]time.Time
	if twoWeeks {
		// Verify startDate value is correct format
		t, err := time.Parse("2006-01-02", startDate)
		check(err)
		// Get reference date
		if diff := tn.Sub(t); diff.Hours() >= 336 {
			mod := int(diff.Hours()) % 336
			dmod, _ := time.ParseDuration(strconv.Itoa(mod) + "h")
			ref = tn.Add(-dmod)
		} else {
			ref = t
		}
	} else {
		// Reference date is either last 15th or 30th or 28th or 29th if Feb
		if int(tn.Month()) == 3 && tn.Day() < 15 {
			if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
				ref, err = time.Parse("2006-01-02", strconv.Itoa(tn.Year())+"-02-29")
				check(err)
			} else {
				ref, err = time.Parse("2006-01-02", strconv.Itoa(tn.Year())+"-02-28")
				check(err)
			}
		} else {
			if tn.Day() < 15 {
				if int(tn.Month()) > 1 {
					ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-30", tn.Year(), int(tn.Month())-1))
					check(err)
				} else {
					ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%d-30", tn.Year()-1, 12))
					check(err)
				}
			} else {
				ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-15", tn.Year(), int(tn.Month())))
				check(err)
			}
		}
	}

	if twoWeeks {
		// Create array of next 5 paydays
		pd1, err := time.ParseDuration("336h")
		check(err)
		pd2, err := time.ParseDuration("672h")
		check(err)
		pd3, err := time.ParseDuration("1008h")
		check(err)
		pd4, err := time.ParseDuration("1344h")
		check(err)
		pd5, err := time.ParseDuration("1680h")
		check(err)
		pdarray = [5]time.Time{ref.Add(pd1), ref.Add(pd2), ref.Add(pd3), ref.Add(pd4), ref.Add(pd5)}
	} else {
		// Create array of next 5 paydays, 15ths and 30ths and 28th or 29th if Feb
		switch int(ref.Month()) {
		case 12:
			if ref.Day() == 15 {
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-12-30", ref.Year()))
				check(err)
				pdarray[0] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-15", ref.Year()+1))
				check(err)
				pdarray[1] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year()+1))
				check(err)
				pdarray[2] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()+1))
				check(err)
				pdarray[3] = tim
				if (ref.Year()+1)%400 == 0 || ((ref.Year()+1)%4 == 0 && (ref.Year()+1)%100 != 0) {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()+1))
					check(err)
					pdarray[4] = tim
				} else {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()+1))
					check(err)
					pdarray[4] = tim
				}
			} else {
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-15", ref.Year()+1))
				check(err)
				pdarray[0] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year()+1))
				check(err)
				pdarray[1] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()+1))
				check(err)
				pdarray[2] = tim
				if (ref.Year()+1)%400 == 0 || ((ref.Year()+1)%4 == 0 && (ref.Year()+1)%100 != 0) {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()+1))
					check(err)
					pdarray[3] = tim
				} else {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()+1))
					check(err)
					pdarray[3] = tim
				}
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()+1))
				check(err)
				pdarray[4] = tim
			}
		case 1:
			if ref.Day() == 15 {
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year()))
				check(err)
				pdarray[0] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()))
				check(err)
				pdarray[1] = tim
				if ref.Year()%400 == 0 || (ref.Year()%4 == 0 && ref.Year()%100 != 0) {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
					check(err)
					pdarray[2] = tim
				} else {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
					check(err)
					pdarray[2] = tim
				}
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
				check(err)
				pdarray[3] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
				check(err)
				pdarray[4] = tim
			} else {
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()))
				check(err)
				pdarray[0] = tim
				if ref.Year()%400 == 0 || (ref.Year()%4 == 0 && ref.Year()%100 != 0) {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
					check(err)
					pdarray[1] = tim
				} else {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
					check(err)
					pdarray[1] = tim
				}
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
				check(err)
				pdarray[2] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
				check(err)
				pdarray[3] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-15", ref.Year()))
				check(err)
				pdarray[4] = tim
			}
		case 2:
			if ref.Day() == 15 {
				if ref.Year()%400 == 0 || (ref.Year()%4 == 0 && ref.Year()%100 != 0) {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
					check(err)
					pdarray[0] = tim
				} else {
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
					check(err)
					pdarray[0] = tim
				}
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
				check(err)
				pdarray[1] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
				check(err)
				pdarray[2] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-15", ref.Year()))
				check(err)
				pdarray[3] = tim
				tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-30", ref.Year()))
				check(err)
				pdarray[4] = tim
			} else {
				for x := 0; x < 5; x++ {
					mon := int((float64(x) / 2.0) + 3.0)
					tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year(), mon, ((x%2)*15)+15))
					check(err)
					pdarray[x] = tim
				}
			}
		default:
			ymod := 0
			if ref.Day() == 30 || ref.Day() == 29 || ref.Day() == 28 {
				for x := 0; x < 5; x++ {
					mon := int((float64(x)/2.0)+float64(ref.Month()+1)) % 12
					if mon == 0 {
						mon = 12
					} else if mon < int(ref.Month()) {
						ymod = 1
					}
					tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year()+ymod, mon, ((x%2)*15)+15))
					if err != nil {
						fmt.Println(err)
					}
					pdarray[x] = tim
				}
			} else {
				for x := 1; x <= 5; x++ {
					mon := int((float64(x)/2.0)+float64(ref.Month())) % 12
					if mon == 0 {
						mon = 12
					} else if mon < int(ref.Month()) {
						ymod = 1
					}
					tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year()+ymod, mon, ((x%2)*15)+15))
					if err != nil {
						fmt.Println(err)
					}
					pdarray[x-1] = tim
				}
			}
		}
	}
	return pdarray
}

func DelayPayments(d0, d1 int, d2 float64) (map[string]string, map[string]float64) {
	// Ask user if any payment's first occurence should be delayed
	// ex.
	// Today == 10
	// dsd0 == 3; dsd1 == 20 expense on 3rd of next month moved to 20th of next month
	// dsd0 == 6; dsd1 == 3 expense on 6th of next month moved to 3rd of following month
	// dsd0 == 6; dsd1 == 7 expense on 6th of next month moved to 7th of next month
	// dsd0 == 12; dsd1 == 3 expense on 12th of this month moved to 3rd of next month
	// dsd0 == 12; dsd1 == 28 expense on 12th of this month moved to 28th of this month
	// dsd0 == 28; dsd1 == 26 expense on 28th of this month moved to 26th of next month

	// Validate input
	if _, ok := exp[strconv.Itoa(d0)]; !ok {
		fmt.Printf("%d is not an expense day\n", d0)
		return map[string]string{}, map[string]float64{}
	}
	if _, ok := dexp[strconv.Itoa(d0)]; ok {
		fmt.Printf("Day %d is already being deferred\nCtrl+C (or Cmd+C) to exit and start again\n", d0)
		return map[string]string{}, map[string]float64{}
	}
	if d0 < 1 || d0 > 31 {
		fmt.Printf("%d is not a valid day\n", d0)
		return map[string]string{}, map[string]float64{}
	}
	if d1 < 1 || d1 > 31 {
		fmt.Printf("%d is not a valid day\n", d1)
		return map[string]string{}, map[string]float64{}
	}
	if d0 == d1 {
		fmt.Println("Days cannot match")
		return map[string]string{}, map[string]float64{}
	}
	if d2 < 0.0 {
		fmt.Println("Amount deferred cannot be negative")
		return map[string]string{}, map[string]float64{}
	}
	if d2 > exp[strconv.Itoa(d0)] {
		fmt.Println("Amount deferred cannot exceed the total payment")
		fmt.Printf("Max for day %d: %.2f\n", d0, exp[strconv.Itoa(d0)])
		return map[string]string{}, map[string]float64{}
	}

	if d0 > tn.Day() {
		// expense of this month
		if d1 > d0 {
			// expense moved to this month
			switch int(tn.Month()) {
			case 2:
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d0 > 29 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
					if d1 > 29 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d0 > 28 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
					if d1 > 28 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				}
			case 4, 6, 9, 11:
				if d0 > 30 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 30 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			default:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			}
		} else {
			// expense moved to next month
			switch int(tn.Month()) {
			case 1:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d1 > 29 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d1 > 28 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				}
			case 2:
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d0 > 29 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d0 > 28 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			case 3, 5, 8, 10:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 30 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			case 4, 6, 9, 11:
				if d0 > 30 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			default:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			}
		}
	} else {
		// expense of next month
		if d1 > d0 {
			// expense moved to next month
			switch int(tn.Month()) {
			case 1:
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d0 > 29 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
					if d1 > 29 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d0 > 28 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
					if d1 > 28 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				}
			case 2, 4, 6, 7, 9, 11, 12:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			default:
				if d0 > 30 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 30 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			}
		} else {
			// expense moved to following month
			switch int(tn.Month()) {
			case 1:
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d0 > 29 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d0 > 28 {
						fmt.Printf("%d is not a valid day\n", d0)
						return map[string]string{}, map[string]float64{}
					}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			case 2, 4, 7, 9:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 30 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			case 3, 5, 8, 10:
				if d0 > 30 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			case 6, 11:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if d1 > 31 {
					fmt.Printf("%d is not a valid day\n", d1)
					return map[string]string{}, map[string]float64{}
				}
			default:
				if d0 > 31 {
					fmt.Printf("%d is not a valid day\n", d0)
					return map[string]string{}, map[string]float64{}
				}
				if tn.Year()%400 == 0 || (tn.Year()%4 == 0 && tn.Year()%100 != 0) {
					if d1 > 29 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				} else {
					if d1 > 28 {
						fmt.Printf("%d is not a valid day\n", d1)
						return map[string]string{}, map[string]float64{}
					}
				}
			}
		}
	}
	dexp[strconv.Itoa(d0)] = strconv.Itoa(d1)
	if d2 == 0.0 {
		vdexp[strconv.Itoa(d0)] = exp[strconv.Itoa(d0)]
	} else {
		vdexp[strconv.Itoa(d0)] = d2
	}
	return dexp, vdexp
}

func AddPayment(e0 int, e1 float64) map[string]float64 {
	if e0 < 1 || e0 > 31 {
		fmt.Printf("%d is not a valid day\n", e0)
		return map[string]float64{}
	}
	if e1 < 0.0 {
		fmt.Println("Amount cannot be negative")
		return map[string]float64{}
	}
	if _, ok := exexp[strconv.Itoa(e0)]; ok {
		exexp[strconv.Itoa(e0)] += e1
	} else {
		exexp[strconv.Itoa(e0)] = e1
	}
	return exexp
}

func ExcludePayment(e2 int, e3 float64) map[string]float64 {
	if e2 < 1 || e2 > 31 {
		fmt.Printf("%d is not a valid day\n", e2)
		return map[string]float64{}
	}
	if e3 < 0.0 {
		fmt.Println("Amount cannot be negative")
		return map[string]float64{}
	}
	if e3 > exp[strconv.Itoa(e2)] {
		fmt.Println("Amount cannot exceed the day's expense")
		return map[string]float64{}
	}
	if _, ok := exp[strconv.Itoa(e2)]; ok {
		nexp[strconv.Itoa(e2)] = e3
	} else {
		fmt.Println("No expenses on day $d", e2)
		return map[string]float64{}
	}
	return nexp
}

func GetPaid(amount float64) float64 {
	if amount < 0.0 {
		fmt.Println("Paycheck cannot be negative")
		return bf
	}
	bf += amount
	return bf
}

func AddMDVIP() map[string]float64 {
	exp["1"] += mdvipprice
	mdvipadded = true
	return exp
}

func SubMDVIP() map[string]float64 {
	exp["1"] -= mdvipprice
	mdvipadded = false
	return exp
}

func UpdateSubLn(d string, payday bool) string {
	dint, err := strconv.Atoi(d)
	check(err)
	if dint < 1 || dint > 31 {
		fmt.Printf("%s is not a valid day\n", d)
		return ""
	}
	if payday {
		subln = ""
	}
	if val, ex := exp[d]; ex {
		if _, ok := dexp[d]; !ok {
			bf -= val
			subln += " - " + fmt.Sprint(val)
		} else {
			diff := val - vdexp[d]
			bf -= diff
			subln += " - " + fmt.Sprint(diff)
			ddexp[dexp[d]] += vdexp[d]
			delete(vdexp, d)
			delete(dexp, d)
		}
	}
	if val, ex := ddexp[d]; ex {
		bf -= val
		subln += " - " + fmt.Sprint(val)
		delete(ddexp, d)
	}
	if val, ex := exexp[d]; ex {
		bf -= val
		subln += " - " + fmt.Sprint(val)
		delete(exexp, d)
	}
	if val, ex := nexp[d]; ex {
		bf += val
		subln += " + " + fmt.Sprint(val)
		delete(nexp, d)
	}
	return subln
}

func CurrentVersion() string {
	return "Version DEVELOPMENT"
}

func main() {
	datePtr := flag.String("d", "2020-07-10", "Last date paid (YYYY-MM-DD)")
	payPtr := flag.Float64("p", 2732.23, "How much are you paid?")
	twoWeekPtr := flag.Bool("twoWeeks", false, "Do you get paid every two weeks?")
	mdvipPtr := flag.Bool("mdvip", false, "Do you pay MDVIP every quarter?")
	versionPtr := flag.Bool("v", false, "Display app version and exit")

	flag.Parse()

	if *payPtr < 0.0 {
		fmt.Println("Paycheck cannot be negative")
		os.Exit(1)
	}

	if *versionPtr {
		fmt.Println(CurrentVersion())
		os.Exit(0)
	}

	tn = time.Now()
	tn, err = time.Parse("2006-01-02", tn.Format("2006-01-02"))
	check(err)
	pda = Paydays(*datePtr, *twoWeekPtr)

	// Read expenses.json and unmarshal data
	dat, err = ioutil.ReadFile(os.Getenv("HOME") + "/.local/etc/bot/expenses.json")
	check(err)
	json.Unmarshal(dat, &exp)

	// Ask user for current account balance
	var b string
	fmt.Println("Current balance?")
	_, err = fmt.Scanln(&b)
	check(err)

	// Ask user if any payment's first occurence should be delayed
	var d0 int
	var d1 int
	var d2 float64
	eodd := false
	sortedexp := make([]int, 0, len(exp))
	for k := range exp {
		ki, err := strconv.Atoi(k)
		check(err)
		sortedexp = append(sortedexp, ki)
	}
	sort.Ints(sortedexp)
	fmt.Println("\nCurrent payments by day:")
	for _, k := range sortedexp {
		fmt.Printf("%d\t%.2f\n", k, exp[strconv.Itoa(k)])
	}
	fmt.Println("\nDelay a payment? (Day of payment followed by new day, e.g. '1 10')")
	fmt.Println("Optionally specify an amount to defer (e.g. '1 10 456.45')")
	fmt.Println("\t0 or blank will use the full payment amount")
	fmt.Println("Specify one day to defer per line and an empty line when done")
	for {
		d0 = 0
		d1 = 0
		d2 = 0.0
		fmt.Print("> ")
		_, err = fmt.Scanln(&d0, &d1, &d2)
		if err != nil {
			switch err.Error() {
			case "unexpected newline":
				if d0 == 0 {
					eodd = true
				}
				if d0 > 0 && d1 == 0 {
					fmt.Println("Invalid input, try again")
					continue
				}
			case "expected integer":
				if d0 > 0 {
					fmt.Println("Second argument is invalid, try again")
					continue
				} else {
					fmt.Println("First argument is invalid, try again")
					continue
				}
			case "expected newline":
				fmt.Println("Too many arguments but that's ok\nYou'll do better next time")
			default:
				fmt.Println("Third argument is invalid, try again")
				continue
			}
		}
		if eodd {
			fmt.Println("")
			break
		}
		_, _ = DelayPayments(d0, d1, d2)
	}

	// Ask user if there are any additional payments (first occurence of day)
	var e0 int
	var e1 float64
	eoad := false
	fmt.Println("\nAny additional payments? (Day of payment followed by amount, e.g. '7 600.00')")
	fmt.Println("Specify one payment per line and an empty line when done")
	for {
		e0 = 0
		e1 = 0.0
		fmt.Print("> ")
		_, err = fmt.Scanln(&e0, &e1)
		if err != nil {
			switch err.Error() {
			case "unexpected newline":
				if e0 == 0 {
					eoad = true
				}
				if e0 > 0 && e1 == 0.0 {
					fmt.Println("Missing amount, try again")
					continue
				}
			case "expected integer":
				fmt.Println("First argument is invalid, try again")
				continue
			case "expected newline":
				fmt.Println("Too many arguments but that's ok\nYou'll do better next time")
			}
		}
		if eoad {
			fmt.Println("")
			break
		}
		_ = AddPayment(e0, e1)
	}

	// Ask user if there are any payments to exclude (first occurence of day)
	var e2 int
	var e3 float64
	neoad := false
	fmt.Println("\nAny payments to exclude? (Day of payment followed by amount, e.g. '7 600.00')")
	fmt.Println("Specify one payment per line and an empty line when done")
	for {
		e2 = 0
		e3 = 0.0
		fmt.Print("> ")
		_, err = fmt.Scanln(&e2, &e3)
		if err != nil {
			switch err.Error() {
			case "unexpected newline":
				if e2 == 0 {
					neoad = true
				}
				if e2 > 0 && e3 == 0.0 {
					fmt.Println("Missing amount, try again")
					continue
				}
			case "expected integer":
				fmt.Println("First argument is invalid, try again")
				continue
			case "expected newline":
				fmt.Println("Too many arguments but that's ok\nYou'll do better next time")
			}
		}
		if neoad {
			fmt.Println("")
			break
		}
		_ = ExcludePayment(e2, e3)
	}

	// Convert account balance to float
	bf, err = strconv.ParseFloat(b, 64)
	check(err)

	// Main loop:
	ft := tn
	var sfd string
	dd, err := time.ParseDuration("24h")
	check(err)
	sm := bf
	smperdiem := 9999.0
	var perdiem float64
	var obf float64
	for _, payday := range pda {
		obf = bf
		for ft.Before(payday) {
			ft = ft.Add(dd)
			sfd = strconv.Itoa(ft.Day())
			if ft.Equal(payday) {
				// Print subln and new bf
				fmt.Println(fmt.Sprintf("%.2f", obf) + subln)
				fmt.Printf("%.2f\n", bf)
				if bf < sm {
					sm = bf
				}
				perdiem = bf / (ft.Sub(tn).Hours() / 24.0)
				if perdiem < smperdiem {
					smperdiem = perdiem
				}
				obf = bf
				_ = GetPaid(*payPtr)
				fmt.Println(fmt.Sprintf("%.2f", obf), "+", *payPtr, "/*", ft.Format("2006-01-02"), "*/", fmt.Sprintf("(%.2f per diem)", perdiem))
				fmt.Printf("%.2f\n", bf)
				_ = UpdateSubLn(sfd, true)
				break
			}
			// Check if sfd exists in exp
			// if yes, subtract exp[sfd] from bf and add to subln
			if *mdvipPtr {
				if int(ft.Month())%3 == 0 && ft.Day() == 1 && !mdvipadded {
					_ = AddMDVIP()
				} else if mdvipadded {
					_ = SubMDVIP()
				}
			}
			_ = UpdateSubLn(sfd, false)
		}
	}
	fmt.Println("")
	fmt.Println("Discretionary funds: $" + fmt.Sprintf("%.2f", sm))
	fmt.Println("")
	fmt.Println("Per diem: $" + fmt.Sprintf("%.2f", smperdiem))
}
