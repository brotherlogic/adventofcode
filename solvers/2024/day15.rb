require 'logger'
require 'date'

class Day15

    def buildData(data)
        map = []
        instructions = []

        data.split("\n").each do |line|
            if line.include?("^")
                line.strip().split("").each do |item|
                    instructions.push(item)
                end
            else
                if line.strip().length() > 0
                    map.push(line.strip().split(""))
                end
            end
        end

        return map, instructions
    end

    def runInstruction(map, instr, rx, ry)
        print "ROBOT ", rx, ",", ry, "\n"
        if instr == "<"
            if map[ry][rx-1] == "#"
                return rx, ry
            end

            if map[ry][rx-1] == "."
                map[ry][rx] = "."
                map[ry][rx-1] = "@"
                return rx-1, ry
            end

            chain = 1
            while map[ry][rx-chain-1] == "O"
                chain+=1
            end

            if map[ry][rx-chain-1] == "#"
                return rx, ry
            else
                map[ry][rx-chain-1 ] = "O"
                map[ry][rx-1] = "@"
                map[ry][rx] = "."
                return rx-1, ry
            end
        elsif instr == ">"
            if map[ry][rx+1] == "#"
                return rx, ry
            end

            if map[ry][rx+1] == "."
                map[ry][rx] = "."
                map[ry][rx+1] = "@"
                return rx+1, ry
            end

            chain = 1
            while map[ry][rx+chain+1] == "O"
                chain+=1
            end

            if map[ry][rx+chain+1] == "#"
                return rx, ry
            else
                map[ry][rx+chain+1] = "O"
                map[ry][rx+1] = "@"
                map[ry][rx] = "."
                return rx+1, ry
            end
        elsif instr == "^"
            if map[ry-1][rx] == "#"
                return rx, ry
            end

            if map[ry-1][rx] == "."
                map[ry][rx] = "."
                map[ry-1][rx] = "@"
                return rx, ry-1
            end

            chain = 1

            while map[ry-chain-1][rx] == "O"
                chain+=1
            end

            if map[ry-chain-1][rx] == "#"
                return rx, ry
            else
                map[ry-chain-1][rx] = "O"
                map[ry-1][rx] = "@"
                map[ry][rx] = "."
                return rx, ry-1
            end
        else
            if map[ry+1][rx] == "#"
                return rx, ry
            end

            if map[ry+1][rx] == "."
                map[ry][rx] = "."
                map[ry+1][rx] = "@"
                return rx, ry+1
            end

            chain = 1
            while map[ry+chain+1][rx] == "O"
                chain+=1
            end

            if map[ry+chain+1][rx] == "#"
                return rx, ry
            else
                map[ry+chain+1][rx] = "O"
                map[ry+1][rx] = "@"
                map[ry][rx] = "."
                return rx, ry+1
            end
        end


        return -1, -1
    end

    def findRobot(map)
        for y in 0..map.length()-1
            for x in 0..map[y].length() - 1
                if map[y][x] == "@"
                    return x, y
                end
            end
        end
    end

    def gps(map)
        sumv = 0
        for y in 0..map.length()-1
            for x in 0..map[y].length() - 1
                if map[y][x] == "O"
                    sumv += 100*y+x
                end
            end
        end

        return sumv
    end

    def printMap(map)
        map.each do |line|
            line.each do |item|
                print item
            end
            print "\n"
        end
    end

    def solvePart1(solve_req)
        map, instructions = buildData(solve_req.data)
        robotx, roboty = findRobot(map)

        while instructions.length() > 0
            print "-----------------\n"
            print "MOVE ", instructions[0], "\n"
            robotx, roboty = runInstruction(map, instructions[0], robotx, roboty)
            printMap(map)
            instructions = instructions[1..instructions.length()-1]
        end

        return gps(map)
    end

    def solvePart2(solve_req)
        return 0
    end
end
