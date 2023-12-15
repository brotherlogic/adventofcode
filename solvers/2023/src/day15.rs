
pub fn solve_day15_part1(data: String) -> i32 {
    let mut total = 0;
    for elem in data.split(",") {
        total += run_hash(elem.to_string());
    }
    return total;
}

fn run_hash(s: String) -> i32 {
    let mut hash = 0;

    return hash;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
    let test_case = "HASH".to_string();

    let score = run_hash(test_case);
    assert_eq!(score, 52)
    }
}