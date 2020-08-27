class Solution:
    def myAtoi(self, str: str) -> int:
        i = 0
        sign = 1
        idx = 0
        while idx < len(str) and str[idx] == ' ':
            idx += 1

        if idx == len(str):
            return 0

        if str[idx] == '-':
            sign = -1
            idx += 1
        elif str[idx] == '+':
            idx += 1
        while idx < len(str) and '0' <= str[idx] <= '9':
            if sign == 1:
                if i > 2 ** 31 // 10 or i == 2 ** 31 // 10 and str[idx] >= '7':
                    return 2 ** 31 - 1
                i = i * 10 + ord(str[idx]) - ord('0')
            else:
                if i < -2 ** 31 // 10 or i == -2 ** 31 // 10 + 1 and str[idx] >= '8':
                    return -2 ** 31
                i = i * 10 - ord(str[idx]) + ord('0')
            idx += 1
        return i


def test():
    sl = Solution()
    for s, i in [
        ("+1", 1),
        (" -42", -42),
        ("4193 with words", 4193),
        ("-91283472332", -2147483648),
        ("-2147483649", -2147483648),
        ("2147483649", 2147483647),
    ]:
        ans = sl.myAtoi(s)
        assert ans == i, f"{s}, {i}, {ans}"


if __name__ == "__main__":
    test()
