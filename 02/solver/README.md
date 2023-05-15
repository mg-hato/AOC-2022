# Rock-Paper-Scissors enumerations and transformations

Here we document the interesting bits about solving this problem of Advent of Code 2022, day 2.
This mainly focuses on weird modulo-enums transformations.

## Enumerations

Here we define values of our enums (for the sake of reference) and their respective enum-index values.

### Left symbol:

- A : 0
- B : 1
- C : 2

### Right symbol

- X : 0
- Y : 1
- Z : 2

### Shapes

- Rock : 0
- Paper : 1
- Scissors : 2

### Game outcomes

- Lose : 0
- Draw : 1
- Win : 2

### Notation

We will write `int(E)` to represent the enum-index of enum `E` e.g. `int(Paper) = 1`.

Similarly the other way around, for enum group `EG` we write `EG.val(v)` to denote the enum of group `EG` whose enum-index has value `v` e.g. `GameOutcome.val(1) = Draw`

Lastly, we use operator `%` as a modulo-operator that yields lowest non-negative values equivalents. E.g. `-1 % 3 = 2`

## Transformations

Here we elaborate the transformations we apply.

### Trivial transformations

There are 4 trivial transformations.
 
 1. Left symbols A ,B and C to shapes: Rock, Paper and Scissors, respectively (same enum-indices)
 
 2. Same as 1. but with Right symbols X, Y and Z (again, same enum-indices)

 3. Right symbols X, Y and Z to desired game outcomes: Lose, Draw and Win, respectively (same enum-indices).
 
 4. Game outcome to the outcome score: As Lose, Draw and Win are worth 0, 3 and 6 points, respectively, the transformation is to take the outcome's enum-index and multiply it with 3.

 ### Game outcome transformation

 Given the opponent's shape `OP` and the player's shape `P` we need to determine the game's outcome from player's perspective. We infer this by looking at the "modulo" difference between shape's enum-indices, concretely:
 
 `(int(P) - int(OP)) % 3`
 
 Possible values of this expression are 0, 1 or 2.
 
 1. For draw, the mod-difference is 0 (same shapes, same enum-indices).
 2. For win, the mod-difference is 1.
 E.g. `P = Rock; OP = Scissors`. Expression becomes

 `(int(Rock) - int(Scissors)) % 3 = (0 - 2) % 3 = -2 % 3 = 1`

 3. For lose, the mod-difference is 2.

 Since the expression values are skewed comparing to their corresponding outcomes' enum-index  (e.g. index of Win is 2, but the mod-difference for Win is 1 etc.), we unskew them by adding 1 w.r.t. modulo 3. Hence the transformation:

 `getGameOutcome(P, OP) = GameOutcome.val((int(P) - int(OP) + 1) % 3)`

 ### Player's shape from desired outcome transformation

 Given the game's outcome `G` and the opponent's shape `OP` and we need to infer what is the player's shape that would yield the outcome `G` for the player against the opponent's shape `OP`. The transformation is:

 `getDesiredShapeForPlayer(G, OP) = Shape.val((int(G) + int(OP) - 1) % 3)`

### Shape score transformation

Finally, a transformation that is also almost trivial. Given a player's shape `P` the chosen shape yields some points regardless of the game's outcome. The mapping is that Rock, Paper and Scissors are worth 1, 2 and 3 points, respectively. The transformation is obvious:

`getShapeScore(P) = int(P) + 1`