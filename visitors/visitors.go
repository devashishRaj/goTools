package visitors

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func TopTenVisitors() {
	f, err := os.Open("log.txt")
	if err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	input := bufio.NewScanner(f)
	uniques := map[string]int{}
	for input.Scan() {
		fields := strings.Fields(input.Text())
		if len(fields) > 0 {
			uniques[fields[0]]++
		}
	}
	type freq struct {
		addr  string
		count int
	}
	freqs := make([]freq, 0, len(uniques))
	for addr, count := range uniques {
		freqs = append(freqs, freq{addr, count})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})
	fmt.Printf("%-16s%s\n", "Address", "Requests")
	for i, freq := range freqs {
		if i > 9 {
			break
		}
		fmt.Printf("%-16s%d\n", freq.addr, freq.count)
	}
}
