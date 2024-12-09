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

    def compactFull(arr)
        biggest = 0
        i = arr.length() - 1
        while i >= 0
            if arr[i] != "."
                biggest = arr[i]
                break
            end
            i -= 1
        end

    
        epoint = arr.length()-1
    

        # Find the first epoint
        while epoint >= 0
            while arr[epoint] != biggest
                epoint -= 1
                if epoint < 0
                    break
                end
            end

            espoint = epoint-1
            while arr[espoint] == arr[epoint]
                espoint -=1
            end
            espoint += 1

         
            numlen = epoint - espoint + 1

            # find a place to put this
            spoint = 0
            sspoint = 0
            while spoint < arr.length()
                if spoint > espoint 
                    break
                end
                if arr[spoint] != "."

                    if spoint - sspoint >= numlen 
                        # We can place
                        for i in sspoint..sspoint+numlen-1
                            arr[i] = arr[espoint]
                        end

                        for i in espoint..epoint
                            arr[i] = "."
                        end

                        break
                    end

                    spoint += 1
                    sspoint = spoint
                else
                    spoint += 1
                end
            end

            biggest -= 1
            epoint = espoint-1
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


    def solvePart2(solve_req)
        fstring = convert(solve_req.data)
        comp = compactFull(fstring)
        return checksum(comp)
    end
end