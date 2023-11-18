def solve21(num1, num2, num3):
    if num1 > num2 and num1 > num3:
        return num2*2+num3*2 + num1*num2*num3
    elif num2 > num1 and num2 > num3:
        return num1*2+num3*2 + num1*num2*num3
    else:
        return num1*2+num2*2 + num1*num2*num3

def SolveDay2Part1(instructions):
    sum = 0
    for line in instructions.splitlines():
        pieces = line.split("x")
        num1 = int(pieces[0])
        num2 = int(pieces[1])
        num3 = int(pieces[2])
        sum += solve21(num1,num2,num3)
    return sum

def SolveDay2Part2(instructions):
    return -1