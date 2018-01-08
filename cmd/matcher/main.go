package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/psykhi/sequencing"
	"github.com/satori/go.uuid"
	"golang.org/x/tools/container/intsets"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

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
	delimiters := make(map[rune]int)
	for i, r := range dels {
		delimiters[r] = i
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		threshold := 0.3
		count := 0
		nwCount := 0
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			count++
			in := scanner.Text()
			vec := make([]int, len(dels))

			for i := 0; i < len(in); i++ {
				pos, ok := delimiters[rune(in[i])]
				if ok {
					vec[pos]++
				}
			}
			words := strings.FieldsFunc(in, func(r rune) bool {
				_, ok := delimiters[r]
				return ok
			})
			minDistance := intsets.MaxInt
			//var minZ []byte
			//var minW []byte
			ratio := math.MaxFloat64
			var oldMatch *Cluster
			for _, old := range clusters {
				//fmt.Printf("%#v\n %#v\n", words, old.OriginalLine)
				distance := sequencing.LevenshteinDistanceStrings(words, old.OriginalLine, nil, nil)
				nwCount++
				if distance < minDistance {
					minDistance = distance
					//minW = w
					//minZ = z
					oldMatch = old
					ratio = float64(minDistance) / float64(len(words))
					//fmt.Printf("score %d, ratio %f", minDistance, ratio)
					if ratio < threshold {
						oldMatch.Matches = append(oldMatch.Matches,
							&Match{minDistance, float64(minDistance) / float64(len(in)), string(in)})
						break
					}
				}
			}

			if ratio > threshold {
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
