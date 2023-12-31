use std::collections::HashMap;
use std::time::Instant;

struct GNode {
    curr_node: String,
    seen: Vec<String>,
}

pub fn solve_day25_part1(data: String) -> i32 {
    let circuit = build_circuit(data);

    // println!("HERE {}", circuit.raw.len());
    let mut table = Vec::new();
    for i in 0..circuit.entries.len() {
        table.push(vec![0; circuit.entries.len()]);
    }

    let mut loop_count = 0;
    for start in &circuit.entries {
        for end in &circuit.entries {
            if start != end {
                loop_count += 1;
                //println!("TRYING {} -> {}", start, end);
                let mut search = Vec::new();
                let mut startv = Vec::new();
                startv.push(start.to_string());
                search.push(GNode {
                    curr_node: start.to_string(),
                    seen: startv,
                });
                let now = Instant::now();
                let mut nodes = 0;
                let mut seen = HashMap::new();
                while search.len() > 0 {
                    nodes += 1;
                    let curr = search.remove(0);
                    seen.insert(curr.curr_node.clone(), true);
                    if curr.curr_node == *end {
                        // Update the table and break out
                        //println!("FOUND {:?}", curr.seen);

                        // Update the table
                        for i in 1..curr.seen.len() {
                            for sn in 0..circuit.entries.len() {
                                if circuit.entries[sn] == curr.seen[i - 1] {
                                    for en in 0..circuit.entries.len() {
                                        if circuit.entries[en] == curr.seen[i] {
                                            table[sn][en] += 1;
                                            table[en][sn] += 1;
                                        }
                                    }
                                }
                            }
                        }

                        break;
                    }

                    for next in &circuit.wires[&curr.curr_node] {
                        if !seen.contains_key(next) {
                            let mut nseen = curr.seen.clone();
                            nseen.push(next.clone());
                            search.push(GNode {
                                curr_node: next.to_string(),
                                seen: nseen,
                            });
                            seen.insert(next.to_string(), true);
                        }
                    }
                }
                //println!("SEARCH {} IN {}", nodes, now.elapsed().as_millis());

                if loop_count % 1000 == 0 {
                    let now2 = Instant::now();
                    let mut scores = Vec::new();
                    for sn in 0..circuit.entries.len() {
                        for en in 0..circuit.entries.len() {
                            scores.push(
                                ((
                                    circuit.entries[sn].clone(),
                                    circuit.entries[en].clone(),
                                    table[sn][en],
                                )),
                            );
                        }
                    }

                    scores.sort_unstable_by(|a, b| b.2.cmp(&a.2));
                    let cnt = count_groups(
                        circuit.clone(),
                        scores[0].clone(),
                        scores[2].clone(),
                        scores[4].clone(),
                    );

                    if cnt > 0 {
                        return cnt;
                    }
                }
            }
        }
    }

    let mut scores = Vec::new();
    for sn in 0..circuit.entries.len() {
        for en in 0..circuit.entries.len() {
            scores.push(
                ((
                    circuit.entries[sn].clone(),
                    circuit.entries[en].clone(),
                    table[sn][en],
                )),
            );
        }
    }

    scores.sort_unstable_by(|a, b| b.2.cmp(&a.2));
    let cnt = count_groups(
        circuit.clone(),
        scores[0].clone(),
        scores[2].clone(),
        scores[4].clone(),
    );

    if cnt > 0 {
        return cnt;
    }

    return 0;
}

fn count_groups<T>(
    c: Circuit,
    s1: (String, String, T),
    s2: (String, String, T),
    s3: (String, String, T),
) -> i32 {
    let mut ccount = Vec::new();
    let mut tcount = HashMap::new();
    let mut mult = 1;
    for (wire, _) in &c.wires {
        ccount.push(wire.clone());
        tcount.insert(wire.clone(), 1);
    }
    let ilen = tcount.len();

    let mut seen = 0;
    while ccount.len() > 0 {
        let mut search = Vec::new();
        search.push(ccount.pop().unwrap());
        seen += 1;

        let mut loops = 1;
        while search.len() > 0 {
            let curr = search.pop().unwrap();
            ccount.retain(|i| *i != curr);

            for key in &c.wires[&curr] {
                let mut found = false;

                for exval in &ccount {
                    if *exval == *key {
                        found = true;
                    }
                }

                if (s1.0 == curr && s1.1 == *key) || (s1.0 == *key && s1.1 == curr) {
                    found = false;
                }

                if (s2.0 == curr && s2.1 == *key) || (s2.0 == *key && s2.1 == curr) {
                    found = false;
                }

                if (s3.0 == curr && s3.1 == *key) || (s3.0 == *key && s3.1 == curr) {
                    found = false;
                }

                if found {
                    search.push(key.to_string());
                    seen += 1;
                    ccount.retain(|i| *i != *key);
                }
            }

            loops += 1;
            if loops > 50 {
                return 0;
            }
        }

        mult *= seen;
        seen = 0;
    }

    if mult != ilen {
        return mult as i32;
    }

    return 0;
}

#[derive(Clone, Debug)]
struct Circuit {
    wires: HashMap<String, Vec<String>>,
    raw: Vec<(String, String)>,
    entries: Vec<String>,
}

fn build_circuit(data: String) -> Circuit {
    let mut twires: Vec<(String, String)> = Vec::new();
    let mut tcount = HashMap::new();

    for line in data.split("\n") {
        let mut pieces = line.trim().split(":");
        let first = pieces.next().unwrap().to_string();
        let second = pieces.next().unwrap().to_string();

        let ff = first.trim();
        for ss in second.split_whitespace() {
            let mut found = false;
            for (se, st) in &twires {
                if (*se == ff.to_string() && *st == ss.to_string())
                    || (*se == ss.to_string() && *st == ff.to_string())
                {
                    found = true;
                }
            }

            if !found {
                twires.push((ff.to_string(), ss.to_string()));
                tcount.insert(ff.to_string(), 1);
                tcount.insert(ss.to_string(), 1);
            }
        }
    }

    let mut fin: HashMap<String, Vec<String>> = HashMap::new();
    let mut entries = Vec::new();
    for (key, _) in tcount {
        let mut nvec = Vec::new();
        for (se, st) in &twires {
            if *se == key {
                nvec.push(st.clone());
            }
            if *st == key {
                nvec.push(se.clone());
            }
        }
        entries.push(key.clone());
        fin.insert(key, nvec);
    }

    return Circuit {
        wires: fin,
        raw: twires,
        entries: entries,
    };
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
