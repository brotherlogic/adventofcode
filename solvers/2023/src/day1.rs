use std::collections::HashMap;

pub fn solve_day1_part2(data: String) -> i32 {
    let mut translate = HashMap::new();
    let mut ndata = data;
    translate.insert("one", "1");
    translate.insert("two", "2");
    translate.insert("three", "3");
    translate.insert("four", "4");
    translate.insert("five", "5");
    translate.insert("six", "6");
    translate.insert("seven", "7");
    translate.insert("eight", "8");
    translate.insert("nine", "9");
    for (key, val) in &translate {
        ndata = ndata.replace(key,&(key.to_string()+val+key));
    }
    return solve_day1_part1(ndata);
}

pub fn solve_day1_part1(data: String) -> i32 {
    let parts = data.split("\n");
    let mut value: i32 = 0;
    for part in parts {
        if part.len() > 0 {
    let mut left: i32 = -1;
    let mut right: i32 = -1;
    for c in part.chars() {
        if c.is_numeric() {
            left = c.to_digit(10).unwrap() as i32;
            break;
        }
    }
    for c in part.chars().rev() {
        if c.is_numeric() {
            right = c.to_digit(10).unwrap() as i32;
            break;
        }
    }
    value += (left*10 + right) as i32;
    println!("VALUE {} from {}", value, part);
    }
}
    return value
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests() {
        let mut test_cases: HashMap<String, i32> = HashMap::new();
        test_cases.insert("1abc2".to_string(), 12);
        test_cases.insert("pqr3stu8vwx".to_string(), 38);
        test_cases.insert("a1b2c3d4e5f".to_string(), 15);
        test_cases.insert("treb7uchet".to_string(), 77);
        test_cases.insert("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet".to_string(), 142);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day1_part1(case.to_string()),answer);
        } 
    }

    #[test]
    fn part2_tests() {
        let mut test_cases: HashMap<String, i32> = HashMap::new();
        test_cases.insert("two1nine".to_string(), 29);
        test_cases.insert("eightwothree".to_string(), 83);
        test_cases.insert("abcone2threexyz".to_string(), 13);
        test_cases.insert("xtwone3four".to_string(), 24);
        test_cases.insert("4nineeightseven2".to_string(), 42);
        test_cases.insert("zoneight234".to_string(),14);
        test_cases.insert("7pqrstsixteen".to_string(),76);
        test_cases.insert("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\noneight234\n7pqrstsixteen".to_string(), 281);

        for (case, answer) in &test_cases {
            assert_eq!(&solve_day1_part2(case.to_string()),answer);
        } 

    }
}