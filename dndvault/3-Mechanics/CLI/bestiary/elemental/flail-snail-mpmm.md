---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/elemental
aliases: ["Flail Snail"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 126, Volo's Guide to Monsters p. 144
---
# [Flail Snail](3-Mechanics\CLI\bestiary\elemental/flail-snail-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 126, Volo's Guide to Monsters p. 144*  

```statblock
"name": "Flail Snail (MPMM)"
"size": "Large"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "5d10 + 25"
"stats":
- !!int "17"
- !!int "5"
- !!int "20"
- !!int "3"
- !!int "10"
- !!int "5"
"speed": "10 ft."
"damage_immunities": "fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 10"
"languages": ""
"cr": "3"
"traits":
- "desc": "The snail has advantage on saving throws against spells, and any creature\
    \ making a spell attack against the snail has disadvantage on the attack roll.\n\
    \nIf the snail succeeds on its saving throw against a spell or a spell's attack\
    \ roll misses it, the snail's shell converts some of the spell's energy into a\
    \ burst of destructive force if the spell is of 1st level or higher; each creature\
    \ within 30 feet of the snail must make a DC 15 Constitution saving throw, taking\
    \ dice:1d6|text(3) (1d6) force damage per level of the spell on a failed save,\
    \ or half as much damage on a successful one."
  "name": "Antimagic Shell"
"actions":
- "desc": "The snail makes five Flail Tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage."
  "name": "Flail Tentacle"
- "desc": "The snail's shell emits dazzling, colored light until the end of the snail's\
    \ next turn. During this time, the shell sheds bright light in a 30-foot radius\
    \ and dim light for an additional 30 feet, and creatures that can see the snail\
    \ have disadvantage on attack rolls against it. In addition, any creature within\
    \ the bright light and able to see the snail when this power is activated must\
    \ succeed on a DC 15 Wisdom saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the light ends."
  "name": "Scintillating Shell (Recharges after a Short or Long Rest)"
- "desc": "The flail snail withdraws into its shell. Until it emerges, it gains a\
    \ +4 bonus to its AC and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ It can emerge from its shell as a bonus action on its turn."
  "name": "Shell Defense"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Flail%20Snail.webp"
```
^statblock

```encounter-table
name: Flail Snail
creatures:
 - 1: Flail Snail
```

## Environment

forest, swamp, underdark