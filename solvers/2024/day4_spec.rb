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
  describe "part 2" do
    it "returns 9" do
      day4 = Day4.new
      data = ".M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
.........."
      expect(day4.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(9)
    end
  end
end