# day10_spec.rb
require 'rspec'
require_relative 'day10'
require_relative 'lib/advent_pb'

RSpec.describe Day10 do
  describe "part 1" do
    it "returns 36" do
      day10 = Day10.new
      data = "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"
      expect(day10.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(36)
    end
  end
  describe "part 2" do
    it "returns 81" do
      day10 = Day10.new
      data = "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"
      expect(day10.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(81)
    end
  end
end