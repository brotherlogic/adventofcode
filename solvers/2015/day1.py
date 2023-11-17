def SolveDay1Part1(instructions):
    cFloor = 0
    for c in instructions:
        if c == '(':
            cFloor+=1
        elif c == ')':
            cFloor-=1
    return cFloor

def SolveDay1Part2(instructions):
    cFloor = 0
    count = 0
    for c in instructions:
        count+=1
        if c == '(':
            cFloor+=1
        elif c == ')':
            cFloor-=1
        if cFloor == -1:
            return count