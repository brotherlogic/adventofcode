# day11_spec.rb
require 'rspec'
require_relative 'day11'
require_relative 'lib/advent_pb'

RSpec.describe Day11 do
  describe "part 1" do
    it "returns 55312" do
      day11 = Day11.new
      data = "125 17"
      expect(day11.fblink(day11.toarr(data), 2)).to eq(4)
      expect(day11.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(55312)
    end
  end
  describe "part 2" do
    it "returns 65601038650482" do
      day11 = Day11.new
      data = "125 17"
      expect(day11.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(65601038650482)
    end
  end
end