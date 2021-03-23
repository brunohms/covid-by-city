package Cities

type Resende struct {}

func (r Resende) Name() string {
	return "resende"
}

func (r Resende) GetLink() string {
	return "https://www.resende.rj.gov.br/boletins-covid-19"
}

func (r Resende) FetchDownloadLinks(downloadLinksFilePath string) {

}

func (r Resende) DownloadFiles(downloadLinksFilePath string) {

}

func (r Resende) ProcessFiles() {
	
}