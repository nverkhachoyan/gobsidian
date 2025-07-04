---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/14
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Inquisitor"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 184
---
# [Drow Inquisitor](3-Mechanics\CLI\bestiary\humanoid/drow-inquisitor-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 184*  

> [!quote]- A quote from Mordenkainen  
> 
> Drow mages can be quite learned and skilled in magic. Some of them can even cast my spells.

## Drow Inquisitor

Drow expect treachery. After all, the Spider Queen encourages it. A certain amount of backstabbing and double-crossing can be managed, but too much can undermine an entire community. To keep some semblance of order and to root out traitors, drow priestesses employ inquisitors. Chosen from the ranks of the priesthood, these female drow possess authority equaled only by the matrons of the noble houses. Anyone they decide is at odds with the hierarchy faces torture and usually an excruciating death.

```statblock
"name": "Drow Inquisitor (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "143"
"hit_dice": "22d8 + 44"
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
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 18).\
    \ She can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [detect\
    \ magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day each: [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md),\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "The drow is a 12th-level spellcaster. Her spellcasting ability is Wisdom\
    \ (spell save DC 18, dice: d20+10 (+10) to hit with spell attacks). She has\
    \ the following cleric spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [resistance](/3-Mechanics/CLI/spells/resistance.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1st level (4 slots): [bane](/3-Mechanics/CLI/spells/bane.md), [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [inflict wounds](/3-Mechanics/CLI/spells/inflict-wounds.md)\n\n2nd level (3\
    \ slots): [blindness/deafness](/3-Mechanics/CLI/spells/blindness-deafness.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md), [spiritual weapon](/3-Mechanics/CLI/spells/spiritual-weapon.md)\n\
    \n3rd level (3 slots): [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [magic circle](/3-Mechanics/CLI/spells/magic-circle.md)\n\
    \n4th level (3 slots): [banishment](/3-Mechanics/CLI/spells/banishment.md),\
    \ [divination](/3-Mechanics/CLI/spells/divination.md), [freedom of movement](/3-Mechanics/CLI/spells/freedom-of-movement.md)\n\
    \n5th level (2 slots): [contagion](/3-Mechanics/CLI/spells/contagion.md),\
    \ [dispel evil and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md), [insect\
    \ plague](/3-Mechanics/CLI/spells/insect-plague.md)\n\n6th level (1 slots):\
    \ [harm](/3-Mechanics/CLI/spells/harm.md), [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Spellcasting"
- "desc": "The drow knows when she hears a creature speak a lie in a language she\
    \ knows."
  "name": "Discern Lie"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "The drow has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes three death lance attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage plus dice:4d8|text(18)\
    \ (4d8) necrotic damage. The target's hit point maximum is reduced by an amount\
    \ equal to the necrotic damage it takes. This reduction lasts until the target\
    \ finishes a long rest. The target dies if its hit point maximum is reduced to\
    \ 0."
  "name": "Death Lance"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20Inquisitor.webp"
```
^statblock

```encounter-table
name: Drow Inquisitor
creatures:
 - 1: Drow Inquisitor
```

## Environment

underdark