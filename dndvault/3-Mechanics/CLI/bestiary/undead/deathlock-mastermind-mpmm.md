---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/8
- monster/environment/urban
- monster/size/medium
- monster/type/undead/warlock
aliases: ["Deathlock Mastermind"]
NoteIcon: monster
BestiaryType: undead (warlock)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 87, Mordenkainen's Tome of Foes p. 129
---
# [Deathlock Mastermind](3-Mechanics\CLI\bestiary\undead/deathlock-mastermind-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 87, Mordenkainen's Tome of Foes p. 129*  

```statblock
"name": "Deathlock Mastermind (MPMM)"
"size": "Medium"
"type": "undead"
"subtype": "warlock"
"alignment": "Typically  Neutral Evil"
"ac": !!int "13"
"hp": !!int "110"
"hit_dice": "20d8 + 20"
"stats":
- !!int "11"
- !!int "16"
- !!int "12"
- !!int "15"
- !!int "12"
- !!int "17"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Intelligence": !!int "5"
"skillsaves":
  "Perception": !!int "4"
  "History": !!int "5"
  "Arcana": !!int "5"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "the languages it knew in life"
"cr": "8"
"traits":
- "desc": "The deathlock casts one of the following spells, using Charisma as the\
    \ spellcasting ability (spell save DC 14):\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\n1/day each:\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Spellcasting"
- "desc": "Magical darkness doesn't impede the deathlock's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The deathlock has advantage on saving throws against any effect that turns\
    \ Undead."
  "name": "Turn Resistance"
- "desc": "The deathlock doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The deathlock makes two Deathly Claw or Grave Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: 13 (dice: 3d6 + 3|avg|noform (3d6 + 3) necrotic damage)."
  "name": "Deathly Claw"
- "desc": "Ranged Spell Attack: dice: d20+6 (+6) to hit, range 120 ft., one\
    \ target. Hit: dice:3d8|text(13) (3d8) necrotic damage. If the target is\
    \ Large or smaller, it must succeed on a DC 16 Strength saving throw or become\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) as shadowy tendrils\
    \ wrap around it for 1 minute. A [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ target can use its action to repeat the saving throw, ending the effect on itself\
    \ on a success."
  "name": "Grave Bolt"
"source":
- "MPMM"
- "MTF"
- "AATM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Deathlock%20Mastermind.webp"
```
^statblock

```encounter-table
name: Deathlock Mastermind
creatures:
 - 1: Deathlock Mastermind
```

## Environment

urban