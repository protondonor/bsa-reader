package classes

var Classes = []string{
	"Mage", "Spellsword", "Battlemage",
	"Sorceror", "Healer", "Nightblade",
	"Bard", "Burglar", "Rogue",
	"Acrobat", "Thief", "Assassin",
	"Monk", "Archer", "Ranger",
	"Barbarian", "Warrior", "Knight",
}

var Skills = []string{
	"Medical", "Etiquette", "Streetwise", "Jumping",
	"Orcish", "Harpy", "Giantish", "Dragonish",
	"Nymph", "Daedric", "Spriggan", "Centaurish",
	"Impish", "Lockpicking", "Mercantile", "Pickpocket",
	"Stealth", "Swimming", "Climbing", "Backstabbing",
	"Dodging", "Running", "Destruction", "Restoration",
	"Illusion", "Alteration", "Thaumaturgy", "Mysticism",
	"Short Blade", "Long Blade", "Hand-to-Hand", "Axe",
	"Blunt Weapon", "Archery", "Critical Strike",
}

type Class struct {
	// name starts at position 28 and seems to end at 51
	// which seems to be a fixed length field rather than a terminated string
	// since they all seem to be 74 bytes long
	Name string

	// starting at 16 seems to be a 3-byte array of primary skills,
	// followed by a 3-byte array of major skills
	// followed by a 6-byte array of minor skills
	// presented in the same order that they are presented in the UI
	// see https://en.uesp.net/wiki/Daggerfall_Mod:BIOG%3F%3FT0.TXT_Files/Skill_Codes
	Skills struct {
		Primary [3]string
		Major   [3]string
		Minor   [6]string
	}
	// position 52 seems to hold max HP per level
	HP uint8

	// Position 5 seems to be spell points; the *higher* this value is the *lower* the SP/INT ratio is:
	//   - 0x14 (20): 0.5
	//   - 0x10 (16): 1
	//   - 0x0C (12): 1.5
	//   - 0x08 (8): 1.75
	//   - 0x04 (4): 2
	//   - 0x00 (0): 3
	// i.e. x -> 3 - x/8 except for 2 and 1.75
	// I've decided to store this multiplied by 4, so that we don't have to use floats.
	SPMultiplier uint8

	// base attributes starting at 58, perplexingly with 0s in between
	Attributes struct {
		Strength     uint8
		Intelligence uint8
		Willpower    uint8
		Agility      uint8
		Endurance    uint8
		Personality  uint8
		Speed        uint8
		Luck         uint8
	}
}

func ReadClass(classCfg []byte) Class {
	class := Class{}

	class.SPMultiplier = getSPMultiplier(classCfg[5])
	class.HP = classCfg[52]

	rawName := classCfg[28:52]
	nameLen := 52 - 28
	for i := 0; i < len(rawName); i++ {
		if (rawName[i]) == 0 {
			nameLen = i
			break
		}
	}
	class.Name = string(rawName[:nameLen])

	class.Attributes.Strength = classCfg[58]
	class.Attributes.Intelligence = classCfg[60]
	class.Attributes.Willpower = classCfg[62]
	class.Attributes.Agility = classCfg[64]
	class.Attributes.Endurance = classCfg[66]
	class.Attributes.Personality = classCfg[68]
	class.Attributes.Speed = classCfg[70]
	class.Attributes.Luck = classCfg[72]

	class.Skills.Primary = [3]string{Skills[classCfg[16]], Skills[classCfg[17]], Skills[classCfg[18]]}
	class.Skills.Major = [3]string{Skills[classCfg[19]], Skills[classCfg[20]], Skills[classCfg[21]]}
	class.Skills.Minor = [6]string{
		Skills[classCfg[22]], Skills[classCfg[23]], Skills[classCfg[24]],
		Skills[classCfg[25]], Skills[classCfg[26]], Skills[classCfg[27]],
	}

	return class
}

func getSPMultiplier(b byte) uint8 {
	if b == 8 {
		return 7 // 4*1.75
	}
	if b == 4 {
		return 4 * 2
	}
	return 12 - b/2
}

// [0, 14]
// - 0 must contain resistances. Only the Monk has this set.
// - 1 may contain immunities. Barbarian and Knight have this set, at different bits (for poison and paralysis)
// - 2, 3, 7, 11 may be used for advantages and/or disadvantages that aren't used by any of the default classes. They are all 0.
// - 4: athleticism and adrenaline rush (2s and 4s position) and either spell absorption or no regen magicka (8s position)
// - 6: rapid healing at the 4s position
// - 8: the 8s position is set for ranger but no one else. I don't think ranger has any special advantages though.
// - 9: the 4s position is set for sorcerer, so it is either spell absorption or no regen magicka
// - 10: the 4s position is bonus to hit humanoids. Other bonuses to hit may live here.
// - 12: material limitations. 1s is Daedric, 2s is Orcish
// - 13: weapon expertise. 32s position is expertise in archery
// [15, 16] seems to be weapon/armor/shield proficiencies
// [53, 56]: this seems to be level speed multiplier, although I don't know what floating point format is used.
// 57 is always 0.
