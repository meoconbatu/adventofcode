package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"sync"
)

type point3D struct {
	x, y, z float64
}

var wg sync.WaitGroup

func day12Part1() {
	file, err := os.Open("input12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos, vel := day12Part1GetState(file, 1000)
	fmt.Println(calcEnegy(pos, vel))
}
func day12Part1GetState(in io.Reader, stepth int) ([]point3D, []point3D) {
	scanner := bufio.NewScanner(in)
	pos := make([]point3D, 4)
	vel := make([]point3D, 4)
	i := 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "<x=%f, y=%f, z=%f>", &pos[i].x, &pos[i].y, &pos[i].z)
		i++
	}

	for n := 0; n < stepth; n++ {
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if pos[i].x < pos[j].x {
					vel[i].x++
					vel[j].x--
				} else if pos[i].x > pos[j].x {
					vel[i].x--
					vel[j].x++
				}
				if pos[i].y < pos[j].y {
					vel[i].y++
					vel[j].y--
				} else if pos[i].y > pos[j].y {
					vel[i].y--
					vel[j].y++
				}
				if pos[i].z < pos[j].z {
					vel[i].z++
					vel[j].z--
				} else if pos[i].z > pos[j].z {
					vel[i].z--
					vel[j].z++
				}
			}
		}
		for i := 0; i < len(pos); i++ {
			pos[i].x += vel[i].x
			pos[i].y += vel[i].y
			pos[i].z += vel[i].z
		}
	}
	return pos, vel
}
func calcEnegy(pos []point3D, vel []point3D) float64 {
	total := 0.0
	for i := 0; i < len(pos); i++ {
		total += (math.Abs(pos[i].x) + math.Abs(pos[i].y) + math.Abs(pos[i].z)) *
			(math.Abs(vel[i].x) + math.Abs(vel[i].y) + math.Abs(vel[i].z))
	}
	return total
}

func day12Part2() {
	file, err := os.Open("input12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(day12Part2GetStepthThread(file))
}

func day12Part2GetStepth(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	pos := make([]point3D, 4)
	vel := make([]point3D, 4)

	i := 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "<x=%f, y=%f, z=%f>", &pos[i].x, &pos[i].y, &pos[i].z)
		i++
	}
	firstPos := make([]point3D, 4)
	copy(firstPos, pos)

	temps := make([]int, 4)
	done := 0
	n := 0
	for {
		n++
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if pos[i].x < pos[j].x {
					vel[i].x++
					vel[j].x--
				} else if pos[i].x > pos[j].x {
					vel[i].x--
					vel[j].x++
				}
				if pos[i].y < pos[j].y {
					vel[i].y++
					vel[j].y--
				} else if pos[i].y > pos[j].y {
					vel[i].y--
					vel[j].y++
				}
				if pos[i].z < pos[j].z {
					vel[i].z++
					vel[j].z--
				} else if pos[i].z > pos[j].z {
					vel[i].z--
					vel[j].z++
				}
			}
		}
		for i := 0; i < len(pos); i++ {
			pos[i].x += vel[i].x
			pos[i].y += vel[i].y
			pos[i].z += vel[i].z
		}
		for i := 0; i < len(temps); i++ {
			if (temps[i] == 0 && reflect.DeepEqual(pos[i], firstPos[i]) && vel[i] == point3D{}) {
				temps[i] = n
				fmt.Print(temps)
				done++
			}
		}
		if done == 4 {
			fmt.Println(temps)
			break
		}

	}
	return lcmArray(temps)
}

func day12Part2GetStepthThread(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	pos := make([]point3D, 4)
	vel := make([]point3D, 4)

	i := 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "<x=%f, y=%f, z=%f>", &pos[i].x, &pos[i].y, &pos[i].z)
		i++
	}

	firstPos := make([]point3D, 4)
	copy(firstPos, pos)

	out := make(chan int, 3)

	wg.Add(3)
	go x(pos, vel, firstPos, out)
	go y(pos, vel, firstPos, out)
	go z(pos, vel, firstPos, out)

	wg.Wait()

	return lcm(<-out, lcm(<-out, <-out))
}
func x(pos, vel, firstPos []point3D, out chan int) {
	n := 0
	for {
		n++
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if pos[i].x < pos[j].x {
					vel[i].x++
					vel[j].x--
				} else if pos[i].x > pos[j].x {
					vel[i].x--
					vel[j].x++
				}
			}
		}
		for i := 0; i < len(pos); i++ {
			pos[i].x += vel[i].x
		}
		d := 0
		for i := 0; i < len(pos); i++ {
			if pos[i].x == firstPos[i].x && vel[i].x == 0 {
				d++
			}
		}
		if d == 4 {
			out <- n
			wg.Done()
			return
		}
	}
}

func y(pos, vel, firstPos []point3D, out chan int) {
	n := 0
	for {
		n++
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if pos[i].y < pos[j].y {
					vel[i].y++
					vel[j].y--
				} else if pos[i].y > pos[j].y {
					vel[i].y--
					vel[j].y++
				}
			}
		}
		for i := 0; i < len(pos); i++ {
			pos[i].y += vel[i].y
		}
		d := 0
		for i := 0; i < len(pos); i++ {
			if pos[i].y == firstPos[i].y && vel[i].y == 0 {
				d++
			}
		}
		if d == 4 {
			out <- n
			wg.Done()
			return
		}
	}
}

func z(pos, vel, firstPos []point3D, out chan int) {
	n := 0
	for {
		n++
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				if pos[i].z < pos[j].z {
					vel[i].z++
					vel[j].z--
				} else if pos[i].z > pos[j].z {
					vel[i].z--
					vel[j].z++
				}
			}
		}
		for i := 0; i < len(pos); i++ {
			pos[i].z += vel[i].z
		}
		d := 0
		for i := 0; i < len(pos); i++ {
			if pos[i].z == firstPos[i].z && vel[i].z == 0 {
				d++
			}
		}
		if d == 4 {
			out <- n
			wg.Done()
			return
		}
	}
}
