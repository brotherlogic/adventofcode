use std::collections::HashMap;

pub fn solve_day20_part1(data: String) -> i64 {
    let computer = build_computer(data);
    let pulse_count = run_program(computer, 1000);
    return pulse_count;
}

fn run_program(computer: Vec<Computer>, reps: i32) -> i64 {
    let mut low_pulse = 0;
    let mut high_pulse = 0;
    let mut state_map = HashMap::new();
    let mut conj_mem = HashMap::new();

    for c in &computer {
        if matches!(c.tp, ComputerType::FlipFlop) {
            state_map.insert(c.name.clone(),  false);
        }
        if matches!(c.tp, ComputerType::Conjunction) {
            let mut conj_settings = HashMap::new();
            for c2 in &computer {
                for d in &c2.destination {
                    if *d == c.name {
                        conj_settings.insert(c2.name.clone().to_string(), false);
                    }
                }
            }
            conj_mem.insert(c.name.clone().to_string(), conj_settings);
        }
    }


    for _ in 0..reps {
    let mut pulses = Vec::new();
    pulses.push(Pulse{
        state: false,
        origin: "Button".to_string(),
        destination: "roadcaster".to_string(),
    });

    while pulses.len() > 0 {
        let cpulse = pulses.remove(0);
      if cpulse.state {
        high_pulse += 1;
      } else {
        low_pulse += 1;
      }
        let dest = &cpulse.destination;

        for comp in &computer {
           // println!("TRYING {:?}", comp);
            if comp.name == *dest {
               // println!("FOUND COMPUTER {:?}", comp);
                if matches!(comp.tp, ComputerType::FlipFlop) && !cpulse.state{
                    let istate = !(state_map.get(&comp.name).unwrap());
                    state_map.insert(comp.name.clone().to_string(), istate);
                    for dest2 in &comp.destination {
                        pulses.push(Pulse{
                            state: istate,
                            origin: comp.name.clone().to_string(),
                            destination: (dest2.clone()).to_string(),
                        })
                    }
                } else if matches!(comp.tp, ComputerType::Conjunction) {
                    let mut mapper = conj_mem.get(&comp.name).unwrap().clone();
                    mapper.insert(cpulse.origin.clone().to_string(), cpulse.state );
                    let mut active = true;
                    for (_, val) in &mapper {
                        if !val {
                            active = false;
                            break;
                        }
                    }
                    conj_mem.insert(comp.name.clone().to_string(), mapper);

                    for dest2 in &comp.destination {
                        pulses.push(Pulse{
                            state: !active,
                            origin: comp.name.clone().to_string(),
                            destination: (dest2.clone()).to_string(),
                        })
                    }    
                } else if matches!(comp.tp, ComputerType::Broadcast) {
                    for dest2 in &comp.destination {
                        pulses.push(Pulse{
                            state: cpulse.state,
                            origin: comp.name.clone().to_string(),
                            destination: (dest2.clone()).to_string(),
                        })
                    }  
                } 
            }
        }
    }
    }


    return low_pulse * high_pulse;
}

#[derive(Clone)]
#[derive(Debug)]
enum ComputerType {
    Broadcast,
    Conjunction,
    FlipFlop,
}

#[derive(Clone)]
#[derive(Debug)]
struct Computer {
    name: String,
    tp: ComputerType,
    destination: Vec<String>,
}

#[derive(Debug)]
struct Pulse {
    state: bool,
    destination: String,
    origin: String,
}


fn build_computer(data: String) -> Vec<Computer> {
    let mut computer = Vec::new();
    for line in data.split("\n") {
        let mut elems = line.trim().split("->");
        let first = elems.next().unwrap().to_string();
        let second = elems.next().unwrap();

        let mut outputs = Vec::new();
        for piece in second.split(",") {
            outputs.push(piece.trim().to_string());
        }

        match first.chars().next().unwrap() {
            '%' => computer.push( Computer{
                name: first[1..].trim().to_string(),
                tp: ComputerType::FlipFlop,
                destination: outputs,
            }),
            '&' => computer.push(Computer{
                name: first[1..].trim().to_string(),
                tp: ComputerType::Conjunction,
                destination: outputs,
            }),
            _ => computer.push(Computer{
                name: first[1..].trim().to_string(),
                tp: ComputerType::Broadcast,
                destination: outputs,
        }),
        }
    }

    return computer;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "broadcaster -> a, b, c
        %a -> b
        %b -> c
        %c -> inv
        &inv -> a".to_string();

        let pulses = solve_day20_part1(test_case);
        assert_eq!(pulses, 32000000)
    }
    #[test]
    fn part1_test_second() {
        let test_case = "broadcaster -> a
        %a -> inv, con
        &inv -> b
        %b -> con
        &con -> output".to_string();

        let pulses = solve_day20_part1(test_case);
        assert_eq!(pulses, 11687500)
    }
}

