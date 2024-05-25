package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func HelloWorld() string {
	return `Hello World`
}

// func printTree(path string, prefix string, isLast bool, depth int) {
// 	files, err := os.ReadDir(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for i, file := range files {
// 		if file.IsDir() {
// 			fmt.Print(prefix)
// 			if isLast {
// 				fmt.Print("└── ")
// 			} else {
// 				fmt.Print("├── ")
// 			}
// 			fmt.Println(file.Name())

//				if depth > 0 {
//					var newPrefix string
//					if isLast {
//						newPrefix = prefix + "    "
//					} else {
//						newPrefix = prefix + "│   "
//					}
//					printTree(filepath.Join(path, file.Name()), newPrefix, i == len(files)-1, depth-1)
//				}
//			} else {
//				fmt.Print(prefix)
//				if isLast {
//					fmt.Print("└── ")
//				} else {
//					fmt.Print("├── ")
//				}
//				fmt.Println(file.Name())
//			}
//		}
//	}
var run = true

func liveTerminal() {
	for i := 0; i < 5001; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Print("\033[H\033[2J") // очистка терминала
		fmt.Printf("Live application running... %d\n", i)
	}
}

func liveTerminalAnimation() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	rand.Seed(time.Now().UTC().UnixNano())
	randomDataAndOffset := func() (data []float64, offset float64) {
		noSlices := 1 + rand.Intn(5)
		data = make([]float64, noSlices)
		for i := range data {
			data[i] = rand.Float64()
		}
		offset = 2.0 * math.Pi * rand.Float64()
		return
	}

	pc := widgets.NewPieChart()
	pc.Title = "Pie Chart"
	pc.SetRect(5, 5, 70, 36)
	pc.Data = []float64{.25, .25, .25, .25}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}

	pause := func() {
		run = !run
		if run {
			pc.Title = "Pie Chart"
		} else {
			pc.Title = "Pie Chart (Stopped)"
		}
		ui.Render(pc)
	}

	ui.Render(pc)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "s":
				pause()
			}
		case <-ticker:
			if run {
				pc.Data, pc.AngleOffset = randomDataAndOffset()
				ui.Render(pc)
			}
		}
	}
}

func liveTerminalTime() {
	for {
		currentTime := time.Now()
		fmt.Printf("Текущее время: %02d:%02d:%02d\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
		fmt.Printf("Текущая дата: %s\n", currentTime.Format("2006-01-02"))
		time.Sleep(1 * time.Second)
	}
}

func main() {

	liveTerminal()
	liveTerminalAnimation()
	liveTerminalTime()
	// var depth int
	// flag.IntVar(&depth, "n", -1, "depth of the tree")
	// flag.Parse()

	// args := flag.Args()
	// if len(args) == 0 {
	// 	fmt.Println("Please provide a directory path.")
	// 	os.Exit(1)
	// }

	// path := args[0]

	// if !strings.HasPrefix(path, "/") {
	// 	wd, err := os.Getwd()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// 	path = filepath.Join(wd, path)
	// }

	// printTree(path, "", true, depth)
}
