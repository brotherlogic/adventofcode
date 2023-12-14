pub fn solve_day14_part1(data: String) -> i32 {
    let mut sum = 0;

    return sum;
}

#[cfg(test)]
mod testsca {
    use super::*;

#[test]
fn part1_test_first() {
   let test_case = "O....#....
   O.OO#....#
   .....##...
   OO.#O....O
   .O.....O#.
   O.#..O.#.#
   ..O..#O..O
   .......O..
   #....###..
   #OO..#....".to_string();

   let score = solve_day14_part1(test_case);
   assert_eq!(score, 136)
}
}