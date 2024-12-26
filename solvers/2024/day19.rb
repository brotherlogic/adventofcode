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
     
        count = 0
        things.each do |thing|
            results =  search(towels, thing, [])
            if results.length() > 0
                count += 1
            end
        end

        return count
    end

    def reverseMap(results, map, sofara)
      
        if results.length() == 0
            return 1
        end

        sofar = 0
        map.each do |key, vals|
            vals.each do |val|
                found = true
                for i in 0..val.length()-1
                    if results[i] != val[i]
                        found = false
                        break
                    end
                end

               
                if found
                    sofar = reverseMap(results[val.length()..results.length()-1], map, sofara + [key])
                end
            end
        end
        return sofar + reverseMap(results[1..results.length()-1], map, sofara + [results[0]])
    end

    def solvePart2(solve_req)
        ftowels, things = buildData(solve_req.data)
        return 0
       
        towels,map = trimTowels(ftowels)
        count = 0
        things.each do |thing|
            print thing, " with", map, "\n"
            results =  search(towels, thing, [])
            results.each do |result|
                sumv = reverseMap(result, map, [])
                count += sumv
            end
        end

        return count
    end
end