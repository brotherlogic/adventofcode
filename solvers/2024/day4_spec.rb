# day1_spec.rb
require 'rspec'
require_relative 'day4'
require_relative 'lib/advent_pb'

RSpec.describe Day4 do
  describe "part 1" do
    it "returns 18" do
      day4 = Day4.new
      data = "MMMSXXMASM
      MSAMXMSMSA
      AMXSXMAAMM
      MSAMASMSMX
      XMASAMXAMM
      XXAMMXXAMA
      SMSMSASXSS
      SAXAMASAAA
      MAMMMXMMMM
      MXMXAXMASX"
      expect(day4.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(18)
    end
  end
#  describe "part 2" do
#    it "returns 48" do
#      day3 = Day3.new
#      data = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
#      expect(day3.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(48)
#    end
#  end
end