package main

import (
	wingetdownload "golangwinget/wingetDownload"
	wingetsearch "golangwinget/wingetSearch"
)

func main() {

	wingetsearch.WingetSearchMonikerSource()
	wingetdownload.WingetDownloadfromJSON()

	/*
		wingetsearch.WingetSearchMonikerSource()
		wingetlist.WingetJSONList()
	*/
}
