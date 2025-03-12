package views

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
	"todo-app/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var filename = "data.txt"

func PersonalTab() *fyne.Container {
	// set a clock
	clock := widget.NewLabel("")
	clock.TextStyle.Bold = true
	clock.SetText(time.Now().Format("Time: 03:04:05"))
	go func() {
		for range time.Tick(time.Second) {
			clock.SetText(time.Now().Format("Time: 03:04:05"))
		}
	}()

	// Task Data and List
	var taskChecked []bool
	var taskData []string
	var taskList *widget.List

	taskList = widget.NewList(
		func() int {
			return len(taskData)
		},
		func() fyne.CanvasObject {
			check := widget.NewCheck("", nil)
			// check.
			deleteTaskBtn := components.NewDangerButton("R", nil)
			return container.NewHBox(
				container.New(layout.NewGridWrapLayout(fyne.NewSize(350, check.MinSize().Height)), check),
				container.New(layout.NewGridWrapLayout(fyne.NewSize(40, 20)), deleteTaskBtn),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			objects := o.(*fyne.Container).Objects
			// update check
			check := objects[0].(*fyne.Container).Objects[0].(*widget.Check)
			check.SetText(taskData[i])
			check.SetChecked(taskChecked[i])
			// log.Println(taskChecked[i])
			// check.Checked = taskChecked[i]
			check.OnChanged = func(b bool) {
				taskChecked[i] = b
				taskList.Refresh()
				fileUpdateLine(i, b)
			}
			// on delete click, remove row
			deleteTaskBtn := objects[1].(*fyne.Container).Objects[0]
			deleteTaskBtn.(*components.CustomButton).OnTapped = func() {
				taskData = append(taskData[:i], taskData[i+1:]...)
				taskList.Refresh()
				fileDeleteLine(i)
			}
			// log.Println(i, reflect.TypeOf(deleteTaskBtn).String())
		},
	)

	// set task input
	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("What do you need to do?...")

	// saveTaskButton
	saveTaskBtn := components.NewPrimaryButton("Add", func() {
		text := taskEntry.Text
		if text == "" {
			return
		}
		taskData = append(taskData, text)
		taskEntry.SetText("")
		taskList.Refresh()
		// save to file
		fileAddLine(text)
	})

	// populate taskData using file lines
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			checked := false
			if strings.Contains(line, ";") {
				lineChars := strings.Split(line, ";")
				if lineChars[0] == "checked" {
					checked = true
				}
				line = lineChars[1]
			}
			taskChecked = append(taskChecked, checked)
			taskData = append(taskData, line)
		}
		if err := scanner.Err(); err != nil {
			log.Println("Error reading file:", err)
		}
	}

	// layout
	return container.NewVBox(
		container.NewHBox(clock),
		container.NewHBox(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(330, taskEntry.MinSize().Height)), taskEntry),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(60, saveTaskBtn.MinSize().Height)), saveTaskBtn),
		),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 300)), taskList),
	)
}

func ProfessionalTab() *fyne.Container {
	return container.NewVBox(widget.NewLabel("Professional Content"))
}

func fileUpdateLine(i widget.ListItemID, checked bool) {
	// update local file
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file and update line
	n := 0
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if n == i {
			if strings.Contains(line, ";") {
				lineChars := strings.Split(line, ";")
				if checked {
					line = "checked;" + lineChars[1]
				} else {
					line = "unchecked;" + lineChars[1]
				}
			} else if checked {
				line = "checked;" + line
			} else {
				line = "unchecked;" + line
			}
		}
		lines = append(lines, line+"\n")
		n++
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error scanning file:", err)
		return
	}

	err = os.WriteFile(filename, []byte(strings.Join(lines, "")), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
	}
}

func fileDeleteLine(i widget.ListItemID) {
	// update local file
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file and store lines except the one to delete
	n := 0
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if n != i {
			lines = append(lines, scanner.Text()+"\n")
		}
		n++
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error scanning file:", err)
		return
	}

	err = os.WriteFile(filename, []byte(strings.Join(lines, "")), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
	}
}

func fileAddLine(text string) {
	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// file does not exist, create it
		err := os.WriteFile(filename, []byte(text+"\n"), 0644)
		if err != nil {
			log.Println("Error writing file:", err)
			return
		}
	} else {
		// file exists, append data
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(text + "\n")
		if err != nil {
			log.Println("Error writing to file:", err)
			return
		}
	}
}
