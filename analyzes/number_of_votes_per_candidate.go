package analyzes

import (
	"log"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GetVotesPerCandidate(filePath string) (map[string]float64, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	votesPerCandidate := make(map[string]float64)

	votesStartColumn := 25

	candidates := []string{
		"Nathalie ARTHAUD",
		"Fabien ROUSSEL",
		"Emmanuel MACRON",
		"Jean LASSALLE",
		"Marine LE PEN",
		"Éric ZEMMOUR",
		"Jean-Luc MÉLENCHON",
		"Anne HIDALGO",
		"Yannick JADOT",
		"Valérie PÉCRESSE",
		"Philippe POUTOU",
		"Nicolas DUPONT-AIGNAN",
	}

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}

			for i, candidate := range candidates {
				if i*7+votesStartColumn >= len(row.Cells) {
					break
				}
				votesStr := row.Cells[i*7+votesStartColumn].String()
				votes, err := strconv.ParseFloat(strings.Replace(votesStr, ",", ".", -1), 64)
				if err != nil {
					log.Printf("Erreur lors de la conversion des voix: %v", err)
					continue
				}
				votesPerCandidate[candidate] += votes
			}
		}
	}

	return votesPerCandidate, nil
}
