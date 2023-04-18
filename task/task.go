package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
)

var blackWords []string
var words []string

func Init() {

	content, err := ioutil.ReadFile("task/blackWords.txt")
	if err != nil {
		log.Panic(err)
	}
	blackWords = strings.Split(string(content), "\n")

	content, err = ioutil.ReadFile("task/words.txt")
	if err != nil {
		log.Panic(err)
	}
	words = strings.Split(string(content), "\n")
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getThreeNumbers(min, max int) [3]int {
	set := make(map[int]bool)

	for len(set) != 3 {
		set[randInt(min, max)] = true
	}

	arr := [3]int{}

	i := 0
	for num := range set {
		arr[i] = num
		i++
	}

	return arr
}

func GetTask() (model.Task, string) {

	rand.Seed(time.Now().UnixNano())

	var variants [4]string

	variants[0] = words[randInt(0, len(words))]
	answer := variants[0]

	numbers := getThreeNumbers(0, len(blackWords))

	for i := 0; i < 3; i++ {
		variants[i+1] = blackWords[numbers[i]]
	}

	rand.Shuffle(len(variants), func(i, j int) {
		variants[i], variants[j] = variants[j], variants[i]
	})

	i := 0
	for variants[i] != answer {
		i++
	}
	task := model.Task{Answer: fmt.Sprint(i + 1), Variants: variants}
	message := "✍️ Укажите вариант ответа, в которых <b>верно</b> выделена буква, обозначающая ударный гласный звук.\n"
	for i, v := range task.Variants {
		message += fmt.Sprintf("%d) %s\n", i+1, v)
	}
	return task, message
}
