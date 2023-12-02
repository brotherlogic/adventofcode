use std::collections::HashMap;

pub fn solve_day1_part1(data: String) -> i32 {
    let parts = data.split("\n");
    let mut value: i32 = 0;
    for part in parts {
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
}