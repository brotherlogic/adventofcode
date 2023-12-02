use tonic::{transport::Server, Request, Response, Status};
use solver::{SolveResponse, SolveRequest, solver_service_server::{SolverService, SolverServiceServer}};
use solver::advent_of_code_internal_service_client::AdventOfCodeInternalServiceClient;
use solver::RegisterRequest;
pub mod solver {
    tonic::include_proto!("adventofcode");
}

mod day1;
mod day2;

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
        Ok(Response::new(SolveResponse{
                string_answer:"hello".to_string(),
                big_answer:0,
                answer:10,
        }))
    }
}

async fn register() -> Result<(), Box<dyn std::error::Error>> {
    println!("Registering 2023 Solver");

    let mut client = AdventOfCodeInternalServiceClient::connect("http://adventofcode.adventofcode:8082").await?;

    let request = tonic::Request::new(RegisterRequest {
        year: 2023,
        callback: "adventofcode-solver-2023.adventofcode:8080".to_string(),
    });

    let response = client.register(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:8080".parse().unwrap();
    let solver = RServer::default();
    println!("Server listening on {}", addr);

    register().await?;

    Server::builder()
        .add_service(SolverServiceServer::new(solver))
        .serve(addr)
        .await?;
    Ok(())
}