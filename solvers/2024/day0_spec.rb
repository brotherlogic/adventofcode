# day0_spec.rb
require 'rspec'
require 'advent_services_pb'
require_relative 'day0'
require_relative 'server'

RSpec.describe Day0 do
  describe "part 0" do
    it "returns 10" do
      day0 = SolverServer.new
      expect(day0.solve(Adventofcode::SolveRequest.new(year:2024, day: 0))).to eq(10)
    end
  end
  describe "part 1" do
    it "returns 15" do
      day0 = Day0.new
      expect(day0.solvePart2("")).to eq(15)
    end
  end
end