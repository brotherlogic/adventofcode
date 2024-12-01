#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'advent_services_pb'

class Registrar
  def Register()
    puts "Registering 2024 Solver"
    hostname = "adventofcode.adventofcode:8082"
    stub = Adventofcode::AdventOfCodeInternalService::Stub.new(hostname, :this_channel_is_insecure)
    begin
      result = stub.register(Adventofcode::RegisterRequest.new(callback: "adventofcode-solver-2024.adventofcode:8080", year: 2024))
      puts result
    rescue GRPC::BadStatus => e
      abort "Register Error: #{e.message}"
    end
  end

  def runRegister()
    puts "Running Register"
    registrar = Registrar.new
    while true 
       sleep 60
       registrar.Register()
    end
  end
end

class SolverServer < Adventofcode::SolverService::Service
  def solve(solve_req, _unused_call)
    # Find the class solver for the given day
    if solve_req.Day == 1 && solve_req.Part == 1
        d1 = d1.new
        Adventofcode::SolveResponse.new(answer: d1.solvePart1(solve_req)) 
    end
  end
end

def main
  puts "Starting Server"
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:8080', :this_port_is_insecure)
  s.handle(SolverServer)

  registar = Registrar.new

  puts "Starting register thread"
  Thread.start {registar.runRegister}

  # Runs the server with SIGHUP, SIGINT and SIGTERM signal handlers to
  #   gracefully shutdown.
  # User could also choose to run server via call to run_till_terminated
  puts "Running in background"
  s.run_till_terminated_or_interrupted([1, 'int', 'SIGTERM'])
end

main
