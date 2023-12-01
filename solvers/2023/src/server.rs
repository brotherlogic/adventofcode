use tonic::{transport::Server, Request, Response, Status};
use solver::{SolveResponse, SolveRequest, solver_service_server::{SolverService, SolverServiceServer}};
pub mod solver {
    tonic::include_proto!("adventofcode");
} 

// defining a struct for our service
#[derive(Default)]
pub struct RServer {}

// implementing rpc for service defined in .proto
#[tonic::async_trait]
impl SolverService for RServer {
    async fn solve(&self,_request:Request<SolveRequest>)->Result<Response<SolveResponse>,Status>{
        Ok(Response::new(SolveResponse{
                string_answer:"hello".to_string(),
                big_answer:0,
                answer:10,
        }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:8080".parse().unwrap();
    let solver = RServer::default();
    println!("Server listening on {}", addr);
// adding our service to our server.
    Server::builder()
        .add_service(SolverServiceServer::new(solver))
        .serve(addr)
        .await?;
    Ok(())
}