package day7

import (
	"strconv"
	"strings"

	"github.com/A-Siam/aoc-go/utils"
	"github.com/A-Siam/aoc-go/utils/datastructures"
)

type File struct {
	Children []*File
	Size     int64
	Name     string
	Parent   *File
}

type Condition func(f File, root File) bool
type Operation func(f File)

func newFile(name string, size int64, parent *File) File {
	return File{
		Children: []*File{},
		Size:     size,
		Name:     name,
		Parent:   parent,
	}
}

type Command struct {
	program string
	arg     string
}

func tokenize(command string) Command {
	splits := strings.Split(command, " ")
	arg := ""
	if len(splits) == 3 {
		arg = splits[2]
	}
	return Command{
		program: splits[1],
		arg:     arg,
	}
}

func Solution(inputPath string, cond Condition, op Operation) {
	root := newFile("/", 0, nil)
	dirStack := datastructures.NewStack([]*File{})
	dirStack.Push(&root)
	filesMap := []*File{}
	filesMap = append(filesMap, &root)
	utils.GetInput(inputPath, func(line string) {
		if isLineCommand(line) {
			command := tokenize(line)
			if command.program == "cd" {
				if command.arg == ".." {
					dirStack.Pop()
				} else if command.arg != "/" {
					// in a child of root
					// assume it will always go to new dirs
					current, _ := dirStack.Top()
					newCurrent := newFile(command.arg, 0, *current)
					newCurrent.Parent = *current
					filesMap = append(filesMap, &newCurrent)
					(*current).Children = append((*current).Children, &newCurrent)
					dirStack.Push(&newCurrent)
				}
			} // we don't care about the ls program in our problem at least for part 1 :smile: I hope this remains in part 2 :rofl:
		} else {
			// inside ls output
			splits := strings.Split(line, " ")
			size, err := strconv.ParseInt(splits[0], 10, 64)
			if err == nil {
				current, _ := dirStack.Top()
				increaseSize(*current, size)
			}
		}
	})

	// traverse the fils
	for _, val := range filesMap {
		if cond(*val, root) {
			op(*val)
		}
	}

}

func isLineCommand(line string) bool {
	return strings.Index(line, "$") == 0
}

func increaseSize(file *File, increment int64) {
	if file == nil {
		return
	}
	file.Size += increment
	increaseSize(file.Parent, increment)
}
