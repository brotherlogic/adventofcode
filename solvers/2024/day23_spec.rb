# day16_spec.rb
require 'rspec'
require_relative 'day23'
require_relative 'lib/advent_pb'

RSpec.describe Day23 do
  describe "part 1" do
    it "returns 7" do
      day23 = Day23.new
      data = "kh-tc
      qp-kh
      de-cg
      ka-co
      yn-aq
      qp-ub
      cg-tb
      vc-aq
      tb-ka
      wh-tc
      yn-cg
      kh-ub
      ta-co
      de-co
      tc-td
      tb-wq
      wh-td
      ta-ka
      td-qp
      aq-cg
      wq-ub
      ub-vc
      de-ta
      wq-aq
      wq-vc
      wh-yn
      ka-de
      kh-ta
      co-tc
      wh-qp
      tb-vc
      td-yn"
      expect(day23.solvePart1(Adventofcode::SolveRequest.new(data: data))).to eq(7)
    end
  end
end