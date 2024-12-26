# day16_spec.rb
require 'rspec'
require_relative 'day19'
require_relative 'lib/advent_pb'

RSpec.describe Day19 do
  describe "part 1" do
    it "returns 6" do
      day19 = Day19.new
      data = "r, wr, b, g, bwu, rb, gb, br

      brwrr
      bggr
      gbbr
      rrbgbr
      ubwu
      bwurrg
      brgr
      bbrgwb"
      #expect(day19.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(6)
    end
  end
  describe "part 2" do
    it "returns 16" do
      day19 = Day19.new
      data = "r, wr, b, g, bwu, rb, gb, br

      brwrr
      bggr
      gbbr
      rrbgbr
      ubwu
      bwurrg
      brgr
      bbrgwb"
      #expect(day19.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(16)
    end
  end
end