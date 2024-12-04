class Day4

    def buildArr(data)
        arr = []

        lines = data.split("\n")
        safe = 0
        lines.each do |line|
            narr = []
            line.strip.each_char do |elem|
                narr.push(elem)
            end
            
            arr.push(narr)
        end

        return arr
    end

    def findXmas(arr, x, y, finder, i, adjx, adjy)

        if x < 0 || y < 0
            return 0
        end

        if x >= arr.length() || y >= arr.length()
            return 0
        end

        if arr[y][x] != finder[i]
            return 0
        end

        if i == finder.length()-1
            return 1
        end

      
        return findXmas(arr, x+adjx, y+adjy, finder, i+1, adjx, adjy)
    end

    def findXmasFirst(arr, x, y, finder, i)
        return findXmas(arr, x, y, finder, i, 0, 1) +
        findXmas(arr, x, y, finder, i, 0, -1) + 
        findXmas(arr, x, y, finder, i, 1, 0) + 
        findXmas(arr, x, y, finder, i, -1, 0) + 
        findXmas(arr, x, y, finder, i, 1, 1) +
        findXmas(arr, x, y, finder, i, 1, -1) +
        findXmas(arr, x, y, finder, i, -1, 1) + 
        findXmas(arr, x, y, finder, i, -1, -1)
    end

    def solvePart1(solve_req)
      arr = buildArr(solve_req.data)

      count = 0

      for y in  0..arr.length()-1
        for x in  0..arr[y].length()-1
            count += findXmasFirst(arr, x, y, "XMAS", 0)
        end
      end

      return count
    end
end