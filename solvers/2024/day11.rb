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

    def rblink(num, count)
        if count == 0
            return 1
        end

        c = num.to_s
        if num == 0
            return rblink(1, count-1)
        elsif c.length() % 2 == 0
            return rblink(c[0..c.length()/2-1].to_i, count-1) + rblink(c[c.length()/2..c.length()-1].to_i, count-1)
        else
            return rblink(2024*num, count-1)
        end

        print "HUH"
    end

    def fblink(arr, count)
        for i in 1..count
            arr = blink(arr)
        end

        return arr.length()
    end

    def frblink(arr, count)
        sumv = 0
        arr.each do |item|
            sumv += rblink(item, count)
        end
        return sumv
    end

    def toarr(data)
        narr = []
        data.strip().split().each do |item|
            narr.push(item.to_i)
        end
        return narr
    end

    def buildMap(arr, val)
        map = Hash.new

        mEntry = Struct.new(:len, :follow)

        while arr.length() > 0
            first = arr[0]
            fval = first
            arr = arr[1..arr.length()-1]

            count = 1
            while first.to_s.length() % 2 != 0
                if first == 0
                    first = 1
                else
                    first *= 2024
                end
                count += 1
            end

            c = first.to_s
            follow1 = c[0..c.length()/2-1].to_i
            follow2 = c[c.length()/2..c.length()-1].to_i
            map[fval] = mEntry.new(count, [follow1,follow2])
            
            if !map.key?(follow1)
                arr.push(follow1)
            end
            if !map.key?(follow2)
                arr.push(follow2)
            end
        end

        return map
    end

    def sresolve(num, map, count, cache)
        found = map[num]


        if cache.key?(num) && cache[num].key?(count)
            return cache[num][count]
        end

        if found.len > count
            if !cache.key?(num)
                cache[num] = Hash.new
            end
            cache[num][count] = 1
            return 1
        else
            val =  sresolve(found.follow[0], map, count-found.len, cache) + sresolve(found.follow[1], map, count-found.len, cache)
            if !cache.key?(num)
                cache[num] = Hash.new
            end
            cache[num][count] = val
            return val
        end
    end

    def resolve(arr, map, count)
        sumv = 0
        arr.each do |num|
            cache = Hash.new
            sumv += sresolve(num, map, count, cache)
        end
        return sumv
    end

    def solvePart1(solve_req)
        map = buildMap(toarr(solve_req.data), 25)
        return resolve(toarr(solve_req.data), map, 25)
    end
    def solvePart2(solve_req)
        map = buildMap(toarr(solve_req.data), 75)
        print map, "\n"
        return resolve(toarr(solve_req.data), map, 75)
    end
end
