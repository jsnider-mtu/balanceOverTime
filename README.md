# balanceOverTime
Go app to calculate checking account balance over time

You'll need a json file in $HOME/.local/etc/bot/expenses.json that looks like this:

`{"1": 1686.11, "2": 145.53, "5": 184.90, "12": 500.00, "18": 25.00, "20": 39.10, "23": 410.86, "25": 251.65, "29": 390.60}`

The keys are days of the month and values are how much gets taken out that day, usually autopay stuff

__The optional flags are:__
```
  -d string
        Last date paid (YYYY-MM-DD) (default "2020-07-10")
  -mdvip
        Do you pay MDVIP every quarter?
  -p float
        How much are you paid? (default 2732.23)
  -twoWeeks
        Do you get paid every two weeks?
```

You can see that output doing `bot.exe --help`

- `-d` is only relevant for the getting paid every two weeks workflow
- `-mdvip` and `-twoWeeks` are booleans that default to `false`
- `-p` is how much you receive per paycheck

__EXAMPLE:__

`bot.exe -mdvip -twoWeeks -d 2022-07-10 -p 2500`
