"""
整数反转：
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

🔗 https://leetcode-cn.com/problems/reverse-integer/
"""


class Solution:
    MAX = (1 << 31) - 1
    MIN = -1 << 31

    def reverse(self, x: int) -> int:
        MAX = 2 ** 31 - 1 if x > 0 else 2 ** 31
        abx_x, res = abs(x), 0
        while abx_x != 0:
            res = res * 10 + abx_x % 10
            abx_x = abx_x // 10
            if res > MAX:
                return 0
        return res if x > 0 else -res

    def reverse_v1(self, x: int) -> int:
        if -9 <= x <= 9:
            return x
        out = ""
        for a in str(x)[::-1]:
            out += a
        if out[0] == "0":
            out = out[1:]

        if out[-1] == "-":
            out = "-" + out[:-1]
        res = int(out)
        return res if self.MIN < res < self.MAX else 0


if __name__ == '__main__':
    print(Solution().reverse(8723872388242))
