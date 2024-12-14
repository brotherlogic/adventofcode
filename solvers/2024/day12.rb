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
        left = x > 0 && map[y][x-1].start_with?(ch)
        right = x < map[0].length()-1 && map[y][x+1].start_with?(ch)
        up = y > 0 && map[y-1][x].start_with?(ch)
        down = y < map.length()-1 && map[y+1][x].start_with?(ch)

      
        count = 1
        perim = 0

        # Account for th 
        
        map[y][x] = ch + "."

        if left && !map[y][x-1].end_with?(".")
            c, p = buildFence(ch, x-1, y, map)
            count += c
            perim += p
        elsif !left
            perim += 1
        end

        if right && !map[y][x+1].end_with?(".")
            c, p = buildFence(ch, x+1, y, map)
            count += c
            perim += p
        elsif !right
            perim += 1
        end

        if up && !map[y-1][x].end_with?(".")
            c, p = buildFence(ch, x, y-1, map)
            count += c
            perim += p
        elsif !up
            perim += 1
        end

        if down && !map[y+1][x].end_with?(".")
            c, p = buildFence(ch, x, y+1, map)
            count += c
            perim += p
        elsif !down
            perim += 1
        end

        return count, perim
    end

    def followSides(ch, x, y, map)
        print "SEARCH ", x, ",", y, "\n" 
        point = "TOPRIGHT" 
        nx = x
        ny = y   
        turns = 1

        while nx != x || ny != y || point != "TOPLEFT"
            print "NOW ", nx, ",", ny , " w ", turns, " -> ", point, "\n"
            if point == "TOPRIGHT"
                if nx < map[0].length() -1 && ny > 0 && map[ny][nx+1] == ch && map[ny-1][nx+1] == ch
                    turns += 1
                    nx += 1
                    ny -=1
                    point = "TOPLEFT"
                elsif nx < map[0].length()-1 && map[ny][nx+1] == ch
                    nx += 1
                else
                    turns += 1
                    point = "BOTTOMRIGHT"
                end
            elsif point == "BOTTOMRIGHT"
                if ny < map.length() -1 && nx < map[0].length()-1 && map[ny+1][nx] == ch && map[ny+1][nx+1] == ch
                    turns += 1
                    nx += 1
                    ny +=1
                    point = "TOPRIGHT"
                elsif ny < map.length()-1 && map[ny+1][nx] == ch
                    ny += 1
                else
                    turns += 1
                    point = "BOTTOMLEFT"
                end
            elsif point == "BOTTOMLEFT"
                if ny < map.length() -1 && nx >0 && map[ny][nx-1] == ch && map[ny+1][nx-1] == ch
                    turns += 1
                    nx -= 1
                    ny +=1
                    point = "BOTTOMRIGHT"
                elsif nx > 0 && map[ny][nx-1] == ch
                    nx -= 1
                else
                    turns += 1
                    point = "TOPLEFT"
                end
            elsif point == "TOPLEFT"
                if ny > 0 -1 && nx > 0 && map[ny-1][nx] == ch && map[ny-1][nx-1] == ch
                    turns += 1
                    nx -= 1
                    ny -=1
                    point = "BOTTOMLEFT"
                elsif ny > 0 && map[ny-1][nx] == ch
                    ny -= 1
                else
                    turns += 1
                    point = "TOPRIGHT"
                end
            end
        end

        return turns
    end


    def buildFence2(ch, x, y, map)
        left = x > 0 && map[y][x-1].start_with?(ch)
        right = x < map[0].length()-1 && map[y][x+1].start_with?(ch)
        up = y > 0 && map[y-1][x].start_with?(ch)
        down = y < map.length()-1 && map[y+1][x].start_with?(ch)

      
        count = 1

        # Account for th 
        
        map[y][x] = ch + "."

        if left && !map[y][x-1].end_with?(".")
            c, p = buildFence2(ch, x-1, y, map)
            count += c
        end

        if right && !map[y][x+1].end_with?(".")
            c, p = buildFence2(ch, x+1, y, map)
            count += c
                end

        if up && !map[y-1][x].end_with?(".")
            c, p = buildFence2(ch, x, y-1, map)
            count += c
        end

        if down && !map[y+1][x].end_with?(".")
            c, p = buildFence2(ch, x, y+1, map)
            count += c
        end

        return count
    end


    def computeFence2(map)
        print map, "\n"
        sumv = 0
        for y in 0..map.length()-1
            for x in 0..map[y].length() - 1
                if !map[y][x].end_with?(".")
                     count = buildFence2(map[y][x], x, y, map)
                     sides = followSides(map[y][x], x, y, map)
                     print map[y][x], " ", count, ",", sides, "\n"
                    sumv += count*sides
                end
            end
        end
        return sumv
    end

    def computeFence(map)
        sumv = 0
        for y in 0..map.length()-1
            for x in 0..map[y].length() - 1
                if !map[y][x].end_with?(".")
                    perim, count = buildFence(map[y][x], x, y, map)
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
        map = buildMap(solve_req.data)
        width = computeFence2(map)
       return width
    end
end
