---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Bard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 59, Volo's Guide to Monsters p. 211
---
# [Bard](3-Mechanics\CLI\bestiary\humanoid/bard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 59, Volo's Guide to Monsters p. 211*  

```statblock
"name": "Bard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "[chain shirt](/3-Mechanics/CLI/items/chain-shirt.md)"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "11"
- !!int "14"
- !!int "12"
- !!int "10"
- !!int "13"
- !!int "14"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "4"
  "Wisdom": !!int "3"
"skillsaves":
  "Perception": !!int "5"
  "Performance": !!int "6"
  "Acrobatics": !!int "4"
"senses": "passive Perception 15"
"languages": "any two languages"
"cr": "2"
"traits":
- "desc": "The bard casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The bard makes two Shortsword or Shortbow attacks. It can replace one attack\
    \ with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortbow"
- "desc": "Each creature in a 15-foot cube originating from the bard must make a DC\
    \ 12 Constitution saving throw. On a failed save, a creature takes dice:2d8|text(9)\
    \ (2d8) thunder damage and is pushed up to 10 feet away from the bard. On a\
    \ successful save, a creature takes half as much damage and isn't pushed."
  "name": "Cacophony (Recharge 4-6)"
"bonus_actions":
- "desc": "The bard targets one creature within 30 feet of it. If the target can hear\
    \ the bard, the target must succeed on a DC 12 Charisma saving throw or have disadvantage\
    \ on ability checks, attack rolls, and saving throws until the start of the bard's\
    \ next turn."
  "name": "Taunt (2/Day)"
"source":
- "MPMM"
- "VGM"
- "AATM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Bard.webp"
```
^statblock

```encounter-table
name: Bard
creatures:
 - 1: Bard
```

## Environment

urban