# Tendermint Challenge
## Task
### Introduction
You are given a map containing the names of cities in the non-existent world of X. The map is in a file,
with one city per line. The city name is first, followed by 1-4 directions (north, south, east, or west). Each
one represents a road to another city that lies in that direction.

For example:
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```
The city and each of the pairs are separated by a single space, and the directions are separated from
their respective cities with an equals (=) sign.

You should create N aliens, where N is specified as a command-line argument.
These aliens start out at random places on the map, and wander around randomly, following links. Each
iteration, the aliens can travel in any of the directions leading out of a city. In our example above, an
alien that starts at Foo can go north to Bar, west to Baz, or south to Qu-ux.

When two aliens end up in the same place, they fight, and in the process kill each other and destroy the
city. When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of
it.

In our example above, if Bar were destroyed the map would now be something like:
```
Foo west=Baz south=Qu-ux
```
Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens getting
"trapped".

### Task
You should create a program that reads in the world map, creates N aliens, and unleashes them. The
program should run until all the aliens have been destroyed, or each alien has moved at least 10,000
times. When two aliens fight, print out a message like:
```
> Bar has been destroyed by alien 10 and alien 34!
```
(If you want to give them names, you may, but it is not required.) Once the program has finished, it
should print out whatever is left of the world in the same format as the input file.

Feel free to make assumptions (for example, that the city names will never contain numeric characters),
but please add comments or assertions describing the assumptions you are making.
We are evaluating your skills to write production level code so please provide tests documentation.

## Assumptions
1. If given aliens number (M) is greater than the cities number (N), then only N aliens will be created in the world.
2. If alien is trapped in some city (can't go anywhere) each iteration increase the number of it's steps.
3. City names can only contain `A-Z`, `a-z`, `0-9`, `-`. The name must begin with a letter and cannot end with a dash.
4. When a map is being loaded, city pairs are created together (north-south, west-east). If there is no line for the
city in the input file, but it is referenced in some direction, then this city will be present in the map.

## How to run
1. Generate a map of size 20x30 (use your values):
```
go run ./testdata -n 20 -m 30 -path ./testdata/map.txt
```
2. Run the world with 500 aliens:
```
go run . -n 500 -path testdata/map.txt
```
