fn get_next_num(nums: Vec<i32>) -> i32 {
    let mut diffs: Vec<i32> = Vec::new();

    let mut iter = nums.iter();
    let mut prev = iter.next().unwrap();
    for num in iter {
        diffs.push(num-prev);
        prev = num
    }

    for d in &diffs {
        if d != &0 {
            return prev+get_next_num(diffs);
        }
    }

    return *prev;
}

fn get_next_num_rev(nums: Vec<i32>) -> i32 {
    let mut diffs: Vec<i32> = Vec::new();

    let mut iter = nums.iter();
    let mut prev = iter.next().unwrap();
    let first = prev;
    for num in iter {
        diffs.push(num-prev);
        prev = num
    }

    for d in &diffs {
        if d != &0 {
            return first-get_next_num(diffs);
        }
    }

    return *first;
}

fn get_next(line: String) -> i32 {
    let mut nums: Vec<i32> = Vec::new();
    for snum in line.split_whitespace() {
        nums.push(snum.parse::<i32>().unwrap());
    }

    return get_next_num(nums);
}

fn get_next_rev(line: String) -> i32 {
    let mut nums: Vec<i32> = Vec::new();
    for snum in line.split_whitespace() {
        nums.push(snum.parse::<i32>().unwrap());
    }

    return get_next_num_rev(nums);
}

pub fn solve_day9_part1(data: String) -> i32 {
    let mut total: i32 = 0;

    for line in data.split("\n") {
        total += get_next(line.to_string());
    }

    return total
}

pub fn solve_day9_part2(data: String) -> i32 {
    let mut total: i32 = 0;

    for line in data.split("\n") {
        total += get_next_rev(line.to_string());
    }

    return total
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_tests_focus() {
       let test_case = "0 3 6 9 12 15".to_string();

       let score = solve_day9_part1(test_case);
       assert_eq!(score, 18)
    }

     #[test]
    fn part1_tests() {
       let test_case = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45".to_string();

       let score = solve_day9_part1(test_case);
       assert_eq!(score, 114)
    }


   // #[test]
    fn part2_tests() {
       let test_case = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45".to_string();

       let score = solve_day9_part2(test_case);
       assert_eq!(score, 2)
    }

    #[test]
    fn part2_tests_focus() {
       let test_case = "1 3 6 10 15 21".to_string();

       let score = solve_day9_part2(test_case);
       assert_eq!(score, 0)
    }
}