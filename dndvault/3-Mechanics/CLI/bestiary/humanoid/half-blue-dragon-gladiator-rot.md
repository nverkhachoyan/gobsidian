---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/rot
- monster/cr/5
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Half-Blue Dragon Gladiator"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: The Rise of Tiamat p. 55, Divine Contention
---
# [Half-Blue Dragon Gladiator](3-Mechanics\CLI\bestiary\humanoid/half-blue-dragon-gladiator-rot.md)
*Source: The Rise of Tiamat p. 55, Divine Contention*  

```statblock
"name": "Half-Blue Dragon Gladiator (RoT)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "112"
"hit_dice": "15d8 + 45"
"stats":
- !!int "18"
- !!int "15"
- !!int "16"
- !!int "10"
- !!int "12"
- !!int "15"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "5"
  "Strength": !!int "7"
  "Constitution": !!int "6"
"skillsaves":
  "Intimidation": !!int "5"
  "Athletics": !!int "10"
"damage_resistances": "lightning"
"senses": "blindsight 10 ft., darkvision 60 ft., passive Perception 11"
"languages": "any one language (usually Common), Draconic"
"cr": "5"
"traits":
- "desc": "The half-dragon has advantage on saving throws against being [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Brave"
- "desc": "A melee weapon deals one extra die of its damage when the half-dragon hits\
    \ with it (included in the attack)."
  "name": "Brute"
"actions":
- "desc": "The half-dragon makes three melee attacks or two ranged attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ and range 20/60 ft., one target. Hit: dice:2d6 + 4|text(11) (2d6 + 4)\
    \ piercing damage, or dice:2d8 + 4|text(13) (2d8 + 4) piercing damage if used\
    \ with two hands to make a melee attack."
  "name": "Spear"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d4 + 4|text(9) (2d4 + 4) bludgeoning damage. If the target\
    \ is a Medium or smaller creature, it must succeed on a DC 15 Strength saving\
    \ throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Shield Bash"
- "desc": "The half-dragon exhales lightning in a 30-foot line that is 5 feet wide.\
    \ Each creature in that line must make a DC 12 Dexterity saving throw, taking\
    \ dice:4d10|text(22) (4d10) lightning damage on a failed save, or half as\
    \ much damage on a successful one."
  "name": "Lightning Breath (Recharge 5-6)"
"reactions":
- "desc": "The half-dragon adds 3 to its AC against one melee attack that would hit\
    \ it. To do so, the half-dragon must see the attacker and be wielding a melee\
    \ weapon."
  "name": "Parry"
"source":
- "RoT"
- "DC"
- "ToD"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/RoT/Half-Blue%20Dragon%20Gladiator.webp"
```
^statblock

```encounter-table
name: Half-Blue Dragon Gladiator
creatures:
 - 1: Half-Blue Dragon Gladiator
```