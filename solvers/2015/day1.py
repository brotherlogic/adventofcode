def SolveDay1Part1(instructions):
    cFloor = 0
    for c in instructions:
        if c == '(':
            cFloor+=1
        elif c == ')':
            cFloor-=1
    return cFloor

def SolveDay1Part2(instructions):
    return -1