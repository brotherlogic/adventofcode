class Day1
    def solvePart1(solve_req)
        arr1 = []
        arr2 = []
        lines = solve_req.data.split("\n")
        lines.each do |line|
            pieces = line.strip.split
            arr1.push(pieces[0].to_i)
            arr2.push(pieces[0].to_i)
        end
        return 5
    end
end