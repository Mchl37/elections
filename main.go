package main

import (
	"fmt"
	"log"
	"main/analyzes"
)

func main() {
	// Chemin du fichier
	filePath := "./excel/resultats-par-niveau-burvot-t1-france-entiere.xlsx"

	// Nombre de votes exprimés
	totalVotes, err := analyzes.GetTotalVotes(filePath)
	if err != nil {
		log.Fatal("Erreur lors de la récupération du nombre total de votes exprimés :", err)
	}

	// Nombre de votes par candidat
	votesPerCandidate, err := analyzes.GetVotesPerCandidate(filePath)
	if err != nil {
		log.Fatal("Erreur lors de la récupération du nombre de votes par candidat :", err)
	}

	// Nombre de votes par candidat et par département
	votesPerCandidateByDepartment, err := analyzes.GetVotesPerCandidateByDepartment(filePath)
	if err != nil {
		log.Fatal("Erreur lors de la récupération du nombre de votes par candidat et par département :", err)
	}

	// Palmarès des départements par nombre de votants
	departmentRankings, err := analyzes.GenerateDepartmentRankings(filePath)
	if err != nil {
		log.Fatal("Erreur lors de la génération du palmarès des départements :", err)
	}

	// Affichage des résultats
	fmt.Printf("Nombre total de votes exprimés : %.2f\n", totalVotes)

	fmt.Println("Nombre de votes par candidat :")
	for candidat, voix := range votesPerCandidate {
		fmt.Printf("- %s : %.2f\n", candidat, voix)
	}

	fmt.Println("Nombre de votes par candidat par département :")
	for department, votesPerCandidate := range votesPerCandidateByDepartment {
		fmt.Printf("Département : %s\n", department)
		for candidat, voix := range votesPerCandidate {
			fmt.Printf("- %s : %.2f\n", candidat, voix)
		}
	}

	fmt.Println("Palmarès des départements par nombre de votants :")
	for _, department := range departmentRankings {
		fmt.Println(department)
	}
}
