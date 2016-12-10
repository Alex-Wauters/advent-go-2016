package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var handover, inputbin *regexp.Regexp
var answer chan int
var bins map[int]int

func main() {
	part1()
	fmt.Printf("Solution to Part 2: %v\n", bins[0]*bins[1]*bins[2])
}

func part1() {
	defer track(time.Now(), "part 1")
	answer = make(chan int)
	commands := make(chan command, 50)
	go readInput(commands)
	go dispatcher(commands)
	droidId := <-answer
	fmt.Printf("The droid you're looking for is %v \n", droidId)
}

type command struct {
	valueCmd, destType        string
	source, destNumber, value int
}

type droid struct {
	input    chan int
	commands chan command
	id       int
	output   chan command
}

//A balance bot can receive up to two values and then pass them on
func bot(droid *droid) {
	val1, val2 := <-droid.input, <-droid.input
	if val2 < val1 {
		val1, val2 = val2, val1
	}
	if answerFound(val1, val2) {
		answer <- droid.id
	}
	cmd1, cmd2 := <-droid.commands, <-droid.commands
	if cmd1.valueCmd == "low" {
		cmd1.value, cmd2.value = val1, val2
	} else {
		cmd1.value, cmd2.value = val2, val1
	}
	cmd1.valueCmd, cmd2.valueCmd = "", ""
	droid.output <- cmd1
	droid.output <- cmd2
}

//Dispatcher sends commands from the input to bots or values (received from drones) to drones/output bins
func dispatcher(commands chan command) {
	droids := make(map[int]droid)
	bins = make(map[int]int)
	for cmd := range commands {
		if cmd.valueCmd != "" { // Send a command to a bot
			bot := getBot(droids, cmd.source, commands)
			bot.commands <- cmd
		} else { // Send a value to a bot
			if cmd.destType == "bot" {
				bot := getBot(droids, cmd.destNumber, commands)
				bot.input <- cmd.value
			} else {
				bins[cmd.destNumber] = cmd.value
			}
		}
	}
}

//Gets the references to a bot if it is already launched
//If the bot hasn't been launched yet, it will be launched
func getBot(droids map[int]droid, id int, output chan command) *droid {
	drone, exists := droids[id]
	if !exists {
		drone = droid{make(chan int, 2), make(chan command, 2), id, output}
		droids[id] = drone
		go bot(&drone) //Launch
	}
	return &drone
}

func answerFound(k, m int) bool {
	return k == 17 && m == 61
}

//Read the input, create one or more commands from each input line
func readInput(commands chan command) {
	handover = regexp.MustCompile(`bot (\d+) gives (low|high) to (bot|output) (\d+) and (low|high) to (bot|output) (\d+)`)
	inputbin = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sendCommands(scanner.Text(), commands)
	}
}

//Send commands to the dispatcher
func sendCommands(line string, commands chan command) {
	cmd := handover.FindStringSubmatch(line) //1=botid 2=loworhigh 3=botoroutput 4=destid 5=loworhighh 6=botoroutput 7=destid2
	if len(cmd) > 0 {
		commands <- command{source: toInt(cmd[1]), valueCmd: cmd[2], destType: cmd[3], destNumber: toInt(cmd[4])}
		commands <- command{source: toInt(cmd[1]), valueCmd: cmd[5], destType: cmd[6], destNumber: toInt(cmd[7])}
	} else if cmd = inputbin.FindStringSubmatch(line); len(cmd) > 0 { //1=value 2=botid
		commands <- command{value: toInt(cmd[1]), destNumber: toInt(cmd[2]), destType: "bot"}
	} else {
		panic("Could not parse: " + line)
	}
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
