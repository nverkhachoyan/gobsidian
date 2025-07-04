---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/
- monster/size/medium
- monster/type/humanoid
aliases: ["Expert"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Expert](3-Mechanics\CLI\bestiary\humanoid/expert-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Expert (SLW)"
"size": "Medium"
"type": "humanoid"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "10"
- !!int "16"
- !!int "12"
- !!int "14"
- !!int "10"
- !!int "14"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "6"
"skillsaves":
  "Sleight of Hand": !!int "6"
  "Stealth": !!int "9"
  "Performance": !!int "5"
  "Acrobatics": !!int "9"
  "Persuasion": !!int "5"
"senses": "passive Perception 10"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The expert can take the Help action as a bonus action, and the creature\
    \ who receives the help gains a dice: 1d6|avg|noform (1d6) bonus to the dice:\
    \ d20|avg|noform (d20) roll. If that roll is an attack roll, the creature can\
    \ forgo adding the bonus to it, and then if the attack hits, the creature can\
    \ add the bonus to the attack's damage roll against one target."
  "name": "Helpful"
- "desc": "The expert has [thieves' tools](/3-Mechanics/CLI/items/thieves-tools.md)\
    \ and a musical instrument."
  "name": "Tools"
"actions":
- "desc": "The expert can attack twice, instead of once, whenever it takes the [Attack](/3-Mechanics/CLI/rules/actions.md#Attack)\
    \ action on its turn."
  "name": "Extra Attack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortsword"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortbow"
"source":
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Expert.webp"
```
^statblock

```encounter-table
name: Expert
creatures:
 - 1: Expert
```