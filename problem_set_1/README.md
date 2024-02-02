# Problem Set 1: Palindrome Pairs

## Problem Description
Given a list of words, find all pairs of distinct indices(i,j) in the list so that the concatenation of the two words, i.e., words[i] + words[j], forms a palindrome.

## Solution Overview
The solution is to iterate through the list comparing the first word to the succeeding words, then saving the combination and its indices if its a palindrome. Then, do the same for the second word, comparing it to its succeeding words on the list. This however, does not take into account the words that can be formed backwards. So, we iterate through it backwards as well and save it again if its a palindrome.

Then, we need to prepare the list of lists (2 x 2) that would output the result.

## Instructions to run the code
```
$ python palindrome_pairs.py
```
