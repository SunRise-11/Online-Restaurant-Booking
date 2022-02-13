package helpers

import (
	"Restobook/delivery/common"
	"fmt"
)

func DaytoNumberConverter(openDay, closeDay []string) (string, string, error) {
	openInt := ""
	closeInt := ""
	for j := 0; j < len(openDay); j++ {
		for k := 0; k < len(common.Daytoint); k++ {
			if openDay[j] == common.Daytoint[k].Day {
				openInt += fmt.Sprintf("%v,", common.Daytoint[k].No)
			}
		}
	}

	for j := 0; j < len(closeDay); j++ {
		for k := 0; k < len(common.Daytoint); k++ {
			if closeDay[j] == common.Daytoint[k].Day {
				closeInt += fmt.Sprintf("%v,", common.Daytoint[k].No)
			}
		}
	}
	return openInt, closeInt, nil
}
