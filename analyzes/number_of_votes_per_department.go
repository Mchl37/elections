package analyzes

import (
	"log"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GetVotesPerCandidateByDepartment(filePath string) (map[string]map[string]float64, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	votesPerCandidateByDepartment := make(map[string]map[string]float64)

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

	votesStartColumn := 25
	departmentIndex := 1

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}

			if len(row.Cells) < votesStartColumn {
				continue
			}

			department := row.Cells[departmentIndex].String()

			if _, ok := votesPerCandidateByDepartment[department]; !ok {
				votesPerCandidateByDepartment[department] = make(map[string]float64)
			}

			for i, candidate := range candidates {
				if i*7+votesStartColumn >= len(row.Cells) {
					break
				}
				votesStr := row.Cells[i*7+votesStartColumn].String()
				if votesStr == "" {
					continue
				}
				votes, err := strconv.ParseFloat(strings.Replace(votesStr, ",", ".", -1), 64)
				if err != nil {
					log.Printf("Erreur lors de la conversion des voix pour le candidat %s dans le département %s: %v", candidate, department, err)
					continue
				}
				votesPerCandidateByDepartment[department][candidate] += votes
			}
		}
	}

	return votesPerCandidateByDepartment, nil
}
