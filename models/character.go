package models

type AdditionalParams struct {
	MaxHP                int
	MaxMP                int
	Power                int
	Vitality             int
	Agility              int
	Weight               int
	Intelligence         int
	Spirit               int
	Attack               int
	Defense              int
	AbilityDefence       int
	AbilityEvasionRate   int
	Magic                int
	Luck                 int
	AccuracyRate         int
	EvasionRate          int
	AbilityDisturbedRate int
	CriticalRate         int
	DamageDirmeter       int
	AbilityDefenseRate   int
	AccuracyCount        int
	EvasionCount         int
	DefenseCount         int
	MagicDefenseCount    int
}

type Character struct {
	AdditionalParams
	ID        int
	JobID     int
	IsEnabled bool
	RootName  string
	Name      string
	Level     int
	Exp       int
	CurrentHP int
	CurrentMP int
}

var (
	Characters     []*Character
	CharacterNames []string
)

func GetCharacter(id int) (c *Character, found bool) {
	for _, c = range Characters {
		if c.ID == id {
			found = true
			break
		}
	}
	return
}

func GetCharacterByName(name string) (c *Character, found bool) {
	for _, c = range Characters {
		if c.Name == name {
			found = true
			break
		}
	}
	return
}

func UpdateCharacterName(old, new string) {
	for i, n := range CharacterNames {
		if n == old {
			CharacterNames[i] = new
			break
		}
	}
}
