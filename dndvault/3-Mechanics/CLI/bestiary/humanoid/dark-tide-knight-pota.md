---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/pota
- monster/cr/3
- monster/size/medium
- monster/type/humanoid/human
aliases: ["Dark Tide Knight"]
NoteIcon: monster
BestiaryType: humanoid (human)
SourceType: Bestiary
BookSource: Princes of the Apocalypse p. 205, Storm Lord's Wrath
---
# [Dark Tide Knight](3-Mechanics\CLI\bestiary\humanoid/dark-tide-knight-pota.md)
*Source: Princes of the Apocalypse p. 205, Storm Lord's Wrath*  

```statblock
"name": "Dark Tide Knight (PotA)"
"size": "Medium"
"type": "humanoid"
"subtype": "human"
"alignment": "Lawful Evil"
"ac": !!int "13"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "17"
- !!int "16"
- !!int "14"
- !!int "10"
- !!int "11"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "7"
  "Stealth": !!int "7"
"senses": "passive Perception 10"
"languages": "Common"
"cr": "3"
"traits":
- "desc": "The knight is magically bound to a beast with an innate swimming speed\
    \ trained to serve as its mount. While mounted on this beast, the knight gains\
    \ the beast's senses and ability to breathe underwater. The bonded mount obeys\
    \ the knight's commands. If its mount dies, the knight can train a new beast to\
    \ serve as its bonded mount, a process requiring a month."
  "name": "Bonded Mount"
- "desc": "The knight deals an extra dice:2d6|text(7) (2d6) damage when it hits\
    \ a target with a weapon attack and has advantage on the attack roll, or when\
    \ the target is within 5 feet of an ally of the knight that isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and the knight doesn't have disadvantage on the attack roll."
  "name": "Sneak Attack"
"actions":
- "desc": "The knight makes two shortsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortsword"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d12 + 3|text(9) (1d12 + 3) piercing damage."
  "name": "Lance"
"reactions":
- "desc": "When an attacker the knight can see hits it with an attack, the knight\
    \ can halve the damage against it."
  "name": "Uncanny Dodge"
"source":
- "PotA"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/PotA/Dark%20Tide%20Knight.webp"
```
^statblock

```encounter-table
name: Dark Tide Knight
creatures:
 - 1: Dark Tide Knight
```