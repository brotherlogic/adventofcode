require 'logger'
require 'date'

class Day12
    def buildMap(data)
        map = []
        data.split("\n").each do |line|
            row = []
            line.strip().split("") do |item|
                row.push(item)
            end
            map.push(row)
        end

        return map
    end

    def buildFence(ch, x, y, map)
        print "FENCE ", x, ",", y, " with ", map[y][x], "\n"
     
        left = x > 0 && map[y][x-1].start_with?(ch)
        right = x < map[0].length()-1 && map[y][x+1].start_with?(ch)
        up = y > 0 && map[y-1][x].start_with?(ch)
        down = y < map.length()-1 && map[y+1][x].start_with?(ch)

      
        count = 1
        perim = 0

        # Account for the edge

        
        map[y][x] = ch + "."

        if left && !map[y][x-1].end_with?(".")
            c, p = buildFence(ch, x-1, y, map)
            count += c
            perim += p
        elsif !left
            print y, "," , x, " LEFT", "\n"
            perim += 1
        end

        if right && !map[y][x+1].end_with?(".")
            c, p = buildFence(ch, x+1, y, map)
            count += c
            perim += p
        elsif !right
            print y, "," , x, " RIGHT", "\n"
            perim += 1
        end

        if up && !map[y-1][x].end_with?(".")
            c, p = buildFence(ch, x, y-1, map)
            count += c
            perim += p
        elsif !up
            print y, "," , x, " UP", "\n"
            perim += 1
        end

        if down && !map[y+1][x].end_with?(".")
            c, p = buildFence(ch, x, y+1, map)
            count += c
            perim += p
        elsif !down
            print y, "," , x, " DOWN", "\n"
            perim += 1
        end

        return count, perim
    end

    def computeFence(map)
        print map, "\n"
        sumv = 0
        for y in 0..map.length()-1
            for x in 0..map[y].length() - 1
                if !map[y][x].end_with?(".")
                    perim, count = buildFence(map[y][x], x, y, map)
                    print sumv, " -> ", perim, ",", count, "\n"
                    sumv += perim*count
                end
            end
        end
        return sumv
    end


    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        width = computeFence(map)
       return width
    end
    def solvePart2(solve_req)
        return 0
    end
end
