class Day3
    def solvePart1(solve_req)
        total = 0
        elems = solve_req.data.scan(/mul\(\d*,\d*\)/i)

        elems.each do |elem|
            num1, num2 = elem.match(/mul\((\d*),(\d*)/i).captures
            total += num1.to_i * num2.to_i
        end

        return total
    end

    def solvePart2(solve_req)
        total = 0
        elems = solve_req.data.scan(/don\'t\(\)|do\(\)|mul\(\d*,\d*\)/i)

       
        zoned = true
        elems.each do |elem|
            if elem == "do()"
                zoned = true
            elsif elem == "don't()"
                zoned = false
            else
                num1, num2 = elem.match(/mul\((\d*),(\d*)/i).captures
                if zoned
                    total += num1.to_i * num2.to_i
                end
            end
        end

        return total
    end
end