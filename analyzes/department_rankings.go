package analyzes

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func GenerateDepartmentRankings(filePath string) ([]string, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %v", err)
	}

	votersPerDepartment := make(map[string]int)

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			if len(row.Cells) >= 11 {
				department := row.Cells[1].String()
				votersStr := row.Cells[10].String()

				voters, err := strconv.Atoi(strings.TrimSpace(votersStr))
				if err != nil {
					return nil, fmt.Errorf("erreur lors de la conversion du nombre de votants : %v", err)
				}

				votersPerDepartment[department] += voters
			}
		}
	}

	type departmentVotes struct {
		department string
		votes      int
	}
	var departmentVotesSlice []departmentVotes
	for department, votes := range votersPerDepartment {
		departmentVotesSlice = append(departmentVotesSlice, departmentVotes{department, votes})
	}

	sort.Slice(departmentVotesSlice, func(i, j int) bool {
		return departmentVotesSlice[i].votes > departmentVotesSlice[j].votes
	})

	var result []string
	for _, departmentVotes := range departmentVotesSlice {
		result = append(result, fmt.Sprintf("%s : %d votants", departmentVotes.department, departmentVotes.votes))
	}

	return result, nil
}
