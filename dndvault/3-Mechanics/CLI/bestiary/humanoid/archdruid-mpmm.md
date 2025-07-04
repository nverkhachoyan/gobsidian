---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/forest
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underwater
- monster/size/medium
- monster/type/humanoid/druid
aliases: ["Archdruid"]
NoteIcon: monster
BestiaryType: humanoid (druid)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 48, Volo's Guide to Monsters p. 210
---
# [Archdruid](3-Mechanics\CLI\bestiary\humanoid/archdruid-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 48, Volo's Guide to Monsters p. 210*  

```statblock
"name": "Archdruid (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "druid"
"alignment": "Any alignment"
"ac": !!int "14"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
"hp": !!int "154"
"hit_dice": "28d8 + 28"
"stats":
- !!int "14"
- !!int "14"
- !!int "12"
- !!int "12"
- !!int "20"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "9"
  "Intelligence": !!int "5"
"skillsaves":
  "Medicine": !!int "9"
  "Nature": !!int "5"
  "Perception": !!int "9"
"senses": "passive Perception 19"
"languages": "Druidic plus any two languages"
"cr": "12"
"traits":
- "desc": "The archdruid casts one of the following spells, using Wisdom as the spellcasting\
    \ ability (spell save DC 17):\n\nAt will: [beast sense](/3-Mechanics/CLI/spells/beast-sense.md),\
    \ [entangle](/3-Mechanics/CLI/spells/entangle.md), [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md)\n\
    \n1/day each: [commune with nature](/3-Mechanics/CLI/spells/commune-with-nature.md)\
    \ (as an action), [mass cure wounds](/3-Mechanics/CLI/spells/mass-cure-wounds.md)\n\
    \n3/day each: [animal messenger](/3-Mechanics/CLI/spells/animal-messenger.md),\
    \ [dominate beast](/3-Mechanics/CLI/spells/dominate-beast.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [tree stride](/3-Mechanics/CLI/spells/tree-stride.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The archdruid makes three Staff or Wildfire attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage plus dice:6d6|text(21)\
    \ (6d6) poison damage."
  "name": "Staff"
- "desc": "Ranged Spell Attack: dice: d20+9 (+9) to hit, range 120 ft., one\
    \ target. Hit: dice:6d6 + 5|text(26) (6d6 + 5) fire damage, and the target\
    \ is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) until the start of\
    \ the druid's next turn."
  "name": "Wildfire"
"bonus_actions":
- "desc": "The archdruid magically transforms into a Beast or an Elemental with a\
    \ challenge rating of 6 or less and can remain in that form for up to 9 hours.\
    \ The archdruid can choose whether its equipment falls to the ground, melds with\
    \ its new form, or is worn by the new form. The archdruid reverts to its true\
    \ form if it dies or falls [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious).\
    \ The archdruid can revert to its true form using a bonus action.\n\nWhile in\
    \ a new form, the archdruid's stat block is replaced by the stat block of that\
    \ form, except the archdruid keeps its current hit points, its hit point maximum,\
    \ this bonus action, its languages and ability to speak, and its Spellcasting\
    \ action.\n\nThe new form's attacks count as magical for the purpose of overcoming\
    \ resistances and immunity to nonmagical attacks."
  "name": "Change Shape (2/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Archdruid.webp"
```
^statblock

```encounter-table
name: Archdruid
creatures:
 - 1: Archdruid
```

## Environment

forest, mountain, swamp, underwater