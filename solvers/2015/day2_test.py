import unittest

from day2 import SolveDay2Part1
from day2 import SolveDay2Part2

class Day2Test(unittest.TestCase):

    def test_part1(self):
        testcases = [
            {"input": "2x3x4", "expected": 34},
            {"input":"1x1x10", "expected": 14},
        ]
        for case in testcases:
            actual = SolveDay1Part1(case["input"])
            self.assertEqual(case["expected"], actual, "failed the actual test {} expected {}, actual {}".format(
                    case["input"], case["expected"], actual
                ),)