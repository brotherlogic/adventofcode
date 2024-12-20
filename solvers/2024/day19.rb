require 'logger'
require 'date'

class Day19
    
    def buildData(data)
        lines = data.split("\n")

        towels = []
        blocks = lines[0].strip().split(",")
        blocks.each do |item|
            towels.push(item.strip())
        end

        things = []
        lines[2..lines.length()-1].each do |item|
            things.push(item.strip())
        end

        return towels, things
    end

    def buildRainbow(sofar, towels, rainbow, mlen)
        if sofar.length() > mlen
            return
        end

        rainbow[sofar] = true

        towels.each do |item|
            buildRainbow(sofar + item, towels, rainbow, mlen)
        end
    end

    def buildRainbowTable(towels, mlen)
        rainbow = Hash.new
        
        towels.each do |item|
            buildRainbow(item, towels, rainbow, mlen)
        end

        return rainbow
    end

    def search(towels, thing, sofar)
        if thing.length() == 0 
            return [sofar]
        end

        results = []
        towels.each do |towel|
            if thing.start_with?(towel)
                result = search(towels, thing[towel.length()..thing.length()-1], sofar + [towel])
                results = results + result
            end
        end

        return results
    end

    def trimTowels(towels)
        ntowels = []
        map = Hash.new()
        res = search(towels[1..towels.length()-1], towels[0], [])
        if res.length() == 0
            ntowels.push(towels[0])
        else
            map[towels[0]] = res
        end

        for i in 1..towels.length()-2
            res =  search(towels[0..i-1]+towels[i+1..towels.length()-1], towels[i], [])
        
            if res.length() == 0
                ntowels.push(towels[i])
            else
                map[towels[i]] = res
            end
        end

        res = search(towels[0..towels.length()-2], towels[towels.length()-1], [])
        if res.length() == 0
            ntowels.push(towels[towels.length()-1])
        else
            map[towels[towels.length()-1]] = res
        end

        return ntowels, map
    end

    def solvePart1(solve_req)
        ftowels, things = buildData(solve_req.data)
       
        towels,map = trimTowels(ftowels)
        print "TRIMMED ", ftowels.length(), " to ", towels.length(), "\n"
        print "POST ", towels, "\n"
        print "MAP ", map, "\n"

        count = 0
        things.each do |thing|
            results =  search(towels, thing, [])
            print thing, " -> ", results, "\n"
            if results.length() > 0
                count += 1
            end
        end

        return count
    end

    def solvePart2(solve_req)
        ftowels, things = buildData(solve_req.data)
       
        towels,map = trimTowels(ftowels)
        count = 0
        things.each do |thing|
            results =  search(towels, thing, [])
            print thing, " -> ", results, "\n"
            results.each do |result|
                countin = 1
                finds = 0
                map.each do |key, val|
                    val.each do |searcher|
                        print "SEARCH", searcher, "\n"
                        for i in 0..result.length()-searcher.length()
                            found = true
                            for j in 0..searcher.length()-1
                                print "COMP ", result[i+j], " and ", searcher[j], "\n"
                                if result[i+j] != searcher[j]
                                    found = false
                                    break
                                end
                            end
                            if found
                                print "FOUND ", searcher, " -> ", result, "\n"
                                finds += 1
                            end
                        end
                    end
                end
                print "COUNTIN ", countin, " -> ", finds, "\n"
                count += countin + finds
            end
        end

        return count
    end
end