# day16_spec.rb
require 'rspec'
require_relative 'day17'
require_relative 'lib/advent_pb'

RSpec.describe Day17 do
  describe "part 1" do
    it "returns 4,6,3,5,6,3,5,2,1,0" do
      day17 = Day17.new
      data = "Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0"
      expect(day17.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq("4,6,3,5,6,3,5,2,1,0")
    end
  end
end