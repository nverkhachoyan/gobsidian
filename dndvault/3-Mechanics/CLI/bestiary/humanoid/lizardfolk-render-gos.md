---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/gos
- monster/cr/3
- monster/size/large
- monster/type/humanoid/lizardfolk
aliases: ["Lizardfolk Render"]
NoteIcon: monster
BestiaryType: humanoid (lizardfolk)
SourceType: Bestiary
BookSource: Ghosts of Saltmarsh p. 241, Storm Lord's Wrath
---
# [Lizardfolk Render](3-Mechanics\CLI\bestiary\humanoid/lizardfolk-render-gos.md)
*Source: Ghosts of Saltmarsh p. 241, Storm Lord's Wrath*  

```statblock
"name": "Lizardfolk Render (GoS)"
"size": "Large"
"type": "humanoid"
"subtype": "lizardfolk"
"alignment": "Neutral"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "7d10 + 14"
"stats":
- !!int "16"
- !!int "10"
- !!int "14"
- !!int "7"
- !!int "12"
- !!int "7"
"speed": "30 ft., swim 30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Perception": !!int "3"
  "Survival": !!int "5"
"senses": "passive Perception 13"
"languages": "Draconic"
"cr": "3"
"traits":
- "desc": "The render has advantage on melee attack rolls against any creature that\
    \ doesn't have all its hit points."
  "name": "Blood Frenzy"
- "desc": "The render can hold its breath for 15 minutes."
  "name": "Hold Breath"
"actions":
- "desc": "The render makes two attacks: one with its claws and one with its bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage."
  "name": "Bite"
- "desc": "The render makes a claw attack against each creature of its choice within\
    \ 10 feet of it. A creature hit by this attack must succeed on a DC 13 Strength\
    \ saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Rend the Field (Recharge 5-6)"
"source":
- "GoS"
- "SLW"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/GoS/Lizardfolk%20Render.webp"
```
^statblock

```encounter-table
name: Lizardfolk Render
creatures:
 - 1: Lizardfolk Render
```