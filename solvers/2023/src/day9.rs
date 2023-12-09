pub fn solve_day9_part1(data: String) -> i32 {
    return 0;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let test_case = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45".to_string();

       let score = solve_day9_part1(test_case);
       assert_eq!(score, 2)
    }
}