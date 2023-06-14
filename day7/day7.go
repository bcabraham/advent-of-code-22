package day7

import (
	"advent-of-code-22/lib"
	"fmt"
	"strings"
)

const (
	TotalDiskSpace   = 70_000_000
	TotalSpaceNeeded = 30_000_000
)

func Run() {
	input, err := lib.ReadLines("file-traverse.txt") // file-traverse-test
	lib.HandleError(err)

	root := Directory{}
	root.Name = "/"
	var cwd *Directory = &root
	for _, line := range input {
		// fmt.Printf("%d: %s\n", i, line)
		cmd := strings.Split(line, " ")

		switch cmd[0] {
		case "$":
			switch cmd[1] {
			case "ls":
				continue
			case "cd":
				cwd = cwd.ChangeDir(cmd[2])
				// fmt.Println(cwd.Path())
			}
		case "dir":
			cwd.AddDirectory(cmd[1])
		default:
			cwd.AddFile(lib.StrToInt(cmd[0]), cmd[1])
		}
	}

	currentSpaceUsed := root.Size()
	unusedSpace := TotalDiskSpace - currentSpaceUsed
	spaceNeeded := TotalSpaceNeeded - unusedSpace

	fmt.Printf(
		"currentSpaceUsed: %d\nunusedSpace: %d\nspaceNeeded: %d\n",
		currentSpaceUsed,
		unusedSpace,
		spaceNeeded)

	dirs := root.GetSize(currentSpaceUsed)

	fmt.Println("dir\t\tsize")
	minMemToDelete := currentSpaceUsed
	for _, d := range dirs {
		s := d.Size()

		if s >= spaceNeeded && s < minMemToDelete {
			minMemToDelete = s
			fmt.Printf("%s\t\t%d\n", d.Path(), s)
		}
	}

	fmt.Println("minMemToDelete:", minMemToDelete)
}

type File struct {
	Size int
	Name string
}

type Directory struct {
	Parent   *Directory
	Name     string
	Files    []File
	Children []*Directory
}

type Callable interface{}

func (d *Directory) AddFile(size int, name string) {
	f := File{size, name}
	d.Files = append(d.Files, f)
}

func (d *Directory) AddDirectory(name string) {
	nd := Directory{}
	nd.Name = name
	nd.Parent = d
	d.Children = append(d.Children, &nd)
}

func (d *Directory) ChangeDir(path string) *Directory {
	if path == ".." {
		return d.Parent
	}

	for _, c := range d.Children {
		if c.Name == path {
			return c
		}
	}

	return d
}

func (d *Directory) Path() string {
	path := d.Name
	if d.Parent != nil {
		path = fmt.Sprintf("%s%s/", d.Parent.Path(), path)
	}

	return path
}

func (d *Directory) Size() int {
	size := 0

	for _, f := range d.Files {
		size += f.Size
	}

	for _, c := range d.Children {
		size += c.Size()
	}

	return size
}

func (d *Directory) GetSize(size int) []*Directory {
	dirs := []*Directory{}

	if d.Size() <= size {
		dirs = append(dirs, d)
	}

	for _, c := range d.Children {
		dirs = append(dirs, c.GetSize(size)...)
	}

	return dirs
}
