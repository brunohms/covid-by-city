package Helper

type City interface {
	Name() string
	GetLink() string
	FetchDownloadLinks(downloadLinksFilePath string)
	DownloadFiles(downloadLinksFilePath string)
	ProcessFiles() []CovidData
}
