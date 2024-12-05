class Day5

    def buildMapping(data)
        lines = data.split("\n")
        maps = []
        lines.each do |line|
            elems = line.strip.split("|")
            if elems.length() == 2
                maps.push([elems[0], elems[1]])
            end
        end

        return maps
    end

    def validate(arr, map)
        cnum = arr[arr.length()-1]

        map.each do |item|
            if item[0] == cnum
                arr.each do |elem|
                    if elem == item[1]
                        print "Found ", item, " in ", arr, "\n"
                        return false
                    end
                end
            end
        end
        return true
    end

    def checkPages(data)
        count = 0

        map = buildMapping(data)

        lines = data.split("\n")
        lines.each do |line|
            elems = line.strip.split(",")
            if elems.length() > 1
                safe = true
                for i in 1..elems.length()-1
                    if !validate(elems[0..i], map)
                        safe = false
                    end
                end
                if safe
                    puts line.strip
                    count += elems[elems.length()/2].to_i
                else
                    print "Unsafe: ", line.strip, "\n"
                end
            end
        end

        return count
    end

    def solvePart1(solve_req)
        return checkPages(solve_req.data)
    end
end