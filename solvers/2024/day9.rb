require 'logger'
require 'date'

class Day9

    def convert(str)
        narr = []
        free = false
        counter = 0
        for char in str.split("")
            if free
                for i in 1..char.to_i
                    narr.push(".")
                end
                free = false
            else
                for i in 1..char.to_i
                    narr.push(counter)
                end
                counter += 1
                free = true
            end 
        end

        return narr
    end

    def compact(arr)
        spoint = 0
        epoint = arr.length()-1

        while spoint <= epoint
            while arr[spoint] != "."
                spoint+=1
            end

            while arr[epoint] == "."
                epoint-=1
            end

            if spoint < epoint
                temp = arr[spoint]
                arr[spoint] = arr[epoint]
                arr[epoint] = temp
            end
        end

        return arr
    end

    def checksum(arr)
        sum = 0
        for i in 0..arr.length()-1
            if arr[i] != "."
                sum += i*arr[i]
            end
        end
        return sum
    end

    def solvePart1(solve_req)
        fstring = convert(solve_req.data)
        comp = compact(fstring)
        return checksum(comp)
    end
end