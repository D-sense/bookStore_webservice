package helper

import "log"

func LogFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
