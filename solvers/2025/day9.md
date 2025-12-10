# Process

Look at each line in the rectangle, and see if the line
intersects the edge of the polygon if there's an intersection,
the shape is out. Otherwise it's either in or out - so just look
at the middle of the rectangle - if that's inside then we're fine.

## Line Crossing

Where does a vertical line cross a horizontal one?

L1: x1,y -> x2, y  (H)
L2: x, y1 -> x, y2 (V)

The vertical line stays at x between y1 and y2

The horizontal line stays at y between x1 and x2

if x is between x1 and x2 and y is between y1 and y2, they cross explicitly.

An exception is if the lines cross where they start so if
x,y = one of the points in L1 and one of the points in L2 they do
not intersect
