use std::collections::HashMap;

pub fn solve_day15_part1(data: String) -> i32 {
    let mut total = 0;
    for elem in data.split(",") {
        total += run_hash(elem.to_string());
    }
    return total;
}

pub fn solve_day15_part2(data: String) -> i32 {
    let mut boxes: HashMap<i32,Vec<(String,i32)>> = HashMap::new();

    for i in 0..256 {
        boxes.insert(i,Vec::new());
    }

    for elem in data.split(",") {
        let mut temp = elem.to_string();
        if temp.contains("=") {
            let mut parts = temp.split("=");
            if elem.to_string().chars().last().unwrap().is_digit(10) {
                let label = parts.next().unwrap();
                let length = parts.next().unwrap();

                let boxc = run_hash(label.to_string());
                let ln = length.parse::<i32>().unwrap();

                // Build out a new vector
                let mut found = false;
                let mut nvec: Vec<(String,i32)> = Vec::new();
                let mut cvec: Vec<(String,i32)> = boxes.get(&boxc).unwrap().to_vec();
                println!("ADDING TO {:?}", cvec);
                for (nl, nv) in &cvec {
                    if nl == label {
                        found = true;
                        println!("REPLACE -> {}", nl);
                        nvec.push((nl.to_string(), ln));
                    } else {
                        nvec.push((nl.to_string(), *nv));
                    }
                }
                if !found {
                    println!("NEW -> {}", label);
                    nvec.push((label.to_string(),ln));
                }
                boxes.insert(boxc,nvec);
            }
        } else if temp.contains("-") {
            let mut parts = temp.split("-");
            let label = parts.next().unwrap();
        
            let boxc = run_hash(label.to_string());
        
            // Build out a new vector
            let mut nvec: Vec<(String,i32)> = Vec::new();
            let mut cvec: Vec<(String,i32)> = boxes.get(&boxc).unwrap().to_vec();
            println!("FORM {:?}", cvec);
           for (nl,nv) in &cvec {
                if nl != label {
                    nvec.push((nl.to_string(), *nv));
                } else {
                    println!("REMOVE {}", nl);
                }
            }
            boxes.insert(boxc,nvec);
        }
    }

    println!("RESULT = {:?}", boxes);

    let mut total = 0;
    for (bv, boxc) in boxes {
        for (slot, (label, number)) in boxc.iter().enumerate() {
            println!("ADDING {} {} {}", bv, slot+1, number);
            total += (bv+1) * (slot as i32+1) * number;
        }
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

    #[test]
    fn part2_test_full() {
        let test_case = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7".to_string();
        let score = solve_day15_part2(test_case);
        assert_eq!(score, 145);
    }
}