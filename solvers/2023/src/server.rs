use tonic::{transport::Server, Request, Response, Status};
use solver::solver__service_server::{Solve, SolveServicerServer};
use solver::{SolveResponse, SolveRequest};
mod solver;  

// defining a struct for our service
#[derive(Default)]
pub struct Server {}

// implementing rpc for service defined in .proto
#[tonic::async_trait]
impl Solve for Server {
    async fn send(&self,request:Request<SolveRequest>)->Result<Response<SolveResponse>,Status>{
        Ok(Response::new(SolveResponse{
                answer:"hello",
        }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
// defining address for our service
    let addr = "[::1]:8080".parse().unwrap();
// creating a service
    let solver = Solver::default();
    println!("Server listening on {}", addr);
// adding our service to our server.
    Server::builder()
        .add_service(SolveServer::new(solver))
        .serve(addr)
        .await?;
    Ok(())
}