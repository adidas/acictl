package generator

type Generator struct {
	Folders Folder
	Sources []string
}

type Folder struct {
	Dot       string
	Templates string
}
