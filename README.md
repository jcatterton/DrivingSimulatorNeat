# DrivingSimulatorNeat
A 2D top-down driving simulator which learns to traverse a track using NEAT. This is simply the program to run the 
simulator, not the NEAT library.

To run, simply use `go run main.go` in terminal.

As it stands, the cars move forward automatically, and the NEAT lib is used only to determine when they should turn left 
or right. Fitness is simply an additive function which increases over time.

The cars see using sight vectors, which are lines radiating away from the car which check for intersection with walls.
The length of the sight vector thus indicates the distance from the car to the nearest wall in a given direction. Only
three sight vectors are used. One directly in front of the car, one facing forward and 45 degrees to the left, and one 
facing forward and 45 degrees to the right. The length of each sight vector is given to the NEAT library, and the output
of the car's corresponding genome is used to determine when the car should turn left or right.

Display Options:
* Enable/Disable Sight Vectors - Click 1
* Enable/Disable genome drawing - Click 2
* Stop Current Generation - Click X

Expected Performance:

It is anticipated that there will be several generations with little to no progress due to the current settings of the
NEAT library. Several cars should complete the first right turn after a few (varies) generations. The first left turn
takes considerably longer. This is to be expected as until this point, cars which favor turning right have survived much
better than cars which favor turning left. Eventually, some cars will pass this point. On average it takes between 25
and 60 generations before the first car completes the track. A larger population (set where `InitPop(i, o, g)` is called)
can improve the amount of generations it takes to complete the track, but will also be much more demanding on your 
system.
