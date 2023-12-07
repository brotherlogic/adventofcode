pub fn solve_day7_part1(data: String) -> i64 {
    return 0;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let test_case = "32T3K 765
       T55J5 684
       KK677 28
       KTJJT 220
       QQQJA 483".to_string();

       let score = solve_day7_part1(test_case);
       assert_eq!(score, 6440)
    }
}