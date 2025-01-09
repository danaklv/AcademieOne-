### Description
This project is meant to create a digital simulation of an ant farm, focusing on optimizing the paths that ants take through a network of rooms and tunnels. Your program, lem-in, will `read data from a file` that describes the ants and the colony structure.

The objective is to `find the quickest path for n ants` to travel from the start room (##start) to the end room (##end) within a colony of rooms and tunnels.


### How it Works

***Input Format***: lem-in reads from a file describing the ants and the colony. The file format is as follows:
***Number of ants***: The first line of the file contains the number of ants.
***Room definitions***: Subsequent lines define the rooms in the colony with their coordinates (e.g., "room_name x y"). The start and end rooms are designated by ##start and ##end respectively, followed by the room definition.
***Links***: Lines containing hyphen-separated room names (e.g., "room1-room2") define the links (tunnels) between rooms.
***Simulation***: Given the number of ants and the colony layout, the program will simulate the ants' movements from ##start to ##end while avoiding traffic jams and optimizing the path.
***Output***: The program will print the content of the input file followed by the sequence of movements the ants make from room to room to achieve the quickest possible route.


### Use:

`go run . filename.txt`

## Author

***dkalykov***