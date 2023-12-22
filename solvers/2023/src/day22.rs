pub fn solve_day22_part1(data: String) -> i32 {
    return 0
}


#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "1,0,1~1,2,1
        0,0,2~2,0,2
        0,2,3~2,2,3
        0,0,4~0,2,4
        2,0,5~2,2,5
        0,1,6~2,1,6
        1,1,8~1,1,9".to_string();

        let pulses = solve_day22_part1(test_case);
        assert_eq!(pulses, 5)
    }
}