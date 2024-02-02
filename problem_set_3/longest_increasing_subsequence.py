from typing import List

def lengthOfLIS(nums: List[int]) -> int:
    
    end = len(nums)
    database = [1] * end

    for idx in range(1, end):
        counter = 0
        while counter != idx:
            if nums[idx] > nums[counter]:
                database[idx] = max(database[idx], database[counter] + 1)
            counter+=1

    return max(database)


if __name__ == '__main__':
    print(lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]))
    print(lengthOfLIS([0, 1, 0, 3, 2, 3]))
    print(lengthOfLIS([1, 1, 1, 1, 1, 1, 1]))