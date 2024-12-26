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
                if !database.has_key?(spieces[0].strip())
                    database[spieces[0].strip()] = -1
                end
                if !database.has_key?(spieces[2].strip())
                database[spieces[2].strip()] = -1
                end
                if !database.has_key?(pieces[1].strip())
                database[pieces[1].strip()] = -1
                end
                mapper[pieces[1].strip()] = spieces
            end
        end

        return database, mapper
    end

    def resolve(target, database, mapper)
        if database[target] != -1
            return database[target]
        end

    
        nval = -1
        lhs = resolve(mapper[target][0], database, mapper)
        rhs = resolve(mapper[target][2], database, mapper)

        if mapper[target][1] == "XOR"
            if (lhs == 1 && rhs == 0) || (lhs == 0 && rhs == 1)
                nval = 1
            else
                nval = 0

            end
        elsif mapper[target][1] == "OR"
            if lhs ==  1 || rhs == 1
                nval = 1
            else
                nval = 0
            end
        elsif mapper[target][1] == "AND"
            if lhs == 1 && rhs == 1
                nval = 1
            else 
                nval = 0
            end
        else
            print "FAILURE"
        end
        
        database[target] = nval
        return nval
    end

    def solvePart1(solve_req)
        database, mapper = buildMapping(solve_req.data)
        str = ""
        for i in 0..99
            if i < 10
                if database.has_key?("z0" + i.to_s)
                    str += resolve("z0" + i.to_s, database, mapper).to_s
                end
            else
                if database.has_key?("z" + i.to_s)
                    str += resolve("z" + i.to_s, database, mapper).to_s
                else
                    break
                end
            end
        end
        print database, "\n"
        print "GOT ", str.reverse, "\n"
       return str.reverse.to_i(2)
    end

    def solvePart2(solve_req)
        return solvePart1(solve_req)
    end
end