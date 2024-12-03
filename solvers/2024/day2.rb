class Day2

    def direction(a, b)
        if a.to_i < b.to_i 
            return -1
        end
        return 1
    end

    def diff(a, b)
        if a.to_i < b.to_i
            return b.to_i - a.to_i
        end
        return a.to_i - b.to_i
    end

    def safe(line)
        nums = line.strip.split

        for i in 1..nums.length()-2
            if direction(nums[0],nums[1]) != direction(nums[i],nums[i+1])
                #puts "unsafe", line
                return false
            end
            if diff(nums[i],nums[i+1]) < 1 || diff(nums[i],nums[i+1]) > 3
                #puts "unsafes", line
                return false
            end
        end

        return true
    end


    def solvePart1(solve_req)
        lines = solve_req.data.split("\n")
        safe = 0
        lines.each do |line|
            if safe(line)
                safe += 1
            end
        end
        return safe
    end
end