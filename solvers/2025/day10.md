# Day 10 Part 2

This is equations

For the irst example
S
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
=======

So we can just solve these using a linear solver

For the second

7 = 1 3 4
5 = 4 5
12 = 1 2 4 5
7 = 1 2 5
2 = 1 3 5

From z3:

 (define-fun s1 () Int
    2)
  (define-fun s3 () Int
    1)
  (define-fun s2 () Int
    5)
  (define-fun s5 () Int
    3)

s1: 3
s2: 1,3
s3: 2
s4: 2,3
s5: 0,2
s6: 0,1

0 + 0 + 0 = 3  
1 + 1 + 1 + 1 + 1 = 5
2 + 2 + 2 + 2 + 4
3 + 3 + 3 + 3 + 3 + 3 + 3 = 7

(
  (define-fun s6 () Int
    0)
  (define-fun s4 () Int
    1)
  (define-fun s3 () Int
    0)
  (define-fun s1 () Int
    1)
  (define-fun s2 () Int
    5)
  (define-fun s5 () Int
    3)
)
0 + 0 + 0
1 + 1 + 1 + 1 + 1
2 + 2 + 2 + 2
3 + 3  + 3 + 3 + 3 + 3 + 3

11 total pushes

I cheated and used z3