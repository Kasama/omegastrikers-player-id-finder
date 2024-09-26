# Omegastrikers Player ID finder

A quick and dirty project to get the player ids from omegastrikers usernames.

## Usage

Have a file with the player names, one each line. `players.txt`

```txt
player1
player2
player3
```

then run this program:
```sh
go run main.go < players.txt > player-ids.txt
```

and the `player-ids.txt` file will contain the player ids of each player, or `Unknown` if the player was not found
