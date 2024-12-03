# day1_spec.rb
require 'rspec'
require_relative 'day2'
require_relative 'lib/advent_pb'

RSpec.describe Day2 do
  describe "part 1" do
    it "returns 2" do
      day2 = Day2.new
      data = "7 6 4 2 1
      1 2 7 8 9
      9 7 6 2 1
      1 3 2 4 5
      8 6 4 4 1
      1 3 6 7 9"
      expect(day2.solvePart1(Adventofcode::SolveRequest.new(year:2024, day: 2, data: data))).to eq(2)
    end
  end
 # describe "part 2" do
 #   it "returns 31" do
 #     day1 = Day1.new
 #     data = "3   4
 #     4   3
 #     2   5
 #     1   3
 #     3   9
 #     3   3"
 #     expect(day1.solvePart2(Adventofcode::SolveRequest.new(year:2024, day: 1, data: data))).to eq(31)
 #   end
 # end
end