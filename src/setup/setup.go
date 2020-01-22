// Create a set of random logs with varied timestamps.
// Format: [timestamp] - [<request method>] - [<gibberish>]

package setup

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Length of log files
const maxLength = 1000
const minLength = 500

// Maximum dt between logs
const maxTimeInterval = 400

// Length of generated text
const lengthText = 5

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Setup generates a log file with given name filename in the current directory.
func Setup(filename string) {
	nextLogTime := time.Now().UTC()
	length := rand.Intn(maxLength-minLength+1) + minLength
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < length; i++ {
		_, err = f.WriteString(generateLog(nextLogTime))
		nextLogTime = nextLogTime.Add(time.Second * time.Duration(rand.Intn(maxTimeInterval+1)))
	}
}

func generateLog(timestamp time.Time) string {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	gibberish := []string{"second", "greater", "belt", "share", "jack", "east",
		"seldom", "drawn", "dot", "orange", "our", "according",
		"exactly", "snow", "include", "state", "length", "private",
		"atomic", "letter", "color", "minute", "throat", "desert",
		"went", "organized", "design", "design", "straight", "bigger",
		"cabin", "simplest", "battle", "today", "drove", "terrible",
		"gravity", "firm", "someone", "gulf", "caught", "warm"}
	method := methods[rand.Int()%len(methods)]
	text := ""

	for i := 0; i < lengthText; i++ {
		text += gibberish[rand.Int()%len(gibberish)] + " "
	}

	return fmt.Sprintf(
		"[%s] - [%s]: \"%s\"\n",
		timestamp.Format(time.RFC3339),
		method,
		text,
	)
}
