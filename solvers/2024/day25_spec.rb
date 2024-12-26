# day16_spec.rb
require 'rspec'
require_relative 'day25'
require_relative 'lib/advent_pb'

RSpec.describe Day25 do
  describe "part 1" do
    it "returns 3" do
      day25 = Day25.new
      data = "#####
      .####
      .####
      .####
      .#.#.
      .#...
      .....
      
      #####
      ##.##
      .#.##
      ...##
      ...#.
      ...#.
      .....
      
      .....
      #....
      #....
      #...#
      #.#.#
      #.###
      #####
      
      .....
      .....
      #.#..
      ###..
      ###.#
      ###.#
      #####
      
      .....
      .....
      .....
      #....
      #.#..
      #.#.#
      #####"
      expect(day25.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(3)
    end
  end
end