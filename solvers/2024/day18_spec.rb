# day16_spec.rb
require 'rspec'
require_relative 'day18'
require_relative 'lib/advent_pb'

RSpec.describe Day18 do
  describe "part 1" do
    it "returns 7036" do
      day18 = Day18.new
      data = "5,4
      4,2
      4,5
      3,0
      2,1
      6,3
      2,4
      1,5
      0,6
      3,3
      2,6
      5,1
      1,2
      5,5
      2,5
      6,5
      1,4
      0,4
      6,4
      1,1
      6,1
      1,0
      0,5
      1,6
      2,0"
      expect(day18.solve(data, 6, 12)).to eq(22)
    end
  end
end