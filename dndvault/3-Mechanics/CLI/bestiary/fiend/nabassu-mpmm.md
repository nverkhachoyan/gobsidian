---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/15
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Nabassu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 188, Mordenkainen's Tome of Foes p. 135
---
# [Nabassu](3-Mechanics\CLI\bestiary\fiend/nabassu-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 188, Mordenkainen's Tome of Foes p. 135*  

```statblock
"name": "Nabassu (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "190"
"hit_dice": "20d8 + 100"
"stats":
- !!int "22"
- !!int "14"
- !!int "21"
- !!int "14"
- !!int "15"
- !!int "17"
"speed": "40 ft., fly 60 ft."
"saves":
  "Dexterity": !!int "7"
  "Strength": !!int "11"
"skillsaves":
  "Perception": !!int "7"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 17"
"languages": "Abyssal, telepathy 120 ft."
"cr": "15"
"traits":
- "desc": "The nabassu darkens the area around its body in a 10-foot radius. Nonmagical\
    \ light can't illuminate this area of dim light."
  "name": "Demonic Shadows"
- "desc": "A nabassu can eat the soul of a creature it has killed within the last\
    \ hour, provided that creature is neither a Construct nor an Undead. The devouring\
    \ requires the nabassu to be within 5 feet of the corpse for at least 10 minutes,\
    \ after which it gains a number of Hit Dice (d8s) equal to half the creature's\
    \ number of Hit Dice. Roll those dice, and increase the nabassu's hit points by\
    \ the numbers rolled. For every 4 Hit Dice the nabassu gains in this way, its\
    \ attacks deal an extra dice:1d6|text(3) (1d6) damage on a hit. The nabassu\
    \ retains these benefits for 6 days. A creature devoured by a nabassu can be restored\
    \ to life only by a [wish](/3-Mechanics/CLI/spells/wish.md) spell."
  "name": "Devour Soul"
- "desc": "The nabassu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The nabassu makes one Bite attack and one Claw attack, and it uses Soul-Stealing\
    \ Gaze."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:5d12 + 6|text(38) (5d12 + 6) necrotic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d10 + 6|text(28) (4d10 + 6) force damage."
  "name": "Claw"
- "desc": "The nabassu targets one creature it can see within 30 feet of it. If the\
    \ target isn't a Construct or an Undead, it must succeed on a DC 16 Charisma saving\
    \ throw or take dice:2d12|text(13) (2d12) necrotic damage. The target's hit\
    \ point maximum is reduced by an amount equal to the necrotic damage dealt, and\
    \ the nabassu regains hit points equal to half that amount. This reduction lasts\
    \ until the target finishes a short or long rest. The target dies if its hit point\
    \ maximum is reduced to 0, and if the target is a Humanoid, it immediately rises\
    \ as a [ghoul](/3-Mechanics/CLI/bestiary/undead/ghoul.md) under the nabassu's\
    \ control."
  "name": "Soul-Stealing Gaze"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Nabassu.webp"
```
^statblock

```encounter-table
name: Nabassu
creatures:
 - 1: Nabassu
```

## Environment

swamp, underdark, urban