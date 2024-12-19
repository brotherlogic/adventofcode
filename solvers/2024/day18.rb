require 'logger'
require 'date'

class Day18

   def buildMap(data, msize)
      map = []
      for y in 0..msize
         tmap = []
         for x in 0..msize
            tmap.push(".")
         end
         map.push(tmap)
      end

      data.each do |coord|
         elems = coord.strip().split(",")
         map[elems[1].to_i][elems[0].to_i] = "#"
      end

      return map
   end

   def printMap(map)
      map.each do |row|
         row.each do |item|
            print item
         end
         print "\n"
      end
   end

   def solveMap(map)
      cstart = [0, 0, 0, [0,0, 0]]
      backlog = [cstart]
      seen = Hash.new
      while backlog.length() > 0
         backlog = backlog.sort { |a,b| a[2] <=> b[2]}
         curr = backlog[0]
         backlog = backlog[1..backlog.length()]

         x = curr[0]
         y = curr[1]
         d = curr[2]
         e = curr[3]

         if x == map.length()-1 && y == map.length() - 1
            #print e, "\n"
            #printMap(map)
            return d
         end

         if !seen.key?(y) || !seen[y].key?(x)
            if !seen.key?(y)
               seen[y] = Hash.new
            end
            seen[y][x] = true

            if y > 0 && map[y-1][x] != "#"
               nn = Marshal.load(Marshal.dump(e))
               nn.push([x,y-1,d+1])
               backlog.push([x, y-1, d+1,nn])
            end

            if x > 0 && map[y][x-1] != "#"
               nn = Marshal.load(Marshal.dump(e))
               nn.push([x-1,y,d+1])

               backlog.push([x-1, y, d+1,nn])
            end

            if x < map.length() - 1 && map[y][x+1] != "#"
               nn = Marshal.load(Marshal.dump(e))
               nn.push([x+1,y,d+1])

               backlog.push([x+1, y, d+1,nn])
            end

            if y < map.length() -1 && map[y+1][x] != "#"
               nn = Marshal.load(Marshal.dump(e))
               nn.push([x,y+1,d+1])

               backlog.push([x, y+1, d+1, nn])
            end
         end
      end

      return -1
   end

   def solve(data, msize, len)
      map = buildMap(data.split("\n")[0..len-1], msize)
      return solveMap(map)
   end

    def solvePart1(solve_req)
       return solve(solve_req.data, 70, 1024)
    end

    def solvePart2(solve_req)
      return solve2(solve_req.data, 70)
   end

    def solve2(data, msize)
        lines = data.split("\n")
        bottom = 0
        top = lines.length()-1

        while top-bottom > 1
            map = buildMap(lines[0..(top+bottom)/2], msize)
            if solveMap(map) > 0
               bottom = (top+bottom)/2
            else
               top = (top+bottom)/2
            end
         end

         return lines[top ].strip()
     end
end
