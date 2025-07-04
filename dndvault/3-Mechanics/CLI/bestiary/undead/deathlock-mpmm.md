---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/urban
- monster/size/medium
- monster/type/undead/warlock
aliases: ["Deathlock"]
NoteIcon: monster
BestiaryType: undead (warlock)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 86, Mordenkainen's Tome of Foes p. 128
---
# [Deathlock](3-Mechanics\CLI\bestiary\undead/deathlock-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 86, Mordenkainen's Tome of Foes p. 128*  

```statblock
"name": "Deathlock (MPMM)"
"size": "Medium"
"type": "undead"
"subtype": "warlock"
"alignment": "Typically  Neutral Evil"
"ac": !!int "12"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "11"
- !!int "15"
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Intelligence": !!int "4"
"skillsaves":
  "History": !!int "4"
  "Arcana": !!int "4"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "the languages it knew in life"
"cr": "4"
"traits":
- "desc": "The deathlock casts one of the following spells, using Charisma as the\
    \ spellcasting ability (spell save DC 13):\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\n\n1/day each: [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [hunger of Hadar](/3-Mechanics/CLI/spells/hunger-of-hadar.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [spider climb](/3-Mechanics/CLI/spells/spider-climb.md)"
  "name": "Spellcasting"
- "desc": "The deathlock has advantage on saving throws against any effect that turns\
    \ Undead."
  "name": "Turn Resistance"
- "desc": "The deathlock doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The deathlock makes two Deathly Claw or Grave Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) necrotic damage."
  "name": "Deathly Claw"
- "desc": "Ranged Spell Attack: dice: d20+5 (+5) to hit, range 120 ft., one\
    \ target. Hit: dice:2d10 + 3|text(14) (2d10 + 3) necrotic damage."
  "name": "Grave Bolt"
"source":
- "MPMM"
- "MTF"
- "AATM"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Deathlock.webp"
```
^statblock

```encounter-table
name: Deathlock
creatures:
 - 1: Deathlock
```

## Environment

urban