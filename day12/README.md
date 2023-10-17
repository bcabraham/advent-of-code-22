# Day 12: Hill Climbing Algorithm
----------------------------------------

You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter, where `a` is the lowest elevation, `b` is the next-lowest, and so on up to the highest elevation, `z`.

Also included on the heightmap are marks for your current position (`S`) and the location that should get the best signal (`E`). Your current position (`S`) has elevation `a`, and the location that should get the best signal (`E`) has elevation `z`.

You'd like to reach `E`, but to save energy, you should do it in _as few steps as possible_. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be _at most one higher_ than the elevation of your current square; that is, if your current elevation is `m`, you could step to elevation `n`, but not to elevation `o`. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

For example:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
```

Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the `e` at the bottom. From there, you can spiral around to the goal:

```
v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^
```

In the above diagram, the symbols indicate whether the path exits each square moving up (`^`), down (`v`), left (`<`), or right (`>`). The location that should get the best signal is still `E`, and `.` marks unvisited squares.

This path reaches the goal in `_31_` steps, the fewest possible.

_What is the fewest steps required to move from your current position to the location that should get the best signal?_
Answer: 408

## Research
This looks like aperfect use case for the A* Pathfinding Algorithm.
I found this video to use as a refresher for implementing it: https://www.youtube.com/watch?v=-L-WgKMFuhE

![A* Algorithm Pseudocode](a-star-pseudo.png)

## --- Part Two ---

As you walk up the hill, you suspect that the Elves will want to turn this into a hiking trail. The beginning isn't very scenic, though; perhaps you can find a better starting point.

To maximize exercise while hiking, the trail should start as low as possible: elevation a. The goal is still the square marked E. However, the trail should still be direct, taking the fewest steps to reach its goal. So, you'll need to find the shortest path from any square at elevation a to the square marked E.

Again consider the example from above:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
```

Now, there are six choices for starting position (five marked a, plus the square marked S that counts as being at elevation a). If you start at the bottom-left square, you can reach the goal most quickly:

```
...v<<<<
...vv<<^
...v>E^^
.>v>>>^^
>^>>>>>^
```

This path reaches the goal in only 29 steps, the fewest possible.

What is the fewest steps required to move starting from any square with elevation a to the location that should get the best signal?


