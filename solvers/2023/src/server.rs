use tonic::{transport::Server, Request, Response, Status};
use solver::{SolveResponse, SolveRequest, solver_service_server::{SolverService, SolverServiceServer}};
use solver::advent_of_code_internal_service_client::AdventOfCodeInternalServiceClient;
use solver::RegisterRequest;
use std::thread;
use tokio::time::{ Duration};
pub mod solver {
    tonic::include_proto!("adventofcode");
}

mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
mod day10;
mod day11;
mod day12;
mod day13;
mod day14;
mod day15;

// defining a struct for our service
#[derive(Default)]
pub struct RServer {}

// implementing rpc for service defined in .proto
#[tonic::async_trait]
impl SolverService for RServer {
    async fn solve(&self,request:Request<SolveRequest>)->Result<Response<SolveResponse>,Status>{
        let rq = request.into_inner();
        if rq.year == 2023 && rq.day == 1 && rq.part == 1 {
            let tanswer = day1::solve_day1_part1(rq.data);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 1 && rq.part == 2 {
            let tanswer = day1::solve_day1_part2(rq.data);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 2 && rq.part == 1 {
            let tanswer = day2::solve_day2_part1(rq.data);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 2 && rq.part == 2 {
            let tanswer = day2::solve_day2_part2(rq.data);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 3 && rq.part == 1 {
            let tanswer = day3::solve_day3_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 3 && rq.part == 2 {
            let tanswer = day3::solve_day3_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 4 && rq.part == 1 {
            let tanswer = day4::solve_day4_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 4 && rq.part == 2 {
            let tanswer = day4::solve_day4_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer as i32,
         }));
        }
        if rq.year == 2023 && rq.day == 5 && rq.part == 1 {
            let tanswer = day5::solve_day5_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 5 && rq.part == 2 {
            let tanswer = day5::path_part_2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
         if rq.year == 2023 && rq.day == 6 && rq.part == 1 {
            let tanswer = day6::solve_day6_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
         if rq.year == 2023 && rq.day == 6 && rq.part == 2 {
            let tanswer = day6::solve_day6_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 7 && rq.part == 1 {
            let tanswer = day7::solve_day7_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
        if rq.year == 2023 && rq.day == 7 && rq.part == 2 {
            let tanswer = day7::solve_day7_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
        if rq.year == 2023 && rq.day == 8 && rq.part == 1 {
            let tanswer = day8::solve_day8_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 8 && rq.part == 2 {
            let tanswer = day8::solve_day8_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
        if rq.year == 2023 && rq.day == 9 && rq.part == 1 {
            let tanswer = day9::solve_day9_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 9 && rq.part == 2 {
            let tanswer = day9::solve_day9_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 10 && rq.part == 1 {
            let tanswer = day10::solve_day10_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 10 && rq.part == 2 {
            let tanswer = day10::solve_day10_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 11 && rq.part == 1 {
            let tanswer = day11::solve_day11_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
        if rq.year == 2023 && rq.day == 11 && rq.part == 2 {
            let tanswer = day11::solve_day11_part2(rq.data, 1000000);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:tanswer,
                answer:0,
         }));
        }
        if rq.year == 2023 && rq.day == 12 && rq.part == 1 {
            let tanswer = day12::solve_day12_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 12 && rq.part == 2 {
            let tanswer = day12::solve_day12_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 13 && rq.part == 1 {
            let tanswer = day13::solve_day13_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 13 && rq.part == 2 {
            let tanswer = day13::solve_day13_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 14 && rq.part == 1 {
            let tanswer = day14::solve_day14_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 14 && rq.part == 2 {
            let tanswer = day14::solve_day14_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 15 && rq.part == 1 {
            let tanswer = day15::solve_day15_part1(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
        if rq.year == 2023 && rq.day == 15 && rq.part == 2 {
            let tanswer = day15::solve_day15_part2(rq.data);
            println!("Returning {}", tanswer);
            return Ok(Response::new(SolveResponse{
                string_answer:"".to_string(),
                big_answer:0,
                answer:tanswer,
         }));
        }
       Err(Status::unimplemented("Solution is not implemented yet"))
    }
}

#[tokio::main]
async fn register() -> Result<(), Box<dyn std::error::Error>> {

    let mut client = AdventOfCodeInternalServiceClient::connect("http://adventofcode.adventofcode:8082").await?;

    let request = tonic::Request::new(RegisterRequest {
        year: 2023,
        callback: "adventofcode-solver-2023.adventofcode:8080".to_string(),
    });

    let _ = client.register(request).await?;


    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:8080".parse().unwrap();
    let solver = RServer::default();
    println!("Server listening on {}", addr);

    println!("Spawning the registration thread");
    thread::spawn(|| {
        println!("Starting");
        loop {
            let _ = register();
            thread::sleep(Duration::from_millis(60000));
        }
    });

    println!("Running server");
    Server::builder()
        .add_service(SolverServiceServer::new(solver))
        .serve(addr)
        .await?;
    Ok(())
}