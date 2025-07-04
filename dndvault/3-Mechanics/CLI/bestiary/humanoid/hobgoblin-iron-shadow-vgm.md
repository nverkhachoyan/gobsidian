---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/goblinoid
aliases: ["Hobgoblin Iron Shadow"]
NoteIcon: monster
BestiaryType: humanoid (goblinoid)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 162
---
# [Hobgoblin Iron Shadow](3-Mechanics\CLI\bestiary\humanoid/hobgoblin-iron-shadow-vgm.md)
*Source: Volo's Guide to Monsters p. 162*  

The Iron Shadows are hobgoblin monks that serve as secret police, scouts, and assassins. Among other hobgoblins, they spy to ferret out treachery, rebellion, and betrayal.

## Trained in Secret

Iron Shadows are recruited from across the hobgoblin ranks. Each member keeps her eyes open for potential recruits, those whose agility and stamina are matched only by an ironclad commitment to Maglubiyet's will.

A candidate for admission undergoes a series of tests designed to reveal any potential for treachery. Those who fail are slain, while those who pass receive secret training in the magical and martial arts. This indoctrination is a slow and arduous process; many aspirants don't finish it, and years might go by during which the Iron Shadows welcome no new members into their ranks. While a recruit is in training, it serves the Iron Shadows by looking for and reporting suspicious behavior.

Their masks also signify the supposed origin of their fighting techniques. The priests of Maglubiyet teach that the Great One stole the secrets of shadows from an archdevil, allowing his followers to conceal their identities, walk between shadows, and craft illusions to confuse and confound their enemies.

## Masters of Shadow and Fist

When a recruit's training is complete, she is ready to wield a deadly combination of unarmed fighting techniques and shadow magic to deceive and defeat her foes. She continues to spy on other hobgoblins, but is now also empowered to conduct assassinations and spy missions, both against enemies and among goblinoids. These missions are ordained by the clerics of Maglubiyet, who keep a careful eye on the goblinoid community to ensure that it functions according to Maglubiyet's will.

## Masked Devils

Iron Shadows on a secret mission wear masks crafted to resemble devils, both to conceal their identities and to strike fear into their foes.

```statblock
"name": "Hobgoblin Iron Shadow (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "goblinoid"
"alignment": "Lawful Evil"
"ac": !!int "15"
"hp": !!int "32"
"hit_dice": "5d8 + 10"
"stats":
- !!int "14"
- !!int "16"
- !!int "15"
- !!int "14"
- !!int "15"
- !!int "11"
"speed": "40 ft."
"skillsaves":
  "Athletics": !!int "4"
  "Stealth": !!int "5"
  "Acrobatics": !!int "5"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Common, Goblin"
"cr": "2"
"traits":
- "desc": "The hobgoblin is a 2nd-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). It has the\
    \ following wizard spells prepared:\n\nCantrips (at will): [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [true strike](/3-Mechanics/CLI/spells/true-strike.md)\n\
    \n1st level (3 slots): [charm person](/3-Mechanics/CLI/spells/charm-person.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md),\
    \ [silent image](/3-Mechanics/CLI/spells/silent-image.md)"
  "name": "Spellcasting"
- "desc": "While the hobgoblin is wearing no armor and wielding no shield, its AC\
    \ includes its Wisdom modifier."
  "name": "Unarmored Defense"
"actions":
- "desc": "The hobgoblin makes four attacks, each of which can be an unarmed strike\
    \ or a dart attack. It can also use Shadow Jaunt once, either before or after\
    \ one of the attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage."
  "name": "Unarmed Strike"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 20/60 ft., one\
    \ target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Dart"
- "desc": "The hobgoblin magically teleports, along with any equipment it is wearing\
    \ or carrying, up to 30 feet to an unoccupied space it can see. Both the space\
    \ it is leaving and its destination must be in dim light or darkness."
  "name": "Shadow Jaunt"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Hobgoblin%20Iron%20Shadow.webp"
```
^statblock

```encounter-table
name: Hobgoblin Iron Shadow
creatures:
 - 1: Hobgoblin Iron Shadow
```

## Environment

grassland, forest, hill