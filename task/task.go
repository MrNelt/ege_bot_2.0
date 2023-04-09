package task

import (
	"math/rand"
	"time"
)

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

func GetTask(blackWords []string, words []string) ([4]string, int) {

	rand.Seed(time.Now().UnixNano())

	var task [4]string

	task[0] = words[randInt(0, len(words))]
	answer := task[0]

	numbers := getThreeNumbers(0, len(blackWords))

	for i := 0; i < 3; i++ {
		task[i+1] = blackWords[numbers[i]]
	}

	rand.Shuffle(len(task), func(i, j int) {
		task[i], task[j] = task[j], task[i]
	})

	i := 0
	for task[i] != answer {
		i++
	}

	return task, i + 1
}

// func main() {

// 	content, err := ioutil.ReadFile("blackWords.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	blackWords := strings.Split(string(content), "\n")

// 	content, err = ioutil.ReadFile("words.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	words := strings.Split(string(content), "\n")

// 	task, answer := GetTask(blackWords, words)

// 	fmt.Println(task, answer)

// }
