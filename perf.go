package baselayergo

import (
	f "fmt"
	"slices"
	"time"
)

//----------------------------------------
type TimeItResult struct {
	Count int64
	Avg int64
	P50 int64
	P90 int64
	P99 int64
	Values []int64
}

func (r *TimeItResult) PrintWithValues() {
	for i := range(len(r.Values)) {
		f.Printf("%d %d\n", i, r.Values[i])
	}
	r.Print()
}

func (r *TimeItResult) Print() {
	f.Printf("avg: %d\n", r.Avg)
	f.Printf("p50: %d\n", r.P50)
	f.Printf("p90: %d\n", r.P90)
	f.Printf("p99: %d\n", r.P99)
}

//----------------------------------------
func TimeItSingle(f func()) int64{
	startTime := time.Now().UnixNano()
	f()
	endTime := time.Now().UnixNano()
	totalTime := endTime - startTime
	return totalTime	
}

//----------------------------------------
func TimeItMany(count int64, f func()) *TimeItResult {
	allTimes := make([]int64, count)

	var sum int64
	for i := range count {
		singleRunTime := TimeItSingle(f)
		sum += singleRunTime
		allTimes[i] = singleRunTime
	}

	avgTime := sum / count

	slices.Sort(allTimes)

	p50Index := int64(float64(count) * 0.5)
	p90Index := int64(float64(count) * 0.9)
	p99Index := int64(float64(count) * 0.99)

	return &TimeItResult {
		count,
		avgTime,
		allTimes[p50Index],
		allTimes[p90Index],
		allTimes[p99Index],
		allTimes,
	}
}