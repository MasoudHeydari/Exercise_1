# Task3 - Calculator
A simple CLI calculator written in `Golang`.

## How to Run?
To run the `Calculator` run below command:
```bash
make run
```

>NOTE: make sure `make` command installed in your system.

Now give your calculations to `Calculator` line by line.
* Examples:
  * Valid inputs:
      ```bash
      make run
      Enter Calculation Lines:
      2
      12, -10
      -5  ,, + 10
      result is:  7
      ```
  * Invalid inputs cause `syntax error`:
    ```bash
    make run
    Enter Calculation Lines:
    2
    8,,10
    12-15
    syntax error - '12-15' not an integer
    ```
    because there is no valid delimiter(plus, comma or space) between `12` and `-15`.

## What was my approach for solving this task? 
Because it was mentioned in problem statement that it's forbidden to use `for` loop, so I used `Recursion` approach.