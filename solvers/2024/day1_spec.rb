# day1_spec.rb
require 'rspec'
require_relative 'day1'
require_relative 'lib/advent_pb'

RSpec.describe Day1 do
  describe "part 1" do
    it "returns 11" do
      day1 = Day1.new
      data = "3   4
      4   3
      2   5
      1   3
      3   9
      3   3"
      expect(day1.solvePart1(Adventofcode::SolveRequest.new(year:2024, day: 1, data: data))).to eq(11)
    end
  end
end