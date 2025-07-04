---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/underdark
- monster/size/medium
- monster/type/undead/mind-flayer
- monster/type/undead/wizard
aliases: ["Alhoon"]
NoteIcon: monster
BestiaryType: undead (mind flayer, wizard)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 43, Volo's Guide to Monsters p. 172
---
# [Alhoon](3-Mechanics\CLI\bestiary\undead/alhoon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 43, Volo's Guide to Monsters p. 172*  

```statblock
"name": "Alhoon (MPMM)"
"size": "Medium"
"type": "undead"
"subtype": "mind flayer, wizard"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "150"
"hit_dice": "20d8 + 60"
"stats":
- !!int "11"
- !!int "12"
- !!int "16"
- !!int "19"
- !!int "17"
- !!int "17"
"speed": "30 ft., fly 15 ft. (hover)"
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "7"
  "Intelligence": !!int "8"
  "Constitution": !!int "7"
"skillsaves":
  "Deception": !!int "7"
  "Stealth": !!int "5"
  "Insight": !!int "7"
  "Perception": !!int "7"
  "History": !!int "8"
  "Arcana": !!int "8"
"damage_resistances": "cold, lightning, necrotic"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 17"
"languages": "Deep Speech, Undercommon, telepathy 120 ft."
"cr": "10"
"traits":
- "desc": "The alhoon casts one of the following spells, requiring no material components\
    \ and using Intelligence as the spellcasting ability (spell save DC 16):\n\nAt\
    \ will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [detect\
    \ magic](/3-Mechanics/CLI/spells/detect-magic.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1/day each:\
    \ [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md), [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [modify memory](/3-Mechanics/CLI/spells/modify-memory.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only), [wall of\
    \ force](/3-Mechanics/CLI/spells/wall-of-force.md)"
  "name": "Spellcasting"
- "desc": "The alhoon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The alhoon has advantage on saving throws against any effect that turns\
    \ Undead."
  "name": "Turn Resistance"
"actions":
- "desc": "The alhoon makes two Chilling Grasp or Arcane Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Spell Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d6|text(14) (4d6) cold damage, and the alhoon regains 14 hit\
    \ points."
  "name": "Chilling Grasp"
- "desc": "Ranged Spell Attack: dice: d20+8 (+8) to hit, range 120 ft., one\
    \ target. Hit: dice:8d6|text(28) (8d6) force damage."
  "name": "Arcane Bolt"
- "desc": "The alhoon magically emits psychic energy in a 60-foot cone. Each creature\
    \ in that area must succeed on a DC 16 Intelligence saving throw or take dice:4d8\
    \ + 4|text(22) (4d8 + 4) psychic damage and be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. A target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Mind Blast (Recharge 5-6)"
"reactions":
- "desc": "The alhoon targets one creature it can see within 60 feet of it that is\
    \ casting a spell. If the spell is 3rd level or lower, the spell fails, but any\
    \ spell slots or charges are not wasted."
  "name": "Negate Spell (3/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Alhoon.webp"
```
^statblock

```encounter-table
name: Alhoon
creatures:
 - 1: Alhoon
```

## Environment

underdark