```
a-b-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-c-c-c-c-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-
a-b-a-a-a-a-a-a-c-c-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-c a-a-a-a-c-a-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-
a-b-a-a-a-a-a-a-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-c-a-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-
a-b-a-a-a-a-a-a-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-a-c-c-a-a-a-c a-a-a-c-c-c-c-c-c-c-c-c*c*c*c*c*c*c-c-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-
a-b-c-c-a-a-a-c-c-a-a-a-a-a-a-c-c-c-a-a-a-a-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-c-c-c-c-a-a-c a-a-a-c-c-c-c-c-c-c-c-c*a-a-c-a-c*c-c-c-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-
a-b-c-c-c-c-c-c-c-a-a-a-a-a-c-c-c-c-a-a-a-a-c-c-c-c-c-a-a-a-a-a-c-c-c-a-a-a-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c*a-a-a-a-c*c*c*c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-c-
a-b-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-c-a-a-a-a-c-c-c-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c*c*k-l-l-l-l-l-c*c*c*c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-
a-b-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c*c*k-k-l-l*l*l*l-l-l-c*c*c-c-c-d-d-c-c-c-c-a-a-c-c-c-c-c-c-
a-b-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c*k-k-k*k*l*s-l*l*l*l-l-c*c*c*d-d-d-d-d-d-a-a-a-c-c-c-c-c-c-
a-b-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-a-a-a-a-a-a-a-a-c-c-c-c-c-c*c*k-k*k*s-s-s-s-s-l*l*l*l-l-c*d*d*d*d*d*d-d-d-a-c-c-c-c-c-c-
a-b-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-c-c-a-a-a-c-a-c-c-c-c-c-c-c*k-k-k*s-s-s-s*s*s*s-s-l*l*l-m-m-m-m-m-d*d*d-d-d-a-a-c-c-c-c-
a-b-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-c-c-a-a-c*c*c*c*c*c*c*c*c*c*k-k*k*s-s*s*s*u-s*s*s-s-l*m*m*m-m-m-m-m-d*d*d*d-d-a-c-c-c-c-
a-b-c-c-c-c-c-c-c-a-a-c-c-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-c-c-c-c-c-c*a-a-a-a-a-c-c-c-k-k-k*r-s*s*u-u-u-u-s*s*s-s-q-m*m*m*m*m*m-m-m-d*d-d-c-c-c-c-c-
a-b-c-c-c-c-c-c-c-a-a-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-a-a-c-c-c-c-a-a-a-a-a-c-c-c-c*c*a-a-a-a-a-c-c-k-k-k*k*r-r*u-u-u*u*u*u-s*s*q-q-q-q-q-q-m*m*m-m-d*d*d-c-c-c-c-c-
a-b-c-c-c-c-a-a-a-a-a-a-a-a-c-a-a-a-a-a-a-c-c-c-c-c-c-c-a-a-a-a-c-c-a-a-c-c-a-c-c-c-c*a-a-a-a-a-a-c-j-k*k*k*r-r-r*u-u-u*x u*u*u-s*q*q*q*q*q*q*q-m*m*m-m-e*e-c-c-c-c-c-
a-b-c-a-a-a-a-a-a-a-a-a-a-a-c-a-a-a-a-a-c-c-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-a-a-c-c-c-c*a-a-a-a-a-j-j-j*j*r-r-r-r*r*u-u*u*x x u*v*v-v-v-v-v-v-q*q-q-m*m-m-e*e-c-c-c-c-c-
a-b-c-a-a-c c c a-a-a-a-c-c-c-c-a-a-a-a-a-a-a-c-c-c-a-a-a-a-a-c-c-c-a-c-a-a-a-c-c-c-c*a-a-a-a-j-j-j*j*r-r-r*r*r*u-u-u*x x x x v*v*v*v*v*v-v-q*q-q-m*m-e*e*e-c-c-c-c-c-
a-b-a-a-a-a-c c a-a-a-a-a-c-c-c-c-c-c-c-a-a-a-c-c-c-c-a-a-a-a-a-c-a-a-a-a-a-a-a-a-c-c*c-a-a-j-j-j*j*r-r*r*r*t-u-u*u*u*x x x y v-y y y v*v-v-q*q-q-n*n-e*e-e-c-c-c-c-c-
a-b-a-a-a-a-a-a-a-a-a-a-a-c-c-c-a-a-a-a-a-a-a-c-c-c-c-a-a-c-a-a-c-a-a-a-a-a-a-a-a-c-c*c-c-c-j-j-j*r-r-r*t-t-t*t*u*x x x O x y y y y y v*v-v-q*q-n*n*n-e*e-e-c-c-c-c-c-
a-b-a-a-a-a-a-a-a-c-c-a-a-c-c-c-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-c-c-c*c-c-c-j-j-j*r-r-r*t-t-t*x x O x O x-O y y y y y v*v-v-q*q-n*n-n-e*e-e-c-c-c-c-c-
S-b-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-a-a-a-a-a-c-c-c-c*c-c-c-j-j-j*r-r-r*t-t-t*x O x-E*z*z*z-O y y y v-v*r-r*r*n-n*n-e*e*e-c-c-c-c-c-c-
a-b-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-c*c*c*c-c-c-j-j-j*q-q-q*t-t-t*O x-x-x-x*y*y-y-O y v-v-v*r-r*r-n-n*n-e*e-e-c-c-c-c-c-c-
a-b-a-a-a-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-c c a-a-c-c-c-c-c-c-c-c-c-c-a-a-c-c*a-a-a-a-a-j-j-j*q-q-q*t-t-t*t*x-x-x-x*y-y-y-y-O y v-v*r-r*r-n-n*n-e*e-e-c-c-c-c-c-c-
a-b-a-a-a-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-c c a-a-a-c-a-a-a-c-c-c-c-c-c-c-c*c*a-a-a-a-a-j-j-j*j*q-q*q*t-t-t*t*t*x-x*y-y-w-O y y y w*v-r*r-n-n*n-f*e-e-c-c-c-c-c-c-
a-b-c-a-a-a-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-c a-a-a-a-a-a-a-c-c-c-c-c-c-c-c*a-a-a-a-a-a-c-i-i-i*i*q-q*q*q*t-t-t*x-w*y-y-w-w-y y w*w*w-r*r-r-n*n-f*f-f-c-c-c-c-c-c-
a-b-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-c c a-a-a-a-a-a-c-c-c-c-c-c-c-c-c*a-a-a-a-a-a-c-c-i-i-i*i*q-q-q*q*t-t*w-w*y-w*w*w*w*w*w*w-w-r*r-r-n*n-f*f-f-c-c-c-c-c-c-
a-b-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-c-c-c-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c*a-a-a-a-a-a-c-c-c-i-i-i*i*q-q-q*t-t*w-w*w*w*w-s-w-w-w-w-r*r*r-r-n*n-f*f-f-c-c-c-c-c-c-
a-b-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-c-c-c-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-c*c-a-a-a-c-c-c-c-c-c-i-i-i*q-q-q*t-s*w-w-w-w-s*s*s*s*r*r*r*r-r-r-o*o-f*f-f-c-c-c-c-c-c-
a-b-c-c-c-c-c-c-a-a-a-a-a-c-a-a-a-a-a-a-c-c-c-a-a-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c*c-c-c-c-c-c-c-c-c-c-i-i-i*q-q-q*s-s*s*w-s*s*s*s-s-s-s-r-r-r-r-o*o*o-f*f-f-a-c-c-c-c-c-
a-b-c-c-c-c-c-c-a-a-a-a-a-c-a-a-c-c-a-a-c-c-c-c-c-c-a-a-a-c-a-a-a-c-c-c-c-c-c-c*c-c-c-c-c-c-c-c-c-c-c-i-i*i*q-q*s-s-s*s*s*s-s-p-o-o-r-r-r-o*o*o*o-o-f*f-f-a-a-c-c-c-c-
a-b-c-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-c-c-c-c-c-c-c-c*c-c-c-c-c-c-c-c-c-c-c-i-i-i*q-p*p*s-s-s-s-s-p*p*p*o*o*o*o*o*o-o-o-f*f*f-f-a-a-c-c-c-c-
a-b-c-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-c-c-c-c*c-c-c-c-c-c-c-c-c-c-c-c-i-i*p-p-p*p*p*p*p*p*p*p-p-o-o-o-o-o-o-o-f*f*f-f-a-a-c-c-c-c-c-
a-b-c-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c*c*c*c*c*c*c*c-c-c-c-c-c-c-c-c-c-c-c-i-i*h*p-p-p-p-p-p-p-p-g*g*g*g*g*g*g*g*g*f*f-f-a-a-a-c-c-c-c-c-
a*b*c*c*c*c*c*c*a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c*c*c*c*c*c*c*c*a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-h-h-h*h*p-p-p-p-p-p-h*g*g-g-g-g-g-g-g-g-g-f-a-a-a-a-c-c-c-c-c-
a b-a-a-a-c-c-c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*c*a-c-c-c-c-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-h-h-h*h*h*h*h*h*h*h*g-g-g-g-g-g-g-g-g-c-a-a-c-c-c-c-c-c-c-c-
a-b-a-a-c-c-a-a-a-c-c-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-a-a-a-c-a-a-c-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-h-h-h-h-h-h-h-h-h-g-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-
a-b-a-a-a-c-a-a-a-c-a-a-c-c-c-c-c-c-c-c-c-a-a-a-c-c-a-a-a-a-c-a-a-a-a-a-a-a-a-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c-h-h-h-h-h-h-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-
a-b-a-a-a-c-c-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-c-a-a-a-a-a-a-a-a-c a-a-a-a-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-a-c-c-c-a-
a-b-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-c a-a-a-a-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-c-c-c-c-a-a-a-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-a-
a-b-c-c-c-a-a-a-a-a-a-a-a-c-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-c-c-c-c-a-a-a-a-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-
a-b-c-c-c-a-a-c a-a-a-a-a-c-c-c-c-c-c-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-a-c-c-c-c-c-a-a-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-c-a-a-a-a-a-
```
Shortest Path: 399
