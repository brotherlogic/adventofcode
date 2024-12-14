# day13_spec.rb
require 'rspec'
require_relative 'day13'
require_relative 'lib/advent_pb'

RSpec.describe Day13 do
  describe "part 1" do
    it "returns 480" do
      day13= Day13.new
      data = "Button A: X+94, Y+34
      Button B: X+22, Y+67
      Prize: X=8400, Y=5400
      
      Button A: X+26, Y+66
      Button B: X+67, Y+21
      Prize: X=12748, Y=12176
      
      Button A: X+17, Y+86
      Button B: X+84, Y+37
      Prize: X=7870, Y=6450
      
      Button A: X+69, Y+23
      Button B: X+27, Y+71
      Prize: X=18641, Y=10279"
      expect(day13.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(480)
    end
  end
  describe "part 1" do
    it "returns 480" do
      day13= Day13.new
      data = "Button A: X+94, Y+34
      Button B: X+22, Y+67
      Prize: X=8400, Y=5400
      
      Button A: X+26, Y+66
      Button B: X+67, Y+21
      Prize: X=12748, Y=12176
      
      Button A: X+17, Y+86
      Button B: X+84, Y+37
      Prize: X=7870, Y=6450
      
      Button A: X+69, Y+23
      Button B: X+27, Y+71
      Prize: X=18641, Y=10279"
      expect(day13.solvePart2(Adventofcode::SolveRequest.new(data: data))).to eq(480)
    end
  end
end