package Cities

import (
	"bufio"
	"covid-crawler/Helper"
	"covid-crawler/LinkData"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type OuroPreto struct {}

type CovidDayInfo struct {
	city string
	id int
	day, month, year int
}

func (op OuroPreto) Name() string {
	return "ouro-preto"
}

func (op OuroPreto) GetLink() string {
	return "https://www.ouropreto.mg.gov.br"
}

func (op OuroPreto) FetchDownloadLinks(downloadLinksFilePath string) {
	c := colly.NewCollector()
	var downloadLinks []LinkData.LinkData

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			var texts, links []string

			tr.ForEach("td", func(i int, element *colly.HTMLElement) {
				pdfLink := element.ChildAttr("a[href]", "href")
				text := element.ChildText("a")
				texts = append(texts, element.Text)
				if len(pdfLink) > 0 && strings.Contains(text, "BOLETIM"){
					links = append(links, op.GetLink() + pdfLink)
				}
			})

			if len(links) > 0 {
				downloadLinks = append(downloadLinks, LinkData.LinkData{
					Link: links[0],
					Date: texts[0],
				})
			}
		})
	})
	c.Visit(op.GetLink() + "/coronavirus")

	f, err := os.Create(downloadLinksFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := len(downloadLinks) - 1 ; i >= 0 ; i-- {
		f.WriteString(downloadLinks[i].String() + "\n")
	}
}

