package demo

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rrkas/ps_practice/models"
	"github.com/rrkas/ps_practice/utils"
)

func InsertDemoQuestions(start int) {
	n := rand.Intn(5) + 5
	log.Printf("%v questions", n)
	for i := start; i < n+start; i++ {
		q := models.Question{
			Title:     fmt.Sprintf("Title %v", i),
			Statement: fmt.Sprintf("Statement %v", i),
			InputFormat: `
1st line contains an int which represents number of records 'n'
Next n lines contain a number, ith line represents a[i]`,
			OutputFormat: `
An int which is sum of all a[i]
`,
			Level: int64(rand.Intn(models.QuestionHard + 1)),
			SampleIOs: []models.IO{
				{
					Input:  "3\n1\n2\n3",
					Output: "6",
				},
			},
			DateTime: time.Now(),
		}
		q.AddInDB(utils.DB)
	}
}
