---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/fey/goblinoid
aliases: ["Hobgoblin Iron Shadow"]
NoteIcon: monster
BestiaryType: fey (goblinoid)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 154, Volo's Guide to Monsters p. 162
---
# [Hobgoblin Iron Shadow](3-Mechanics\CLI\bestiary\fey/hobgoblin-iron-shadow-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 154, Volo's Guide to Monsters p. 162*  

```statblock
"name": "Hobgoblin Iron Shadow (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "goblinoid"
"alignment": "Typically  Lawful Neutral"
"ac": !!int "15"
"ac_class": "Unarmored Defense"
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
- "desc": "The hobgoblin casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 12):\n\nAt will: [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1/day each:\
    \ [charm person](/3-Mechanics/CLI/spells/charm-person.md), [disguise self](/3-Mechanics/CLI/spells/disguise-self.md),\
    \ [silent image](/3-Mechanics/CLI/spells/silent-image.md)"
  "name": "Spellcasting"
- "desc": "While the hobgoblin is wearing no armor and wielding no shield, its AC\
    \ includes its Wisdom modifier."
  "name": "Unarmored Defense"
"actions":
- "desc": "The hobgoblin makes four attacks, each of which can be an Unarmed Strike\
    \ or a Dart attack. It can also use\n\nShadow Jaunt once, either before or after\
    \ one of the attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage."
  "name": "Unarmed Strike"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 20/60 ft., one\
    \ target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Dart"
- "desc": "The hobgoblin teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see. Both the space it leaves and\
    \ its destination must be in dim light or darkness."
  "name": "Shadow Jaunt"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Hobgoblin%20Iron%20Shadow.webp"
```
^statblock

```encounter-table
name: Hobgoblin Iron Shadow
creatures:
 - 1: Hobgoblin Iron Shadow
```

## Environment

forest, grassland, hill