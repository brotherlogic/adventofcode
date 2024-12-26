require 'logger'
require 'date'

class Day25

    def convert(arr)
        narr = []
        for x in 0..4
            count = 0
            for y in  0..6
                if arr[y][x] == "#"
                    count+=1
                end
            end
            narr.push(count-1)
        end
        return narr
    end

    def buildKeysAndLocks(data)
        inkey = 0
        arr = []
        keys = []
        locks = []
        data.split("\n").each do |nline|
            line = nline.strip()
            if inkey == 0 && line == "#####"
                inkey = 1
                arr.push(line)
            elsif line.strip().length() == 0
                if inkey == 1
                    locks.push(convert(arr))
                else
                    keys.push(convert(arr))
                end
                arr = []
                inkey = 0
            elsif inkey == 0 && line == "....."
                inkey = 2
                arr.push(line)
            else
                arr.push(line)
            end
        end

        if inkey == 1
            locks.push(convert(arr))
        else
            keys.push(convert(arr))
        end

        return keys, locks
    end

    def solvePart1(solve_req)
        keys, locks = buildKeysAndLocks(solve_req.data)
        print "LOCKS ", locks, "\n"
        print "KEYS ", keys, "\n"

        count = 0
        locks.each do |lock|
            keys.each do |key|
                print lock, " -> ", key, "\n"
                found = true
                for i in 0..4
                    if lock[i] + key[i] >= 6
                        print "MISS ", lock[i], " and ", key[i], "\n"
                        found = false
                    end
                end
                if found
                    count += 1
                end
            end
        end

        return count
    end
end