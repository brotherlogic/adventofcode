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
        backlog.push([x,y,"EAST", 0, [[x,y,"EAST"]]])

        seen = Hash.new

        # Record the best paths into a given position
        best = Hash.new
        for y in 0..map.length()
            best[y] = Hash.new
            for x in 0..map.length()
                best[y][x] = Hash.new
                best[y][x]["NORTH"] = [99999999, []]
                best[y][x]["EAST"] = [99999999, []]
                best[y][x]["SOUTH"] = [99999999, []]
                best[y][x]["WEST"] = [99999999, []]
            end
        end

        foundv = 9999999999999999999
        foundx = []
        foundy = []
        foundd = []

        while backlog.length() > 0
            backlog = backlog.sort { |a,b| a[3] <=> b[3]}
            csearch = backlog[0]


            cbest = best[csearch[1]][csearch[0]][csearch[2]]
            if cbest[0] == csearch[3]
                cbest[1].push(csearch[4])
            elsif csearch[3] < cbest[0]
                cbest[0] = csearch[3]
                cbest[1] = [Marshal.load(Marshal.dump(csearch[4]))]
            end
          
            
            # We've seen all the paths at this point
            if csearch[3] > foundv
                return best, foundx, foundy, foundd
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
                #print "SOLUTION ", csearch, "\n"
                foundv = csearch[3]
                foundx.push(csearch[0])
                foundy.push(csearch[1])
                foundd.push(csearch[2])
            end

       
            if !seenv
                if csearch[2] == "EAST"
                    if map[csearch[1]][csearch[0]+1] != "#"
                        backlog.push([csearch[0]+1, csearch[1], "EAST", csearch[3]+1, [csearch[0], csearch[1], "EAST"]])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, [csearch[0], csearch[1], "EAST"]])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, [csearch[0], csearch[1], "EAST"]])
                    end
                elsif csearch[2] == "SOUTH"
                    if map[csearch[1]+1][csearch[0]] != "#"
                      
                        backlog.push([csearch[0], csearch[1]+1, "SOUTH", csearch[3]+1, [csearch[0], csearch[1], "SOUTH"]])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, [csearch[0], csearch[1], "SOUTH"]])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, [csearch[0], csearch[1], "SOUTH"]])
                    end
                elsif csearch[2] == "WEST"
                    if map[csearch[1]][csearch[0]-1] != "#"
                      
                        backlog.push([csearch[0]-1, csearch[1], "WEST", csearch[3]+1, [csearch[0], csearch[1], "WEST"]])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, [csearch[0], csearch[1], "WEST"]])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, [csearch[0], csearch[1], "WEST"]])
                    end
                elsif csearch[2] == "NORTH"
                    if map[csearch[1]-1][csearch[0]] != "#"
                      
                        backlog.push([csearch[0], csearch[1]-1, "NORTH", csearch[3]+1, [csearch[0], csearch[1], "NORTH"]])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, [csearch[0], csearch[1], "NORTH"]])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, [csearch[0], csearch[1], "NORTH"]])
                    end
                end
            end

            backlog = backlog[1..backlog.length() - 1]
        end
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        best, fx, fy, c = runSearch(map)
        return best[fy[0]][fx[0]][c[0]][0]
    end

    def fillPassed(x, y, d, best, passed)
        if d 
        passed[y][x] = true
        if best[y][x][d]
            best[y][x][d][1].each do |item|
                fillPassed(item[0], item[1], item[2], best, passed)
            end
        end
    end
    end

    def solvePart2(solve_req)
        map = buildMap(solve_req.data)

        passed = [Hash.new]
        for y in 0..map.length()-1
            passed[y] = Hash.new
            for x in 0..map[y].length()-1
                passed[y][x] = false
            end
        end

        best, fxi, fyi, directioni = runSearch(map)
        for i in 0..fxi.length()-1
            fx, fy, direction = fxi[i], fyi[i], directioni[i]
            fillPassed(fx, fy, direction, best, passed)
        end

        countv = 0
        passed.each do |row|
            row.each do |key, item|
                if item
                    countv += 1
                end
            end
        end

        return countv
    end
end
