package helper

import "log"

func Contains(item string, itemList []string) (int, bool)  {
	for i, n := range itemList {
		if item == n {
			return i, true
		}
	}
	return -1, false
}

func ErrorFatal(err error, message string) {
	if err != nil {
		log.Fatalf(message, err.Error())
	}
}

func ErrorPrintln(err error, message string) error {
	if err != nil {
		log.Println(message, err.Error())
		return err
	}
	return nil
}

