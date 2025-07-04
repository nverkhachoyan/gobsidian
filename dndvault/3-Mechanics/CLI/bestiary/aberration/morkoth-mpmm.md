---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/coastal
- monster/environment/underwater
- monster/size/large
- monster/type/aberration
aliases: ["Morkoth"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 186, Volo's Guide to Monsters p. 177
---
# [Morkoth](3-Mechanics\CLI\bestiary\aberration/morkoth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 186, Volo's Guide to Monsters p. 177*  

```statblock
"name": "Morkoth (MPMM)"
"size": "Large"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "165"
"hit_dice": "22d10 + 44"
"stats":
- !!int "14"
- !!int "14"
- !!int "14"
- !!int "20"
- !!int "15"
- !!int "13"
"speed": "25 ft., swim 50 ft."
"saves":
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
  "Intelligence": !!int "9"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "10"
  "History": !!int "9"
  "Arcana": !!int "9"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 20"
"languages": "telepathy 120 ft."
"cr": "11"
"traits":
- "desc": "The morkoth casts one of the following spells, requiring no material components\
    \ and using Intelligence as the spellcasting ability (spell save DC 17):\n\nAt\
    \ will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\n\
    \n3/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md), [sending](/3-Mechanics/CLI/spells/sending.md)"
  "name": "Spellcasting"
- "desc": "The morkoth can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "The morkoth makes either two Bite attacks and one Tentacles attack or three\
    \ Bite attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) slashing damage plus dice:3d6|text(10)\
    \ (3d6) psychic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 15 ft., one target.\
    \ Hit: dice:3d8 + 2|text(15) (3d8 + 2) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14)\
    \ if it is a Large or smaller creature. Until this grapple ends, the target is\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) and takes dice:3d8\
    \ + 2|text(15) (3d8 + 2) bludgeoning damage at the start of each of its turns,\
    \ and the morkoth can't use its tentacles on another target."
  "name": "Tentacles"
- "desc": "The morkoth projects a 30-foot cone of magical energy. Each creature in\
    \ that area must make a DC 17 Wisdom saving throw. On a failed save, the creature\
    \ is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the morkoth for\
    \ 1 minute. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) in this\
    \ way, the target tries to get as close to the morkoth as possible, using its\
    \ actions to [Dash](/3-Mechanics/CLI/rules/actions.md#Dash) until it is within\
    \ 5 feet of the morkoth. A [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target can repeat the saving throw at the end of each of its turns and whenever\
    \ it takes damage, ending the effect on itself on a success. If a creature's saving\
    \ throw is successful or the effect ends for it, the creature has advantage on\
    \ saving throws against the morkoth's Hypnosis for 24 hours."
  "name": "Hypnosis"
"reactions":
- "desc": "If the morkoth makes a successful saving throw against a spell or a spell\
    \ attack misses it, the morkoth can choose another creature (including the spellcaster)\
    \ it can see within 120 feet of it. The spell targets the chosen creature instead\
    \ of the morkoth. If the spell forced a saving throw, the chosen creature makes\
    \ its own save. If the spell was an attack, the attack roll is rerolled against\
    \ the chosen creature."
  "name": "Spell Reflection"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Morkoth.webp"
```
^statblock

```encounter-table
name: Morkoth
creatures:
 - 1: Morkoth
```

## Environment

coastal, underwater