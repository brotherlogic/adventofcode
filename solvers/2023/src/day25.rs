use std::collections::HashMap;

pub fn solve_day25_part1(data: String) -> i32 {
    let circuit = build_circuit(data);

    for x in 0..circuit.wires.len() {
        for y in x + 1..circuit.wires.len() {
            for z in y + 1..circuit.wires.len() {
                let mut nwires = Vec::new();
                for i in 0..circuit.wires.len() {
                    if i != x && i != y && i != z {
                        nwires.push(circuit.wires[i].clone());
                    }
                }

                if count_groups(Circuit {
                    wires: nwires.clone(),
                }) > 0
                {
                    return count_groups(Circuit {
                        wires: nwires.clone(),
                    });
                }
            }
        }
    }

    return 0;
}

fn count_groups(c: Circuit) -> i32 {
    println!("COUNTING GROUPS");
    let mut ccount = Vec::new();
    let mut tcount = HashMap::new();
    let mut mult = 1;
    for wire in &c.wires {
        ccount.push(wire.0.clone());
        ccount.push(wire.1.clone());
        tcount.insert(wire.0.clone(), 1);
        tcount.insert(wire.1.clone(), 1);
    }
    let ilen = tcount.len();

    let mut seen = 0;
    while ccount.len() > 0 {
        let mut search = Vec::new();
        search.push(ccount.pop().unwrap());
        seen += 1;

        let mut loops = 1;
        while search.len() > 0 {
            println!("SEARCHING {:?}", search);
            println!("CCOUNT {:?}", ccount);
            let curr = search.pop().unwrap();
            ccount.retain(|i| *i != curr);

            for (key, val) in &c.wires {
                if *key == curr {
                    let mut found = false;
                    for exval in &ccount {
                        if exval == val {
                            found = true;
                        }
                    }

                    if found {
                        search.push(val.to_string());
                        seen += 1;
                        ccount.retain(|i| i != val);
                    }
                }

                if *val == curr {
                    let mut found = false;
                    for exval in &ccount {
                        if exval == key {
                            found = true;
                        }
                    }

                    if found {
                        search.push(key.to_string());
                        seen += 1;
                        ccount.retain(|i| i != key);
                    }
                }
            }

            loops += 1;
            if loops > 3 {
                //åreturn 0;
            }
        }

        println!("HERE {}", seen);
        mult *= seen;
        seen = 0;
    }

    println!("HERE {} {}", mult, ilen);
    if mult != ilen {
        return mult as i32;
    }

    return 0;
}

struct Circuit {
    wires: Vec<(String, String)>,
}

fn build_circuit(data: String) -> Circuit {
    let mut twires = Vec::new();

    for line in data.split("\n") {
        let mut pieces = line.trim().split(":");
        let first = pieces.next().unwrap().to_string();
        let second = pieces.next().unwrap().to_string();

        let ff = first.trim();
        for ss in second.split_whitespace() {
            twires.push((ff.to_string(), ss.to_string()));
        }
    }

    return Circuit { wires: twires };
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "jqt: rhn xhk nvd
        rsh: frs pzl lsr
        xhk: hfx
        cmg: qnr nvd lhk bvb
        rhn: xhk bvb hfx
        bvb: xhk hfx
        pzl: lsr hfx nvd
        qnr: nvd
        ntq: jqt hfx bvb xhk
        nvd: lhk
        lsr: lhk
        rzs: qnr cmg lsr rsh
        frs: qnr lhk lsr"
            .to_string();

        let removes = solve_day25_part1(test_case);
        assert_eq!(removes, 54)
    }
}
