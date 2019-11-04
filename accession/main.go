package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	simple_util "github.com/liserjrqlxue/simple-util"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// os
var (
	ex, _  = os.Executable()
	exPath = filepath.Dir(ex)
)

// flag
var (
	acc = flag.String(
		"acc",
		filepath.Join(exPath, "chr_accessions_GRCh37.p13"),
		"path of chr_accessions_GRCh37.p13",
	)
	mod = flag.String(
		"mod",
		"nc2chr",
		"convert type",
	)
	input = flag.String(
		"input",
		"",
		"input file",
	)
	output = flag.String(
		"output",
		"",
		"output file",
	)
)

func main() {
	flag.Parse()
	if *input == "" || *output == "" {
		flag.Usage()
		log.Print("-input and -output is required")
		os.Exit(1)
	}
	accInfo, _ := simple_util.File2MapArray(*acc, "\t", nil)
	var nc2chr = make(map[string]string)
	for _, item := range accInfo {
		chr := item["#Chromosome"]
		nc := item["RefSeq Accession.version"]
		nc2chr[nc] = chr
	}
	switch *mod {
	case "nc2chr":
		convertAcc(*input, *output, "\t", nc2chr)

	}
}

func convertAcc(input, output, sep string, k2v map[string]string) {
	file, err := os.Open(input)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(file)

	gr, err := gzip.NewReader(file)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(gr)

	out, err := os.Create(output)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(out)

	scanner := bufio.NewScanner(gr)

	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, sep)
		v, ok := k2v[array[0]]
		if ok {
			array[0] = v
		}
		_, err := fmt.Fprintln(out, strings.Join(array, sep))
		simple_util.CheckErr(err)
	}
	simple_util.CheckErr(scanner.Err())
}
