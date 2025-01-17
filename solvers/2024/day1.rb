class Day1
 
    def differ(v1, v2)
        if v1 > v2
            return v1-v2
        end
        return v2-v1
    end

    def solvePart1(solve_req)
        arr1 = []
        arr2 = []
        lines = solve_req.data.split("\n")
        lines.each do |line|
            pieces = line.strip.split
            arr1.push(pieces[0].to_i)
            arr2.push(pieces[1].to_i)
        end

        sarr1 = arr1.sort
        sarr2 = arr2.sort

   
        diff = 0
        sarr1.each_with_index do |val, index|
            diff += differ(sarr1[index], sarr2[index])
        end

        return diff
    end

    def solvePart2(solve_req)
        arr1 = []
        arr2 = []
        lines = solve_req.data.split("\n")
        lines.each do |line|
            pieces = line.strip.split
            arr1.push(pieces[0].to_i)
            arr2.push(pieces[1].to_i)
        end

        ocount = 0
        arr1.each do |num|
            count = 0
            arr2.each do |onum|
                if onum == num
                    count+=1
                end
            end
            ocount += count*num 
        end
        return ocount
    end
end