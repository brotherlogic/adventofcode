# day1_spec.rb
require 'rspec'
require_relative 'day5'
require_relative 'lib/advent_pb'

RSpec.describe Day5 do
  describe "part 1" do
    it "returns 143" do
      day5 = Day5.new
      data = "47|53
      97|13
      97|61
      97|47
      75|29
      61|13
      75|53
      29|13
      97|29
      53|29
      61|53
      97|53
      61|29
      47|13
      75|47
      97|75
      47|61
      75|61
      47|29
      75|13
      53|13
      
      75,47,61,53,29
      97,61,53,29,13
      75,29,13
      75,97,47,61,53
      61,13,29
      97,13,75,29,47"
      expect(day5.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(143)
    end
  end
  describe "part 2" do
    it "returns 123" do
      day5 = Day5.new
      data = "47|53
      97|13
      97|61
      97|47
      75|29
      61|13
      75|53
      29|13
      97|29
      53|29
      61|53
      97|53
      61|29
      47|13
      75|47
      97|75
      47|61
      75|61
      47|29
      75|13
      53|13
      
      75,47,61,53,29
      97,61,53,29,13
      75,29,13
      75,97,47,61,53
      61,13,29
      97,13,75,29,47"
      expect(day5.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(123)
    end
  end
end