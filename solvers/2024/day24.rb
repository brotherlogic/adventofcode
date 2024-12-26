require 'logger'
require 'date'

class Day24

    def buildMapping(data)
        database = Hash.new
        mapper = Hash.new
        data.split("\n").each do |line|
            if line.include?(":")
                elems = line.strip().split(":")
                database[elems[0].strip()] = elems[1].strip().to_i
            elsif line.include?("->")
                pieces = line.strip().split("->")
                spieces = pieces[0].strip().split()
                database[spieces[0]] = -1
                database[spieces[2]] = -1
                database[pieces[1]] = -1
                mapper[pieces[1]] = spieces
            end
        end

        return database, mapper
    end

    def resolve(target, database, mapper)
        if database[target] != -1
            return database[target]
    end

    def solvePart1(solve_req)
        database, mapper = buildMapping(solve_req.data)
        print resolve("z01", database, mapper), "\n"
       return 0
    end

    def solvePart2(solve_req)
        return solvePart1(solve_req)
    end
end