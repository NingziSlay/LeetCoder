# coding: utf-8
"""
字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

🔗 https://leetcode-cn.com/problems/multiply-strings/
"""

class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        if num1 == "1":
            return num2
        if num2 == "1":
            return num1
        multiplier = 10
        res = 0
        for i, x in enumerate(num1[::-1]):
            m1 = multiplier ** i
            x = self.str2int(x) * m1
            for i1, y in enumerate(num2[::-1]):
                m2 = multiplier ** i1
                y = self.str2int(y) * m2
                res += (x * y)
        return str(res)

    @staticmethod
    def str2int(char):
        return ord(char) - ord("0")


if __name__ == '__main__':
    print(Solution().multiply("12", "123"))
