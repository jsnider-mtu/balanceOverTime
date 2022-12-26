package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
    "sort"
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
    mdvipPtr := flag.Bool("mdvip", false, "Do you pay MDVIP every quarter?")

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
            ymod := 0
            if ref.Day() == 30 || ref.Day() == 29 || ref.Day() == 28 {
                for x := 0; x < 5; x++ {
                    mon := int((float64(x)/2.0)+float64(ref.Month() + 1)) % 12
                    if mon == 0 {
                        mon = 12
                    } else if mon < int(ref.Month()) {
                        ymod = 1
                    }
                    tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year() + ymod, mon, ((x%2)*15)+15))
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
                        ymod = 1
                    }
                    tim, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%d", ref.Year() + ymod, mon, ((x%2)*15)+15))
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

    // Ask user if any payment's first occurence should be delayed
    var d0 int
    var d1 int
    var d2 float64
    dexp := make(map[string]string)
    ddexp := make(map[string]float64)
    vdexp := make(map[string]float64)
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
        fmt.Println(fmt.Sprintf("%d\t%.2f", k, exp[strconv.Itoa(k)]))
    }
    fmt.Println("\nDelay a payment? (Day of payment followed by new day, e.g. '1 10.00')")
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
        // Validate input
        if _, ok := exp[strconv.Itoa(d0)]; !ok {
            fmt.Println(fmt.Sprintf("%d is not an expense day", d0))
            continue
        }
        if _, ok := dexp[strconv.Itoa(d0)]; ok {
            fmt.Println(fmt.Sprintf("Day %d is already being deferred\nCtrl+C (or Cmd+C) to exit and start again", d0))
            continue
        }
        if d0 < 1 || d0 > 31 {
            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
            continue
        }
        if d1 < 1 || d1 > 31 {
            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
            continue
        }
        if d0 == d1 {
            fmt.Println("Days cannot match")
            continue
        }
        if d2 < 0.0 {
            fmt.Println("Amount deferred cannot be negative")
            continue
        }
        if d2 > exp[strconv.Itoa(d0)] {
            fmt.Println("Amount deferred cannot exceed the total payment")
            fmt.Println(fmt.Sprintf("Max for day %d: %.2f", d0, exp[strconv.Itoa(d0)]))
            continue
        }
        // ex.
        // Today == 10
        // dsd0 == 3; dsd1 == 20 expense on 3rd of next month moved to 20th of next month
        // dsd0 == 6; dsd1 == 3 expense on 6th of next month moved to 3rd of following month
        // dsd0 == 6; dsd1 == 7 expense on 6th of next month moved to 7th of next month
        // dsd0 == 12; dsd1 == 3 expense on 12th of this month moved to 3rd of next month
        // dsd0 == 12; dsd1 == 28 expense on 12th of this month moved to 28th of this month
        // dsd0 == 28; dsd1 == 26 expense on 28th of this month moved to 26th of next month
        if d0 > tn.Day() {
            // expense of this month
            if d1 > d0 {
                // expense moved to this month
                switch int(tn.Month()) {
                case 2:
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d0 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                        if d1 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    } else {
                        if d0 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                        if d1 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    }
                case 4,6,9,11:
                    if d0 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                default:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                }
            } else {
                // expense moved to next month
                switch int(tn.Month()) {
                case 1:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d1 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    } else {
                        if d1 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    }
                case 2:
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d0 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                    } else {
                        if d0 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                case 3,5,8,10:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                case 4,6,9,11:
                    if d0 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                default:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                }
            }
        } else {
            // expense of next month
            if d1 > d0 {
                // expense moved to next month
                switch int(tn.Month()) {
                case 1:
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d0 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                        if d1 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    } else {
                        if d0 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                        if d1 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    }
                case 2,4,6,7,9,11,12:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                default:
                    if d0 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                }
            } else {
                // expense moved to following month
                switch int(tn.Month()) {
                case 1:
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d0 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                    } else {
                        if d0 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                            continue
                        }
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                case 2,4,7,9:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                case 3,5,8,10:
                    if d0 > 30 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                case 6,11:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if d1 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                        continue
                    }
                default:
                    if d0 > 31 {
                        fmt.Println(fmt.Sprintf("%d is not a valid day", d0))
                        continue
                    }
                    if tn.Year() % 400 == 0 || (tn.Year() % 4 == 0 && tn.Year() % 100 != 0) {
                        if d1 > 29 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
                        }
                    } else {
                        if d1 > 28 {
                            fmt.Println(fmt.Sprintf("%d is not a valid day", d1))
                            continue
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
    }

    // Ask user if there are any additional payments (first occurence of day)
    var e0 int
    var e1 float64
    exexp := make(map[string]float64)
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
        if e0 < 1 || e0 > 31 {
            fmt.Println(fmt.Sprintf("%d is not a valid day", e0))
            continue
        }
        if e1 < 0.0 {
            fmt.Println("Amount cannot be negative")
            continue
        }
        if _, ok := exexp[strconv.Itoa(e0)]; ok {
            exexp[strconv.Itoa(e0)] += e1
            continue
        }
        exexp[strconv.Itoa(e0)] = e1
    }

    // Convert account balance to float
    bf, err := strconv.ParseFloat(b, 64)
    check(err)

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
            if *mdvipPtr {
                if int(ft.Month()) % 3 == 0 && ft.Day() == 1 && !mdvipadded {
                    exp["1"] += 450.0
                    mdvipadded = true
                //} else if int(ft.Month()) % 3 == 0 && ft.Day() == 2 && mdvipadded {
                } else if mdvipadded {
                    exp["1"] -= 450.0
                    mdvipadded = false
                }
            }
            if val, ex = exp[sfd]; ex {
                if _, ok := dexp[sfd]; !ok {
                    bf -= val
                    subln += " - " + fmt.Sprint(val)
                } else {
                    val2 := vdexp[sfd]
                    diff := val - val2
                    bf -= diff
                    subln += " - " + fmt.Sprint(diff)
                    ddexp[dexp[sfd]] += val2
                    delete(vdexp, sfd)
                    delete(dexp, sfd)
                }
            }
            if val, ex = ddexp[sfd]; ex {
                bf -= val
                subln += " - " + fmt.Sprint(val)
                delete(ddexp, sfd)
            }
            if val, ex = exexp[sfd]; ex {
                bf -= val
                subln += " - " + fmt.Sprint(val)
                delete(exexp, sfd)
            }
        }
    }
    fmt.Println("")
    fmt.Println("Discretionary funds: $" + fmt.Sprintf("%.2f", sm))
}
