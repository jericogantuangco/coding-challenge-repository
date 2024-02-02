# Coding Challenge


## Description
A coding challenge to solve three problems that vary in difficulty from easy to moderately challenging.
The code used to implement the solutions for the coding problem are Python and GoLang.


## Table of Contents
- [Prerequisites](#prerequisites)
- [Directory Structure](#directory-structure)
- [Usage](#usage)



## Prerequisites
Python 3.12.1 was used to implement the Python solution. For Golang, version go1.21.6 windows/amd64 was used.
Both versions need to be installed to be able to run the solution for the problem sets.
Refer to this (Python installation)[https://docs.python.org/3/using/windows.html] and this (GoLang installation)[https://go.dev/doc/install] on how to install it in your system.


## Directory Structure
```
| README.md
|
|--problem_set_1
|    |    palindrom_pairs.go
|    |    palindrom_pairs.py
|    |    README.md
|--problem_set_2
|    |    valid_parentheses.go
|    |    valid_parentheses.py
|    |    README.md
|--problem_set_3
|    |    longest_increasing_subsequence.go
|    |    longest_increasing_subsequence.py
|    |    README.md
```

## Usage
To run the Python solution sets, simply type in 'python <file>' in the problem set directory you want to try to run.
For example, if you want to run the Palindrome Pairs problem solution, go to the problem_set_1 directory and execute the command as seen below.
```
$python palindrome_pairs.py
```

It should have a similiar out below
```
[[0, 1], [1, 0]]
[[0, 1], [2, 4], [3, 2], [1, 0]]
[[0, 1], [1, 0]]
[[0, 1], [2, 3], [3, 2]]
[[0, 3], [3, 0], [2, 1]]
```

For Golang, it is typically the same thing of navigating to the problem set and executing the command below.
```
$go run palindrome_pairs.go
```

And, the output below.
```
[[0 1] [1 0]]
[[3 2] [1 0] [0 1] [2 4]]
[[0 1] [1 0]]
[[2 3] [3 2] [0 1]]
[[2 1] [0 3] [3 0]]
```
