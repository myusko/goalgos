package goalgos

import (
	"sort"
	"strconv"
	"strings"
)

func getClosest(times []int) int {
	valueOne := times[0]
	valueTwo := times[1]

	minimumDiff := valueTwo - valueOne

	for i := 1; i < len(times)-1; i++ {
		if times[i+1]-times[i] < minimumDiff {
			valueOne = times[i]
			valueTwo = times[i+1]
			minimumDiff = valueTwo - valueOne
		}
	}

	result := valueTwo - valueOne

	if 0 > result {
		return -result
	}

	return result
}

func TimeDifference(n int, times string) int {
	splitedTimes := strings.Split(times, " ")
	convertedTimes := make([]int, 0, cap(splitedTimes))

	for i := 0; i < n; i++ {
		splitedHM := strings.Split(splitedTimes[i], ":")
		convertedH, err := strconv.Atoi(splitedHM[0])
		if err != nil {
			return -1
		}

		convertedM, err := strconv.Atoi(splitedHM[1])
		if err != nil {
			return -1
		}

		if convertedH == 0 {
			convertedH = 24
		}

		convertedTimes = append(convertedTimes, (convertedH*60)+convertedM)
	}

	sort.Slice(convertedTimes, func(i, j int) bool {
		return convertedTimes[i] < convertedTimes[j]
	})

	getClosest := getClosest(convertedTimes)

	return getClosest
}
