# day0_spec.rb
require 'rspec'
require_relative 'day0'

RSpec.describe Day0 do
  describe "part 0" do
    it "returns 10" do
      day0 = Day0.new
      expect(day0.solvePart1("")).to eq(10)
    end
  end
  describe "part 1" do
    it "returns 15" do
      day0 = Day0.new
      expect(day0.solvePart2("")).to eq(15)
    end
  end

end