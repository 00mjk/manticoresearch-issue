package main

import (
	"bufio"
	"os"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/manticoresoftware/go-sdk/manticore"
	"github.com/spf13/pflag"
)

var (
	manticorePort uint16
	manticoreHost string
	inputFile     string
	query         string
	indexName     string
	isIndex       bool
	isSearch      bool
	help          bool
)

func main() {
	pflag.StringVarP(&manticoreHost, "manticore-host", "m", "localhost", "input-file")
	pflag.Uint16VarP(&manticorePort, "manticore-port", "p", 9312, "input-file")
	pflag.StringVarP(&query, "query", "q", "Fulda", "query string")
	pflag.StringVarP(&indexName, "index-name", "n", "rt_pneus_illico_catalog2", "rt_index name")
	pflag.StringVarP(&inputFile, "input-file", "f", "./data/pneus-illico-dump.sql", "input-file")
	pflag.BoolVarP(&isIndex, "index", "i", false, "index data")
	pflag.BoolVarP(&isSearch, "search", "s", false, "search query")
	pflag.BoolVarP(&help, "help", "h", false, "display help")
	pflag.Parse()
	if help {
		pflag.PrintDefaults()
		os.Exit(1)
	}

	cl, _, err := initSphinx(manticoreHost, manticorePort)
	check(err)

	if isIndex {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			query := scanner.Text()
			log.Info("query:", query)
			resp, err := cl.Sphinxql(query)
			if err != nil {
				log.Println("query error: ", query)
				log.Fatalln(err)
			}
			if resp[0].Msg != "" {
				log.Println("query msg: ", query)
				log.Fatalln(resp[0].Msg)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if isSearch {
		fmt.Println("query: ", query)
		fmt.Println("indexName: ", indexName)
		res2, err2 := cl.Query(query, indexName)
		fmt.Println(res2, err2)
	}
}

func initSphinx(host string, port uint16) (manticore.Client, bool, error) {
	cl := manticore.NewClient()
	cl.SetServer(host, port)
	status, err := cl.Open()
	if err != nil {
		return cl, status, err
	}
	return cl, status, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
