require 'logger'
require 'date'

class Computer
    @A = 0
    @B = 0
    @C = 0
    @output = []

    def initialize
        @A = 0
        @B = 0
        @C = 0
        @output = []
    end

    def printC()
        print @A, ",", @B, ",", @C, "\n"
    end

    def getCombo(combo)
        if combo < 4
            return combo
        elsif combo == 4
            return @A
        elsif combo == 5
            return @B
        elsif combo == 6
            return @C
        end
    end

    def runInstruction(opcode, literal, pointer)
        combo = getCombo(literal)
        print "RUN ", opcode, " -> ", combo, "\n"
        if opcode == 0
            numerator = @A
            denominator = 2**combo
            @A = numerator / denominator

            return pointer + 2
        elsif opcode == 1
            @B = @B ^ literal

            return pointer + 2
        elsif opcode == 2
            @B = combo %8
            return pointer + 2
        elsif opcode == 3
            if @A == 0
                return pointer + 2
            end

            return literal
        elsif opcode == 4
            @B = @B ^ @C

            return pointer + 2
        elsif opcode == 5
            @output.push(combo%8)

            return pointer + 2
        elsif opcode == 6
            numerator = @A
            denominator = 2**combo
            @B = numerator / denominator

            return pointer + 2
        elsif opcode == 7
            numerator = @A
            denominator = 2**combo
            @C = numerator / denominator

            return pointer + 2
        end
        return 0
    end

    def setA(val)
        @A = val
    end
    def setB(val)
        @B = val
    end
    def setC(val)
        @C = val
    end
    

    def getOutput()
        return @output
    end
end

class Day17

    def runComputer(data)
        c = Computer.new
        program = ""
        data.split("\n").each do |line|
            pieces = line.strip().split()
            if pieces[0] == "Register"
                if pieces[1] == "A:"
                    c.setA(pieces[2].to_i)
                elsif pieces[2] == "B:"
                    c.setB(pieces[2].to_i)
                elsif pieces[3] == "C:"
                    c.setC(pieces[3].to_i)
                end
            elsif pieces[0] == "Program:"
                program = pieces[1].split(",")
            end
        end

        pointer = 0
        while pointer < program.length()
            pointer = c.runInstruction(program[pointer].to_i, program[pointer+1].to_i, pointer)
            c.printC()
            print "POINTER ", pointer, "\n"
        end
      
        return c.getOutput()
    end

    def solvePart1(solve_req)
        solution = runComputer(solve_req.data)
        print "SOLUTION ", solution, "\n"
        return solution.join(",")
    end

    def solvePart2(solve_req)
        return solvePart1(solve_req)
    end
end