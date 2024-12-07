require 'logger'
require 'date'

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
            return map,x, y,false
        end

        # Have we hit a wall ?
        if map[ny][nx] == "#"
            return step(map, x, y, newDirection(direction))
        end

        # Have we been in this square before
        if map[y][x][direction]
            return map, 0, 0, 0, false, true
        end

        if map[y][x] == "." || map[y][x] == "^"
            map[y][x] = "X-" + direction
        else
            map[y][x] += direction
        end

        return map, nx, ny, direction, true
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)

        x,y = findMan(map)
        legal = true
        direction = "NORTH"
        while legal
            nmap, x, y, direction, legal, before = step(map, x, y, direction)
        end

        print nmap, "\n"

        count = 0
        nmap.each do |line|
            line.each do |item|
                if item[0] == "X"
                    count += 1
                end
            end
        end

        return count + 1
    end

    def solveMap(map)
        x,y = findMan(map)
        legal = true
        before = false
        direction = "NORTH"
        while legal && !before
            nmap, x, y, direction, legal, before = step(map, x, y, direction)
        end

       return before
    end


    def solvePart2(solve_req)
        logger = Logger.new($stdout) 

        map = buildMap(solve_req.data)
        tmap = Marshal.load(Marshal.dump(map))
        x,y = findMan(tmap)
        legal = true
        before = false
        direction = "NORTH"
        while legal && !before
          tmap ,x,y,direction, legal, before = step(tmap, x, y, direction)
        end
        print "HERE ", x, y
        tmap[y][x] = "E"
        print "TMAP ", tmap, "\n"

        count = 0

        logger.info("STARTED")
        start = Time.now
        for y in 0..map.length() - 1
            for x in 0..map[y].length() - 1 
                nmap = Marshal.load(Marshal.dump(map))
                if nmap[y][x] != "^" && tmap[y][x] != "#" && tmap[y][x] != "."
                    nmap[y][x] = "#"
                    if solveMap(nmap)
                        count += 1
                    end
                end
            end
        end
        logger.info("FINISHED")
        print "FINISHED IN ", Time.now.to_i - start.to_i, "\n"
        return count
    end
end