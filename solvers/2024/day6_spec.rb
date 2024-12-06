# day1_spec.rb
require 'rspec'
require_relative 'day6'
require_relative 'lib/advent_pb'

RSpec.describe Day6 do
  describe "part 1" do
    it "returns 41" do
      day6 = Day6.new
      data = "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."
      expect(day6.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(41)
    end
  end
end