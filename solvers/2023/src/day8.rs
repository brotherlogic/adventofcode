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


#[derive(Debug)]
struct Looper {
    round: i64,
}

impl Looper {
    fn next(&self, i: i64) -> i64 {
       if i%&self.round == 0 {
        return self.round;
       }
       return i%self.round;
    }

    fn in_z(&self, i: i64) -> bool {
       return i%self.round == 0;
    }
}

fn build_looper(start: String, path: String, routes: &Vec<Route>) -> Looper {
    let mut done: Vec<(String, i64, i64)> = Vec::new();
    let mut steps = 0;
    let mut nstate = "".to_string()+&start;

    let  _loops = 0;
    let  _jump = 0;
    loop {
        if nstate[2..3] == *"Z" {
            for (_pos, (st, nu, _)) in done.iter().enumerate() {
                if st == &nstate && nu == &(steps%path.len() as i64){
                    for (_npos, (_st, _nu2, steps)) in done.iter().enumerate() {
                        return Looper{round: *steps};
                    }
                }
            }
            done.push((nstate.clone(), steps%path.len() as i64, steps));
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

pub fn solve_day8_part2(data: String) -> i64 {
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
        loopers.push(looper);
    }

    let mut pointer:i64 = 0;
    let mut _loops = 0;
    loop {
        _loops += 1;
     
        let mut highest = 0;
        for looper in &loopers {
            let val = looper.next(pointer);
            if val > highest {
                highest = val;
            }
        }
    
  
        pointer += highest;
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