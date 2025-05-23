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

        for i in 1..nums.length()-1
            if direction(nums[0],nums[1]) != direction(nums[i-1],nums[i])
                return false
            end
            if diff(nums[i-1],nums[i]) < 1 || diff(nums[i-1],nums[i]) > 3
                return false
            end
        end

        return true
    end


    def safe2(line)
        if safe(line)
            return true
        end

        nums = line.strip.split

        for drop in 0..nums.length()-1
            narr = []
            for i in 0..nums.length()-1
                if i != drop
                    narr.push(nums[i])
                end
            end
            if safe(narr.join(" "))
                return true
            end
        end

        return false
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

    def solvePart2(solve_req)
        lines = solve_req.data.split("\n")
        safe = 0
        lines.each do |line|
            if safe2(line)
                safe += 1
            end
        end
        return safe
    end
end