require 'logger'
require 'date'

class Day11

    def blink(arr)
        narr = []
        arr.each do |elem|
            c = elem.to_s
            if elem == 0
                narr.push(1)
            elsif c.length() % 2 == 0
                narr.push(c[0..c.length()/2-1].to_i)
                narr.push(c[c.length()/2..c.length()-1].to_i)
            else
                narr.push(2024*elem)
            end
        end

     
        return narr
    end

    def fblink(arr, count)
        for i in 1..count
            arr = blink(arr)
        end

        return arr.length()
    end

    def toarr(data)
        narr = []
        data.strip().split().each do |item|
            narr.push(item.to_i)
        end
        return narr
    end

    def solvePart1(solve_req)
        return fblink(toarr(solve_req.data), 25)
    end
end
