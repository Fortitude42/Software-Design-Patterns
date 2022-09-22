package main

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (j *JobSite) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) removeVacancy(vacancy string) {
	for i, v := range j.vacancies {
		if vacancy == v {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
			break
		}
	}
	j.sendAll()
}

func (j *JobSite) subscribe(observer Observer) {
	j.subscribers = append(j.subscribers, observer)
}

func (j JobSite) sendAll() {
	for _, subscriber := range j.subscribers {
		subscriber.handleEvent(j.vacancies)
	}
}
