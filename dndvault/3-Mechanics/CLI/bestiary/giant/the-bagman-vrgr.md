---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/5
- monster/size/large
- monster/type/giant
aliases: ["The Bagman"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 225
---
# [The Bagman](3-Mechanics\CLI\bestiary\giant/the-bagman-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 225*  

## Beware the Bagman

The Bagman is an urban legend about an adventurer who sought to escape doom by abandoning his party and hiding inside a [bag of holding](/3-Mechanics/CLI/items/bag-of-holding.md). When he tried to leave, though, he became lost amid a constantly increasing number of extradimensional storage spaces. Over time, the strange forces of this magical in-between place transformed the adventurer into a monstrous creature. Now, every night, the Bagman slips out from a random [bag of holding](/3-Mechanics/CLI/items/bag-of-holding.md). If he doesn't find his home, he drags someone back into the bag with him and leaves behind some trinket from his hidden kingdom of lost junk. Some say that if you speak too loudly over an open [bag of holding](/3-Mechanics/CLI/items/bag-of-holding.md) or whisper "follow my voice" into a magical storage space three times, the Bagman will come for you.

Any character might know the story of the Bagman. What the Bagman is and how you use this urban legend is up to you. Is there truly a Bagman, or is he just a story? If an object vanishes overnight or if someone finds something that isn't theirs in a [bag of holding](/3-Mechanics/CLI/items/bag-of-holding.md), is the Bagman to blame? Is the Bagman just a monster that preys on adventurers, or is he the Darklord of his own hidden domain? The possibilities for horror adventures are endless, and nowhere—especially not adventurers' gear—is safe.

```statblock
"name": "The Bagman (VRGR)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "18"
- !!int "13"
- !!int "20"
- !!int "7"
- !!int "9"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Giant"
"cr": "5"
"traits":
- "desc": "The Bagman has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "The Bagman regains 10 hit points at the start of its turn. If the Bagman\
    \ takes acid or fire damage, this trait doesn't function at the start of the Bagman's\
    \ next turn. The Bagman dies only if it starts its turn with 0 hit points and\
    \ doesn't regenerate."
  "name": "Regeneration"
- "desc": "The Bagman has advantage on attack rolls against any creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by it."
  "name": "Grappler"
- "desc": "The Bagman can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If a creature tries to read the Bagman's thoughts, that creature must succeed\
    \ on a DC 8 Intelligence saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. The [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) creature\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Alien Mind"
"actions":
- "desc": "The Bagman makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage."
  "name": "Claw"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/The%20Bagman.webp"
```
^statblock

```encounter-table
name: The Bagman
creatures:
 - 1: The Bagman
```