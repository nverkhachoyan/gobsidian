---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Warlord"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 257, Volo's Guide to Monsters p. 220
---
# [Warlord](3-Mechanics\CLI\bestiary\humanoid/warlord-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 257, Volo's Guide to Monsters p. 220*  

```statblock
"name": "Warlord (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "229"
"hit_dice": "27d8 + 108"
"stats":
- !!int "20"
- !!int "16"
- !!int "18"
- !!int "12"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "7"
  "Strength": !!int "9"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "8"
  "Athletics": !!int "9"
  "Perception": !!int "5"
  "Persuasion": !!int "8"
"senses": "passive Perception 15"
"languages": "any two languages"
"cr": "12"
"traits":
- "desc": "The warlord can reroll a saving throw it fails. It must use the new roll."
  "name": "Indomitable (3/Day)"
- "desc": "The warlord regains 10 hit points at the start of its turn if it has fewer\
    \ than half its hit points but at least 1 hit point."
  "name": "Survivor"
"actions":
- "desc": "The warlord makes two Greatsword or Short bow attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) slashing damage."
  "name": "Greatsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortbow"
"legendary_actions":
- "desc": "The warlord targets one ally it can see within 30 feet of it. if the target\
    \ can see and hear the warlord, the target can make one weapon attack as a reaction\
    \ and gains advantage on the attack roll."
  "name": "Command Ally"
- "desc": "The warlord makes one Greatsword or Shortbow attack."
  "name": "Weapon Attack"
- "desc": "The warlord targets one creature it can see within 30 feet of it. If the\
    \ target can see and hear it, the target must succeed on a DC 16 Wisdom saving\
    \ throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until\
    \ the end of warlord's next turn."
  "name": "Frighten Foe (Costs 2 Actions)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Warlord.webp"
```
^statblock

```encounter-table
name: Warlord
creatures:
 - 1: Warlord
```

## Environment

urban