use std::process;
use std::collections::HashMap;

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

#[derive(Debug)]
struct Looper {
    preamble: Vec<i32>,
    round: Vec<i32>,
}

impl Looper {
    fn next(&self, i: i32) -> i32 {
        let mut sofar = i;
        if self.preamble.len() > 0 && i < self.preamble[self.preamble.len()-1] {
            for val in &self.preamble {
                if sofar > *val {
                    return sofar - val;
                }
                sofar -= val;
            }
        }

        println!("HUH1 {:?} -> {} {}", &self, sofar, i);
     
        sofar = sofar % self.round[self.round.len()-1];

        println!("HUH2 {:?} -> {} {}", &self, sofar, i);
        for val in &self.round {
            println!("VALUE {} {}", val, sofar);
            if sofar < *val {
                println!("RESOLVE {}", val-sofar);
                return val - sofar;
            }
            sofar -= val;
        }

        return 0;
    }

    fn in_z(&self, i: i32) -> bool {
        if self.preamble.len() > 0 && i < self.preamble[self.preamble.len()-1] {
            for val in &self.preamble {
                if *val == i {
                    return true;
                }
            }
            return false;
        }

        let sofar = i % self.round[self.round.len()-1];
        println!("CHECKING {}", sofar);
        if sofar == 0 {
            return true;
        }
        for val in &self.round {
            if *val == sofar {
                return true;
            }
        }

        return false;
    }
}

fn build_looper(start: String, path: String, routes: &Vec<Route>) -> Looper {
    let mut done: Vec<(String, i32, i32)> = Vec::new();
    let mut steps = 0;
    let mut nstate = "".to_string()+&start;

    let mut loops = 0;
    loop {
        if nstate[2..3] == *"Z" {
            for (pos, (st, nu, _)) in done.iter().enumerate() {
                if st == &nstate && nu == &(steps%path.len() as i32){

                    let mut preamble: Vec<i32> = Vec::new();
                    let mut round: Vec<i32> = Vec::new();

                    for (npos, (_st, nu2, steps)) in done.iter().enumerate() {
                        if npos < pos {
                            preamble.push(*steps);
                        } else {
                            round.push(*steps);
                        }
                    }
                    
                    return Looper{
                        preamble: preamble, 
                        round: round,
                    }
                }
            }
            done.push((nstate.clone(), steps%path.len() as i32, steps));
        }

        for route in routes {
            if route.name == nstate {
                if path.as_bytes()[steps as usize%path.len()] as char == 'L' {
                    nstate = "".to_string() + &route.left;
                } else {
                    nstate = "".to_string()+&route.right;
                }
                break;
            }
        }

        steps += 1;
    }
}

pub fn solve_day8_part2(data: String) -> i32 {
    let (path, routes) = build_routes(data);

    let mut starts: Vec<String> = Vec::new();
    for route in &routes {
        if &route.name[2..3] == "A" {
            starts.push("".to_string() + &route.name);
        }
    }

    let mut loopers: Vec<Looper> = Vec::new();
    for start in starts {
        let looper = build_looper(start, "".to_string()+&path, &routes);
        println!("LOOPER {:?}", looper);
        loopers.push(looper);
    }

    let mut pointer = 0;
    let mut loops = 0;
    loop {
        println!("POINTER {}", pointer);
        loops += 1;
       
        let mut lowest = i32::MAX;
        for looper in &loopers {
            let val = looper.next(pointer);
            if val < lowest {
                lowest = val;
            }
        }

        println!("LOWEST {} -> {}", lowest, pointer+lowest);

        pointer += lowest;
        let mut found = false;
        for looper in &loopers {
            if !looper.in_z(pointer) {
                found = true;
                break;
            }
        }

        if !found {
            return pointer;
        }
    }
   return -1;
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