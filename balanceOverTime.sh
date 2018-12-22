#!/bin/bash

###
# Balance Over Time
###
#
# TODO:
# - If an expense day is on a non-existant day (i.e. day 29 of February) then it
#   should be applied on the last day of the month instead.
#

# https://stackoverflow.com/a/29754866
! getopt --test > /dev/null
if [[ ${PIPESTATUS[0]} -ne 4 ]]
then
  echo "getopt --test failed in this environment"
  exit 1
fi

OPTIONS=c:d:e:f:p:
LONGOPTS=current-balance:,payday:,expenses:,following-month-end-day:,pay:

! PARSED=`getopt --options=$OPTIONS --longoptions=$LONGOPTS --name "$0" -- "$@"`
if [[ ${PIPESTATUS[0]} -ne 0 ]]
then
  exit 2
fi

eval set -- "$PARSED"

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

BALANCE=
DIFM=7
DOM=`date +%d`
DOW=`date +%w`
EXPENSES=()
PAY=
PAYDAY=

while true
do
  case "$1" in
    -c|--current-balance)
      BALANCE="$2"
      shift 2
      ;;
    -d|--payday)
      PAYDAY="$2"
      shift 2
      ;;
    -e|--expenses)
      EXPENSES="$2"
      shift 2
      ;;
    -f|--following-month-end-day)
      DIFM="$2"
      shift 2
      ;;
    -p|--pay)
      PAY="$2"
      shift 2
      ;;
    --)
      shift
      break
      ;;
    *)
      echo "Programming error"
      exit 3
      ;;
  esac
done

while [[ ! $PAYDAY =~ ^[0-6]$ ]]
do
  echo -e "\033[32mWhat day of the week do you get paid? (0..6); 0 is Sunday\033[0m"
  read -e -p 'Payday> ' PAYDAY CRUFT
done

while [[ ! $PAY =~ ^[0-9]{3,4}\.[0-9]{2}$ ]]
do
  echo -e "\n\033[32mHow much do you get paid a week? (xxxx.xx)\033[0m"
  read -e -p 'Pay> ' PAY CRUFT
done

while [[ ! $BALANCE =~ ^[0-9]{1,}\.[0-9]{2}$ ]]
do
  echo -e "\n\033[32mCurrent account balance? (x.xx .. xxxx.xx)\033[0m"
  read -e -p 'Current Balance> ' BALANCE CRUFT
done

while true
do
  TAINTED=1
  if [[ ${#EXPENSES} -eq 0 ]]
  then
    TAINTED=0
  fi
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
    echo -e "\n\033[32mExpenses? (dd,x.xx dd,x.xx ..)\033[0m"
    read -e -p 'Expenses> ' -a EXPENSES
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
    echo -e "Day: $COUNTER\nDifference: \033[32m+$PAY\033[0m\nBalance: \033[32m$BALANCE\033[0m\n"
  fi
  for x in ${EXPENSES[@]}
  do
    EDOM=`echo $x | cut -d',' -f1`
    ECOST=`echo $x | cut -d',' -f2`
    if [ $COUNTER -eq $EDOM ]
    then
      BALANCE=`echo "$BALANCE - $ECOST" | bc`
      echo -e "Day: $COUNTER\nDifference: \033[31m-$ECOST\033[0m\nBalance: \033[31m$BALANCE\033[0m\n"
    fi
  done
done

COUNTER=0
while [ $COUNTER -lt $DIFM ]
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
    echo -e "Day: $COUNTER\nDifference: \033[32m+$PAY\033[0m\nBalance: \033[32m$BALANCE\033[0m\n"
  fi
  for x in ${EXPENSES[@]}
  do
    EDOM=`echo $x | cut -d',' -f1`
    ECOST=`echo $x | cut -d',' -f2`
    if [ $COUNTER -eq $EDOM ]
    then
      BALANCE=`echo "$BALANCE - $ECOST" | bc`
      echo -e "Day: $COUNTER\nDifference: \033[31m-$ECOST\033[0m\nBalance: \033[31m$BALANCE\033[0m\n"
    fi
  done
done