func (op OuroPreto) DownloadFiles(downloadLinksFilePath string) {
	file, err := os.Open(downloadLinksFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var downloadLinks []LinkData.LinkData

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		downloadLinks = append(downloadLinks, LinkData.NewLinkData(scanner.Text()))
	}

	total := strconv.Itoa(len(downloadLinks))
	for i, downloadInfo := range downloadLinks {
		index := strconv.Itoa(i + 1)
		for len(total) > len(index) {
			index = "0" + index
		}

		fileName := op.Name() + "-" + index + "-" + downloadInfo.Date + ".pdf"

		fmt.Println("Downloading to file:", fileName, "from link:", downloadInfo.Link)
		err = Helper.DownloadFile("downloads/" + op.Name() + "/" + fileName, downloadInfo.Link)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getDataFromFileName(filename string) CovidDayInfo {
	last := strings.LastIndex(filename, "/")
	file := filename[last+1:]
	dot := strings.Index(file, ".")
	file = file[:dot]

	split := strings.Split(file, "-")
	dayInfo := CovidDayInfo{}

	dayInfo.city = strings.Join(split[:2], "-")
	id, err := strconv.Atoi(split[2])
	if err != nil {
		fmt.Println(err)
	} else {
		dayInfo.id = id
	}
	dayInfo.day, err = strconv.Atoi(split[3])
	if err != nil {
		fmt.Println(err)
	}

	dayInfo.month, err = strconv.Atoi(split[4])
	if err != nil {
		fmt.Println(err)
	}

	dayInfo.year, err = strconv.Atoi(split[5])
	if err != nil {
		fmt.Println(err)
	}

	return dayInfo
}

func printer(numbers []int) Helper.CovidData {
	for i, n := range numbers {
		fmt.Print("[", i, "] -> ", n, " ")
	}
	fmt.Println()
	return Helper.CovidData{}
}

func getProcessorFromId(id int) func([]int) Helper.CovidData {
	if id == 1 {
		return func(numbers []int) Helper.CovidData {
			var data Helper.CovidData
			data.NotifiedCases = numbers[6]
			data.InvestigatingCases = numbers[3]
			data.DiscardedCases = numbers[4]
			data.ConfirmedCases = numbers[5]
			return data
		}
	}
	if id <= 15 {
		return func(numbers []int) Helper.CovidData {
			var data Helper.CovidData
			data.NotifiedCases = numbers[0]
			data.InvestigatingCases = numbers[6]
			data.DiscardedCases = numbers[1]
			data.ConfirmedCases = numbers[5]
			return data
		}
	}
	if id <= 28 {
		return printer
	}
	if id <= 43 {
		return func(numbers []int) Helper.CovidData {
			fmt.Println(numbers)
			var data Helper.CovidData
			data.NotifiedCases = numbers[10]
			data.InvestigatingCases = numbers[7]
			data.ConfirmedCases = numbers[4]
			data.DiscardedCases = numbers[9]
			data.ExcludedCases = numbers[8]
			return data
		}
	}
	if id <= 47 {
		return func(numbers []int) Helper.CovidData {
			var data Helper.CovidData
			data.NotifiedCases = numbers[9]
			data.InvestigatingCases = numbers[5]
			data.ConfirmedCases = numbers[4]
			data.DiscardedCases = numbers[7]
			data.ExcludedCases = numbers[6]

			data.InvestigatingDeaths = numbers[8]

			fluSymptomsStr := strconv.Itoa(numbers[14])
			fluSymptoms, err := strconv.Atoi(fluSymptomsStr[2:])
			if err != nil {
				fmt.Println("Error parsing:", fluSymptomsStr, err)
				return Helper.CovidData{}
			}
			data.FluSymptoms = fluSymptoms
			return data
		}
	}
	if id <= 60 {
		return func(numbers []int) Helper.CovidData {
			var data Helper.CovidData
			data.NotifiedCases = numbers[9]
			data.InvestigatingCases = numbers[5]
			data.ConfirmedCases = numbers[4]
			data.DiscardedCases = numbers[7]
			data.ExcludedCases = numbers[6]

			data.InvestigatingDeaths = numbers[8]

			data.FluSymptoms = numbers[14]
			return data
		}
	}
	if id <= 63 {
		return func(numbers []int) Helper.CovidData {
			var data Helper.CovidData
			data.NotifiedCases = numbers[7]
			data.InvestigatingCases = numbers[4]
			data.ConfirmedCases = numbers[14]
			data.DiscardedCases = numbers[6]
			data.ExcludedCases = numbers[5]

			data.ConfirmedPatients = numbers[8]

			data.InvestigatingDeaths = numbers[16]
			data.ConfirmedDeath = numbers[15]

			data.FluSymptoms = numbers[13]
			return data
		}
	}

	return printer
}

func (op OuroPreto) ProcessFiles() []Helper.CovidData {
	var files []string

	err := filepath.Walk("./downloads/" + op.Name(), func (path string, info os.FileInfo, _ error) error {
		if strings.HasSuffix(path, ".pdf") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	path, _ := os.Getwd()
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	var txtToProcess []string
	errorProcessing := make(map[string]bool)

	for _, file := range files {
		//fmt.Println("Extracting text from pdf of file:", file)
		fileTxt := strings.Replace(file, ".pdf", ".txt", 1)
		txtToProcess = append(txtToProcess, fileTxt)

		if Helper.FileExists(fileTxt) {
			continue
		}

		output, err := exec.Command("java", "-jar", path + "pdf-reader.jar", path + file).CombinedOutput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(output) > 0 {
			errorProcessing[fileTxt] = true
		}
	}

	re := regexp.MustCompile("[0-9]+(\\.[0-9]+)?")

	var covidData []Helper.CovidData

	fmt.Println("Processing:", len(txtToProcess))
	for _, file := range txtToProcess {
		//fmt.Println("Processing:", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", file)
			continue
		}

		numbers := re.FindAllString(string(data), -1)
		if len(numbers) == 0 {
			errorProcessing[file] = true
			continue
		}

		var numbersParsed []int
		for _, n := range numbers {
			aux := strings.Replace(n, ".", "", -1)
			nParsed, err := strconv.Atoi(aux)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numbersParsed = append(numbersParsed, nParsed)
		}
		dayInfo := getDataFromFileName(file)

		processor := getProcessorFromId(dayInfo.id)
		if processor == nil {
			errorProcessing[file] = true
			continue
		}

		covidDataOfTheDay := processor(numbersParsed)

		covidDataOfTheDay.Day = dayInfo.day
		covidDataOfTheDay.Month = dayInfo.month
		covidDataOfTheDay.Year = dayInfo.year

		covidData = append(covidData, covidDataOfTheDay)

		if dayInfo.id > 60 {
			fmt.Println(dayInfo, numbersParsed, covidDataOfTheDay)
		}
	}

	fmt.Println("Files with problem in processing:", len(errorProcessing))
	for invalidFile := range errorProcessing {
		last := strings.LastIndex(invalidFile, "/")
		fmt.Println("Error processing:", invalidFile[last+1:])
	}

	return covidData
}