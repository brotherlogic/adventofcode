class Day3
    def solvePart1(solve_req)
        total = 0
        elems = solve_req.data.scan(/mul\(\d*,\d*\)/i)

        puts solve_req.data
        elems.each do |elem|
            num1, num2 = elem.match(/mul\((\d*),(\d*)/i).captures
            total += num1.to_i * num2.to_i
        end

        return total
    end
end