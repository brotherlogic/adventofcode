# Day 10

## Optimal solution

We have some linear equations to solve:

3 = s5 + s6
5 = s2 + s6
4 = s3 + s4 + s5
7 = s1 + s2 + s4

where s1+s2+s3+s4+s5+s6 is minimal

the overall number of solutions may be high

Two approaches:

1. Get close to the answer and do a deep search to find the exact
   solution
2. Build all possible solutions, find the minimal of these?

Slots approach:

7 0 0
7 1 0
7 0 1

this is too many to effectively search - anything with more than
two options becomes immediately prohibitive

Brute force is out

## Reduction approach

3,5,4,7

O,O,E,O

(3) (1,3) (2) (2,3) (0,2) (0,1)

E, E, E, E

2, 4, 4, 6

reduces to

1, 2, 2, 3
