package structs

import (
	"fmt"
	"strconv"
)

type Directory struct {
	Pointer    *Directory
	Name       string
	TotalSize  int
	Files      []File
	Directorys []Directory
}

type File struct {
	Name string
	Size int
}

func PrintDirectory(d Directory, j int) {
	index := ""
	for i := 0; i < j; i++ {
		index += " "
	}
	fmt.Println(index+"- "+d.Name+" (dir) Totalsize: ", d.TotalSize)
	index += " "
	for i := 0; i < len(d.Directorys); i++ {
		PrintDirectory(d.Directorys[i], j+1)
	}

	for i := 0; i < len(d.Files); i++ {
		fmt.Println(index + "- " + d.Files[i].Name + " (file, size=" + strconv.FormatInt(int64(d.Files[i].Size), 10) + ")")
	}
}

func CalculateTotalSizes(d *Directory) {

	for i := 0; i < len(d.Directorys); i++ {
		CalculateTotalSizes(&d.Directorys[i])
	}

	for i := 0; i < len(d.Files); i++ {
		d.TotalSize += d.Files[i].Size
	}
	for i := 0; i < len(d.Directorys); i++ {
		d.TotalSize += d.Directorys[i].TotalSize
	}
}

func GetDirsToErase(d Directory, e *Directory) {

	for i := 0; i < len(d.Directorys); i++ {
		GetDirsToErase(d.Directorys[i], e)
	}

	if d.TotalSize < 100000 {
		e.Directorys = append(e.Directorys, d)
	}

}

func GetDirsFreeSpaceEnough(d Directory, e *Directory, spaceNeeded int) {

	for i := 0; i < len(d.Directorys); i++ {
		GetDirsFreeSpaceEnough(d.Directorys[i], e, spaceNeeded)
	}

	if d.TotalSize+spaceNeeded > -1 {
		e.Directorys = append(e.Directorys, d)
	}

}
