---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tftyp
- monster/cr/8
- monster/size/huge
- monster/type/beast
aliases: ["Huge Giant Crab"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Tales from the Yawning Portal p. 103, Storm Lord's Wrath
---
# [Huge Giant Crab](3-Mechanics\CLI\bestiary\beast/huge-giant-crab-tftyp.md)
*Source: Tales from the Yawning Portal p. 103, Storm Lord's Wrath*  

```statblock
"name": "Huge Giant Crab (TftYP)"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "161"
"hit_dice": "14d12 + 70"
"stats":
- !!int "20"
- !!int "15"
- !!int "20"
- !!int "1"
- !!int "9"
- !!int "3"
"speed": "30 ft., swim 30 ft."
"skillsaves":
  "Stealth": !!int "5"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)"
"senses": "blindsight 30 ft., passive Perception 9"
"languages": ""
"cr": "8"
"traits":
- "desc": "On one of its claws, the crab wears a rune-covered copper band that makes\
    \ it immune to being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed).\
    \ The copper band is worthless as a treasure, as the magic is keyed to this crab."
  "name": "Banded Claw"
- "desc": "The crab can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d10 + 5|text(27) (4d10 + 5) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), escape DC 14.\
    \ The crab has two claws, each of which can grapple only one target."
  "name": "Claw"
"source":
- "TftYP"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TftYP/Huge%20Giant%20Crab.webp"
```
^statblock

```encounter-table
name: Huge Giant Crab
creatures:
 - 1: Huge Giant Crab
```