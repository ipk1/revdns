package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./cidr_resolver <CIDR or file>")
		return
	}

	arg := os.Args[1]
	var cidrs []string

	if strings.Contains(arg, "/") {
		cidrs = append(cidrs, arg)
	} else {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			cidrs = append(cidrs, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	outputFile, err := os.Create("resolved_domains.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for _, cidr := range cidrs {
		runTool("hakrevdns", cidr, outputFile)
		runTool("hakip2host", cidr, outputFile)
		runTool("cero", cidr, outputFile)
		runTool("gdn", cidr, outputFile)
	}

	runHttprobe("resolved_domains.txt")
}

func runTool(tool, cidr string, outputFile *os.File) {
	cmd := exec.Command(tool, cidr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running %s: %v\n", tool, err)
		return
	}

	outputFile.WriteString(string(out))
}

func runHttprobe(inputFile string) {
	cmd := exec.Command("cat", inputFile)
	pipe, _ := cmd.StdoutPipe()

	cmd2 := exec.Command("httprobe")
	cmd2.Stdin = pipe
	cmd2.Stdout = os.Stdout

	cmd.Start()
	cmd2.Start()
	cmd.Wait()
	cmd2.Wait()
}
