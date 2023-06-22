package helpers

import(
	"github.com/ShivamIITK21/cflockout-backend/models"
	"sort"
)

type sortCmp []models.Submission

func(a sortCmp) Len() int {return len(a)}
func (a sortCmp) Swap(i, j int)	{ a[i], a[j] = a[j], a[i] }
func (a sortCmp) Less(i, j int) bool {
	if a[i].ContestId != a[j].ContestId {
		return a[i].ContestId > a[j].ContestId
	}
	return *a[i].Index < *a[j].Index
}

func SortSubmissions(data *[]models.Submission){
	sort.Sort(sortCmp(*data))
}

type sortByTime []models.Submission

func(a sortByTime) Len() int {return len(a)}
func(a sortByTime) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func(a sortByTime) Less(i, j int) bool {
	return a[i].TimeCreated < a[j].TimeCreated
}

func SortSubmissionsByTime(data *[]models.Submission){
	sort.Sort(sortByTime(*data))
}