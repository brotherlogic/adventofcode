# November

We need to prep in November to be ready to roll by December. That means that
come December 1st and the first puzzle, we're getting the issues created and
all we need to do is write the problem solving code.

This compresses down to four milestones, corresponding to the four weeks we have
available to run things.

## Week 1

Week 1 we set out the framework for building out the scaffolding for supporting
problem solving. Week 1 triggers a bug highlighting this "Scaffolding for YEAR"
required. This bug is closed once the solver service is able to respond to a
Solve request for the given year (with no day or month params).

## Week 2

Once we've established the scaffolding, we need to ensure that a dummy puzzle
can be solved. We file a puzzle bug but for Day 0 Part 1 - this is a simple puzzle
whose solution is any integer number greater than zero. The solution server should
be able to solve this internally with a test case. We trigger this with a single bug
to support Day 0 problem for YEAR, and close it out with a github push with passing
tests.

## Week 3

In week 3 we expand on the solver in week 2 and have the server request the solution
and confirm that this works - the solver triggers a solution bug that needs to be solved.
The solution is injected when the bug triggers.

## Week 4

In week 4 we trigger a bug to validate that the server components are working and running
and that everything is looking good.

## Day 1 December

In day 1 we need to ensure that bugs are filled and that we can respond to them correctly,
most likely requiring some code fixes along the way.
