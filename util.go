package gobless

import "fmt"

func shortenInt(number int) string {

	switch true {
	case number > 1000000000000:
		return fmt.Sprintf("%1fT", float64(number)/1000000000000.0)
	case number > 1000000000:
		return fmt.Sprintf("%1fG", float64(number)/1000000000.0)
	case number > 1000000:
		return fmt.Sprintf("%1fM", float64(number)/1000000.0)
	case number > 1000:
		return fmt.Sprintf("%1fK", float64(number)/1000.0)
	}

	return fmt.Sprintf("%d", number)
}
