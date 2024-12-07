# day7_spec.rb
require 'rspec'
require_relative 'day7'
require_relative 'lib/advent_pb'

RSpec.describe Day7 do
  describe "part 1" do
    it "returns 3749" do
      day7 = Day7.new
      data = "190: 10 19
      3267: 81 40 27
      83: 17 5
      156: 15 6
      7290: 6 8 6 15
      161011: 16 10 13
      192: 17 8 14
      21037: 9 7 18 13
      292: 11 6 16 20"
      expect(day7.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(3749)
    end
  end
end