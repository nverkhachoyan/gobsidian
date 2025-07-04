---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/dip
- monster/cr/1-2
- monster/size/medium
- monster/type/humanoid/human
aliases: ["Don-Jon Raskin"]
NoteIcon: monster
BestiaryType: humanoid (human)
SourceType: Bestiary
BookSource: Dragon of Icespire Peak p. 56
---
# [Don-Jon Raskin](3-Mechanics\CLI\bestiary\npc/don-jon-raskin-dip.md)
*Source: Dragon of Icespire Peak p. 56*  

Adventurers who undertake the Mountain's Toe Quest meet Don-Jon Raskin, a fearless troubleshooter who has experienced more than his fair share of adventures.

```statblock
"name": "Don-Jon Raskin (DIP)"
"size": "Medium"
"type": "humanoid"
"subtype": "human"
"alignment": "Neutral"
"ac": !!int "10"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "11"
- !!int "10"
- !!int "13"
- !!int "12"
- !!int "10"
- !!int "14"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "2"
  "Constitution": !!int "3"
"skillsaves":
  "Deception": !!int "4"
  "Persuasion": !!int "4"
"senses": "passive Perception 10"
"languages": "Common, Dwarvish"
"cr": "1/2"
"traits":
- "desc": "Don-Jon has advantage on saving throws against being [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Brave"
- "desc": "If damage reduces Don-Jon to 0 hit points, he drops to 1 hit point instead\
    \ and gains advantage on attack rolls until the end of his next turn."
  "name": "Not Dead Yet (Recharges after a Long Rest)"
"actions":
- "desc": "Don-Jon makes three melee attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4|text(2) (1d4) piercing damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+2 (+2) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d4|text(2) (1d4) bludgeoning damage."
  "name": "Sling"
"source":
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/DIP/Don-Jon%20Raskin.webp"
```
^statblock

```encounter-table
name: Don-Jon Raskin
creatures:
 - 1: Don-Jon Raskin
```