---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/urban
- monster/size/medium
- monster/type/aberration
aliases: ["Star Spawn Seer"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 230, Mordenkainen's Tome of Foes p. 236
---
# [Star Spawn Seer](3-Mechanics\CLI\bestiary\aberration/star-spawn-seer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 230, Mordenkainen's Tome of Foes p. 236*  

```statblock
"name": "Star Spawn Seer (MPMM)"
"size": "Medium"
"type": "aberration"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "153"
"hit_dice": "18d8 + 72"
"stats":
- !!int "14"
- !!int "12"
- !!int "18"
- !!int "22"
- !!int "19"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "6"
  "Wisdom": !!int "9"
  "Intelligence": !!int "11"
"skillsaves":
  "Perception": !!int "9"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 19"
"languages": "Common, Deep Speech, Undercommon"
"cr": "13"
"traits":
- "desc": "The seer can move through other creatures and objects as if they were difficult\
    \ terrain, and its movement doesn't provoke opportunity attacks.\n\nEach creature\
    \ it moves through takes dice:1d10|text(5) (1d10) psychic damage; no creature\
    \ can take this damage more than once per turn.\n\nThe seer takes dice:1d10|text(5)\
    \ (1d10) force damage if it ends its turn inside an object."
  "name": "Out-Of-Phase Movement"
"actions":
- "desc": "The seer makes two Comet Staff or Psychic Orb attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d8 + 6|text(10) (1d8 + 6) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) psychic damage, and if the target is a creature, it must succeed on\
    \ a DC 19 Constitution saving throw or be [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ until the end of its next turn."
  "name": "Comet Staff"
- "desc": "Ranged Spell Attack: dice: d20+11 (+11) to hit, range 120 feet, one\
    \ creature. Hit: dice:5d10|text(27) (5d10) psychic damage."
  "name": "Psychic Orb"
- "desc": "The seer warps space around one creature it can see within 30 feet of it.\
    \ That creature must make a DC 19 Wisdom saving throw. On a failed save, the target,\
    \ along with any equipment it is wearing or carrying, is teleported up to 60 feet\
    \ to an unoccupied space the seer can see, and then each creature within 10 feet\
    \ of the target's original space takes dice:6d12|text(39) (6d12) psychic damage.\
    \ On a successful save, the target takes dice:3d12|text(19) (3d12) psychic\
    \ damage and isn't teleported."
  "name": "Collapse Distance (Recharge 6)"
"reactions":
- "desc": "When the seer would be hit by an attack roll, it teleports, along with\
    \ any equipment it is wearing or carrying, exchanging positions with another star\
    \ spawn it can see within 60 feet of it. The other star spawn is hit by the attack\
    \ instead."
  "name": "Bend Space"
"source":
- "MPMM"
- "MTF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Star%20Spawn%20Seer.webp"
```
^statblock

```encounter-table
name: Star Spawn Seer
creatures:
 - 1: Star Spawn Seer
```

## Environment

mountain, swamp, urban