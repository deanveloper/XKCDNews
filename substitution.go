package xkcdnews

import "strings"

var subs = [][2]string{
	// Substitutions 1 - https://xkcd.com/1288/
	{"witnesses", "these dudes i know"},
	{"allegedly", "kinda probably"},
	{"new study", "tumblr post"},
	{"rebuild", "avenge"},
	{"space", "spaaace"},
	{"google glass", "virtual boy"},
	{"smartphone", "pokédex"},
	{"electric", "atomic"},
	{"senator", "elf-lord"},
	{"car", "cat"},
	{"election", "eating contest"},
	{"congressional leaders", "river spirits"},
	{"homeland security", "homestar runner"}, // i miss this site
	{"could not be reacehd for comment", "is guilty and everyone knows it"},
	// substitutions 2 - https://xkcd.com/1625/
	{"debate", "dance-off"},
	{"self driving", "uncontrollably swerving"},
	{"poll", "psychic reading"},
	{"candidate", "airbender"},
	{"drone", "dog"},
	{"vows to", "probably won't"},
	{"at large", "very large"},
	{"successfully", "suddenly"},
	{"expands", "physically expands"},
	{"first-degree", "friggin' awful"},
	{"second-degree", "friggin' awful"},
	{"third-degree", "friggin' awful"},
	{"an unknown number", "like hundreds"},
	{"front runner", "blade runner"},
	{"global", "spherical"},
	{"minutes", "!m!inu!tes!"}, // minutes -> weird thing
	{"years", "minutes"},       // years -> minutes
	{"!m!inu!tes!", "years"},   // weird thing -> years (these three steps are required!)
	{"no indication", "lots of signs"},
	{"urged restraint by", "drunkenly egged on"},
	{"horsepower", "tons of horsemeat"},
	// Substitutions 3 - https://xkcd.com/1679/
	{"gaffe", "magic spell"},
	{"ancient", "haunted"},
	{"star-studded", "blood-soaked"},
	{"remains to be seen", "will never be known"},
	{"silver bullet", "way to kill werewolves"},
	{"subway system", "tunnels i found"},
	{"surprising", "surprising (but not to me)"},
	{"war of words", "interplanetary war"},
	{"tension", "sexual tension"},
	{"cautiously optimistic", "delusional"},
	{"doctor who", "the big bang theory"},
	{"win votes", "find pokémon"},
	{"behind the headlines", "beyond the grave"},
	{"email", "poem"},
	{"facebook post", "poem"},
	{"tweet", "poem"},
	{"facebook ceo", "this guy"},
	{"latest", "final"},
	{"disrupt", "destroy"},
	{"meeting", "ménag à trois"},
	{"scientists", "channing tatum and his friends"},
	{"you won't believe", "i'm really sad about"},
}

func Substitute(s string) (string, int) {
	replaced := 0
	s1 := s
	for _, entry := range subs {
		old := s1
		s1 = strings.Replace(s1, entry[0], entry[1], -1)
		if s1 != old {
			replaced++
		}
	}

	return s1, replaced
}
