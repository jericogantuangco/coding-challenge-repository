
from typing import List


def isPalindrome(s: str ) -> bool:
    """
    This functions checks whether a given string is a palindrome
    """
    return s == s[::-1]


def palindromePairs(words: List[str]) -> List[str]:
    """
    This function returns a list of indices that is a concatenation of two words and form a palindrome
    """
    database = {}
    for idx, word in enumerate(words):
        starter = idx + 1
        end = len(words)    
        
        for jdx in range(starter, end, 1):

            for_comparison = words[jdx]
            concatenation = word + for_comparison

            if isPalindrome(concatenation) and concatenation not in database:
                database[concatenation] = [idx, jdx]
            elif isPalindrome(concatenation) and concatenation in database:
                database[concatenation].extend([idx, jdx])


    outer_starter = len(words) -1
    end = -1
    for kdx in range(outer_starter, end, -1):
        starter = kdx - 1
        end = -1
        
        for ldx in range(starter, end, -1):

            for_comparison = words[ldx]
            concatenation = words[kdx] + for_comparison

            if isPalindrome(concatenation) and concatenation not in database:
                database[concatenation] = [kdx, ldx]
            elif isPalindrome(concatenation) and concatenation in database:
                database[concatenation].extend([kdx, ldx])

    result = []
    
    for item in database:
        
        if len(database[item]) > 2:
            for idx in range(0, len(database[item]), 2):
                result.append([database[item][idx], database[item][idx + 1]])
        else:
            result.append(database[item])

    return result
    
            



if __name__ == '__main__':
    
    print(palindromePairs(['bat', 'tab', 'cat']))
    print(palindromePairs(["abcd","dcba","lls","s","sssll"]))
    print(palindromePairs(["a", ""]))
    print(palindromePairs(['ci', 'vic', 'no', 'on']))
    print(palindromePairs(["baab", "abaabaa", "aaba", ""]))
    ## add test case here