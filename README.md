# Advent of code 2021

In Go.

## Days

### Day 1

Part 1 - complete. Straightforward, basically just iterating around a list of ints checking if
current int is larger than previous.

Part 2 - complete. More complicated, had to some previous 3 days. Good chance to get a bit more
familiar with slices in Go.

### Day 2

Part 1 - complete. Straightforward. Again just iterating and adding numbers.

Part 2 - complete. Very minor change to part 1.

### Day 3

Part 1 - complete. Definitely harder and took me a while to get going on this.

Part 2 - complete. Misread the instructions so spent a lot of time with the wrong answer. Initially,
I had thought I was supposed to use the most common character across the entire set rather than
work out the most common character in just the current filtered list. Once I understood that it was
reasonably straightforward to adapt the first part.

### Day 4

Part 1 - complete. Difficulty still ramping up though probably the worst bit was manipulating the
slices. This would definitely be a lot easier in a language with a bigger standard library that 
had better collections types. Spent a lot of time writing code to manually filter slices when in 
some other languages I'd have been able to do something like `.remove`. I'm also creating a lot
of pointless object - doubt this is very idiomatic (or at least I hope not).

Part 2 - complete. Very minor variation on part 1 so quite quick.

### Day 5

Part 1 - complete. Found this a bit easier, probably cause it was quite similar to day 4.

Part 2 - complete. Struggled more with this than I should have. Turns out I really don't understand
diagonals. Probably should've written it down on a piece of paper, in retrospect.

### Day 6

Part 1 - complete. Quite easy, did it in what I thought was a nice object oriented style. Also wrote
a test first for the sample data. First ever go test!

Part 2 - complete. Turns out creating an object per fish didn't scale! Then tried to pre-calculate
the total per starting age of fish and my computer could even do that. Switched it all around to
calculate the number of fish per day and just shift values in arrays and it turns out it can pretty
much be done instantly with zero memory. Still preferred the first approach...

### Day 7

Part 1 - complete. I feel there's probably a smarter solution than this but since this one is
achievable in sensible time I've just gone for calculating the distances and adding them.

Part 2 - complete. Fairly straightforward variation of first part.

### Day 8

Part 1 - complete. This was very easy once I understood the question.

Part 2 - incomplete. Done some work on this but haven't quite thought about the logic enough yet.

### Day 9

Part 1 - complete. Pretty straightforward.

Part 2 - complete. Had to think about this more and ended up writing recursive code for the first
time in a long time. In retrospect this could be more efficient if I'd just started from the low
points rather than testing everything.

### Day 10

Part 1 - complete. Quite enjoyable because it was obvious what to do and didn't feel like you needed
to know some sort of trick like the previous few days.

Part 2 - complete. Reasonably straightforward variation on part 1. Should probably tidy up the code
but it's Friday night!
