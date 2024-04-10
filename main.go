package main

import (
	wingetdownload "golangwinget/wingetDownload"
	wingetlist "golangwinget/wingetList"
)

func main() {

	//wingetlist.WingetJSONList()

	//wingetlist.WingetListtag()

	wingetlist.WingetMonikerSource()
	wingetdownload.WingetDownloadfromJSON()
}
