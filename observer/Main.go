package main

func main() {
	sheSaidNo := Person{"Azhar"}
	svyazi := JobSite{}

	svyazi.addVacancy("Jaksi zhumys, ote persprektivnyi")

	svyazi.subscribe(sheSaidNo)

	svyazi.addVacancy("birinshi zhumys ushin norm")
	svyazi.removeVacancy("birinshi zhumys ushin norm")

}
