class Day6

    def buildMap(data)
        map = []
        lines = data.split("\n")

        lines.each do |line|
            mline = line.strip().split("")
            map.push(mline)
        end

        return map
    end

    def findMan(map)
        for y in  0..map.length()
            for x in  0..map[y].length()
                if map[y][x] == "^"
                    return x,y
                end
            end
        end
    end

    def newDirection(direction)
        if direction == "NORTH"
            return "EAST"
        elsif direction == "EAST"
            return "SOUTH"
        elsif direction == "SOUTH"
            return "WEST"
        else
            return "NORTH"
        end
    end 

    def step(map, x, y, direction)
        nx, ny = x, y
        if direction == "NORTH"
            ny = ny-1
        elsif direction == "EAST"
            nx = nx +1
        elsif direction == "SOUTH"
            ny = ny+1
        else
            nx = nx - 1
        end

        # Have we gone out of bounds?
        if nx < 0 || nx >= map[0].length() || ny < 0 || ny >= map.length()
            return map,0,false
        end

        # Have we hit a wall ?
        if map[ny][nx] == "#"
            return step(map, x, y, newDirection(direction))
        end

        map[y][x] = "X"
        map[ny][nx] = "^"

        return map, nx, ny, direction, true
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)

        x,y = findMan(map)
        legal = true
        direction = "NORTH"
        while legal
            nmap, x, y, direction, legal = step(map, x, y, direction)
        end

        print nmap, "\n"

        count = 0
        nmap.each do |line|
            line.each do |item|
                if item == "X"
                    count += 1
                end
            end
        end

        return count + 1
    end
end