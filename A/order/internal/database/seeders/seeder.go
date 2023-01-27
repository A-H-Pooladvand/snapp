package seeders

var Seeders = []Seeder{}

func Run() {
	for _, seeder := range Seeders {
		seeder.Run()
	}
}
