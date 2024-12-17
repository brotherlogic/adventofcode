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
        backlog.push([x,y,"EAST", 0, [[x,y]]])

        seen = Hash.new

        # Record the best paths into a given position
        best = Hash.new
        for y in 0..map.length()
            best[y] = Hash.new
            for x in 0..map.length()
                best[y][x] = [99999999, []]
            end
        end

        foundv = 9999999999999999999
        foundx = -1
        foundy = -1

        while backlog.length() > 0
            backlog = backlog.sort { |a,b| a[3] <=> b[3]}
            csearch = backlog[0]

            cbest = best[csearch[1]][csearch[0]]
            if cbest[0] == csearch[3]
                csearch[4].each do |item|
                    cbest[1].push(item)
                end
            end
            
            # We've seen all the paths at this point
            if csearch[3] > foundv
                return best[foundy][foundx] 
            end
          
            seenv = false
            if !seen.key?(csearch[0])
                seen[csearch[0]] = Hash.new
                seen[csearch[0]][csearch[1]] = Hash.new
                seen[csearch[0]][csearch[1]][csearch[2]] = true
            elsif !seen[csearch[0]].key?(csearch[1])
                seen[csearch[0]][csearch[1]] = Hash.new
                seen[csearch[0]][csearch[1]][csearch[2]] = true
            elsif !seen[csearch[0]][csearch[1]].key?(csearch[2])
                seen[csearch[0]][csearch[1]][csearch[2]] = true
            else
                seenv = true
            end

       
            if map[csearch[1]][csearch[0]] == "E"
                foundv = csearch[3]
                foundx = csearch[0]
                foundy = csearch[1]
            end

       
            if !seenv
                if csearch[2] == "EAST"
                    if map[csearch[1]][csearch[0]+1] != "#"
                        nn = Marshal.load(Marshal.dump(csearch[4]))
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
                        nn = Marshal.load(Marshal.dump(csearch[4]))
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
                        nn = Marshal.load(Marshal.dump(csearch[4]))
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
                        nn = Marshal.load(Marshal.dump(csearch[4]))
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
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        return runSearch(map)
    end

    def solvePart2(solve_req)
        return 0
    end
end
