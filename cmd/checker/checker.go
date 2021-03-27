package checker

import (
	"encoding/csv"
	"fmt"
	"github.com/aeramu/ia-itb-scraper/internal/entity"
	"github.com/aeramu/ia-itb-scraper/internal/files"
	"github.com/aeramu/ia-itb-scraper/internal/request"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	maxRequest = 100
	outputFile = "out.csv"
	inputFile = "input.csv"
)

func Run() {
	arr, err := files.ReadCSV(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	writer := csv.NewWriter(file)
	if err := writer.Write([]string{"Nama", "Keyword Nama", "Prodi", "Tahun Daftar", "Status", "Error"}); err != nil {
		log.Fatalln(err)
	}
	writer.Flush()

	c := make(chan *entity.Alumnee, maxRequest)
	go scrape(c, arr)

	for res := range c {
		fmt.Println(res)
		if err := writer.Write([]string{res.Name, res.KeywordName, res.Major, res.Generation, res.Status, res.Error}); err != nil {
			log.WithFields(log.Fields{
				"err": err,
				"name": res.Name,
				"keyword name": res.KeywordName,
				"major": res.Major,
				"generation": res.Generation,
				"status": res.Status,
				"error": res.Error,
			}).Errorln("Failed write to file")
		}
		writer.Flush()
	}
}

func scrape(c chan *entity.Alumnee, arr [][]string) {
	resultCh := make(chan *entity.Alumnee, maxRequest)
	fetchCount := 0

	for i, data := range arr {
		name, major, generation := data[1], data[3], data[4]
		go func() {
			res, err := request.Search(name, major, generation)
			if err != nil {
				log.WithFields(log.Fields{
					"err": err,
					"name": name,
					"major": major,
					"generation": generation,
				}).Errorln("Failed search data")
			}
			resultCh <- res
		}()
		fetchCount++
		for wait := true; wait; {
			select {
			case result := <-resultCh:
				if result != nil {
					c <- result
				}
				fetchCount--
			default:
				if fetchCount <= maxRequest && i + 1 < len(arr) {
					wait = false
				}
				if fetchCount == 0 && i + 1 == len(arr) {
					wait = false
				}
			}
		}
	}
	close(c)
}