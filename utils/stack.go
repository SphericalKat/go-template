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
		file := strings.Replace(frame.File, basePath, "", 1)
		frameStr := fmt.Sprintf("%s%s%d", file, color.BlueString(":"), frame.LineNumber)

		source, err := frame.SourceLine()
		boldBlack := color.New(color.FgHiBlack, color.Bold).SprintFunc()
		stack += boldBlack(fmt.Sprintf("%d: ", frameCount-i))
		if err != nil {
			stack += fmt.Sprintf("%s\n", frameStr)
		} else {
			stack += fmt.Sprintf("%s\n    %s\n", frameStr, color.HiBlackString(source))
		}
	}

	return color.RedString("[error]") + " " + err.Error() + "\n" + stack
}
