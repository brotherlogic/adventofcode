require 'logger'
require 'date'

class Day23

    def buildMatch(data)
        map = Hash.new
        data.split("\n").each do |line|
            elems = line.strip().split("-")
            
            if !map.key?(elems[0])
                map[elems[0]] = [elems[1]]
            else
                map[elems[0]].push(elems[1])
            end

            if !map.key?(elems[1])
                map[elems[1]] = [elems[0]]
            else
                map[elems[1]].push(elems[0])
            end
        end

        return map
    end

    def runSearch(st, map, sofar, seen)
        sofar.push(st)

        if sofar.length() == 3
            if map[sofar[0]].include?(sofar[2])
                #print "Seen ", sofar, "\n"
                seen[sofar.sort.join("")] = true
                #print "Res ", seen, "\n"
            end
            return
        end

        map[st].each do |item|
            found = true
            sofar.each do |seen|
                if seen == item 
                    found = false
                    break
                end
            end
            if found
                ns = Marshal.load(Marshal.dump(sofar))
                runSearch(item, map, ns, seen)
            end
        end
    end

    def solvePart1(solve_req)
        map = buildMatch(solve_req.data)

        seen = Hash.new
        map.each do |key, val|
            if key.start_with?("t")
                runSearch(key, map, [], seen)
            end
        end

        #print "FINAL ", seen, "\n"
       return seen.length()
    end

    def solvePart2(solve_req)
        return solvePart1(solve_req)
    end
end