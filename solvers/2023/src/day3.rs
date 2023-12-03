pub fn solve_day3_part1(board: String) -> u32 {
    return 0
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let board = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..".to_string();
       let answer = solve_day3_part1(board);
       assert_eq!(answer, 4361);
    }



}