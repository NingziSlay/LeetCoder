# coding: utf-8
from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        a = dict()
        for i, x in enumerate(nums):
            f = a.get(x)
            if f is not None and i - f <= k:
                return True
            a[x] = i
        return False


if __name__ == '__main__':
    print(Solution().containsNearbyDuplicate([1, 0, 1, 1], 1))
