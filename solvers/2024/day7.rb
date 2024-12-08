require 'logger'
require 'date'

class Day7

    def search(goal, current, remainder)
        if remainder.length() == 0
            return goal == current
        end

        ne = remainder[0].to_i
        return search(goal, current+ne, remainder[1..remainder.length()]) || search(goal, current*ne, remainder[1..remainder.length()])
    end

    def join(a, b)
        return "" + a.to_s + b.to_s
    end

    def search2(goal, current, remainder)
        if remainder.length() == 0
            return goal == current
        end

        ne = remainder[0].to_i
        return search2(goal, current+ne, remainder[1..remainder.length()]) || search2(goal, current*ne, remainder[1..remainder.length()]) || search2(goal, join(current, ne).to_i, remainder[1..remainder.length()])
    end

    def resolve(line)
        elems = line.strip().split()

        goal = elems[0][0..-1].to_i

        startv = elems[1].to_i
        
        if search(goal, startv, elems[2..elems.length()])
            return goal
        else
            return 0
        end
    end


    def resolve2(line)
        elems = line.strip().split()

        goal = elems[0][0..-1].to_i

        startv = elems[1].to_i
        
        if search2(goal, startv, elems[2..elems.length()])
            return goal
        else
            return 0
        end
    end

    def solvePart1(solve_req)
       lines = solve_req.data.split("\n")
       sum = 0
       lines.each do |line|
        sum += resolve(line)
       end

       return sum
    end

    def solvePart2(solve_req)
        lines = solve_req.data.split("\n")
        sum = 0
        lines.each do |line|
         sum += resolve2(line)
        end
 
        return sum
     end
end