package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/psykhi/sequencing"
	"github.com/satori/go.uuid"
	"golang.org/x/tools/container/intsets"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func similarity(b byte, b2 byte) int {
	if b == b2 {
		return 1
	}
	return -1
}
func similarityStrings(a string, b string) int {
	if a == b {
		return 1
	}
	return -1
}

type Cluster struct {
	id           uuid.UUID
	OriginalLine []string
	Matches      []*Match
}

type Match struct {
	Score int
	Ratio float64
	Line  string
}

func NewCluster(line []string) *Cluster {
	matches := make([]*Match, 0)
	id, _ := uuid.NewV4()
	return &Cluster{
		id,
		line,
		matches,
	}
}

var dels = []rune{'/', ',', ':', ' ', ')', '(', '|', '{', '}', '=', '.', '"', '\'', '[', ']', '-'}

func main() {
	clusters := make([]*Cluster, 0)
	delimiters := make(map[rune]bool)
	for _, r := range dels {
		delimiters[r] = true
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		count := 0
		nwCount := 0
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			count++
			in := scanner.Text()
			words := strings.FieldsFunc(in, func(r rune) bool {
				ok, _ := delimiters[r]
				return ok
			})
			maxScore := intsets.MinInt
			//var minZ []byte
			//var minW []byte
			ratio := 0.0
			var oldMatch *Cluster
			for _, old := range clusters {
				//fmt.Printf("%#v\n %#v\n", words, old.OriginalLine)
				_, _, score := sequencing.NeedlemanWunschReuseWords(words, old.OriginalLine, -1, similarityStrings)
				nwCount++
				if score > maxScore {
					maxScore = score
					//minW = w
					//minZ = z
					oldMatch = old
					ratio = float64(maxScore) / float64(len(words))
					//fmt.Printf("score %d, ratio %f", maxScore, ratio)
					if ratio > 0.5 {
						oldMatch.Matches = append(oldMatch.Matches,
							&Match{maxScore, float64(maxScore) / float64(len(in)), string(in)})
						break
					}
				}
			}

			if ratio < 0.5 {
				// New cluster
				clusters = append(clusters, NewCluster(words))
				fmt.Printf("%d clusters, %d lines read, %d NW\n", len(clusters), count, nwCount)
			}

		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for range signals {
		// Write plain text
		f, err := os.OpenFile("clusters.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}
		for _, cluster := range clusters {
			for _, word := range cluster.OriginalLine {
				f.WriteString(word)
			}
			f.WriteString("\n")
			for _, m := range cluster.Matches {
				f.WriteString(m.Line)
				f.WriteString("\n")
			}
			f.WriteString("\n")
		}
		err = f.Close()
		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}

		// Write as JSON
		js, err := json.Marshal(clusters)

		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}
		f, err = os.OpenFile("clusters.json", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}
		_, err = f.WriteString(string(js))
		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}

		err = f.Close()
		if err != nil {
			panic(fmt.Sprintf("err %s", err))
		}
		fmt.Printf("File written. Exiting...")
		os.Exit(0)
	}

}
