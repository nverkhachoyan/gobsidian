---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/26
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Orcus"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 204, Mordenkainen's Tome of Foes p. 153
---
# [Orcus](3-Mechanics\CLI\bestiary\npc/orcus-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 204, Mordenkainen's Tome of Foes p. 153*  

```statblock
"name": "Orcus (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor; 20 from with the [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md)"
"hp": !!int "405"
"hit_dice": "30d12 + 210"
"stats":
- !!int "27"
- !!int "14"
- !!int "25"
- !!int "20"
- !!int "20"
- !!int "25"
"speed": "40 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "13"
  "Constitution": !!int "15"
"skillsaves":
  "Perception": !!int "13"
  "Arcana": !!int "13"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "necrotic; poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Orcus casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 23):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day: [time\
    \ stop](/3-Mechanics/CLI/spells/time-stop.md)\n\n3/day: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md)"
  "name": "Spellcasting"
- "desc": "While holding the [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md),\
    \ Orcus casts one of the following spells (spell save DC 18), some of which require\
    \ charges; the wand has 7 charges to fuel these spells, and it regains dice:\
    \ 1d4 + 3|avg|noform (1d4 + 3) charges daily at dawn:\n\nAt will: [animate\
    \ dead](/3-Mechanics/CLI/spells/animate-dead.md) (as an action), [blight](/3-Mechanics/CLI/spells/blight.md),\
    \ [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)"
  "name": "Wand Spellcasting"
- "desc": "If Orcus fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Orcus has advantage on saving throws against spells and other magical effects."
  "name": "Magic Resistance"
- "desc": "Orcus can cast [animate dead](/3-Mechanics/CLI/spells/animate-dead.md)\
    \ (at will) and [create undead](/3-Mechanics/CLI/spells/create-undead.md) (3/day).\
    \ He chooses the level at which the spells are cast, and the creatures created\
    \ by them remain under his control indefinitely. Additionally, he can cast [create\
    \ undead](/3-Mechanics/CLI/spells/create-undead.md) even when it isn't night."
  "name": "Master of Undeath"
- "desc": "Orcus wields the [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md)."
  "name": "Special Equipment"
"actions":
- "desc": "Orcus makes three Wand of Orcus, Tail, or Necrotic Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+19 (+19) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 11|text(24) (3d8 + 11) bludgeoning damage plus\
    \ dice:2d12|text(13) (2d12) necrotic damage."
  "name": "Wand of Orcus"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) force damage plus dice:2d8|text(9)\
    \ (2d8) poison damage."
  "name": "Tail"
- "desc": "Ranged Spell Attack: dice: d20+15 (+15) to hit, range 120 ft., one\
    \ target. Hit: dice:5d8 + 7|text(29) (5d8 + 7) necrotic damage."
  "name": "Necrotic Bolt"
- "desc": "While holding the [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md),\
    \ Orcus conjures Undead creatures whose combined average hit points don't exceed\
    \ 500. These creatures magically rise up from the ground or otherwise form in\
    \ unoccupied spaces within 300 feet of Orcus and obey his commands until they\
    \ are destroyed or until he dismisses them as an action."
  "name": "Conjure Undead (1/Day)"
"legendary_actions":
- "desc": "Orcus can take 3 legendary actions, choosing from the options below. Only\
    \ one legendary action option can be used at a time and only at the end of another\
    \ creature's turn. Orcus regains spent legendary actions at the start of his turn."
  "name": ""
- "desc": "Orcus makes one Tail or Necrotic Bolt attack."
  "name": "Attack"
- "desc": "Orcus chooses a point on the ground that he can see within 100 feet of\
    \ him. A cylinder of swirling necrotic energy 60 feet tall and with a 10-foot\
    \ radius rises from that point and lasts until the end of Orcus's next turn. Creatures\
    \ in that area have vulnerability to necrotic damage."
  "name": "Creeping Death (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Orcus.webp"
```
^statblock

```encounter-table
name: Orcus
creatures:
 - 1: Orcus
```