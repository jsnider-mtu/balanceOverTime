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

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    datePtr := flag.String("d", "2020-07-10", "Last date paid (YYYY-MM-DD)")
    payPtr := flag.Float64("p", 2905.63, "How much are you paid?")

    flag.Parse()

    // Verify datePtr value is correct format
    t, err := time.Parse("2006-01-02", *datePtr)
    check(err)

    // Get reference date
    var ref time.Time
    tn := time.Now()
    tn, err = time.Parse("2006-01-02", tn.Format("2006-01-02"))
    if diff := tn.Sub(t); diff.Hours() >= 336 {
        mod := int(diff.Hours()) % 336
        dmod, _ := time.ParseDuration(strconv.Itoa(mod) + "h")
        ref = tn.Add(-dmod)
    } else {
        ref = t
    }

    // Create array of next 5 paydays
    pd1, err := time.ParseDuration("336h")
    pd2, err := time.ParseDuration("672h")
    pd3, err := time.ParseDuration("1008h")
    pd4, err := time.ParseDuration("1344h")
    pd5, err := time.ParseDuration("1680h")
    pda := [5]time.Time{ref.Add(pd1), ref.Add(pd2), ref.Add(pd3), ref.Add(pd4), ref.Add(pd5)}

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
    var ex bool
    var pde float64
    var val float64
    var obf float64
    var subln string
    for _, payday := range pda {
        obf = bf
        for ft.Before(payday) {
            ft = ft.Add(dd)
            sfd = strconv.Itoa(ft.Day())
            if ft.Equal(payday) {
                // Print subln and new bf
                fmt.Println(fmt.Sprintf("%.2f", obf) + subln)
                fmt.Println(fmt.Sprintf("%.2f", bf))
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
            if val, ex = exp[sfd]; ex {
                bf -= val
                subln += " - " + fmt.Sprint(val)
            }
        }
    }
}
