package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/psykhi/sequencing"
	"golang.org/x/tools/container/intsets"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func similarity(b byte, b2 byte) int {
	if b == b2 {
		return 1
	}
	return -1
}

type Cluster struct {
	OriginalLine string
	Matches      []*Match
}

type Match struct {
	Score int
	Ratio float64
	Line  string
}

func NewCluster(line []byte) *Cluster {
	matches := make([]*Match, 0)
	return &Cluster{
		string(line),
		matches,
	}
}

func main() {
	clusters := make([]*Cluster, 0)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		count := 0
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			count++
			in := scanner.Bytes()
			maxScore := intsets.MinInt
			//var minZ []byte
			//var minW []byte
			ratio := 0.0
			var oldMatch *Cluster
			for _, old := range clusters {
				_, _, score := sequencing.NeedlemanWunsch(in, []byte(old.OriginalLine), -1, similarity)
				if score > maxScore {
					maxScore = score
					//minW = w
					//minZ = z
					oldMatch = old
					ratio = float64(maxScore) / float64(len(in))
					if ratio > 0.7 {
						oldMatch.Matches = append(oldMatch.Matches,
							&Match{maxScore, float64(maxScore) / float64(len(in)), string(in)})
						break
					}
				}
			}

			if ratio < 0.7 {
				// New cluster
				clusters = append(clusters, NewCluster(in))
				fmt.Printf("%d clusters, %d lines read\n", len(clusters), count)
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
			f.WriteString(cluster.OriginalLine)
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
