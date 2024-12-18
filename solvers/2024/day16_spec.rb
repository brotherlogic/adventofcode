# day16_spec.rb
require 'rspec'
require_relative 'day16'
require_relative 'lib/advent_pb'

RSpec.describe Day16 do
  describe "part 1" do
    it "returns 7036" do
      day16 = Day16.new
      data = "###############
      #.......#....E#
      #.#.###.#.###.#
      #.....#.#...#.#
      #.###.#####.#.#
      #.#.#.......#.#
      #.#.#####.###.#
      #...........#.#
      ###.#.#####.#.#
      #...#.....#.#.#
      #.#.#.###.#.#.#
      #.....#...#.#.#
      #.###.#.#.#.#.#
      #S..#.....#...#
      ###############"
      #expect(day16.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(7036)
    end
  end
  describe "part 2" do
    it "returns 45" do
      day16 = Day16.new
      data = "###############
      #.......#....E#
      #.#.###.#.###.#
      #.....#.#...#.#
      #.###.#####.#.#
      #.#.#.......#.#
      #.#.#####.###.#
      #...........#.#
      ###.#.#####.#.#
      #...#.....#.#.#
      #.#.#.###.#.#.#
      #.....#...#.#.#
      #.###.#.#.#.#.#
      #S..#.....#...#
      ###############"
      expect(day16.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(45)
    end
  end
  describe "part 2" do
    it "returns 45" do
      day16 = Day16.new
      data = "#####
      #..E#
      #.#.#
      #...#
      #S#.#
      #####"
      #expect(day16.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(8)
    end
  end
  describe "part 2" do
    it "returns 45" do
      day16 = Day16.new
      data = "#####
      ###E#
      #...#
      #.#.#
      #...#
      #S#.#
      #####"
      #expect(day16.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(10)
    end
  end
end