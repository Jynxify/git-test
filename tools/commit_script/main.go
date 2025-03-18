package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

// Fonction pour exécuter une commande et retourner la sortie ou l'erreur
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Fonction pour obtenir le nom de la branche Git actuelle
func getCurrentBranch() (string, error) {
	output, err := runCommand("git", "branch", "--show-current")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// Fonction pour vérifier s'il y a des modifications à valider
func checkChanges() (bool, error) {
	output, err := runCommand("git", "status", "--short")
	if err != nil {
		return false, err
	}
	return len(output) > 0, nil
}

// Fonction pour afficher les changements en cours de validation
func showStagedChanges() string {
	output, err := runCommand("git", "diff", "--cached")
	if err != nil {
		return "Erreur lors de l'affichage des changements validés."
	}
	return output
}

// Fonction de confirmation (oui/non)
func confirm(prompt string) bool {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := strings.ToLower(scanner.Text())
	return answer == "y" || answer == "yes"
}

// Fonction pour réinitialiser le dépôt Git à son dernier commit
func resetChanges() error {
	_, err := runCommand("git", "reset")
	return err
}

// Fonction pour réinitialiser complètement le dépôt Git (reset dur)
func resetHardChanges() error {
	_, err := runCommand("git", "reset", "--hard")
	return err
}

func waitForExit() {
	fmt.Println("Appuyez sur une touche pour quitter...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	// Style de couleur avec Lipgloss
	var (
		titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("62"))
		sectionStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("34"))
		errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("160"))
		successStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("44"))
		questionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("228"))
	)

	// Icônes
	var (
		checkIcon = lipgloss.NewStyle().SetString("✔").Foreground(lipgloss.Color("42")).String()
		crossIcon = lipgloss.NewStyle().SetString("✖").Foreground(lipgloss.Color("160")).String()
		// arrowIcon = lipgloss.NewStyle().SetString("➜").Foreground(lipgloss.Color("228")).String()
		infoIcon = lipgloss.NewStyle().SetString("ℹ").Foreground(lipgloss.Color("34")).String()
	)

	// Créer un spinner
	spinnerModel := spinner.New(spinner.WithSpinner(spinner.Dot))
	spinnerModel.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	// Vérifier si le répertoire est un dépôt Git
	_, err := runCommand("git", "rev-parse", "--is-inside-work-tree")
	if err != nil {
		fmt.Println(crossIcon, errorStyle.Render("Erreur : Ce n'est pas un dépôt Git."))
		waitForExit()
		return
	}

	// Afficher le statut Git actuel
	fmt.Println(titleStyle.Render("==============================="))
	fmt.Println(titleStyle.Render("STATUT GIT AVANT COMMIT"))
	fmt.Println(titleStyle.Render("==============================="))
	status, err := runCommand("git", "status", "--short")
	if err != nil {
		fmt.Println(crossIcon, errorStyle.Render("Erreur lors de l'affichage du statut Git :"), err)
		waitForExit()
		return
	}
	fmt.Println(sectionStyle.Render(status))

	// Demander confirmation pour continuer
	if !confirm(questionStyle.Render("Voulez-vous continuer avec le commit et push ? (y/n) : ")) {
		fmt.Println(crossIcon, "Opération annulée.")
		waitForExit()
		return
	}

	// Demander le message de commit
	fmt.Print(questionStyle.Render("Entrez le message de commit : "))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commitMsg := scanner.Text()

	// Obtenir le nom de la branche actuelle
	branch, err := getCurrentBranch()
	if err != nil {
		fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la récupération de la branche actuelle :"), err)
		waitForExit()
		return
	}

	// Vérifier s'il y a des changements à valider
	changes, err := checkChanges()
	if err != nil {
		fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la vérification des changements :"), err)
		waitForExit()
		return
	}

	if changes {
		// Stage tous les changements
		fmt.Println(sectionStyle.Render("==============================="))
		fmt.Println(sectionStyle.Render("STAGING DES CHANGEMENTS"))
		fmt.Println(sectionStyle.Render("==============================="))

		// Utiliser un spinner pendant le processus de staging
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					spinnerModel.Tick()
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
		_, err := runCommand("git", "add", ".")
		cancel()
		if err != nil {
			fmt.Println(crossIcon, errorStyle.Render("Erreur lors du staging des changements :"), err)
			waitForExit()
			return
		}

		// Afficher les changements mis en scène
		fmt.Println(sectionStyle.Render("==============================="))
		fmt.Println(sectionStyle.Render("CHANGEMENTS MISE EN SCÈNE"))
		fmt.Println(sectionStyle.Render("==============================="))
		fmt.Println(showStagedChanges())

		// Confirmer le commit
		if !confirm(questionStyle.Render("Procéder avec le commit ? (y/n) : ")) {
			fmt.Println(crossIcon, "Commit annulé.")
			err := resetChanges()
			if err != nil {
				fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la réinitialisation des changements :"), err)
			}
			waitForExit()
			return
		}

		// Committer les changements
		ctx, cancel = context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					spinnerModel.Tick()
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
		_, err = runCommand("git", "commit", "-m", commitMsg)
		cancel()
		if err != nil {
			fmt.Println(crossIcon, errorStyle.Render("Erreur lors du commit des changements :"), err)
			err = resetChanges()
			if err != nil {
				fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la réinitialisation des changements :"), err)
			}
			waitForExit()
			return
		}

		fmt.Println(successStyle.Render("==============================="))
		fmt.Println(successStyle.Render(checkIcon + " COMMIT RÉUSSI"))
		fmt.Println(successStyle.Render("==============================="))
	} else {
		fmt.Println(infoIcon, "Aucun changement à valider.")
		waitForExit()
		return
	}

	// Confirmer avant de pousser
	if !confirm(questionStyle.Render("Voulez-vous pousser sur la branche '" + branch + "' ? (y/n) : ")) {
		fmt.Println(crossIcon, "Push annulé.")
		err := resetChanges()
		if err != nil {
			fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la réinitialisation des changements :"), err)
		}
		waitForExit()
		return
	}

	// Pousser les changements
	fmt.Println(sectionStyle.Render("==============================="))
	fmt.Println(sectionStyle.Render("PUSH VERS LA BRANCHE DISTANTE"))
	fmt.Println(sectionStyle.Render("==============================="))
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				spinnerModel.Tick()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	_, err = runCommand("git", "push", "origin", branch)
	cancel()
	if err != nil {
		fmt.Println(crossIcon, errorStyle.Render("Erreur lors du push des changements :"), err)
		err = resetChanges()
		if err != nil {
			fmt.Println(crossIcon, errorStyle.Render("Erreur lors de la réinitialisation des changements :"), err)
		}
		waitForExit()
		return
	}

	fmt.Println(successStyle.Render("==============================="))
	fmt.Println(successStyle.Render(checkIcon + " PUSH RÉUSSI"))
	fmt.Println(successStyle.Render("==============================="))

	// Attendre la pression d'une touche avant de quitter
	waitForExit()
}
