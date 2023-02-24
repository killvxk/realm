package namegen

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

var (
	adjectives = [...]string{
		"absolute",
		"accomplished",
		"accurate",
		"acrobatic",
		"admiring",
		"adoring",
		"adorable",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"ballin",
		"beautiful",
		"beloved",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"celebrated",
		"charming",
		"clever",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"cool",
		"cranky",
		"crazy",
		"dazzling",
		"determined",
		"distracted",
		"dreamy",
		"eager",
		"ecstatic",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"epic",
		"exciting",
		"fervent",
		"festive",
		"flamboyant",
		"focused",
		"friendly",
		"frosty",
		"funny",
		"gallant",
		"gifted",
		"goofy",
		"gracious",
		"great",
		"happy",
		"hardcore",
		"heuristic",
		"hopeful",
		"hungry",
		"infallible",
		"inspiring",
		"intelligent",
		"interesting",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"laughing",
		"loving",
		"lucid",
		"magical",
		"modest",
		"musing",
		"mystifying",
		"naughty",
		"nervous",
		"nice",
		"nifty",
		"nostalgic",
		"objective",
		"optimistic",
		"peaceful",
		"pedantic",
		"pensive",
		"practical",
		"priceless",
		"quantum",
		"quality",
		"quirky",
		"quizzical",
		"rare",
		"recursing",
		"relaxed",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"silly",
		"sleepy",
		"splendiforous",
		"spooky",
		"stoic",
		"strange",
		"stupefied",
		"suspicious",
		"sweet",
		"tender",
		"thirsty",
		"trusting",
		"unruffled",
		"upbeat",
		"unusual",
		"vibrant",
		"vigilant",
		"vigorous",
		"wizardly",
		"wonderful",
		"xenodochial",
		"youthful",
		"zealous",
		"zen",
	}

	nouns = [...]string{
		"aboleth",
		"acolyte",
		"adult-black-dragon",
		"adult-blue-dragon",
		"adult-brass-dragon",
		"adult-bronze-dragon",
		"adult-copper-dragon",
		"adult-gold-dragon",
		"adult-green-dragon",
		"adult-red-dragon",
		"adult-silver-dragon",
		"adult-white-dragon",
		"air-elemental",
		"allosaurus",
		"ancient-black-dragon",
		"ancient-blue-dragon",
		"ancient-brass-dragon",
		"ancient-bronze-dragon",
		"ancient-copper-dragon",
		"ancient-gold-dragon",
		"ancient-green-dragon",
		"ancient-red-dragon",
		"ancient-silver-dragon",
		"ancient-white-dragon",
		"androsphinx",
		"animated-armor",
		"ankheg",
		"ankylosaurus",
		"ape",
		"archmage",
		"assassin",
		"awakened-shrub",
		"awakened-tree",
		"axe-beak",
		"azer",
		"baboon",
		"badger",
		"balor",
		"bandit",
		"bandit-captain",
		"banshee",
		"barbed-devil",
		"basilisk",
		"bat",
		"bearded-devil",
		"behir",
		"berserker",
		"black-bear",
		"black-dragon-wyrmling",
		"black-pudding",
		"blink-dog",
		"blood-hawk",
		"blue-dragon-wyrmling",
		"boar",
		"bone-devil",
		"brass-dragon-wyrmling",
		"bronze-dragon-wyrmling",
		"brown-bear",
		"bugbear",
		"bulette",
		"camel",
		"cat",
		"centaur",
		"chain-devil",
		"chimera",
		"chuul",
		"clay-golem",
		"cloaker",
		"cloud-giant",
		"cockatrice",
		"commoner",
		"constrictor-snake",
		"copper-dragon-wyrmling",
		"couatl",
		"crab",
		"crocodile",
		"cult-fanatic",
		"cultist",
		"cyclops",
		"darkmantle",
		"death-dog",
		"deep-gnome-(svirfneblin)",
		"deer",
		"deva",
		"dire-wolf",
		"djinni",
		"doppelganger",
		"draft-horse",
		"dragon-turtle",
		"dretch",
		"drider",
		"drow",
		"druid",
		"dryad",
		"duergar",
		"dust-mephit",
		"eagle",
		"earth-elemental",
		"efreeti",
		"elephant",
		"elk",
		"erinyes",
		"ettercap",
		"ettin",
		"fire-elemental",
		"fire-giant",
		"flameskull",
		"flesh-golem",
		"flying-snake",
		"flying-sword",
		"frog",
		"frost-giant",
		"gargoyle",
		"gelatinous-cube",
		"ghast",
		"ghost",
		"ghoul",
		"giant-ape",
		"giant-badger",
		"giant-bat",
		"giant-boar",
		"giant-centipede",
		"giant-constrictor-snake",
		"giant-crab",
		"giant-crocodile",
		"giant-eagle",
		"giant-elk",
		"giant-fire-beetle",
		"giant-frog",
		"giant-goat",
		"giant-hyena",
		"giant-lizard",
		"giant-octopus",
		"giant-owl",
		"giant-poisonous-snake",
		"giant-rat",
		"giant-scorpion",
		"giant-sea-horse",
		"giant-shark",
		"giant-spider",
		"giant-toad",
		"giant-vulture",
		"giant-wasp",
		"giant-weasel",
		"giant-wolf-spider",
		"gibbering-mouther",
		"glabrezu",
		"gladiator",
		"gnoll",
		"goat",
		"goblin",
		"gold-dragon-wyrmling",
		"gorgon",
		"gray-ooze",
		"green-dragon-wyrmling",
		"green-hag",
		"grick",
		"griffon",
		"grimlock",
		"guard",
		"guardian-naga",
		"gynosphinx",
		"half-red-dragon-veteran",
		"harpy",
		"hawk",
		"hell-hound",
		"hezrou",
		"hill-giant",
		"hippogriff",
		"hobgoblin",
		"homunculus",
		"horned-devil",
		"hunter-shark",
		"hydra",
		"hyena",
		"ice-devil",
		"ice-mephit",
		"imp",
		"invisible-stalker",
		"iron-golem",
		"jackal",
		"killer-whale",
		"knight",
		"kobold",
		"kraken",
		"lamia",
		"lemure",
		"lich",
		"lion",
		"lizard",
		"lizardfolk",
		"mage",
		"magma-mephit",
		"magmin",
		"mammoth",
		"manticore",
		"marilith",
		"mastiff",
		"medusa",
		"merfolk",
		"merrow",
		"mimic",
		"minotaur",
		"minotaur-skeleton",
		"mule",
		"mummy",
		"mummy-lord",
		"nalfeshnee",
		"night-hag",
		"nightmare",
		"noble",
		"nothic",
		"ochre-jelly",
		"octopus",
		"ogre",
		"ogre-zombie",
		"oni",
		"orc",
		"otyugh",
		"owl",
		"owlbear",
		"panther",
		"pegasus",
		"phase-spider",
		"pit-fiend",
		"planetar",
		"plesiosaurus",
		"poisonous-snake",
		"polar-bear",
		"pony",
		"priest",
		"pseudodragon",
		"pteranodon",
		"purple-worm",
		"quasit",
		"quipper",
		"rakshasa",
		"rat",
		"raven",
		"red-dragon-wyrmling",
		"reef-shark",
		"remorhaz",
		"rhinoceros",
		"riding-horse",
		"roc",
		"roper",
		"rug-of-smothering",
		"rust-monster",
		"saber-toothed-tiger",
		"sahuagin",
		"salamander",
		"satyr",
		"scorpion",
		"scout",
		"sea-hag",
		"sea-horse",
		"shadow",
		"shambling-mound",
		"shield-guardian",
		"shrieker",
		"silver-dragon-wyrmling",
		"skeleton",
		"solar",
		"spectator",
		"specter",
		"spider",
		"spirit-naga",
		"sprite",
		"spy",
		"steam-mephit",
		"stirge",
		"stone-giant",
		"stone-golem",
		"storm-giant",
		"strahd",
		"succubus/incubus",
		"swarm-of-bats",
		"swarm-of-insects",
		"swarm-of-poisonous-snakes",
		"swarm-of-quippers",
		"swarm-of-rats",
		"swarm-of-ravens",
		"tarrasque",
		"thug",
		"tiger",
		"tortle",
		"treant",
		"tribal-warrior",
		"triceratops",
		"troll",
		"twig-blight",
		"tyrannosaurus-rex",
		"unicorn",
		"vampire",
		"vampire-spawn",
		"veteran",
		"violet-fungus",
		"vrock",
		"vulture",
		"warhorse",
		"warhorse-skeleton",
		"water-elemental",
		"weasel",
		"werebear",
		"wereboar",
		"wererat",
		"weretiger",
		"werewolf",
		"white-dragon-wyrmling",
		"wight",
		"will-o'-wisp",
		"winter-wolf",
		"wolf",
		"worg",
		"wraith",
		"wyvern",
		"xorn",
		"yeti",
		"young-black-dragon",
		"young-blue-dragon",
		"young-brass-dragon",
		"young-bronze-dragon",
		"young-copper-dragon",
		"young-gold-dragon",
		"young-green-dragon",
		"young-red-dragon",
		"young-silver-dragon",
		"young-white-dragon",
		"zombie",
	}
)

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective-surname". For example 'focused-turing'.
func GetRandomName() string {
	adj1Index := newRandInt(int64(len(adjectives)))
	adj2Index := newRandInt(int64(len(adjectives)))
	nounIndex := newRandInt(int64(len(nouns)))
	randNum := newRandInt(10000000)
	return fmt.Sprintf("%s-%s-%s-%d", adjectives[adj1Index], adjectives[adj2Index], nouns[nounIndex], randNum)
}

// cryptoRandSecure is not always secure, if it errors we return 1337 % max
func newRandInt(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Printf("failed to generate random number: %v", err)
		return 1337 % max
	}
	return nBig.Int64()
}
