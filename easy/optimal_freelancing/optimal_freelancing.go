package optimalfreelancing

import (
	"sort"
)

type Job struct {
	time    int
	paymant int
}

func OptimalWeek(jobs []Job) int {
	if jobs == nil || len(jobs) == 0 {
		return 0
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].paymant > jobs[j].paymant
	})

	sch := make([]bool, 7)
	total := 0

	for i := range jobs {
		j := 6
		if jobs[i].time-1 < j {
			j = jobs[i].time - 1
		}

		for ; j >= 0; j-- {

			if !sch[j] {
				sch[j] = true
				total += jobs[i].paymant

				break
			}
		}
	}

	return total
}
