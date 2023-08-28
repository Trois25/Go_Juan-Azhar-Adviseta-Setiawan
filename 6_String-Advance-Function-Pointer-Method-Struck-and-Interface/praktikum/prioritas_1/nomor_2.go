package main

import "fmt"

type Student struct {
	name string
	score float64
}

func avg(students []Student) {
	var total,average float64
	min,max := students[0].score,students[0].score
	minstudent,maxstudent := 0,0
	for i := 0; i < len(students); i++ {
		total += students[i].score
		if min < students[i].score {
			min = students[i].score
			minstudent = i
		}
		if max > students[i].score {
			max = students[i].score
			maxstudent = i
		}
	}
	average = total/float64(len(students))

	fmt.Printf(" Average Score %g: \n Max Score of Student : %s (%g) \n Min Score of Student : %s (%g)", average,students[minstudent].name,min,students[maxstudent].name,max)
}

func main() {
	students := []Student{
		Student{
			name: "Juan",
			score: 80,
		},
		Student{
			name: "Azhar",
			score: 70,
		},
		Student{
			name: "Adviseta",
			score: 60,
		},
		Student{
			name: "Setiawan",
			score: 90,
		},
		Student{
			name: "wawa",
			score: 50,
		},

	}

	avg(students)

}