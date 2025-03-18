package readfiles

func GetFileList(path string) {
	myFolderReader := createFolderReader(path)
	myFolderReader.init()
	myFolderReader.printFiles()

}
