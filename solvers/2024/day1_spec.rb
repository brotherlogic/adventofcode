# day1_spec.rb
require 'rspec'
require_relative 'day1'

RSpec.describe Day1 do
  describe "part 1" do
    it "returns 10" do
      day1 = Day1.new
      expect(day1.solve(Adventofcode::SolveRequest.new(year:2024, day: 0))).to eq(10)
    end
  end
end