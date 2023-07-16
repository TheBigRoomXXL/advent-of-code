package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var result string
	part := flag.Int("part", 0, "part number (1 or 2)")
	flag.Parse()
	switch *part {
	case 1:
		result = part1(os.Stdin)
	case 2:
		result = part2(os.Stdin)
	default:
		panic("invalid part")
	}
	fmt.Println(result)
}

func part1(input io.Reader) string {
	root := parseInput(input)
	small_dirs := root.findDirLesserThan(100000)
	total := 0
	for _, file := range small_dirs {
		total += file.size()
	}
	return fmt.Sprint(total)
}

func part2(input io.Reader) string {
	root := parseInput(input)
	limit := 30000000 - (70000000 - root.size())
	min := 70000000
	small_dirs := root.findDirGreaterThan(limit)
	for _, file := range small_dirs {
		x := file.size()
		if x < min {
			min = x
		}
	}
	return fmt.Sprint(min)
}

type file struct {
	name      string
	base_size int
	childs    []*file
	parent    *file
}

func (f *file) size() int {
	size := f.base_size
	for _, childfile := range f.childs {
		size += childfile.size()
	}
	return size
}

func (f *file) child(name string) *file {
	for _, childfile := range f.childs {
		if childfile.name == name {
			return childfile
		}
	}
	panic(fmt.Sprintf("No child file match name %s", name))
}

func (f *file) findDirLesserThan(limit int) []*file {
	files := []*file{}
	if f.base_size == 0 && f.size() <= limit {
		files = append(files, f)
	}
	for _, childfile := range f.childs {
		files = append(files, childfile.findDirLesserThan(limit)...)
	}
	return files
}

func (f *file) findDirGreaterThan(limit int) []*file {
	files := []*file{}
	if f.base_size == 0 && f.size() >= limit {
		files = append(files, f)
	}
	for _, childfile := range f.childs {
		files = append(files, childfile.findDirGreaterThan(limit)...)
	}
	return files
}

func (f *file) format(deep int) string {
	str := fmt.Sprintf("%s - %s, %d\n", strings.Repeat(" ", deep*2), f.name, f.size())
	for _, childfile := range f.childs {
		str = str + childfile.format(deep+1)
	}
	return str
}

func parseInput(input io.Reader) *file {
	root := &file{"root", 0, []*file{}, nil}
	head := root
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			command := line[2:]
			switch {
			case command == "cd /":
				head = root
			case command == "cd ..":
				head = head.parent
			case command[:2] == "cd":
				dir_name := ""
				n, err := fmt.Sscanf(command, "cd %s", &dir_name)
				if n != 1 || err != nil {
					panic(fmt.Sprintf("cannot parse command: %s", command))
				}
				head = head.child(dir_name)
			}
		} else {
			switch {
			case line[:3] == "dir":
				dir_name := ""
				n, err := fmt.Sscanf(line, "dir %s", &dir_name)
				if n != 1 || err != nil {
					panic(fmt.Sprintf("cannot parse line: %s (%d match, %s)", line, n, err))
				}
				head.childs = append(head.childs, &file{dir_name, 0, []*file{}, head})
			default:
				file_name := ""
				file_size := 0
				n, err := fmt.Sscanf(line, "%d %s", &file_size, &file_name)
				if n != 2 || err != nil {
					panic(fmt.Sprintf("cannot parse line: %s (%d match, %s)", line, n, err))
				}
				head.childs = append(head.childs, &file{file_name, file_size, []*file{}, head})
			}
		}
	}
	return root
}
