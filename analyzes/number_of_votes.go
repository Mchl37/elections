package analyzes

import (
	"log"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GetTotalVotes(filePath string) (float64, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	totalVotes := 0.0

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}

			cells := row.Cells

			voixStr := strings.Replace(cells[18].String(), ",", ".", -1)
			voix, err := strconv.ParseFloat(voixStr, 64)
			if err != nil {
				log.Printf("Erreur lors de la conversion des voix exprimées : %v", err)
				continue
			}

			totalVotes += voix
		}
	}

	return totalVotes, nil
}
