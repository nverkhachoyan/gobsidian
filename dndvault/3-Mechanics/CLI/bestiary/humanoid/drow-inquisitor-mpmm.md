---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/14
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/cleric
- monster/type/humanoid/elf
aliases: ["Drow Inquisitor"]
NoteIcon: monster
BestiaryType: humanoid (cleric, elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 102, Mordenkainen's Tome of Foes p. 184
---
# [Drow Inquisitor](3-Mechanics\CLI\bestiary\humanoid/drow-inquisitor-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 102, Mordenkainen's Tome of Foes p. 184*  

```statblock
"name": "Drow Inquisitor (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "cleric, elf"
"alignment": "Typically  Neutral Evil"
"ac": !!int "16"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "149"
"hit_dice": "23d8 + 46"
"stats":
- !!int "11"
- !!int "15"
- !!int "14"
- !!int "16"
- !!int "21"
- !!int "20"
"speed": "30 ft."
"saves":
  "Charisma": !!int "10"
  "Wisdom": !!int "10"
  "Constitution": !!int "7"
"skillsaves":
  "Stealth": !!int "7"
  "Religion": !!int "8"
  "Insight": !!int "10"
  "Perception": !!int "10"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 20"
"languages": "Elvish, Undercommon"
"cr": "14"
"traits":
- "desc": "The drow's casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 18):\n\nAt will:\
    \ [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1/day each: [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md), [levitate](/3-Mechanics/CLI/spells/levitate.md)\
    \ (self only), [silence](/3-Mechanics/CLI/spells/silence.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Spellcasting"
- "desc": "The drow discerns when a creature in earshot speaks a lie in a language\
    \ the drow knows."
  "name": "Discern Lie"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes three Death Lance attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage plus dice:4d8|text(18)\
    \ (4d8) necrotic damage. The target's hit point maximum is reduced by an amount\
    \ equal to the necrotic damage taken. This reduction lasts until the target finishes\
    \ a long rest. The target dies if its hit point maximum is reduced to 0."
  "name": "Death Lance"
"bonus_actions":
- "desc": "The drow conjures a floating, spectral dagger within 60 feet of itself.\
    \ The drow can make a melee spell attack (dice: d20+10 (+10) to hit) against\
    \ one creature within 5 feet of the dagger. On a hit, the target takes dice:1d8\
    \ + 5|text(9) (1d8 + 5) force damage.\n\nThe dagger lasts for 1 minute. As\
    \ a bonus action on later turns, the drow can move the dagger up to 20 feet and\
    \ repeat the attack against one creature within 5 feet of the dagger."
  "name": "Spectral Dagger (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Drow%20Inquisitor.webp"
```
^statblock

```encounter-table
name: Drow Inquisitor
creatures:
 - 1: Drow Inquisitor
```

## Environment

underdark