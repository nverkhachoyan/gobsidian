---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/
- monster/size/large
- monster/type/construct
aliases: ["Sacred Statue"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 114, Mordenkainen's Tome of Foes p. 194
---
# [Sacred Statue](3-Mechanics\CLI\bestiary\construct/sacred-statue-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 114, Mordenkainen's Tome of Foes p. 194*  

```statblock
"name": "Sacred Statue (MPMM)"
"size": "Large"
"type": "construct"
"alignment": "as the eidolon's alignment"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "95"
"hit_dice": "10d10 + 40"
"stats":
- !!int "19"
- !!int "8"
- !!int "19"
- !!int "14"
- !!int "19"
- !!int "16"
"speed": "25 ft."
"saves":
  "Wisdom": !!int "8"
"damage_resistances": "acid; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "the languages the [eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mpmm.md)\
  \ knew in life"
"traits":
- "desc": "If the statue is motionless at the start of combat, it has advantage on\
    \ its initiative roll. Moreover, if a creature hasn't observed the statue move\
    \ or act, that creature must succeed on a DC 18 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check to discern that the statue isn't an object."
  "name": "False Appearance"
- "desc": "The [eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mpmm.md) that enters\
    \ the statue remains inside it until the statue drops to 0 hit points, the eidolon\
    \ uses a bonus action to move out of the statue, or the eidolon is turned or forced\
    \ out by an effect such as the [dispel evil and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md)\
    \ spell. When the eidolon leaves the statue, it appears in an unoccupied space\
    \ within 5 feet of the statue."
  "name": "Ghostly Inhabitant"
- "desc": "Without an [eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mpmm.md)\
    \ inside, the statue is an object."
  "name": "Inert"
- "desc": "The statue doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The statue makes two Slam or Rock attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:6d12 + 4|text(43) (6d12 + 4) bludgeoning damage."
  "name": "Slam"
- "desc": "Ranged Weapon Attack: dice: d20+8 (+8) to hit, range 60 ft./240 ft.,\
    \ one target. Hit: dice:6d10 + 4|text(37) (6d10 + 4) bludgeoning damage."
  "name": "Rock"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Sacred%20Statue.webp"
```
^statblock

```encounter-table
name: Sacred Statue
creatures:
 - 1: Sacred Statue
```