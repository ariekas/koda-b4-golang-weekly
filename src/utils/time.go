package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func DefualtTime(key string, defualValue string) string {
	godotenv.Load(".env")
	val, exitis := os.LookupEnv(key)
	if !exitis {
		return  defualValue
	}
	return  val
}
func Time() int {
	time, _ := strconv.Atoi(DefualtTime("timeLimit", "15"))
	return time
}
