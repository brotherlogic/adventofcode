#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'advent_services_pb'
require_relative 'day1'
require_relative 'day2'
require_relative 'day3'
require_relative 'day4'
require_relative 'day5'
require_relative 'day6'
require_relative 'day7'
require_relative 'day8'
require_relative 'day9'
require_relative 'day10'

include GRPC::Core::StatusCodes

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
    if solve_req.day == 1 && solve_req.part == 1
        d1 = Day1.new
        return Adventofcode::SolveResponse.new(answer: d1.solvePart1(solve_req)) 
    end
    if solve_req.day == 1 && solve_req.part == 2
      d1 = Day1.new
      return Adventofcode::SolveResponse.new(answer: d1.solvePart2(solve_req)) 
    end

    if solve_req.day == 2 && solve_req.part == 1
      d1 = Day2.new
      return Adventofcode::SolveResponse.new(answer: d1.solvePart1(solve_req)) 
    end
    if solve_req.day == 2 && solve_req.part == 2
      d1 = Day2.new
      return Adventofcode::SolveResponse.new(answer: d1.solvePart2(solve_req)) 
    end
    
    if solve_req.day == 3 && solve_req.part == 1
      d3 = Day3.new
      return Adventofcode::SolveResponse.new(answer: d3.solvePart1(solve_req)) 
    end
    if solve_req.day == 3 && solve_req.part == 2
      d3 = Day3.new
      return Adventofcode::SolveResponse.new(answer: d3.solvePart2(solve_req)) 
    end

    if solve_req.day == 4 && solve_req.part == 1
      d4 = Day4.new
      return Adventofcode::SolveResponse.new(answer: d4.solvePart1(solve_req)) 
    end
    if solve_req.day == 4 && solve_req.part == 2
      d4 = Day4.new
      return Adventofcode::SolveResponse.new(answer: d4.solvePart2(solve_req)) 
    end
    
    if solve_req.day == 5 && solve_req.part == 1
      d5 = Day5.new
      return Adventofcode::SolveResponse.new(answer: d5.solvePart1(solve_req)) 
    end
    if solve_req.day == 5 && solve_req.part == 2
      d5 = Day5.new
      return Adventofcode::SolveResponse.new(answer: d5.solvePart2(solve_req)) 
    end

    if solve_req.day == 6 && solve_req.part == 1
      d6 = Day6.new
      return Adventofcode::SolveResponse.new(answer: d6.solvePart1(solve_req)) 
    end
    if solve_req.day == 6 && solve_req.part == 2
      d6 = Day6.new
      return Adventofcode::SolveResponse.new(answer: d6.solvePart2(solve_req)) 
    end

    if solve_req.day == 7 && solve_req.part == 1
      d7 = Day7.new
      return Adventofcode::SolveResponse.new(big_answer: d7.solvePart1(solve_req)) 
    end
    if solve_req.day == 7 && solve_req.part == 2
      d7 = Day7.new
      return Adventofcode::SolveResponse.new(big_answer: d7.solvePart2(solve_req)) 
    end
  
    if solve_req.day == 8 && solve_req.part == 1
      d8 = Day8.new
      return Adventofcode::SolveResponse.new(big_answer: d8.solvePart1(solve_req)) 
    end
    if solve_req.day == 8 && solve_req.part == 2
      d8 = Day8.new
      return Adventofcode::SolveResponse.new(answer: d8.solvePart2(solve_req)) 
    end

    if solve_req.day == 9 && solve_req.part == 1
      d9 = Day9.new
      return Adventofcode::SolveResponse.new(big_answer: d9.solvePart1(solve_req)) 
    end
    if solve_req.day == 9 && solve_req.part == 2
      d9 = Day9.new
      return Adventofcode::SolveResponse.new(big_answer: d9.solvePart2(solve_req)) 
    end

    if solve_req.day == 10 && solve_req.part == 1
      d10 = Day10.new
      return Adventofcode::SolveResponse.new(big_answer: d10.solvePart1(solve_req)) 
    end
  
    raise GRPC::BadStatus.new_status_exception(UNIMPLEMENTED, details = 'Solution is not ready')
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

STDOUT.sync = true
main
