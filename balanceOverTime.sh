#!/bin/bash

###
# Balance Over Time
###
#
# Grab current date, current account balance, payday (start with weekly, expand later),
# expense days (day of month/amount), and normal pay amount
#
# Prompt for all of the above (minus date) due to this being a container image

case `date +%m` in
01|03|05|07|08|10|12)
  DICM=31
  ;;
02)
  DICM=28
  ;;
*)
  DICM=30
  ;;
esac

DOM=`date +%d`
DOW=`date +%w`

echo -e "\033[32mWhat day of the week do you get paid? (0..6); 0 is Sunday\033[0m"
read -e -p 'Payday> ' PAYDAY CRUFT
while [[ ! $PAYDAY =~ ^[0-6]$ ]]
do
  read -e -p 'Payday> ' PAYDAY CRUFT
done

echo -e "\n\033[32mHow much do you get paid a week? (xxxx.xx)\033[0m"
read -e -p 'Pay> ' PAY CRUFT
while [[ ! $PAY =~ ^[0-9]{3,4}\.[0-9]{2}$ ]]
do
  read -e -p 'Pay> ' PAY CRUFT
done

echo -e "\n\033[32mCurrent account balance? (x.xx .. xxxx.xx)\033[0m"
read -e -p 'Current Balance> ' BALANCE CRUFT
while [[ ! $BALANCE =~ ^[0-9]{1,}\.[0-9]{2}$ ]]
do
  read -e -p 'Current Balance> ' BALANCE CRUFT
done

echo -e "\n\033[32mExpenses? (dd,x.xx dd,x.xx ..)\033[0m"
while true
do
  TAINTED=1
  read -e -p 'Expenses> ' -a EXPENSES
  for i in ${EXPENSES[@]}
  do
    if [[ ! $i =~ , ]]
    then
      TAINTED=0
      break
    fi
    EDOM=`echo $i | cut -d',' -f1`
    ECOST=`echo $i | cut -d',' -f2`
    if [[ ! $EDOM =~ ^[0-3][0-9]{0,1}$ ]]
    then
      TAINTED=0
      break
    fi
    if [[ ! $EDOM -le 31 ]]
    then
      TAINTED=0
      break
    fi
    if [[ ! $ECOST =~ ^[0-9]{1,}\.[0-9]{2}$ ]]
    then
      TAINTED=0
      break
    fi
  done
  if [[ $TAINTED = 0 ]]
  then
    continue
  else
    break
  fi
done

COUNTER=$DOM
COUNTERW=$DOW
echo -e "\n\033[32mCurrent day: $DOM\nCurrent balance: $BALANCE\n\033[0m"
while [ $COUNTER -lt $DICM ]
do
  let COUNTER++
  if [ $COUNTERW -lt 6 ]
  then
    let COUNTERW++
  else
    let COUNTERW=0
  fi
  if [ $COUNTERW -eq $PAYDAY ]
  then
    BALANCE=`echo "$BALANCE + $PAY" | bc`
    echo -e "Day: $COUNTER\nBalance: \033[32m$BALANCE\033[0m\n"
  fi
  for x in ${EXPENSES[@]}
  do
    EDOM=`echo $x | cut -d',' -f1`
    ECOST=`echo $x | cut -d',' -f2`
    if [ $COUNTER -eq $EDOM ]
    then
      BALANCE=`echo "$BALANCE - $ECOST" | bc`
      echo -e "Day: $COUNTER\nBalance: \033[31m$BALANCE\033[0m\n"
    fi
  done
done

COUNTER=0
while [ $COUNTER -lt 7 ]
do
  let COUNTER++
  if [ $COUNTERW -lt 6 ]
  then
    let COUNTERW++
  else
    let COUNTERW=0
  fi
  if [ $COUNTERW -eq $PAYDAY ]
  then
    BALANCE=`echo "$BALANCE + $PAY" | bc`
    echo -e "Day: $COUNTER\nBalance: \033[32m$BALANCE\033[0m\n"
  fi
  for x in ${EXPENSES[@]}
  do
    EDOM=`echo $x | cut -d',' -f1`
    ECOST=`echo $x | cut -d',' -f2`
    if [ $COUNTER -eq $EDOM ]
    then
      BALANCE=`echo "$BALANCE - $ECOST" | bc`
      echo -e "Day: $COUNTER\nBalance: \033[31m$BALANCE\033[0m\n"
    fi
  done
done