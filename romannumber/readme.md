## roman nummbers
### Instructions
Write a program called rn. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

The program should have a limit of 4000. In case of an invalid number, for example "hello" or 0 the program should print ERROR: cannot convert to roman digit.

Roman Numerals reminder:

| Roman Numeral | Value |
|---------------|-------|
| I             | 1     |
| V             | 5     |
| X             | 10    |
| L             | 50    |
| C             | 100   |
| D             | 500   |
| M             | 1000  |


For example, the number 1732 would be denoted MDCCXXXII in Roman numerals. However, Roman numerals are not a purely additive number system. In particular, instead of using four symbols to represent a 4, 40, 9, 90, etc. (i.e., IIII, XXXX, VIIII, LXXXX, etc.), such numbers are instead denoted by preceding the symbol for 5, 50, 10, 100, etc., with a symbol indicating subtraction. For example, 4 is denoted IV, 9 as IX, 40 as XL, etc.

The following table gives the Roman numerals for the first few positive integers.

| 1 | I  | 2  | II  | 3  | III  |
|---|----|----|-----|----|------|
| 4 | IV | 5  | V   | 6  | VI   |
| 7 | VII| 8  | VIII| 9  | IX   |
|10 | X  | 11 | XI  | 12 | XII  |
|13 | XIII | 14 | XIV | 15 | XV |
|16 | XVI | 17 | XVII | 18 | XVIII |
|19 | XIX | 20 | XX | 21 | XXI  |
|22 | XXII | 23 | XXIII | 24 | XXIV |
|25 | XXV | 26 | XXVI | 27 | XXVII |
|28 | XXVIII | 29 | XXIX | 30 | XXX |

## Usage
```
$ go run . hello
ERROR: cannot convert to roman digit
$ go run . 123
C+X+X+I+I+I
CXXIII
$ go run . 999
(M-C)+(C-X)+(X-I)
CMXCIX
$ go run . 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
$ go run . 4000
ERROR: cannot convert to roman digit
$
```