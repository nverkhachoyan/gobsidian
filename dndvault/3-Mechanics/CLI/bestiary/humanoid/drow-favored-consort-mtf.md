---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/18
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Favored Consort"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 183
---
# [Drow Favored Consort](3-Mechanics\CLI\bestiary\humanoid/drow-favored-consort-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 183*  

> [!quote]- A quote from Mordenkainen  
> 
> Many creatures enjoy torture, but the dark elves have made it into an exquisite art.

## Drow Favored Consort

Nearly all priestesses of Lolth, including the powerful matron mothers, take attractive drow as their consorts. Often these individuals serve no purpose beyond pleasure, breeding, or both, but sometimes consorts can gain the ear of their priestess and be relied on to provide useful advice. No position of consort is assured for long; priestesses are infamous for being fickle with their favor, which are they are especially glad to lavish on a consort who combines beauty with magical might.

```statblock
"name": "Drow Favored Consort (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "18 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "225"
"hit_dice": "30d8 + 90"
"stats":
- !!int "15"
- !!int "20"
- !!int "16"
- !!int "18"
- !!int "15"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "10"
  "Dexterity": !!int "11"
  "Constitution": !!int "9"
"skillsaves":
  "Athletics": !!int "8"
  "Stealth": !!int "11"
  "Perception": !!int "8"
  "Acrobatics": !!int "11"
"senses": "darkvision 120 ft., passive Perception 18"
"languages": "Elvish, Undercommon"
"cr": "18"
"traits":
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 18).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md)\n\n\
    1/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "The drow is a 11th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 18, dice: d20+10 (+10) to hit with spell attacks). It has\
    \ the following wizard spells prepared:\n\nCantrips (at will): [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md), [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\
    \n1st level (4 slots): [burning hands](/3-Mechanics/CLI/spells/burning-hands.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [magic missile](/3-Mechanics/CLI/spells/magic-missile.md),\
    \ [shield](/3-Mechanics/CLI/spells/shield.md)\n\n2nd level (3 slots): [gust\
    \ of wind](/3-Mechanics/CLI/spells/gust-of-wind.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [shatter](/3-Mechanics/CLI/spells/shatter.md)\n\
    \n3rd level (3 slots): [counterspell](/3-Mechanics/CLI/spells/counterspell.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [haste](/3-Mechanics/CLI/spells/haste.md)\n\
    \n4th level (3 slots): [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [Otiluke's resilient sphere](/3-Mechanics/CLI/spells/otilukes-resilient-sphere.md)\n\
    \n5th level (2 slots): [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md)\n\
    \n6th level (1 slots): [chain lightning](/3-Mechanics/CLI/spells/chain-lightning.md)"
  "name": "Spellcasting"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "When the drow uses its action to cast a spell, it can make one weapon attack\
    \ as a bonus action."
  "name": "War Magic"
"actions":
- "desc": "The drow makes three scimitar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) slashing damage plus dice:4d8|text(18)\
    \ (4d8) poison damage. In addition, the target has disadvantage on the next\
    \ saving throw it makes against a spell the drow casts before the end of the drow's\
    \ next turn."
  "name": "Scimitar"
- "desc": "Ranged Weapon Attack: dice: d20+11 (+11) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage, and the\
    \ target must succeed on a DC 13 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour. If the saving throw fails by 5 or more, the target is also [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way.\
    \ The target regains consciousness if it takes damage or if another creature takes\
    \ an action to shake it."
  "name": "Hand Crossbow"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20Favored%20Consort.webp"
```
^statblock

```encounter-table
name: Drow Favored Consort
creatures:
 - 1: Drow Favored Consort
```

## Environment

underdark