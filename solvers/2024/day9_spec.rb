# day9_spec.rb
require 'rspec'
require_relative 'day9'
require_relative 'lib/advent_pb'

RSpec.describe Day9 do
  describe "part 1" do
    it "returns 1928" do
      day9 = Day9.new
      data = "2333133121414131402"
      expect(day9.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(1928)
    end
  end
  describe "part 2" do
    it "returns 2858" do
      day9 = Day9.new
      data = "2333133121414131402"
      expect(day9.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(2858)
    end
  end
end