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