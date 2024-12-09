require 'logger'
require 'date'

class Day8

    def buildMap(data)
        map = []
        finds = []
        data.split("\n").each do |line|
            map.push(line.strip.split(""))
            blank = []
            line.strip.split("").each do |piece|
                blank.push(" ")
            end
            finds.push(blank)
        end

        return map, finds
    end

    def addToMap(x, y, nx, ny, finds)
        print x,",",y, " -> ",nx,",",ny,"\n"

        a1x = x-(nx-x)
        a1y = y-(ny-y)

        if a1x >= 0 && a1y >= 0 && a1y < finds.length() && a1x < finds[0].length()
            finds[a1y][a1x] = "#"
        end

        print "A1 ", a1x, ",", a1y, "\n"

        a2x = nx+(nx-x)
        a2y = ny+(ny-y)

        if a2x >= 0 && a2y >= 0 && a2y < finds.length() && a2x < finds[0].length()
            finds[a2y][a2x] = "#"
        end

        print "A2 ", a2x,",",a2y,"\n"
    end

    def solvePart1(solve_req)
       map, finds = buildMap(solve_req.data)

   
       for y in 0..map.length()-1
            for x in 0..map[y].length()-1
                if map[y][x] != "."
                    # Found an antenna, look for other ones
                    for ny in y..map.length()-1
                        for nx in 0..map[y].length()-1
                            if ny != y || nx > x
                                if map[ny][nx] == map[y][x]
                                    addToMap(x, y, nx, ny, finds)
                                end
                            end
                        end
                    end
                end
            end
        end

        count = 0
        finds.each do |line|
            line.each do |elem|
                if elem == "#"
                    count += 1
                end
            end
        end

        return count
    end
end