require 'logger'
require 'date'

class Day16


    def buildMap(data)
        map = []
        data.split("\n").each do |line|
            map.push(line.strip().split(""))
        end
        return map
    end

    def findStart(map)
        for y in 0..map.length()-1
            for x in 0..map[y].length()-1
                if map[y][x] == "S"
                    return x, y
                end
            end
        end
    end

    def runSearch(map)
        backlog = []

        x, y = findStart(map)
        backlog.push([x,y,"EAST", 0])

        while backlog.length() > 0
            backlog = backlog.sort { |a,b| a[3] <=> b[3]}
            csearch = backlog[0]


            #print "BACKLOG ", csearch, " ", backlog.length(), "\n"

            if map[csearch[1]][csearch[0]] == "E"
                return csearch[3]
            end

            if csearch[2] == "EAST"
                if map[csearch[1]][csearch[0]+1] != "#"
                    backlog.push([csearch[0]+1, csearch[1], "EAST", csearch[3]+1])
                end

                if map[csearch[1]+1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000])
                end

                if map[csearch[1]-1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000])
                end
            elsif csearch[2] == "SOUTH"
                if map[csearch[1]+1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1]+1, "SOUTH", csearch[3]+1])
                end

                if map[csearch[1]][csearch[0]+1] != "#"
                    backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000])
                end

                if map[csearch[1]][csearch[0]-1] != "#"
                    backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000])
                end
            elsif csearch[2] == "WEST"
                if map[csearch[1]][csearch[0]-1] != "#"
                    backlog.push([csearch[0]-1, csearch[1], "WEST", csearch[3]+1])
                end

                if map[csearch[1]+1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000])
                end

                if map[csearch[1]-1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000])
                end
            elsif csearch[2] == "NORTH"
                if map[csearch[1]-1][csearch[0]] != "#"
                    backlog.push([csearch[0], csearch[1]-1, "NORTH", csearch[3]+1])
                end

                if map[csearch[1]][csearch[0]+1] != "#"
                    backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000])
                end

                if map[csearch[1]][csearch[0]-1] != "#"
                    backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000])
                end
            end

            backlog = backlog[1..backlog.length() - 1]
        end
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        return runSearch(map)
    end

    def solvePart2(solve_req)
        return 0
    end
end
