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

    def search(towels, thing)
        if thing.length() == 0 
            return true
        end

        towels.each do |towel|
            if thing.start_with?(towel)
                result = search(towels, thing[towel.length()..thing.length()-1])
                if result
                    return result
                end
            end
        end

        return false
    end

    def solvePart1(solve_req)
        towels, things = buildData(solve_req.data)
       
        count = 0
        things.each do |thing|
            if search(towels, thing)
                count += 1
            end
        end

        return count
    end

    def solvePart2(solve_req)
        return solvePart1(solve_req)
    end
end