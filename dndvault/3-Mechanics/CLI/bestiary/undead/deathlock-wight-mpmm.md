---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/size/medium
- monster/type/undead/warlock
aliases: ["Deathlock Wight"]
NoteIcon: monster
BestiaryType: undead (warlock)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 87, Mordenkainen's Tome of Foes p. 129
---
# [Deathlock Wight](3-Mechanics\CLI\bestiary\undead/deathlock-wight-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 87, Mordenkainen's Tome of Foes p. 129*  

```statblock
"name": "Deathlock Wight (MPMM)"
"size": "Medium"
"type": "undead"
"subtype": "warlock"
"alignment": "Typically  Neutral Evil"
"ac": !!int "12"
"hp": !!int "37"
"hit_dice": "5d8 + 15"
"stats":
- !!int "11"
- !!int "14"
- !!int "16"
- !!int "12"
- !!int "14"
- !!int "16"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Perception": !!int "4"
  "Arcana": !!int "3"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "the languages it knew in life"
"cr": "3"
"traits":
- "desc": "The deathlock casts one of the following spells, using Charisma as the\
    \ spellcasting ability (spell save DC 13):\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)\n\
    \n1/day each: [fear](/3-Mechanics/CLI/spells/fear.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md)"
  "name": "Spellcasting"
- "desc": "While in sunlight, the deathlock has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "The deathlock doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The deathlock makes two Life Drain or Grave Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) necrotic damage. The target must succeed\
    \ on a DC 13 Constitution saving throw, or its hit point maximum is reduced by\
    \ an amount equal to the damage taken. This reduction lasts until the target finishes\
    \ a long rest. The target dies if its hit point maximum is reduced to 0.\n\nA\
    \ Humanoid slain by this attack rises 24 hours later as a [zombie](/3-Mechanics/CLI/bestiary/undead/zombie.md)\
    \ under the deathlock's control, unless the Humanoid is restored to life or its\
    \ body is destroyed. The deathlock can have no more than twelve zombies under\
    \ its control at one time."
  "name": "Life Drain"
- "desc": "Ranged Spell Attack: dice: d20+5 (+5) to hit, range 60 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) necrotic damage."
  "name": "Grave Bolt"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Deathlock%20Wight.webp"
```
^statblock

```encounter-table
name: Deathlock Wight
creatures:
 - 1: Deathlock Wight
```