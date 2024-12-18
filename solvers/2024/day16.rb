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
        foundx = -1
        foundy = -1
        foundd = ""

        while backlog.length() > 0
            backlog = backlog.sort { |a,b| a[3] <=> b[3]}
            csearch = backlog[0]

            #print csearch[0], ",", csearch[1], " -> (", csearch[3], ") ", csearch[4], " -> ", backlog, "\n"

            cbest = best[csearch[1]][csearch[0]][csearch[2]]
            if cbest[0] == csearch[3]
                csearch[4].each do |item|
                    cbest[1].push(item)
                end
            elsif csearch[3] < cbest[0]
                cbest[0] = csearch[3]
                cbest[1] = Marshal.load(Marshal.dump(csearch[4]))
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
                foundv = csearch[3]
                foundx = csearch[0]
                foundy = csearch[1]
                foundd = csearch[2]
            end

       
            if !seenv
                if csearch[2] == "EAST"
                    if map[csearch[1]][csearch[0]+1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0]+1, csearch[1], "EAST"])
                        backlog.push([csearch[0]+1, csearch[1], "EAST", csearch[3]+1, nn2])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "SOUTH"])
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, nn2])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "NORTH"])
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, nn2])
                    end
                elsif csearch[2] == "SOUTH"
                    if map[csearch[1]+1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1]+1, "SOUTH"])
                      
                        backlog.push([csearch[0], csearch[1]+1, "SOUTH", csearch[3]+1, nn2])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "EAST"])
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, nn2])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "WEST"])
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, nn2])
                    end
                elsif csearch[2] == "WEST"
                    if map[csearch[1]][csearch[0]-1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0]-1, csearch[1], "WEST"])
                      
                        backlog.push([csearch[0]-1, csearch[1], "WEST", csearch[3]+1, nn2])
                    end

                    if map[csearch[1]+1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "SOUTH"])
                        backlog.push([csearch[0], csearch[1], "SOUTH", csearch[3]+1000, nn2])
                    end

                    if map[csearch[1]-1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "NORTH"])
                        backlog.push([csearch[0], csearch[1], "NORTH", csearch[3]+1000, nn2])
                    end
                elsif csearch[2] == "NORTH"
                    if map[csearch[1]-1][csearch[0]] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1]-1, "NORTH"])
                      
                        backlog.push([csearch[0], csearch[1]-1, "NORTH", csearch[3]+1, nn2])
                    end

                    if map[csearch[1]][csearch[0]+1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "EAST"])
                        backlog.push([csearch[0], csearch[1], "EAST", csearch[3]+1000, nn2])
                    end

                    if map[csearch[1]][csearch[0]-1] != "#"
                        nn2 = Marshal.load(Marshal.dump(csearch[4]))
                        nn2.push([csearch[0], csearch[1], "WEST"])
                        backlog.push([csearch[0], csearch[1], "WEST", csearch[3]+1000, nn2])
                    end
                end
            end

            backlog = backlog[1..backlog.length() - 1]
        end
    end

    def solvePart1(solve_req)
        map = buildMap(solve_req.data)
        best, fx, fy = runSearch(map)
        return best[fy][fx][0]
    end

    def solvePart2(solve_req)
        map = buildMap(solve_req.data)

        passed = Hash.new
        for y in 0..map.length()-1
            passed[y] = Hash.new
            for x in 0..map[y].length()-1
                passed[y][x] = true
            end
        end

        best, fx, fy, direction = runSearch(map)
        print "BEST ", fx, ",", fy, "=>", direction, "\n"
        print "SO ", best[fy][fx][direction], "\n"
        print "BUT ", best[7][5]["EAST"], "\n"
        
        best[fy][fx][direction][1].each do |path|
            #print "HERE ", path, " -> ", "\n"
            print "CHECK ", best[path[1]][path[0]], "\n"

            bestv = best[path[1]][path[0]]["NORTH"][0]
            if best[path[1]][path[0]]["EAST"][0] < bestv
                bestv = best[path[1]][path[0]]["EAST"][0]
            end
            if best[path[1]][path[0]]["SOUTH"][0] < bestv
                bestv = best[path[1]][path[0]]["SOUTH"][0]
            end
            if best[path[1]][path[0]]["WEST"][0] < bestv
                bestv = best[path[1][path[0]]]["WEST"][0]
            end

            #print "NOW ", bestv, "\n"

            best[path[1]][path[0]].each do |direction, val|
                if val[0] == bestv
                    print "BUTT ", path[0], ",", path[1], " ", direction, " -> ", val, "\n"
                   val[1].each do |item|
                        passed[item[0]][item[1]] = true
                    end
                end
            end
        end

        print passed, "\n"

        countv = 0
        passed.each do |row|
            row.each do |item|
                if item
                    countv += 1
                end
            end
        end

        return countv
    end
end
