#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'advent_services_pb'

class SolverServer < Adventofcode::SolverService::Service
    def solve(solve_req, _unused_call)
      Adventofcode::SolveResponse.new(solution: 1)
    end
  end

  def main
    s = GRPC::RpcServer.new
    s.add_http2_port('0.0.0.0:8080', :this_port_is_insecure)
    s.handle(SolverServer)

    hostname = "adventofcode.adventofcode:8082"
    stub = Adventofcode::AdventOfCodeInternalService::Stub.new(hostname, :this_channel_is_insecure)
    begin
      message = stub.Register(Adventofcode::RegisterRequest.new(year: 2024)).message
      p "Greeting: #{message}"
    rescue GRPC::BadStatus => e
      abort "ERROR: #{e.message}"
    end
  

    # Runs the server with SIGHUP, SIGINT and SIGTERM signal handlers to
    #   gracefully shutdown.
    # User could also choose to run server via call to run_till_terminated
    s.run_till_terminated_or_interrupted([1, 'int', 'SIGTERM'])
  end
  
  main