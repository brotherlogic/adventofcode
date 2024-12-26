# day16_spec.rb
require 'rspec'
require_relative 'day24'
require_relative 'lib/advent_pb'

RSpec.describe Day24 do
  describe "part 1" do
    it "returns 4" do
      day24 = Day24.new
      data = "x00: 1
      x01: 1
      x02: 1
      y00: 0
      y01: 1
      y02: 0
      
      x00 AND y00 -> z00
      x01 XOR y01 -> z01
      x02 OR y02 -> z02"
      expect(day24.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(4)
    end
  end
end