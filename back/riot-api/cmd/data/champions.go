package data

type Champion struct {
	Name       string
	Roles      []string
	ChampionID int
}

func GetChampions() []Champion {
	champions := []Champion{
		{ChampionID: 266, Name: "Aatrox", Roles: []string{"top"}},
		{ChampionID: 103, Name: "Ahri", Roles: []string{"mid"}},
		{ChampionID: 84, Name: "Akali", Roles: []string{"top", "mid"}},
		{ChampionID: 166, Name: "Akshan", Roles: []string{"mid", "adc"}},
		{ChampionID: 12, Name: "Alistar", Roles: []string{"support"}},
		{ChampionID: 32, Name: "Amumu", Roles: []string{"jungle", "support"}},
		{ChampionID: 34, Name: "Anivia", Roles: []string{"mid"}},
		{ChampionID: 1, Name: "Annie", Roles: []string{"mid"}},
		{ChampionID: 22, Name: "Ashe", Roles: []string{"adc"}},
		{ChampionID: 136, Name: "Aurelion Sol", Roles: []string{"mid"}},
		{ChampionID: 268, Name: "Azir", Roles: []string{"mid"}},
		{ChampionID: 432, Name: "Bard", Roles: []string{"support"}},
		{ChampionID: 200, Name: "Bel'Veth", Roles: []string{"jungle"}},
		{ChampionID: 53, Name: "Blitzcrank", Roles: []string{"support"}},
		{ChampionID: 63, Name: "Brand", Roles: []string{"support", "mid"}},
		{ChampionID: 201, Name: "Braum", Roles: []string{"support"}},
		{ChampionID: 233, Name: "Briar", Roles: []string{"jungle", "top"}},
		{ChampionID: 51, Name: "Caitlyn", Roles: []string{"adc"}},
		{ChampionID: 164, Name: "Camille", Roles: []string{"top"}},
		{ChampionID: 69, Name: "Cassiopeia", Roles: []string{"mid", "top"}},
		{ChampionID: 31, Name: "Cho'Gath", Roles: []string{"top", "mid"}},
		{ChampionID: 42, Name: "Corki", Roles: []string{"mid", "adc"}},
		{ChampionID: 122, Name: "Darius", Roles: []string{"top"}},
		{ChampionID: 131, Name: "Diana", Roles: []string{"mid", "jungle"}},
		{ChampionID: 36, Name: "Dr. Mundo", Roles: []string{"top", "jungle"}},
		{ChampionID: 119, Name: "Draven", Roles: []string{"adc"}},
		{ChampionID: 245, Name: "Ekko", Roles: []string{"mid", "jungle"}},
		{ChampionID: 60, Name: "Elise", Roles: []string{"jungle"}},
		{ChampionID: 28, Name: "Evelynn", Roles: []string{"jungle"}},
		{ChampionID: 81, Name: "Ezreal", Roles: []string{"adc"}},
		{ChampionID: 9, Name: "Fiddlesticks", Roles: []string{"jungle"}},
		{ChampionID: 114, Name: "Fiora", Roles: []string{"top"}},
		{ChampionID: 105, Name: "Fizz", Roles: []string{"mid"}},
		{ChampionID: 3, Name: "Galio", Roles: []string{"mid", "support"}},
		{ChampionID: 41, Name: "Gangplank", Roles: []string{"top"}},
		{ChampionID: 86, Name: "Garen", Roles: []string{"top"}},
		{ChampionID: 150, Name: "Gnar", Roles: []string{"top"}},
		{ChampionID: 79, Name: "Gragas", Roles: []string{"top", "jungle"}},
		{ChampionID: 104, Name: "Graves", Roles: []string{"jungle", "top"}},
		{ChampionID: 887, Name: "Gwen", Roles: []string{"top"}},
		{ChampionID: 120, Name: "Hecarim", Roles: []string{"jungle"}},
		{ChampionID: 74, Name: "Heimerdinger", Roles: []string{"mid", "top"}},
		{ChampionID: 420, Name: "Illaoi", Roles: []string{"top"}},
		{ChampionID: 39, Name: "Irelia", Roles: []string{"top", "mid"}},
		{ChampionID: 427, Name: "Ivern", Roles: []string{"jungle"}},
		{ChampionID: 40, Name: "Janna", Roles: []string{"support"}},
		{ChampionID: 59, Name: "Jarvan IV", Roles: []string{"jungle"}},
		{ChampionID: 24, Name: "Jax", Roles: []string{"top", "jungle"}},
		{ChampionID: 126, Name: "Jayce", Roles: []string{"top", "mid"}},
		{ChampionID: 202, Name: "Jhin", Roles: []string{"adc"}},
		{ChampionID: 222, Name: "Jinx", Roles: []string{"adc"}},
		{ChampionID: 145, Name: "Kai'Sa", Roles: []string{"adc"}},
		{ChampionID: 429, Name: "Kalista", Roles: []string{"adc"}},
		{ChampionID: 43, Name: "Karma", Roles: []string{"support", "mid"}},
		{ChampionID: 30, Name: "Karthus", Roles: []string{"jungle", "mid"}},
		{ChampionID: 38, Name: "Kassadin", Roles: []string{"mid"}},
		{ChampionID: 55, Name: "Katarina", Roles: []string{"mid"}},
		{ChampionID: 10, Name: "Kayle", Roles: []string{"top"}},
		{ChampionID: 141, Name: "Kayn", Roles: []string{"jungle"}},
		{ChampionID: 85, Name: "Kennen", Roles: []string{"top"}},
		{ChampionID: 121, Name: "Kha'Zix", Roles: []string{"jungle"}},
		{ChampionID: 203, Name: "Kindred", Roles: []string{"jungle"}},
		{ChampionID: 240, Name: "Kled", Roles: []string{"top"}},
		{ChampionID: 96, Name: "Kog'Maw", Roles: []string{"adc"}},
		{ChampionID: 7, Name: "LeBlanc", Roles: []string{"mid"}},
		{ChampionID: 64, Name: "Lee Sin", Roles: []string{"jungle"}},
		{ChampionID: 89, Name: "Leona", Roles: []string{"support"}},
		{ChampionID: 876, Name: "Lillia", Roles: []string{"jungle"}},
		{ChampionID: 127, Name: "Lissandra", Roles: []string{"mid"}},
		{ChampionID: 236, Name: "Lucian", Roles: []string{"adc", "mid"}},
		{ChampionID: 117, Name: "Lulu", Roles: []string{"support"}},
		{ChampionID: 99, Name: "Lux", Roles: []string{"support", "mid"}},
		{ChampionID: 54, Name: "Malphite", Roles: []string{"top", "support"}},
		{ChampionID: 90, Name: "Malzahar", Roles: []string{"mid"}},
		{ChampionID: 57, Name: "Maokai", Roles: []string{"support", "top"}},
		{ChampionID: 11, Name: "Master Yi", Roles: []string{"jungle"}},
		{ChampionID: 21, Name: "Miss Fortune", Roles: []string{"adc"}},
		{ChampionID: 82, Name: "Mordekaiser", Roles: []string{"top"}},
		{ChampionID: 25, Name: "Morgana", Roles: []string{"support", "mid"}},
		{ChampionID: 267, Name: "Nami", Roles: []string{"support"}},
		{ChampionID: 75, Name: "Nasus", Roles: []string{"top"}},
		{ChampionID: 111, Name: "Nautilus", Roles: []string{"support"}},
		{ChampionID: 518, Name: "Neeko", Roles: []string{"mid", "support"}},
		{ChampionID: 76, Name: "Nidalee", Roles: []string{"jungle"}},
		{ChampionID: 56, Name: "Nocturne", Roles: []string{"jungle", "mid"}},
		{ChampionID: 20, Name: "Nunu & Willump", Roles: []string{"jungle"}},
		{ChampionID: 2, Name: "Olaf", Roles: []string{"jungle", "top"}},
		{ChampionID: 61, Name: "Orianna", Roles: []string{"mid"}},
		{ChampionID: 516, Name: "Ornn", Roles: []string{"top"}},
		{ChampionID: 80, Name: "Pantheon", Roles: []string{"top", "support"}},
		{ChampionID: 78, Name: "Poppy", Roles: []string{"top", "support"}},
		{ChampionID: 555, Name: "Pyke", Roles: []string{"support"}},
		{ChampionID: 246, Name: "Qiyana", Roles: []string{"mid"}},
		{ChampionID: 133, Name: "Quinn", Roles: []string{"top"}},
		{ChampionID: 497, Name: "Rakan", Roles: []string{"support"}},
		{ChampionID: 33, Name: "Rammus", Roles: []string{"jungle"}},
		{ChampionID: 421, Name: "Rek'Sai", Roles: []string{"jungle"}},
		{ChampionID: 526, Name: "Rell", Roles: []string{"support"}},
		{ChampionID: 58, Name: "Renekton", Roles: []string{"top"}},
		{ChampionID: 107, Name: "Rengar", Roles: []string{"jungle", "top"}},
		{ChampionID: 92, Name: "Riven", Roles: []string{"top"}},
		{ChampionID: 68, Name: "Rumble", Roles: []string{"top", "mid"}},
		{ChampionID: 13, Name: "Ryze", Roles: []string{"mid"}},
		{ChampionID: 360, Name: "Samira", Roles: []string{"adc"}},
		{ChampionID: 113, Name: "Sejuani", Roles: []string{"jungle"}},
		{ChampionID: 235, Name: "Senna", Roles: []string{"adc", "support"}},
		{ChampionID: 147, Name: "Seraphine", Roles: []string{"mid", "support"}},
		{ChampionID: 875, Name: "Sett", Roles: []string{"top", "support"}},
		{ChampionID: 35, Name: "Shaco", Roles: []string{"jungle"}},
		{ChampionID: 98, Name: "Shen", Roles: []string{"top"}},
		{ChampionID: 102, Name: "Shyvana", Roles: []string{"jungle"}},
		{ChampionID: 27, Name: "Singed", Roles: []string{"top"}},
		{ChampionID: 14, Name: "Sion", Roles: []string{"top"}},
		{ChampionID: 15, Name: "Sivir", Roles: []string{"adc"}},
		{ChampionID: 72, Name: "Skarner", Roles: []string{"jungle"}},
		{ChampionID: 37, Name: "Sona", Roles: []string{"support"}},
		{ChampionID: 16, Name: "Soraka", Roles: []string{"support"}},
		{ChampionID: 50, Name: "Swain", Roles: []string{"support", "mid"}},
		{ChampionID: 517, Name: "Sylas", Roles: []string{"mid", "top"}},
		{ChampionID: 134, Name: "Syndra", Roles: []string{"mid"}},
		{ChampionID: 223, Name: "Tahm Kench", Roles: []string{"support", "top"}},
		{ChampionID: 163, Name: "Taliyah", Roles: []string{"mid", "jungle"}},
		{ChampionID: 91, Name: "Talon", Roles: []string{"mid"}},
		{ChampionID: 44, Name: "Taric", Roles: []string{"support"}},
		{ChampionID: 17, Name: "Teemo", Roles: []string{"top"}},
		{ChampionID: 412, Name: "Thresh", Roles: []string{"support"}},
		{ChampionID: 18, Name: "Tristana", Roles: []string{"adc"}},
		{ChampionID: 48, Name: "Trundle", Roles: []string{"jungle", "top"}},
		{ChampionID: 23, Name: "Tryndamere", Roles: []string{"top"}},
		{ChampionID: 4, Name: "Twisted Fate", Roles: []string{"mid"}},
		{ChampionID: 29, Name: "Twitch", Roles: []string{"adc", "jungle"}},
		{ChampionID: 77, Name: "Udyr", Roles: []string{"jungle"}},
		{ChampionID: 6, Name: "Urgot", Roles: []string{"top"}},
		{ChampionID: 110, Name: "Varus", Roles: []string{"adc"}},
		{ChampionID: 67, Name: "Vayne", Roles: []string{"adc", "top"}},
		{ChampionID: 161, Name: "Vel'Koz", Roles: []string{"mid", "support"}},
		{ChampionID: 254, Name: "Vi", Roles: []string{"jungle"}},
		{ChampionID: 234, Name: "Viego", Roles: []string{"jungle"}},
		{ChampionID: 112, Name: "Viktor", Roles: []string{"mid"}},
		{ChampionID: 8, Name: "Vladimir", Roles: []string{"mid", "top"}},
		{ChampionID: 106, Name: "Volibear", Roles: []string{"top", "jungle"}},
		{ChampionID: 19, Name: "Warwick", Roles: []string{"jungle"}},
		{ChampionID: 62, Name: "Wukong", Roles: []string{"top", "jungle"}},
		{ChampionID: 498, Name: "Xayah", Roles: []string{"adc"}},
		{ChampionID: 101, Name: "Xerath", Roles: []string{"mid", "support"}},
		{ChampionID: 5, Name: "Xin Zhao", Roles: []string{"jungle"}},
		{ChampionID: 157, Name: "Yasuo", Roles: []string{"mid", "top"}},
		{ChampionID: 777, Name: "Yone", Roles: []string{"mid", "top"}},
		{ChampionID: 83, Name: "Yorick", Roles: []string{"top"}},
		{ChampionID: 350, Name: "Yuumi", Roles: []string{"support"}},
		{ChampionID: 154, Name: "Zac", Roles: []string{"jungle"}},
		{ChampionID: 238, Name: "Zed", Roles: []string{"mid"}},
		{ChampionID: 115, Name: "Ziggs", Roles: []string{"mid"}},
		{ChampionID: 26, Name: "Zilean", Roles: []string{"support", "mid"}},
		{ChampionID: 142, Name: "Zoe", Roles: []string{"mid"}},
		{ChampionID: 143, Name: "Zyra", Roles: []string{"support", "mid"}},
	}
	return champions
}
