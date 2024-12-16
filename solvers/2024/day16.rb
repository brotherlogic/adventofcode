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
        solutions = []

        x, y = findStart(map)
        backlog.push([x,y,"EAST", 0, [x, y]])

        seen = Hash.new
        foundv = 99999999999999999
        foundx = -1
        foundy = -1

        while backlog.length() > 0
            backlog = backlog.sort { |a,b| a[3] <=> b[3]}
            csearch = backlog[0]

            if csearch[3] > foundv
                print "BREAK ", csearch[3], " and ", foundv
                break
            end

            seenv = 999999999999999
            if !seen.key?(csearch[0])
                seen[csearch[0]] = Hash.new
                seen[csearch[0]][csearch[1]] = Hash.new
                seen[csearch[0]][csearch[1]][csearch[2]] = csearch[3]
            elsif !seen[csearch[0]].key?(csearch[1])
                seen[csearch[0]][csearch[1]] = Hash.new
                seen[csearch[0]][csearch[1]][csearch[2]] = csearch[3]
            elsif !seen[csearch[0]][csearch[1]].key?(csearch[2])
                seen[csearch[0]][csearch[1]][csearch[2]] = csearch[3]
            end

            sendv = seen[csearch[0]][csearch[1]][csearch[2]] 


            #print "BACKLOG ", csearch, " ", backlog.length(), "\n"

            if map[csearch[1]][csearch[0]] == "E"
                foundv = csearch[3]
                foundx = csearch[0]
                foundy = csearch[1] 
                solutions.push(csearch)
            end

            if csearch[3] <= seenv
                nn = Marshal.load(Marshal.dump(csearch[4]))
                if csearch[2] == "EAST"
                    if map[csearch[1]][csearch[0]+1] != "#"
                        nn.push([csearch[0]+1, csearch[1]])
                        backlog.push([csearch[0]+1, csearch[1], "EAST", csearch[3]+1, nn])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, nn])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, nn])
                    end
                elsif csearch[2] == "SOUTH"
                    if map[csearch[1]+1][csearch[0]] != "#"
                        nn.push([csearch[0], csearch[1]+1])
                        backlog.push([csearch[0], csearch[1]+1, "SOUTH", csearch[3]+1, nn])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, nn])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, nn])
                    end
                elsif csearch[2] == "WEST"
                    if map[csearch[1]][csearch[0]-1] != "#"
                        nn.push([csearch[0]-1, csearch[1]])
                        backlog.push([csearch[0]-1, csearch[1], "WEST", csearch[3]+1, nn])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, nn])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, nn])
                    end
                elsif csearch[2] == "NORTH"
                    if map[csearch[1]-1][csearch[0]] != "#"
                        nn.push([csearch[0], csearch[1]-1])
                        backlog.push([csearch[0], csearch[1]-1, "NORTH", csearch[3]+1, nn])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, nn])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, nn])
                    end
                end
            end

            backlog = backlog[1..backlog.length() - 1]
        end

        return solutions
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        return runSearch(map)[0][3]
    end

    def solvePart2(solve_req)
        seen = Hash.new
        map = buildMap(solve_req.data)
        solutions = runSearch(map)
        print "FOUND ", solutions.length(), "\n"
        countv = 0
        seen = Hash.new
        solutions.each do |solution|
            solution[4].each do |entry|
                if !seen.has_key?(entry[0])
                    seen[entry[0]] = Hash.new
                    seen[entry[0]][entry[1]] = true
                    countv+=1
                elsif !seen[entry[0]].has_key?(entry[1])
                    seen[entry[0]][entry[1]] = true
                    countv+=1
                end
            end
        end
        return countv
    end
end
