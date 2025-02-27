# Description
Lemin finds the optimal route combination for ants to move from start to finish. Returns the moves in a format of a line per turn.    
`L1-roomName` L is followed by the ant name and - is followed by the room it moves to. 

# Usage
It accepts input in the format of a txt file located in `./examples` folder. Do NOT navigate to examples simply prompt:
`$letmein example01.txt`
#### Output

```
10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n

L1-0 L2-h L3-t
L1-o L2-A L3-E L4-0 L5-h L6-t
L1-n L2-c L3-a L4-o L5-A L6-E L7-0 L8-h L9-t
L1-e L2-k L3-m L4-n L5-c L6-a L7-o L8-A L9-E L10-0
L1-end L2-end L3-end L4-e L5-k L6-m L7-n L8-c L9-a L10-o
L4-end L5-end L6-end L7-e L8-k L9-m L10-n
L7-end L8-end L9-end L10-e
L10-end

Total turns: 8
```
# Instalation
### Build bin:   
`make build`   
### Run all examples
`make run`
### Test
`make test`
### Visualize
`make vis f="<file_name>"`
### Remove bin and output.txt:   
`make clean`

# Modules
### Lemin
Overlooks the process and returns results and errors to main call
### Handle
Handle module checks txt file for errors and if none found it creates an instance of struct Colony
### Colony
Colony provides basic functions to construct colony data base
### Paths
Paths provides functions to analize the colony dB and finds the optimal routes for ants to move
### Moves
Moves provides functions to move the ants through the provided paths by paths
# Dependencies
go 1.22.0 
# Authors
- Vagelis Stefanopoulos
- Sotiris Masteas   
- Alexandros Zachos   