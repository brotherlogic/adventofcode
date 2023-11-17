import unittest

from day1 import SolveDay1Part1

class Day1Test(unittest.TestCase):

    def test_part1(self):
        testcases = [
            {"input": "(())", "expected": 0},
            {"input":"()()", "expected": 0},
            {"input":"(((", "expected": 3},
            {"input":"(()(()(", "expected": 3},
            {"input":"))(((((", "expected": 3},
            {"input":"())", "expected": -1},
            {"input":"))(", "expected": -1},
            {"input":")))", "expected": -3},
            {"input":")())())", "expected": -4},
        ]
        for case in testcases:
            actual = SolveDay1Part1(case["input"])
            self.assertEqual(case["expected"], actual, "failed the actual test {} expected {}, actual {}".format(
                    case["input"], case["expected"], actual
                ),)

if __name__ == "__main__":
    unittest.main()