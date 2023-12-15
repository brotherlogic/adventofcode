
pub fn solve_day15_part1(data: String) -> i32 {
    let mut total = 0;
    for elem in data.split(",") {
        total += run_hash(elem.to_string());
    }
    return total;
}

fn run_hash(s: String) -> i32 {
    let mut hash = 0;

    for c in s.chars() {
        let mut val = c.to_ascii_lowercase() as i32;
        if !c.is_ascii_lowercase() {
            val = c.to_ascii_uppercase() as i32;
        }
        hash += val;
        hash *= 17;
        hash %= 256;

        println!("DONE -> {}", hash);
    }


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

    #[test]
    fn part1_test_full() {
        let test_case = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7".to_string();
        let score = solve_day15_part1(test_case);
        assert_eq!(score, 1320);
    }
}