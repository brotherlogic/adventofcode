struct Route {
    name: String,
    left: String,
    right: String,
}

pub fn solve_day8_part1(data: String) -> i32 {
    return 0;
}


#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
       let test_case = "RL

       AAA = (BBB, CCC)
       BBB = (DDD, EEE)
       CCC = (ZZZ, GGG)
       DDD = (DDD, DDD)
       EEE = (EEE, EEE)
       GGG = (GGG, GGG)
       ZZZ = (ZZZ, ZZZ)".to_string();

       let score = solve_day8_part1(test_case);
       assert_eq!(score, 2)
    }

    #[test]
    fn part1_tests_other() {
       let test_case = "LLR

       AAA = (BBB, BBB)
       BBB = (AAA, ZZZ)
       ZZZ = (ZZZ, ZZZ)".to_string();

       let score = solve_day8_part1(test_case);
       assert_eq!(score, 6)
    }
}