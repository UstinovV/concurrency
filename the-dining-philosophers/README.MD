### The Dining Philosophers

Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers.

Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks.

Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.

![The dining philosophers](https://assets.leetcode.com/uploads/2019/09/24/an_illustration_of_the_dining_philosophers_problem.png)

Problem on [leetcode](https://leetcode.com/problems/the-dining-philosophers/)

The idea is to implement solution of this problem, described in CSP-language, taken from this [paper work](https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf), in Go

**Solution in CSP** 
---
The behavior of the i'th philosopher may be described as follows:

```
PHIL = *[
    ... during ith lifetime ... ->
    THINK;
    room!enter();
    fork(i)!pickup(); fork((i+ 1) mod 5)!pickup( );
    EAT;
    fork(i)!putdown(); fork((i+ 1) mod 5)!putdown( );
    room!exit()
]
```

The fate of the i'th fork is to be picked up and put down
by a philosopher sitting on either side of it

```
FORK =
    *[
        phil(i)?pickup()-> phil(i)?putdown() |
        phil((i - 1)mod 5)?pickup() -> phil((i- l) mod 5)?putdown()
    ]
```

The story of the room may be simply told:

```
ROOM = occupancy:integer; occupancy := 0;
*[
    (i:0..4)phil(i)?enter() -> occupancy := occupancy + 1 |
    (i:0..4)phil(i)?exit() -> occupancy := occupancy - 1
]
```

All these components operate in parallel:

```
[room::ROOM || [fork( i:0..4)::FORK || phil( i:0..4)::PHIL].
```

**Notes**: The solution given above does not prevent all five philosophers from entering the room, each picking up his left fork, and starving to death because he cannot
pick up his right fork.

**Exercise**: Adapt the above program to avert this sad possibility. Hint: Prevent more than four philosophers from entering the room. (Solution due to E. W. Dijkstra). 

---

### Notes to the solution

* Because in Go there no named processes, assumed that structs(Room, Fork, Philosopher) are represents those processes
* CSP theory does not imply buffered channels, but they used here to simplify logic
* This is not fastest or cleanest solution, this is just an attempt to implement the solution closest to the CSP-language