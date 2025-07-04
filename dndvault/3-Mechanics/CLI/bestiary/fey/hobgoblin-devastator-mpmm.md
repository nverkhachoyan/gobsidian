---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/fey/goblinoid
aliases: ["Hobgoblin Devastator"]
NoteIcon: monster
BestiaryType: fey (goblinoid)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 153, Volo's Guide to Monsters p. 161
---
# [Hobgoblin Devastator](3-Mechanics\CLI\bestiary\fey/hobgoblin-devastator-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 153, Volo's Guide to Monsters p. 161*  

```statblock
"name": "Hobgoblin Devastator (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "goblinoid"
"alignment": "Typically  Lawful Neutral"
"ac": !!int "13"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "45"
"hit_dice": "7d8 + 14"
"stats":
- !!int "13"
- !!int "12"
- !!int "14"
- !!int "16"
- !!int "13"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Arcana": !!int "5"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Goblin"
"cr": "4"
"traits":
- "desc": "The hobgoblin casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 13):\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n2/day each:\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [gust of wind](/3-Mechanics/CLI/spells/gust-of-wind.md),\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)"
  "name": "Spellcasting"
- "desc": "When the hobgoblin casts a spell that causes damage or that forces other\
    \ creatures to make a saving throw, it can choose itself and any number of allies\
    \ to be immune to the damage caused by the spell and to succeed on the required\
    \ saving throw."
  "name": "Army Arcana"
"actions":
- "desc": "The hobgoblin makes two Quarterstaff or Devastating Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) bludgeoning damage, or dice:1d8 +\
    \ 1|text(5) (1d8 + 1) bludgeoning damage if used with two hands, plus dice:3d8|text(13)\
    \ (3d8) force damage."
  "name": "Quarterstaff"
- "desc": "Ranged Spell Attack: dice: d20+5 (+5) to hit, range 60 ft., one target.\
    \ Hit: dice:4d8 + 3|text(21) (4d8 + 3) force damage, and the target is knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Devastating Bolt"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Hobgoblin%20Devastator.webp"
```
^statblock

```encounter-table
name: Hobgoblin Devastator
creatures:
 - 1: Hobgoblin Devastator
```

## Environment

forest, grassland, hill