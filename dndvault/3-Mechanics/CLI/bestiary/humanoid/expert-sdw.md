---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/sdw
- monster/cr/
- monster/size/medium
- monster/type/humanoid
aliases: ["Expert"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Sleeping Dragon's Wake
---
# [Expert](3-Mechanics\CLI\bestiary\humanoid/expert-sdw.md)
*Source: Sleeping Dragon's Wake*  

```statblock
"name": "Expert (SDW)"
"size": "Medium"
"type": "humanoid"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "55"
"hit_dice": "10d8 + 10"
"stats":
- !!int "10"
- !!int "18"
- !!int "12"
- !!int "14"
- !!int "10"
- !!int "14"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "8"
"skillsaves":
  "Sleight of Hand": !!int "8"
  "Stealth": !!int "12"
  "Performance": !!int "6"
  "Acrobatics": !!int "12"
  "Persuasion": !!int "6"
"senses": "passive Perception 10"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The expert can take the Help action as a bonus action, and the creature\
    \ who receives the help gains a dice: 1d6|avg|noform (1d6) bonus to the dice:\
    \ d20|avg|noform (d20) roll. If that roll is an attack roll, the creature can\
    \ forgo adding the bonus to it, and then if the attack hits, the creature can\
    \ add the bonus to the attack's damage roll against one target."
  "name": "Helpful"
- "desc": "When the expert is not [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and subjected to an effect that allows it to make a Dexterity saving throw to\
    \ take only half damage, it instead takes no damage if it succeeds on the saving\
    \ throw, and only half damage if it failed."
  "name": "Evasion"
- "desc": "The expert has thieves' tools and a musical instrument."
  "name": "Tools"
"actions":
- "desc": "The expert can attack twice, instead of once, whenever it takes the attack\
    \ action on its turn."
  "name": "Extra Attack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Shortsword"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 4|text(6) (1d4 + 4) piercing\
    \ damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+8 (+8) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Shortbow"
"source":
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SDW/Expert.webp"
```
^statblock

```encounter-table
name: Expert
creatures:
 - 1: Expert
```