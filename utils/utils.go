package utils

import "log"

func CheckErr(err error, messages ...string) {
	if err != nil {
		if len(messages) > 0 {
			for _, msg := range messages {
				log.Fatalln(msg)
			}
		} else {
			log.Fatal(err)
		}
	}
}