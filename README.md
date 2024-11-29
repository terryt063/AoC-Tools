# AoC-Tools
Tools for Advent of Code

This tool will grab the details for a specific day and year of the Advent of Code puzzle.

You will need to have been logged on, and will need to have a session token for the logon for this to work.

## Session Token

The easiest way to find this is log on to AoC https://adventofcode.com/ and then use developer tools to inspect the cookie. This long text field needs to be added to a .env file as below:-

```
SESSION_ID=..........................
```
The session tokens are incredibly long lived, so this should only need setting once. If the `SESSION_ID` isn't set, the program will crash.

## Usage

Pull the executable from the [releases page](https://github.com/terryt063/AoC-Tools/releases) and use the flags as below to get day 1 from 2023 puzzle set. 

```shell
./aoc-tools -d 1 -y 2023
```

Will grab the puzzle input and the challenge text converted from HTML to markdown. If the challenge part 1 has not been completed, the second challenge will not be in the `.md` file.

If you wish to try some older puzzles, it will take `all` as a day input and that will grab all the puzzles for that year.

```shell
./aoc-tools -d all -y 2023
```