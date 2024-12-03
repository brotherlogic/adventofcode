# day1_spec.rb
require 'rspec'
require_relative 'day3'
require_relative 'lib/advent_pb'

RSpec.describe Day3 do
  describe "part 1" do
    it "returns 161" do
      day3 = Day3.new
      data = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
      expect(day3.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(161)
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