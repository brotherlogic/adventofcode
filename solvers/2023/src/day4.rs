use std::collections::HashMap;


pub fn solve_day4_part1(data: String) -> u32 {
    return 0
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
        let mut test_cases: HashMap<String, u32> = HashMap::new();
        test_cases.insert("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53".to_string(), 8);
        test_cases.insert("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19".to_string(), 2);
        test_cases.insert("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1".to_string(), 2);
        test_cases.insert("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83".to_string(), 1);
        test_cases.insert("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36".to_string(), 0);
        test_cases.insert("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".to_string(), 0);
        test_cases.insert("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".to_string(), 13);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day4_part1(case.to_string()),answer);
        } 
    }
}