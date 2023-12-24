pub fn solve_day24_part1(data: String) -> i64 {
    return 0;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "19, 13, 30 @ -2,  1, -2
        18, 19, 22 @ -1, -1, -2
        20, 25, 34 @ -2, -2, -4
        12, 31, 28 @ -1, -2, -1
        20, 19, 15 @  1, -5, -3"
            .to_string();

        let pulses = solve_day24_part1(test_case);
        assert_eq!(pulses, 2)
    }
}
