"""
两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

🔗 https://leetcode-cn.com/problems/two-sum/
"""


class Solution(object):
    def twoSum(self, nums, target):
        """
        暴力破解
        """
        length = len(nums)
        for x in range(length):
            for y in range(x + 1, length):
                if nums[x] + nums[y] == target:
                    return [x, y]


class Solution_v1(object):
    def twoSum(self, nums, target):
        """
        `in` 和 list.index 的方法
        """
        length = len(nums)
        for x in range(length):
            y = target - nums[x]
            if y in nums and nums.index(y) != x:
                return [x, nums.index(y)]


class Solution_v2(object):
    def twoSum(self, nums, target):
        """
        使用哈希表
        """
        temp = {x: i for i, x in enumerate(nums)}
        for x in range(len(nums)):
            val = target - nums[x]
            if val in temp and x != temp.get(val):
                return [x, temp.get(val)]


class Solution_v3(object):
    def twoSum(self, nums, target):
        """一次遍历解决
        """
        a = {}
        for i, x in enumerate(nums):
            if target - x in a:
                return [a.get(target - x), i]
            a[x] = i


if __name__ == "__main__":
    print(Solution_v2().twoSum([2, 5, 5, 11], 10))
