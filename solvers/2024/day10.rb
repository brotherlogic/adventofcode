require 'logger'
require 'date'

class Day10

    def buildMap(data)
        map = []
        found = []
        data.split("\n").each do |line|
            mline = []
            fline = []
            line.strip().split("").each do |item|
                mline.push(item.to_i)
                fline.push(".")
            end

            map.push(mline)
            found.push(fline)
        end

        return map, found
    end

    def search(map, found, x, y)
        isearch(map, found, x, y)

    
        sum = 0
        found.each do |line|
            line.each do |elem|
                if elem == "#"
                    sum+=1
                end
            end
        end
        return sum
    end

    def search2(map, found, x, y)
        return isearch2(map, found, x, y)
    end

    def isearch(map, found, x, y)
        if map[y][x] == 9
            found[y][x] = "#"
            return
        end

        if x > 0 && map[y][x-1] == map[y][x] + 1
            isearch(map,found, x-1, y)
        end
        if x < map[0].length()-1 && map[y][x+1] == map[y][x] + 1
            isearch(map, found, x+1, y)
        end
        if y > 0 && map[y-1][x] == map[y][x] + 1
            isearch(map,found, x, y-1)
        end
        if y < map.length()-1 && map[y+1][x] == map[y][x] + 1
            isearch(map, found, x, y+1)
        end
    end

    def isearch2(map, found, x, y)
        if map[y][x] == 9
            return 1
        end

        sumv = 0
        if x > 0 && map[y][x-1] == map[y][x] + 1
            sumv += isearch2(map,found, x-1, y)
        end
        if x < map[0].length()-1 && map[y][x+1] == map[y][x] + 1
            sumv += isearch2(map, found, x+1, y)
        end
        if y > 0 && map[y-1][x] == map[y][x] + 1
            sumv += isearch2(map,found, x, y-1)
        end
        if y < map.length()-1 && map[y+1][x] == map[y][x] + 1
            sumv += isearch2(map, found, x, y+1)
        end

        return sumv
    end

    
    def solvePart1(solve_req)
       map, found = buildMap(solve_req.data)

       sum = 0
        for y in 0..map.length()-1
            for x in 0..map[y].length()-1
                if map[y][x] == 0
                    val = search(map, Marshal.load(Marshal.dump(found)), x, y)
                    sum += val
                end
            end
        end

        return sum
    end

    def solvePart2(solve_req)
        map, found = buildMap(solve_req.data)
 
        sum = 0
         for y in 0..map.length()-1
             for x in 0..map[y].length()-1
                 if map[y][x] == 0
                     val = search2(map, Marshal.load(Marshal.dump(found)), x, y)
                     sum += val
                 end
             end
         end
 
         return sum
     end
end