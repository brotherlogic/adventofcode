require 'logger'
require 'date'

class Day13

   def resolve(x1, x2, y1, y2, gx, gy)
      for y in 1..100
         top = (gx - x2*y)
         bot = x1
         if top % bot == 0
            x = top / bot

            if y1*x + y2*y == gy
               return x,y
            end
         end
      end

      return 0,0
   end

   def buildPuzzles(data)
      puzzles = []
      current = [0,0,0,0,0,0]
      data.split("\n").each do |line|
         if line.include?("Button A")
            elems = line.strip().split(",")
            v1 = elems[0].split("+")[1].to_i
            v2 = elems[1].split("+")[1].to_i
            current[0] = v1
            current[2] = v2
         elsif line.include?("Button B")
            elems = line.strip().split(",")
            v1 = elems[0].split("+")[1].to_i
            v2 = elems[1].split("+")[1].to_i
            current[1] = v1
            current[3] = v2
         elsif line.include?("Prize")
            elems = line.strip().split(",")
            v1 = elems[0].split("=")[1].to_i
            v2 = elems[1].split("=")[1].to_i
            current[4] = v1
            current[5] = v2
            puzzles.push(current)
            current = [0,0,0,0,0,0]
         end
      end

      return puzzles
   end

    def solvePart1(solve_req)
      sumv = 0
      
      puzzles = buildPuzzles(solve_req.data)
      puzzles.each do |puzzle|
         print "PUZZLE ", puzzle, "\n"
         x,y = resolve(puzzle[0],puzzle[1], puzzle[2], puzzle[3], puzzle[4], puzzle[5])
         sumv += x*3 + y
      end
      return sumv
    end
    def solvePart2(solve_req)
       return solvePart1(solve_req)
    end
end
