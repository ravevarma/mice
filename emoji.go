package main

var emoji_code = []string {
	/* (Smileys & Emotion :: 1..166 emojis)*/
	/* Smileys & Emotion - face-smiling (1..14) */
	"\U0001F600", // grinning face
	"\U0001F603", // grinning face with big eyes
	"\U0001F604", // grinning face with smiling eyes
	"\U0001F601", // beaming face with smiling eyes
	"\U0001F606", // grinning squinting face
	"\U0001F605", // grinning face with sweat
	"\U0001F923", // rolling on the floor laughing
	"\U0001F602", // face with tears of joy
	"\U0001F642", // slightly smiling face
	"\U0001F643", // upside-down face
	"\U0001FAE0", // melting face
	"\U0001F609", // winking face
	"\U0001F60A", // smiling face with smiling eyes
	"\U0001F607", // smiling face with halo

	/* Smileys & Emotion - face-affection (15..23) */
	"\U0001F970", // smiling face with hearts
	"\U0001F60D", // smiling face with heart-eyes
	"\U0001F929", // star-struck
	"\U0001F618", // face blowing a kiss
	"\U0001F617", // kissing face
	"\U0000263A", // smiling face
	"\U0001F61A", // kissing face with closed eyes
	"\U0001F619", // kissing face with smiling eyes
	"\U0001F972", // smiling face with tear

	/* Smileys & Emotion - face-tongue (24..29) */
	"\U0001F60B", // face savoring food
	"\U0001F61B", // face with tongue
	"\U0001F61C", // winking face with tongue
	"\U0001F92A", // zany face
	"\U0001F61D", // squinting face with tongue
	"\U0001F911", // money-mouth face

	/* Smileys & Emotion - face-hand (30..36) */
	"\U0001F917", // smiling face with open hands
	"\U0001F92D", // face with hand over mouth
	"\U0001FAE2", // face with open eyes and hand over mouth
	"\U0001FAE3", // face with peeking eye
	"\U0001F92B", // shushing face
	"\U0001F914", // thinking face
	"\U0001FAE1", // saluting face

	/* Smileys & Emotion - face-neutral-skeptical (37..50) */
	"\U0001F910", // zipper-mouth face
	"\U0001F928", // face with raised eyebrow
	"\U0001F610", // neutral face
	"\U0001F611", // expressionless face
	"\U0001F636", // face without mouth
	"\U0001FAE5", // dotted line face
	"\U0001F636\U0000200D\U0001F32B\U0000FE0F", // face in clouds
	"\U0001F60F", // smirking face
	"\U0001F612", // unamused face
	"\U0001F644", // face with rolling eyes
	"\U0001F62C", // grimacing face
	"\U0001F62E\U0000200D\U0001F4A8", // face exhaling
	"\U0001F925", // lying face
	"\U0001FAE8", // shaking face

	/* Smileys & Emotion - face-sleepy (51..55) */
	"\U0001F60C", // relieved face
	"\U0001F614", // pensive face
	"\U0001F62A", // sleepy face
	"\U0001F924", // drooling face
	"\U0001F634", // sleeping face

	/* Smileys & Emotion - face-unwell (56..67) */
	"\U0001F637", // face with medical mask
	"\U0001F912", // face with thermometer
	"\U0001F915", // face with head-bandage
	"\U0001F922", // nauseated face
	"\U0001F92E", // face vomiting
	"\U0001F927", // sneezing face
	"\U0001F975", // hot face
	"\U0001F976", // cold face
	"\U0001F974", // woozy face
	"\U0001F635", // face with crossed-out eyes
	"\U0001F635\U0000200D\U0001F4AB", // face with spiral eyes
	"\U0001F92F", // exploding head

	/* Smileys & Emotion - face-hat (68..70) */
	"\U0001F920", // cowboy hat face
	"\U0001F973", // partying face
	"\U0001F978", // disguised face

	/* Smileys & Emotion - face-glasses (71..73) */
	"\U0001F60E", // smiling face with sunglasses
	"\U0001F913", // nerd face
	"\U0001F9D0", // face with monocle

	/* Smileys & Emotion - face-concerned (74..99) */
	"\U0001F615", // confused face
	"\U0001FAE4", // face with diagonal mouth
	"\U0001F61F", // worried face
	"\U0001F641", // slightly frowning face
	"\U00002639", // frowning face
	"\U0001F62E", // face with open mouth
	"\U0001F62F", // hushed face
	"\U0001F632", // astonished face
	"\U0001F633", // flushed face
	"\U0001F97A", // pleading face
	"\U0001F979", // face holding back tears
	"\U0001F626", // frowning face with open mouth
	"\U0001F627", // anguished face
	"\U0001F628", // fearful face
	"\U0001F630", // anxious face with sweat
	"\U0001F625", // sad but relieved face
	"\U0001F622", // crying face
	"\U0001F62D", // loudly crying face
	"\U0001F631", // face screaming in fear
	"\U0001F616", // confounded face
	"\U0001F623", // persevering face
	"\U0001F61E", // disappointed face
	"\U0001F613", // downcast face with sweat
	"\U0001F629", // weary face
	"\U0001F62B", // tired face
	"\U0001F971", // yawning face

	/* Smileys & Emotions - face-negative (100..107) */
	"\U0001F624", // face with steam from nose
	"\U0001F621", // enraged face
	"\U0001F620", // angry face
	"\U0001F92C", // face with symbols on mouth
	"\U0001F608", // smiling face with horns
	"\U0001F47F", // angry face with horns
	"\U0001F480", // skull
	"\U00002620", // skull and crossbones

	/* Smileys & Emotion - face-costume (108..115) */
	"\U0001F4A9", // pile of poo
	"\U0001F921", // clown face
	"\U0001F479", // ogre
	"\U0001F47A", // goblin
	"\U0001F47B", // ghost
	"\U0001F47D", // alien
	"\U0001F47E", // alien monster
	"\U0001F916", // robot

	/* Smileys & Emotion - cat-face (116..124) */
	"\U0001F63A", // grinning cat
	"\U0001F638", // grinning cat with smiling eyes
	"\U0001F639", // cat with tears of joy
	"\U0001F63B", // smiling cat with heart-eyes
	"\U0001F63C", // cat with wry smile
	"\U0001F63D", // kissing cat
	"\U0001F640", // weary cat
	"\U0001F63F", // crying cat
	"\U0001F63E", // pouting cat

	/* Smileys & Emotion - monkey-face (125..127) */
	"\U0001F648", // see-no-evil monkey
	"\U0001F649", // hear-no-evil monkey
	"\U0001F64A", // speak-no-evil monkey

	/* Smileys & Emotion - heart (128..152) */
	"\U0001F48C", // love letter
	"\U0001F498", // heart with arrow
	"\U0001F49D", // heart with ribbon
	"\U0001F496", // sparkling heart
	"\U0001F497", // growing heart
	"\U0001F493", // beating heart
	"\U0001F49E", // revolving hearts
	"\U0001F495", // two hearts
	"\U0001F49F", // heart decoration
	"\U00002763", // heart exclamation
	"\U0001F494", // broken heart
	"\U00002764\U0000FE0F\U0000200D\U0001F525", // heart on fire
	"\U00002764\U0000FE0F\U0000200D\U0001FA79", // mending heart
	"\U00002764", // red heart
	"\U0001FA77", // pink heart
	"\U0001F9E1", // orange heart
	"\U0001F49B", // yellow heart
	"\U0001F49A", // green heart
	"\U0001F499", // blue heart
	"\U0001FA75", // light blue heart
	"\U0001F49C", // purple heart
	"\U0001F90E", // brown heart
	"\U0001F5A4", // black heart
	"\U0001FA76", // grey heart
	"\U0001F90D", // white heart

	/* Smileys & Emotion - emotion (153..166) */
	"\U0001F48B", // kiss mark
	"\U0001F4AF", // hundred points
	"\U0001F4A2", // anger symbol
	"\U0001F4A5", // collision
	"\U0001F4AB", // dizzy
	"\U0001F4A6", // sweat droplets
	"\U0001F4A8", // dashing away
	"\U0001F573", // hole
	"\U0001F4AC", // speech balloon
	"\U0001F441\U0000FE0F\U0000200D\U0001F5E8\U0000FE0F", // eye in speech bubble
	"\U0001F5E8", // left speech bubble
	"\U0001F5EF", // right anger bubble
	"\U0001F4AD", // thought balloon
	"\U0001F4A4", // zzz

	/* People & Body :: */
	/* People & Body - hand-fingers-open (167..177) */
	"\U0001F44B", // waving hand
	"\U0001F91A", // raised back of hand
	"\U0001F590", // hand with fingers splayed
	"\U0000270B", // raised hand
	"\U0001F596", // vulcan salute
	"\U0001FAF1", // rightwards hand
	"\U0001FAF2", // leftwards hand
	"\U0001FAF3", // palm down hand
	"\U0001FAF4", // palm up hand
	"\U0001FAF7", // leftwards pushing hand
	"\U0001FAF8", // rightwards pushing hand

	/* People & Body - hand-fingers-partial (178..186) */
	"\U0001F44C", // ok hand
	"\U0001F90C", // pinched fingers
	"\U0001F90F", // pinching hand
	"\U0000270C", // victory hand
	"\U0001F91E", // crossed fingers
	"\U0001FAF0", // hand with index finger and thumb crossed
	"\U0001F91F", // love-you gesture
	"\U0001F918", // sign of the horns
	"\U0001F919", // call me hand

	/* People & Body - hand-single-finger (187..193) */
	"\U0001F448", // backhand index pointing left
	"\U0001F449", // backhand index pointing right
	"\U0001F446", // backhand index pointing up
	"\U0001F595", // middle finger
	"\U0001F447", // backhand index pointing down
	"\U0000261D", // index pointing up
	"\U0001FAF5", // index pointing at the viewer

	/* People & Body - hand-fingers-closed (194..199) */
	"\U0001F44D", // thumbs up
	"\U0001F44E", // thumbs down
	"\U0000270A", // raised fist
	"\U0001F44A", // oncoming fist
	"\U0001F91B", // left-facing fist
	"\U0001F91C", // right-facing fist

	/* People & Body - hands (200..206) */
	"\U0001F44F", // clapping hands
	"\U0001F64C", // raising hands
	"\U0001FAF6", // heart hands
	"\U0001F450", // open hands
	"\U0001F932", // palms up together
	"\U0001F91D", // handshake
	"\U0001F64F", // folded hands

	/* People & Body - hand-prop (207..209) */
	"\U0000270D", // writing hand
	"\U0001F485", // nail polish
	"\U0001F933", // selfie

	/* People & Body - body-parts (210..227) */
	"\U0001F4AA", // flexed biceps
	"\U0001F9BE", // mechanical arm
	"\U0001F9BF", // mechanical leg
	"\U0001F9B5", // leg
	"\U0001F9B6", // foot
	"\U0001F442", // ear
	"\U0001F9BB", // ear with hearing aid
	"\U0001F443", // nose
	"\U0001F9E0", // brain
	"\U0001FAC0", // anatomical heart
	"\U0001FAC1", // lungs
	"\U0001F9B7", // tooth
	"\U0001F9B4", // bone
	"\U0001F440", // eyes
	"\U0001F441", // eye
	"\U0001F445", // tongue
	"\U0001F444", // mouth
	"\U0001FAE6", // biting lip

	/* People & Body - person (228..255) */
	"\U0001F476", // baby
	"\U0001F9D2", // child
	"\U0001F466", // boy
	"\U0001F467", // girl
	"\U0001F9D1", // person
	"\U0001F471", // person: blond hair
	"\U0001F468", // man
	"\U0001F9D4", // person: beard
	"\U0001F9D4\U0000200D\U00002642\U0000FE0F", // man: beard
	"\U0001F9D4\U0000200D\U00002640\U0000FE0F", // woman: beard
	"\U0001F468\U0000200D\U0001F9B0", // man: red hair
	"\U0001F468\U0000200D\U0001F9B1", // man: curly hair
	"\U0001F468\U0000200D\U0001F9B3", // man: white hair
	"\U0001F468\U0000200D\U0001F9B2", // man: bald
	"\U0001F469", // woman
	"\U0001F469\U0000200D\U0001F9B0", // woman: red hair
	"\U0001F9D1\U0000200D\U0001F9B0", // person: red hair
	"\U0001F469\U0000200D\U0001F9B1", // woman: curly hair
	"\U0001F9D1\U0000200D\U0001F9B1", // person: curly hair
	"\U0001F469\U0000200D\U0001F9B3", // woman: white hair
	"\U0001F9D1\U0000200D\U0001F9B3", // person: white hair
	"\U0001F469\U0000200D\U0001F9B2", // woman: bald
	"\U0001F9D1\U0000200D\U0001F9B2", // person: bald
	"\U0001F471\U0000200D\U00002640\U0000FE0F", // woman: blond hair
	"\U0001F471\U0000200D\U00002642\U0000FE0F", // man: blond hair
	"\U0001F9D3", // older person
	"\U0001F474", // old man
	"\U0001F475", // old woman

	/* People & Body - person-gesture (256..285) */
	"\U0001F64D", // person frowning
	"\U0001F64D\U0000200D\U00002642\U0000FE0F", // man frowning
	"\U0001F64D\U0000200D\U00002640\U0000FE0F", // woman frowning
	"\U0001F64E", // person pouting
	"\U0001F64E\U0000200D\U00002642\U0000FE0F", // man pouting
	"\U0001F64E\U0000200D\U00002640\U0000FE0F", // woman pouting
	"\U0001F645", // person gesturing no
	"\U0001F645\U0000200D\U00002642\U0000FE0F", // man gesturing no
	"\U0001F645\U0000200D\U00002640\U0000FE0F", // woman gesturing no
	"\U0001F646", // person gesturing ok
	"\U0001F646\U0000200D\U00002642\U0000FE0F", // man gesturing ok
	"\U0001F646\U0000200D\U00002640\U0000FE0F", // woman gesturing ok
	"\U0001F481", // person tipping hand
	"\U0001F481\U0000200D\U00002642\U0000FE0F", // man tipping hand
	"\U0001F481\U0000200D\U00002640\U0000FE0F", // woman tipping hand
	"\U0001F64B", // person raising hand
	"\U0001F64B\U0000200D\U00002642\U0000FE0F", // man raising hand
	"\U0001F64B\U0000200D\U00002640\U0000FE0F", // woman raising hand
	"\U0001F9CF", // deaf person
	"\U0001F9CF\U0000200D\U00002642\U0000FE0F", // deaf man
	"\U0001F9CF\U0000200D\U00002640\U0000FE0F", // deaf woman
	"\U0001F647", // person bowing
	"\U0001F647\U0000200D\U00002642\U0000FE0F", // man bowing
	"\U0001F647\U0000200D\U00002640\U0000FE0F", // woman bowing
	"\U0001F926", // person facepalming
	"\U0001F926\U0000200D\U00002642\U0000FE0F", // man facepalming
	"\U0001F926\U0000200D\U00002640\U0000FE0F", // woman facepalming
	"\U0001F937", // person shrugging
	"\U0001F937\U0000200D\U00002642\U0000FE0F", // man shrugging
	"\U0001F937\U0000200D\U00002640\U0000FE0F", // woman shrugging

	/* People & Body - person-role (286..367) */
	"\U0001F9D1\U0000200D\U00002695\U0000FE0F", // health worker
	"\U0001F468\U0000200D\U00002695\U0000FE0F", // man health worker
	"\U0001F469\U0000200D\U00002695\U0000FE0F", // woman health worker
	"\U0001F9D1\U0000200D\U0001F393", // student
	"\U0001F468\U0000200D\U0001F393", // man student
	"\U0001F469\U0000200D\U0001F393", // woman student
	"\U0001F9D1\U0000200D\U0001F3EB", // teacher
	"\U0001F468\U0000200D\U0001F3EB", // man teacher
	"\U0001F469\U0000200D\U0001F3EB", // woman teacher
	"\U0001F9D1\U0000200D\U00002696\U0000FE0F", // judge
	"\U0001F468\U0000200D\U00002696\U0000FE0F", // man judge
	"\U0001F469\U0000200D\U00002696\U0000FE0F", // woman judge
	"\U0001F9D1\U0000200D\U0001F33E", // farmer
	"\U0001F468\U0000200D\U0001F33E", // man farmer
	"\U0001F469\U0000200D\U0001F33E", // woman farmer
	"\U0001F9D1\U0000200D\U0001F373", // cook
	"\U0001F468\U0000200D\U0001F373", // man cook
	"\U0001F469\U0000200D\U0001F373", // woman cook
	"\U0001F9D1\U0000200D\U0001F527", // mechanic
	"\U0001F468\U0000200D\U0001F527", // man mechanic
	"\U0001F469\U0000200D\U0001F527", // woman mechanic
	"\U0001F9D1\U0000200D\U0001F3ED", // factory worker
	"\U0001F468\U0000200D\U0001F3ED", // man factory worker
	"\U0001F469\U0000200D\U0001F3ED", // woman factory worker
	"\U0001F9D1\U0000200D\U0001F4BC", // office worker
	"\U0001F468\U0000200D\U0001F4BC", // man office worker
	"\U0001F469\U0000200D\U0001F4BC", // woman office worker
	"\U0001F9D1\U0000200D\U0001F52C", // scientist
	"\U0001F468\U0000200D\U0001F52C", // man scientist
	"\U0001F469\U0000200D\U0001F52C", // woman scientist
	"\U0001F9D1\U0000200D\U0001F4BB", // technologist
	"\U0001F468\U0000200D\U0001F4BB", // man technologist
	"\U0001F469\U0000200D\U0001F4BB", // woman technologist
	"\U0001F9D1\U0000200D\U0001F3A4", // singer
	"\U0001F468\U0000200D\U0001F3A4", // man singer
	"\U0001F469\U0000200D\U0001F3A4", // woman singer
	"\U0001F9D1\U0000200D\U0001F3A8", // artist
	"\U0001F468\U0000200D\U0001F3A8", // man artist
	"\U0001F469\U0000200D\U0001F3A8", // woman artist
	"\U0001F9D1\U0000200D\U00002708\U0000FE0F", // pilot
	"\U0001F468\U0000200D\U00002708\U0000FE0F", // man pilot
	"\U0001F469\U0000200D\U00002708\U0000FE0F", // woman pilot
	"\U0001F9D1\U0000200D\U0001F680", // astronaut
	"\U0001F468\U0000200D\U0001F680", // man astronaut
	"\U0001F469\U0000200D\U0001F680", // woman astronaut
	"\U0001F9D1\U0000200D\U0001F692", // firefighter
	"\U0001F468\U0000200D\U0001F692", // man firefighter
	"\U0001F469\U0000200D\U0001F692", // woman firefighter
	"\U0001F46E", // police officer
	"\U0001F46E\U0000200D\U00002642\U0000FE0F", // man police officer
	"\U0001F46E\U0000200D\U00002640\U0000FE0F", // woman police officer
	"\U0001F575", // detective
	"\U0001F575\U0000FE0F\U0000200D\U00002642\U0000FE0F", // man detective
	"\U0001F575\U0000FE0F\U0000200D\U00002640\U0000FE0F", // woman detective
	"\U0001F482", // guard
	"\U0001F482\U0000200D\U00002642\U0000FE0F", // man guard
	"\U0001F482\U0000200D\U00002640\U0000FE0F", // woman guard
	"\U0001F977", // ninja
	"\U0001F477", // construction worker
	"\U0001F477\U0000200D\U00002642\U0000FE0F", // man construction worker
	"\U0001F477\U0000200D\U00002640\U0000FE0F", // woman construction worker
	"\U0001FAC5", // person with crown
	"\U0001F934", // prince
	"\U0001F478", // princess
	"\U0001F473", // person wearing turban
	"\U0001F473\U0000200D\U00002642\U0000FE0F", // man wearing turban
	"\U0001F473\U0000200D\U00002640\U0000FE0F", // woman wearing turban
	"\U0001F472", // person with skullcap
	"\U0001F9D5", // woman with headscarf
	"\U0001F935", // person in tuxedo
	"\U0001F935\U0000200D\U00002642\U0000FE0F", // man in tuxedo
	"\U0001F935\U0000200D\U00002640\U0000FE0F", // woman in tuxedo
	"\U0001F470", // person with veil
	"\U0001F470\U0000200D\U00002642\U0000FE0F", // man with veil
	"\U0001F470\U0000200D\U00002640\U0000FE0F", // woman with veil
	"\U0001F930", // pregnant woman
	"\U0001FAC3", // pregnant man
	"\U0001FAC4", // pregnant person
	"\U0001F931", // breast-feeding
	"\U0001F469\U0000200D\U0001F37C", // woman feeding baby
	"\U0001F468\U0000200D\U0001F37C", // man feeding baby
	"\U0001F9D1\U0000200D\U0001F37C", // person feeding baby

	/* People & Body - person-fantasy (368..399) */
	"\U0001F47C", // baby angel
	"\U0001F385", // santa claus
	"\U0001F936", // mrs.claus
	"\U0001F9D1\U0000200D\U0001F384", // mx claus
	"\U0001F9B8", // superhero
	"\U0001F9B8\U0000200D\U00002642\U0000FE0F", // man superhero
	"\U0001F9B8\U0000200D\U00002640\U0000FE0F", // woman superhero
	"\U0001F9B9", // supervillain
	"\U0001F9B9\U0000200D\U00002642\U0000FE0F", // man supervillain
	"\U0001F9B9\U0000200D\U00002640\U0000FE0F", // woman supervillain
	"\U0001F9D9", // mage
	"\U0001F9D9\U0000200D\U00002642\U0000FE0F", // man mage
	"\U0001F9D9\U0000200D\U00002640\U0000FE0F", // woman mage
	"\U0001F9DA", // fairy
	"\U0001F9DA\U0000200D\U00002642\U0000FE0F", // man fairy
	"\U0001F9DA\U0000200D\U00002640\U0000FE0F", // woman fairy
	"\U0001F9DB", // vampire
	"\U0001F9DB\U0000200D\U00002642\U0000FE0F", // man vampire
	"\U0001F9DB\U0000200D\U00002640\U0000FE0F", // woman vampire
	"\U0001F9DC", // merperson
	"\U0001F9DC\U0000200D\U00002642\U0000FE0F", // merman
	"\U0001F9DC\U0000200D\U00002640\U0000FE0F", // mermaid
	"\U0001F9DD", // elf
	"\U0001F9DD\U0000200D\U00002642\U0000FE0F", // man elf
	"\U0001F9DD\U0000200D\U00002640\U0000FE0F", // woman elf
	"\U0001F9DE", // genie
	"\U0001F9DE\U0000200D\U00002642\U0000FE0F", // man genie
	"\U0001F9DE\U0000200D\U00002640\U0000FE0F", // woman genie
	"\U0001F9DF", // zombie
	"\U0001F9DF\U0000200D\U00002642\U0000FE0F", // man zombie
	"\U0001F9DF\U0000200D\U00002640\U0000FE0F", // woman zombie
	"\U0001F9CC", // troll

	/* People & Body - person-activity (400..438) */

	/* People & Body - person-sport (439..481) */

	/* People & Body - person-resting (482..486) */

	/* People & Body - family (487..524) */

	/* People & Body - person-symbol (525..529) */
	"\U0001F5E3", // speaking head
	"\U0001F464", // bust in silhouette
	"\U0001F465", // busts in silhouette
	"\U0001FAC2", // people hugging
	"\U0001F463", // footprints

	/* Component (530..533) */
	/* Component - hair-style (530..533) */
	"\U0001F9B0", // red hair
	"\U0001F9B1", // curly hair
	"\U0001F9B3", // white hair
	"\U0001F9B2", // bald
}

var emoji_cldr = []string {

}