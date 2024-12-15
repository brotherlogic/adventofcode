# day12_spec.rb
require 'rspec'
require_relative 'day12'
require_relative 'lib/advent_pb'

RSpec.describe Day12 do
  describe "part 1" do
    it "returns 1930" do
      day12 = Day12.new
      data = "RRRRIICCFF
      RRRRIICCCF
      VVRRRCCFFF
      VVRCCCJFFF
      VVVVCJJCFE
      VVIVCCJJEE
      VVIIICJJEE
      MIIIIIJJEE
      MIIISIJEEE
      MMMISSJEEE"
      expect(day12.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(1930)
    end
  end
  describe "part 1" do
    it "returns 1206" do
      day12 = Day12.new
      data = "RRRRIICCFF
      RRRRIICCCF
      VVRRRCCFFF
      VVRCCCJFFF
      VVVVCJJCFE
      VVIVCCJJEE
      VVIIICJJEE
      MIIIIIJJEE
      MIIISIJEEE
      MMMISSJEEE"
      expect(day12.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(1206)
    end
  end
  describe "part 2" do
    it "returns 368" do
      day12 = Day12.new
      data = "AAAAAA
      AAABBA
      AAABBA
      ABBAAA
      ABBAAA
      AAAAAA"
      expect(day12.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(368)
    end
  end
end

