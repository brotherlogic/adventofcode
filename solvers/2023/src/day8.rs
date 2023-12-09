use std::process;

#[derive(Debug)]
struct Route {
    name: String,
    left: String,
    right: String,
}

fn build_routes(data: String) -> (String, Vec<Route>) {
    let mut elems = data.split("\n");
    let path = elems.next().unwrap().to_string();

    // Skip the blank line
    elems.next();

    let mut routes = Vec::new();
    for line in elems {
        let name = line[0..3].to_string();
        let left = line[7..10].to_string();
        let right = line[12..15].to_string();
        routes.push(Route{name: name, left: left, right: right});
    }

    return (path, routes)
}

pub fn solve_day8_part1(data: String) -> i32 {
    let (path, routes) = build_routes(data);

    let mut steps = 0;
    let mut first = "AAA".to_string();
    while first != "ZZZ" {
        for route in &routes {
            if route.name == first {
                match path.as_bytes()[steps%path.len()] as char {
                    'L' => first = "".to_string() + &route.left,
                    'R' => first = "".to_string() + &route.right,
                    _ => process::exit(1),
                  }
                  break;
                }
            }
        steps+=1;
    }
 
    return steps.try_into().unwrap();
}

fn allz(starts: &Vec<String>) -> bool {
    for start in starts {
        if &start[2..3] != "Z" {
            return false;
        }
    }
    return true;
}

pub fn solve_day8_part2(data: String) -> i32 {
    let (path, routes) = build_routes(data);

    let mut steps = 0;
    let mut starts: Vec<String> = Vec::new();
    for route in &routes {
        if &route.name[2..3] == "A" {
            starts.push("".to_string() + &route.name);
        }
    }

    println!("FOUND {} STARTS", starts.len());

    while !allz(&starts) {
        let mut next: Vec<String> = Vec::new();
        while starts.len() > 0 {
            let val = starts.pop().unwrap();
            for route in &routes {
                if route.name == val {
                    match path.as_bytes()[steps%path.len()] as char {
                        'L' => next.push("".to_string() + &route.left),
                        'R' => next.push("".to_string() + &route.right),
                        _ => process::exit(1),
                    }
                      break;
                }
            }
        }
        for v in next {
            starts.push(v);
        }
        steps+=1;
    }
    return steps.try_into().unwrap();
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

    #[test]
    fn part2_tests() {
        let test_case = "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)".to_string();

let score = solve_day8_part2(test_case);
       assert_eq!(score, 6)
    }
}