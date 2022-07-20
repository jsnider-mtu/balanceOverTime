package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
    "strconv"
    "time"
)

var (
    err error
    tim time.Time
    pda [5]time.Time
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    datePtr := flag.String("d", "2020-07-10", "Last date paid (YYYY-MM-DD)")
    payPtr := flag.Float64("p", 2732.23, "How much are you paid?")
    twoWeekPtr := flag.Bool("twoWeeks", false, "Do you get paid every two weeks?")

    flag.Parse()

    var ref time.Time
    tn := time.Now()
    tn, err = time.Parse("2006-01-02", tn.Format("2006-01-02"))
    check(err)

    if *twoWeekPtr {
        // Verify datePtr value is correct format
        t, err := time.Parse("2006-01-02", *datePtr)
        check(err)
    
        // Get reference date
        //var ref time.Time
        //tn := time.Now()
        //tn, err = time.Parse("2006-01-02", tn.Format("2006-01-02"))
        if diff := tn.Sub(t); diff.Hours() >= 336 {
            mod := int(diff.Hours()) % 336
            dmod, _ := time.ParseDuration(strconv.Itoa(mod) + "h")
            ref = tn.Add(-dmod)
        } else {
            ref = t
        }
    } else {
        // Reference date is either last 15th or 30th or 28th or 29th if Feb
        if tn.Month() == 3 && tn.Day() < 15 {
            if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                ref, err = time.Parse("2006-01-02", strconv.Itoa(tn.Year()) + "-02-29")
                check(err)
            } else {
                ref, err = time.Parse("2006-01-02", strconv.Itoa(tn.Year()) + "-02-28")
                check(err)
            }
        } else {
            if tn.Day() < 15 {
                if int(tn.Month()) > 1 {
                    ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-30", tn.Year(), int(tn.Month()) - 1))
                    check(err)
                } else {
                    ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%d-30", tn.Year() - 1, 12))
                    check(err)
                }
            } else {
                ref, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-15", tn.Year(), int(tn.Month())))
                check(err)
            }
        }
    }

    if *twoWeekPtr {
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
        pda = [5]time.Time{ref.Add(pd1), ref.Add(pd2), ref.Add(pd3), ref.Add(pd4), ref.Add(pd5)}
    } else {
        // Create array of next 5 paydays, 15ths and 30ths and 28th or 29th if Feb
        switch int(ref.Month()) {
        case 12:
            if ref.Day() == 15 {
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-12-30", ref.Year()))
                check(err)
                pda[0] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-15", ref.Year() + 1))
                check(err)
                pda[1] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year() + 1))
                check(err)
                pda[2]= tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year() + 1))
                check(err)
                pda[3] = tim
                if ref.Year() + 1 % 400 == 0 || (ref.Year() + 1 % 4 == 0 && ref.Year() + 1 % 100 != 0) {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year() + 1))
                    check(err)
                    pda[4] = tim
                } else {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year() + 1))
                    check(err)
                    pda[4] = tim
                }
            } else {
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-15", ref.Year() + 1))
                check(err)
                pda[0] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year() + 1))
                check(err)
                pda[1] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year() + 1))
                check(err)
                pda[2] = tim
                if ref.Year() + 1 % 400 == 0 || (ref.Year() + 1 % 4 == 0 && ref.Year() + 1 % 100 != 0) {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year() + 1))
                    check(err)
                    pda[3] = tim
                } else {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year() + 1))
                    check(err)
                    pda[3] = tim
                }
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year() + 1))
                check(err)
                pda[4] = tim
            }
        case 1:
            if ref.Day() == 15 {
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-01-30", ref.Year()))
                check(err)
                pda[0] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()))
                check(err)
                pda[1] = tim
                if ref.Year() % 400 == 0 || (ref.Year() % 4 == 0 && ref.Year() % 100 != 0) {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
                    check(err)
                    pda[2] = tim
                } else {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
                    check(err)
                    pda[2] = tim
                }
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
                check(err)
                pda[3] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
                check(err)
                pda[4] = tim
            } else {
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-15", ref.Year()))
                check(err)
                pda[0] = tim
                if ref.Year() % 400 == 0 || (ref.Year() % 4 == 0 && ref.Year() % 100 != 0) {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
                    check(err)
                    pda[1] = tim
                } else {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
                    check(err)
                    pda[1] = tim
                }
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
                check(err)
                pda[2] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
                check(err)
                pda[3] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-15", ref.Year()))
                check(err)
                pda[4] = tim
            }
        case 2:
            if ref.Day() == 15 {
                if ref.Year() % 400 == 0 || (ref.Year() % 4 == 0 && ref.Year() % 100 != 0) {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-29", ref.Year()))
                    check(err)
                    pda[0] = tim
                } else {
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-02-28", ref.Year()))
                    check(err)
                    pda[0] = tim
                }
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-15", ref.Year()))
                check(err)
                pda[1] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-03-30", ref.Year()))
                check(err)
                pda[2] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-15", ref.Year()))
                check(err)
                pda[3] = tim
                tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-04-30", ref.Year()))
                check(err)
                pda[4] = tim
            } else {
                for x := 0; x < 5; x++ {
                    mon, err := strconv.Atoi(fmt.Sprintf("%.0f", (float64(x) / 2.0) + 3.0))
                    check(err)
                    tim, err = time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year(), mon, ((x % 2) * 15) + 15))
                    check(err)
                    pda[x] = tim
                }
            }
        default:
            if ref.Day() == 30 || ref.Day() == 29 || ref.Day() == 28 {
                for x := 0; x < 5; x++ {
                    mon := int((float64(x)/2.0)+float64(ref.Month())) % 12
                    if mon == 0 {
                        mon = 12
                    } else if mon < int(ref.Month()) {
                        ref = ref.Add(time.Hour * 24 * 365)
                    }
                    tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year(), mon, ((x%2)*15)+15))
                    if err != nil {
                        fmt.Println(err)
                    }
                    pda[x] = tim
                }
            } else {
                for x := 1; x <= 5; x++ {
                    mon := int((float64(x)/2.0)+float64(ref.Month())) % 12
                    if mon == 0 {
                        mon = 12
                    } else if mon < int(ref.Month()) {
                        ref = ref.Add(time.Hour * 24 * 365)
                    }
                    tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year(), mon, ((x%2)*15)+15))
                    if err != nil {
                        fmt.Println(err)
                    }
                    pda[x-1] = tim
                }
            }
        }
    }

    // Read expenses.json and unmarshal data
    dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.local/etc/bot/expenses.json")
    check(err)
    var exp map[string]float64
    json.Unmarshal(dat, &exp)

    // Ask user for current account balance
    var b string
    fmt.Println("Current balance?")
    _, err = fmt.Scanln(&b)
    check(err)

    // Convert account balance to float
    bf, err := strconv.ParseFloat(b, 64)
    check(err)

    fmt.Println("")

    // Main loop:
    ft := tn
    sfd := strconv.Itoa(ft.Day())
    dd, err := time.ParseDuration("24h")
    sm := bf
    var ex bool
    var pde float64
    var val float64
    var obf float64
    var subln string
    var mdvipadded bool = false
    for _, payday := range pda {
        obf = bf
        for ft.Before(payday) {
            ft = ft.Add(dd)
            sfd = strconv.Itoa(ft.Day())
            if ft.Equal(payday) {
                // Print subln and new bf
                fmt.Println(fmt.Sprintf("%.2f", obf) + subln)
                fmt.Println(fmt.Sprintf("%.2f", bf))
                if bf < sm {
                    sm = bf
                }
                pde = 0
                obf = bf
                bf += *payPtr
                fmt.Println(fmt.Sprintf("%.2f", obf), "+", *payPtr, "/*", ft.Format("2006-01-02"), "*/")
                fmt.Println(fmt.Sprintf("%.2f", bf))
                // if sfd in exp, set pde to exp[sfd]
                if val, ex = exp[sfd]; ex {
                    pde = val
                }
                // if pde != 0: start new subln
                // else: reset subln
                if pde == 0 {
                    subln = ""
                } else {
                    subln = " - " + fmt.Sprint(pde)
                }
                break
            }
            // Check if sfd exists in exp
            // if yes, subtract exp[sfd] from bf and add to subln
            if int(ft.Month()) % 3 == 0 && ft.Day() == 1 && !mdvipadded {
                exp["1"] += 450.0
                mdvipadded = true
            } else if int(ft.Month()) % 3 == 0 && ft.Day() == 2 && mdvipadded {
                exp["1"] -= 450.0
                mdvipadded = false
            }
            if val, ex = exp[sfd]; ex {
                bf -= val
                subln += " - " + fmt.Sprint(val)
            }
        }
    }
    fmt.Println("")
    fmt.Println("Discretionary funds: $" + fmt.Sprintf("%.2f", sm))
}
