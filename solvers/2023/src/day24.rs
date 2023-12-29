use num_bigfloat::BigFloat;
use num_bigfloat::ZERO;

pub fn solve_day24_part1(data: String, low: f64, high: f64) -> i64 {
    let mut hailstones = build_hailstones(data);

    println!("Found {} hailstones", hailstones.len());

    let mut ccount = 0;
    while hailstones.len() > 0 {
        let chail = hailstones.remove(0);

        for hailstone in &hailstones {
            let crossover = cross(chail.clone(), hailstone.clone());
            if crossover >= BigFloat::from_f64(low) && crossover <= BigFloat::from_f64(high) {
                println!("{:?} crosses {:?} at {}", chail, hailstone, crossover);
                ccount += 1;
            }
        }
    }

    return ccount;
}

fn cross(h1: Hailstone, h2: Hailstone) -> BigFloat {
    let top =
        h2.dx * h2.y * h1.dx - h2.dy * h2.x * h1.dx + h1.dy * h1.x * h2.dx - h1.dx * h1.y * h2.dx;
    let bottom = h2.dx * h1.dy - h1.dx * h2.dy;

    if bottom == ZERO {
        println!("NO CROSS {:?} {:?}", h1, h2);
        return ZERO;
    }

    let cross_point = top / bottom;
    let tval = (cross_point - h1.x) / h1.dx;
    let tval2 = (cross_point - h2.x) / h2.dx;
    println!("CROSS {} -> {} {}", cross_point, tval, tval2);
    if tval > ZERO && tval2 > ZERO {
        return top / bottom;
    }

    return ZERO;
}

#[derive(Clone, Debug)]
struct Hailstone {
    x: BigFloat,
    y: BigFloat,
    z: BigFloat,
    dx: BigFloat,
    dy: BigFloat,
    dz: BigFloat,
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
            x: BigFloat::from_i64(x),
            y: BigFloat::from_i64(y),
            z: BigFloat::from_i64(z),
            dx: BigFloat::from_i64(dx),
            dy: BigFloat::from_i64(dy),
            dz: BigFloat::from_i64(dz),
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
