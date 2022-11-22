Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers.

Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks.

Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.

Design a discipline of behaviour (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.

The philosophers' ids are numbered from 0 to 4 in a clockwise order. Implement the function `void wantsToEat(philosopher, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)` where:

* `philosopher` is the id of the philosopher who wants to eat.
* `pickLeftFork` and `pickRightFork` are functions you can call to pick the corresponding forks of that philosopher.
* `eat` is a function you can call to let the philosopher eat once he has picked both forks.
* `putLeftFork` and `putRightFork` are functions you can call to put down the corresponding forks of that philosopher.
* The philosophers are assumed to be thinking as long as they are not asking to eat (the function is not being called with their number).
Five threads, each representing a philosopher, will simultaneously use one object of your class to simulate the process. The function may be called for the same philosopher more than once, even before the last call ends.



```
*Input*: n = 1
*Output*: [
    [4,2,1],[4,1,1],[0,1,1],[2,2,1],[2,1,1],
    [2,0,3],[2,1,2],[2,2,2],[4,0,3],[4,1,2],
    [0,2,1],[4,2,2],[3,2,1],[3,1,1],[0,0,3],
    [0,1,2],[0,2,2],[1,2,1],[1,1,1],[3,0,3],
    [3,1,2],[3,2,2],[1,0,3],[1,1,2],[1,2,2]
]
*Explanation* :
n is the number of times each philosopher will call the function.
The output array describes the calls you made to the functions controlling the forks and the eat function, its format is:
output[i] = [a, b, c] (three integers)
- a is the id of a philosopher.
- b specifies the fork: {1 : left, 2 : right}.
- c specifies the operation: {1 : pick, 2 : put, 3 : eat}.
```

*Solution in CSP*

The behavior of the ith philosopher may be described as follows:

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

The fate of the ith fork is to be picked up and put down
by a philosopher sitting on either side of it

```
FORK =
    *[
        phil(i)?pickup()-> phil(i)?putdown() |
        phil((i - 1)mod 5)?pickup() -> phil((i- l) mod 5)?putdown()
    ]

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

Notes: (1) The solution given above does not prevent all
five philosophers from entering the room, each picking
up his left fork, and starving to death because he cannot
pick up his right fork. (2) Exercise: Adapt the above
program to avert this sad possibility. Hint: Prevent more
than four philosophers from entering the room. (Solution
due to E. W. Dijkstra). 