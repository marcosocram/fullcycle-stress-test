package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "fullcycle-stress-test",
		Short: "Sistema de Stress Test em Go",
		Run: func(cmd *cobra.Command, args []string) {
			executeStressTest()
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Número total de requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Número de chamadas simultâneas")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")

	arg := os.Args[:]
	fmt.Println(arg)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func executeStressTest() {
	var wg sync.WaitGroup
	statusCount := make(map[int]int)
	statusCountLock := sync.Mutex{}

	requestsPerWorker := requests / concurrency
	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("Erro ao fazer request:", err)
					continue
				}
				defer resp.Body.Close()

				statusCountLock.Lock()
				statusCount[resp.StatusCode]++
				statusCountLock.Unlock()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	// Gerar o relatório
	report := generateReport(totalTime, statusCount)

	// Exibir o relatório no terminal
	fmt.Println(report)

	// Salvar o relatório em um arquivo
	saveReportToFile(report)
}

func generateReport(totalTime time.Duration, statusCount map[int]int) string {
	report := fmt.Sprintf("\nRelatório de Teste:\n")
	report += fmt.Sprintf("URL Testada: %s\n", url)
	report += fmt.Sprintf("Tempo Total: %v\n", totalTime)
	report += fmt.Sprintf("Requests Totais: %d\n", requests)
	report += fmt.Sprintf("Status 200: %d\n", statusCount[200])

	for status, count := range statusCount {
		if status != 200 {
			report += fmt.Sprintf("Status %d: %d\n", status, count)
		}
	}

	return report
}

func saveReportToFile(report string) {
	outputDir := "/app/reports"
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("%s/report-%s.txt", outputDir, timestamp)

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Erro ao criar o diretório de relatório: %v\n", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo de relatório: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(report)
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo de relatório: %v\n", err)
		return
	}

	fmt.Printf("Relatório salvo em: %s\n", filename)
}
