pub fn solve_day11_part1(data: String) -> i64 {
    return 0;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
    let test_case = "...#......
    .......#..
    #.........
    ..........
    ......#...
    .#........
    .........#
    ..........
    .......#..
    #...#.....".to_string();

    let score = solve_day11_part1(test_case);
    assert_eq!(score, 374)
    }
}