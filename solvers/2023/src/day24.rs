pub fn solve_day24_part1(data: String, low: f64, high: f64) -> i64 {
    let mut hailstones = build_hailstones(data);

    println!("Found {} hailstones", hailstones.len());

    let mut ccount = 0;
    while hailstones.len() > 0 {
        let chail = hailstones.remove(0);

        for hailstone in &hailstones {
            let crossover = cross(chail.clone(), hailstone.clone());
            if crossover >= low && crossover <= high {
                println!("{:?} crosses {:?} at {}", chail, hailstone, crossover);
                ccount += 1;
            }
        }
    }

    return ccount;
}

fn cross(h1: Hailstone, h2: Hailstone) -> f64 {
    let top =
        h2.dx * h2.y * h1.dx - h2.dy * h2.x * h1.dx + h1.dy * h1.x * h2.dx - h1.dx * h1.y * h2.dx;
    let bottom = h2.dx * h1.dy - h1.dx * h2.dy;

    if bottom == 0 {
        println!("NO CROSS {:?} {:?}", h1, h2);
        return 0.0;
    }

    let cross_point = top as f64 / bottom as f64;
    let tval = (cross_point - h1.x as f64) / h1.dx as f64;
    let tval2 = (cross_point - h2.x as f64) / h2.dx as f64;
    println!("CROSS {} -> {} {}", cross_point, tval, tval2);
    if tval > 0.0 && tval2 > 0.0 {
        return top as f64 / bottom as f64;
    }

    return 0.0;
}

#[derive(Clone, Debug)]
struct Hailstone {
    x: i64,
    y: i64,
    z: i64,
    dx: i64,
    dy: i64,
    dz: i64,
}

fn build_hailstones(data: String) -> Vec<Hailstone> {
    let mut hailstones = Vec::new();

    for line in data.split("\n") {
        let mut parts = line.trim().split("@");
        let first = parts.next().unwrap().to_string();
        let second = parts.next().unwrap().to_string();

        let mut coords = first.split(",");
        let x = coords
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();
        let y = coords
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();
        let z = coords
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();

        let mut velocity = second.split(",");
        let dx = velocity
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();
        let dy = velocity
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();
        let dz = velocity
            .next()
            .unwrap()
            .to_string()
            .trim()
            .parse::<i64>()
            .unwrap();

        hailstones.push(Hailstone {
            x: x,
            y: y,
            z: z,
            dx: dx,
            dy: dy,
            dz: dz,
        })
    }

    return hailstones;
}

#[cfg(test)]
mod testsca {
    use super::*;

    #[test]
    fn part1_test_first() {
        let test_case = "19, 13, 30 @ -2,  1, -2
        18, 19, 22 @ -1, -1, -2
        20, 25, 34 @ -2, -2, -4
        12, 31, 28 @ -1, -2, -1
        20, 19, 15 @  1, -5, -3"
            .to_string();

        let pulses = solve_day24_part1(test_case, 7.0, 27.0);
        assert_eq!(pulses, 2)
    }
}
