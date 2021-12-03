package main;

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    // Load the file and split up the data
    dat, err := os.ReadFile("input");
    if err != nil {
        panic(err);
    }

    cmds := strings.Split(string(dat), "\n");

    // Part 1
    hor := 0;
    dep := 0;

    for i:=0; i < len(cmds) - 1; i++ {
        cmd := strings.Split(cmds[i], " ");
        val, _ := strconv.Atoi(cmd[1])
        switch {
        case cmd[0] == "forward":
            hor += val;
        case cmd[0] == "down":
            dep += val;
        case cmd[0] == "up":
            dep -= val;
        }
    }
    fmt.Println("Part 1:")
    fmt.Println(hor * dep)

    // Part 2
    hor = 0;
    dep = 0;
    aim := 0;

    for i:=0; i < len(cmds) - 1; i++ {
        cmd := strings.Split(cmds[i], " ");
        val, _ := strconv.Atoi(cmd[1])
        switch {
        case cmd[0] == "forward":
            hor += val;
            dep += aim * val;
        case cmd[0] == "down":
            aim += val;
        case cmd[0] == "up":
            aim -= val;
        }
    }
    fmt.Println("Part 2:")
    fmt.Println(hor * dep)


}
