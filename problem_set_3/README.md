# Problem Set 1: Longest Increasing Subsequence

## Problem Description
Given an unsorted array of integers, find the length of the longest increasing subsequence.

## Solution Overview
Made use of a list that will save the sequence and take note of the longest increasing one by comparing it to previous items in the list. It should iterate through the list (outer point) and iterate through the previous number (inner point). If the outer point is greater than the inner point, we add plus one to the inner point value. Then, we save the greater value between the outer point value and the inner point value plus one. 

After all iterations, we then find the highest value in the list and return it.

## Instructions to run the code
```
$ python longest_increasing_subsequence.py
```
