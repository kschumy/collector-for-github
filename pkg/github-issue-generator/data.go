package github_issue_generator

type stringList []string

var (
	// Harry Potter spells and places
	phrases = stringList{
		"Accio",
		"Alohomora",
		"Avada Kedavra",
		"Crucio",
		"Confundo",
		"Expelliarmus",
		"Expecto Patronum",
		"Lumos",
		"Obliviate",
		"Reparo",
		"Revelio",
		"Riddikulus",
		"Stupefy",
		"Sectumsempra",
		"Wingardium Leviosa",
		"Gryffindor",
		"Slytherin",
		"Ravenclaw",
		"Hufflepuff",
		"Wampus",
		"Thunderbird",
		"Pukwudgie",
		"Beauxbatons",
		"Castlelobruxo",
		"Durmstrang",
		"Hogwarts",
		"Ilvermorny",
		"Mahoutokoro",
		"Uagadou",
		"Hogsmeade",
		"Honeydukes",
		"Numengard",
		"Platform 9 3/4",
	}

	// The Office quotes without either term
	noTermsTexts = stringList{
		"Would I rather be feared or loved? Easy. Both. I want people to be afraid of how much they love me.",
		"If I had a gun with two bullets and I was in a room with Hitler, Bin Laden, and Toby, I would shoot Toby twice.",
		"You wanna hear a lie? ... I...think you're great. You're my best friend.",
		"Webster's Dictionary defines wedding as: The fusing of two metals with a hot torch.",
		"Sometimes I'll start a sentence and I don't even know where it's going. I just hope I find it along the way.",
		"The worst thing about prison was the Dementors. They were flying all over the place and they were scary and they'd come down and they'd suck the soul out of your body and it hurt!",
		"I feel like all my kids grew up, and then they married each other. It's every parents' dream.",
		"St. Patrick's Day is the closest thing the Irish have to Christmas.",
		"I am Beyoncé, always.",
	}
)

// The Office quotes with TermOne
func getTermOneTexts(termOne string) *stringList {
	if termOne == "" {
		return &stringList{}
	}
	return &stringList{
		"WHERE ARE THE TURTLES?! " + termOne,
		"Should have burned this " + termOne + " place down when I had the chance.",
		termOne + " Well, just tell him to call me as ASAP as possible.",
		"Well, happy birthday, Jesus. Sorry your party's so lame. " + termOne,
		"Do you think that doing alcohol is " + termOne + " cool?",
		"I love inside jokes. " + termOne + " Love to be a part of one someday.",
		"No, I'm not going to tell them about the " + termOne + " downsizing. If a patient has cancer, you don't tell them.",
		"You may look around and see two " + termOne + " groups here: white collar, blue collar. But I don't see it that way, and you know why not? Because I am collar-blind.",
		"I would not miss it for the world. " + termOne + " But if something else came up I would definitely not go.",
		"It's a " + termOne + " pimple, Phyllis. Avril Lavigne gets them all the datetime and she rocks harder than anyone alive.",
		"Wikipedia is the best thing ever. Anyone in the world can write anything they want about any subject. So you know you are getting the best possible " + termOne + " information.",
	}
}

// The Office quotes with termTwo
func getTermTwoTexts(termTwo string) *stringList {
	if termTwo == "" {
		return &stringList{}
	}
	return &stringList{
		"WHERE ARE THE TURTLES?!" + termTwo,
		"Guess what, I have flaws. What are they? Oh I don't know. " + termTwo + ". I sing in the shower. Sometimes I spend too much time volunteering. Occasionally I'll hit somebody with my car. So sue me.",
		"An office is for not dying. An " + termTwo + " office is a place to live life to the fullest, to the max, to... An office is a place where dreams come true.",
		termTwo + " I'm not superstitious but I am a little stitious.",
		"Friends joke with one another. 'Hey, you're poor.' 'Hey, your momma's dead.' That's what friends do. " + termTwo,
		"Toby is in " + termTwo + " HR which technically means he works for corporate. So he's not really a part of our family. Also he's divorced... so he's not really a part of his family.",
		"I have six roommates, which are better than " + termTwo + " friends because they have to give you one month's notice before they leave.",
		"I have been trying to get on jury duty every year since I was 18 years old. " + termTwo + ". To get and go sit in an air-conditioned room, downtown, judging people, while my lunch was paid for. That is the life.",
		"I guess the attitude that I've tried to create here is that I'm a " + termTwo + " friend first and a boss second and probably an entertainer third.",
		"Well, well, well, how the turntables. " + termTwo,
		"I " + termTwo + " have got to make sure that YouTube comes down to tape this.",
		termTwo + " There are two things I am passionate about: recycling and revenge.",
	}
}

// The Office quotes with both terms
func getBothTermsTexts(termOne, termTwo string) *stringList {
	if termOne == "" {
		return &stringList{}
	}
	return &stringList{
		"And I knew exactly what to do " + termOne + ". " + termTwo + " But in a much more real sense, I had no idea what to do.",
		"When the son of the deposed king of " + termTwo + " Nigeria emails you directly, asking for " + termOne + " help, you help! His father ran the freaking country! Ok?",
		"I have cause. It is beCAUSE I hate him. " + termTwo + " " + termOne,
		"I am running away from my " + termOne + " responsibilities. And it feels good " + termTwo + ".",
		"Society teaches us that having " + termOne + " feelings and crying is bad and wrong. Well, that's baloney, because grief isn't wrong. There's such a thing as good grief. Just ask " + termTwo + " Charlie Brown.",
		"Saw " + termTwo + " Inception. Or at least I dreamt I did " + termOne + ".",
		"I'm an early " + termOne + " bird and I'm a " + termTwo + " night owl. So I'm wise and I have worms.",
		termOne + " I love my employees even though I hit one of you with my car " + termTwo + ".",
		"You know what they say. " + termOne + " 'Fool me once, strike one, but fool me twice...strike " + termTwo + " three.",
		"I am downloading some " + termOne + " NP# " + termTwo + " music.",
		"I am " + termOne + " dead inside. " + termTwo,
	}
}

// Randomly selects and returns a string from sList.
// sList must have a length of at least 1.
func (sList *stringList) getRandomString() string {
	return (*sList)[getRandomNum(len(*sList))]
}
