package main

import (
	"sort"
)

type TaskDetils struct {
	key                   int
	minimumEngergyToStart int
	actualEngergyNeeded   int
	minAndActualDiff      int
}

func minimumEffort(tasks [][]int) int {
	n := len(tasks)
	taskDetails := make([]TaskDetils, 0, n)

	for i, task := range tasks {
		taskDetail := TaskDetils{
			key:                   i,
			minimumEngergyToStart: task[1],
			actualEngergyNeeded:   task[0],
			minAndActualDiff:      task[1] - task[0],
		}
		taskDetails = append(taskDetails, taskDetail)
	}

	sort.Slice(taskDetails, func(i, j int) bool {
		return taskDetails[i].minAndActualDiff > taskDetails[j].minAndActualDiff
	})

	sort.Slice(taskDetails, func(i, j int) bool {
		return taskDetails[i].minimumEngergyToStart > taskDetails[j].minimumEngergyToStart && taskDetails[i].minAndActualDiff == taskDetails[j].minAndActualDiff
	})

	minimumEnergyNeeded := taskDetails[0].minimumEngergyToStart
	energyConsume := taskDetails[0].actualEngergyNeeded
	engergyLeft := taskDetails[0].minimumEngergyToStart - taskDetails[0].actualEngergyNeeded
	for i, taskDetail := range taskDetails {
		if i != 0 {
			energyConsume += taskDetail.actualEngergyNeeded
			if engergyLeft < taskDetail.minimumEngergyToStart {
				minimumEnergyNeeded += taskDetail.minimumEngergyToStart - engergyLeft
				engergyLeft = minimumEnergyNeeded - energyConsume
			} else {
				engergyLeft = minimumEnergyNeeded - energyConsume
			}
		}
	}

	return minimumEnergyNeeded
}
