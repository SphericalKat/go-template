package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/go-errors/errors"
)

func GetStack(err *errors.Error) string {
	basePath := GetBasePath() + "/"
	stack := ""

	allFrames := err.StackFrames()
	frames := make([]errors.StackFrame, 0, len(allFrames))
	for _, frame := range allFrames {
		if !strings.Contains(frame.Package, "runtime") {
			frames = append(frames, frame)
		}
	}
	frameCount := len(frames)

	for i, frame := range frames {
		var file string
		if strings.Contains(frame.File, basePath) {
			file = strings.Replace(frame.File, basePath, "", 1)
		} else {
			file = color.BlackString(frame.File)
		}

		frameStr := fmt.Sprintf("%s%s%d", file, color.BlueString(":"), frame.LineNumber)

		source, err := frame.SourceLine()
		boldBlack := color.New(color.FgHiBlack, color.Bold).SprintFunc()
		stack += boldBlack(fmt.Sprintf("%2d: ", frameCount-i))
		if err != nil {
			stack += fmt.Sprintf("%s\n", frameStr)
		} else {
			var arrow string
			var fmtdSource string
			if strings.Contains(frame.File, basePath) {
				arrow = color.CyanString(">> ")
				fmtdSource = source
			} else {
				arrow = color.BlackString("> ")
				fmtdSource = color.BlackString(source)
			}

			stack += fmt.Sprintf("%s\n    %s%s\n", frameStr, arrow, fmtdSource)
		}
	}

	return " " + err.Error() + "\n" + stack
}
