---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Necromancer Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 264, Volo's Guide to Monsters p. 217
---
# [Necromancer Wizard](3-Mechanics\CLI\bestiary\humanoid/necromancer-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 264, Volo's Guide to Monsters p. 217*  

```statblock
"name": "Necromancer Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "110"
"hit_dice": "20d8 + 20"
"stats":
- !!int "9"
- !!int "14"
- !!int "12"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "7"
"skillsaves":
  "History": !!int "7"
  "Arcana": !!int "7"
"damage_resistances": "necrotic"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "9"
"traits":
- "desc": "The necromancer casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 15):\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [circle of death](/3-Mechanics/CLI/spells/circle-of-death.md)\n\
    \n2/day each: [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [web](/3-Mechanics/CLI/spells/web.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The necromancer makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:4d10 + 3|text(25) (4d10 + 3) necrotic\
    \ damage."
  "name": "Arcane Burst"
"bonus_actions":
- "desc": "The necromancer magically summons five [skeletons](/3-Mechanics/CLI/bestiary/undead/skeleton.md)\
    \ or [zombies](/3-Mechanics/CLI/bestiary/undead/zombie.md). The summoned creatures\
    \ appear in unoccupied spaces within 60 feet of the necromancer, whom they obey.\
    \ They take their turns immediately after the necromancer. Each lasts for 1 hour,\
    \ until it or the necromancer dies, or until the necromancer dismisses it as a\
    \ bonus action."
  "name": "Summon Undead (1/Day)"
"reactions":
- "desc": "When the necromancer kills a creature with necrotic damage, the necromancer\
    \ regains dice:2d8|text(9) (2d8) hit points. "
  "name": "Grim Harvest (1/Turn)"
"source":
- "MPMM"
- "VGM"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Necromancer%20Wizard.webp"
```
^statblock

```encounter-table
name: Necromancer Wizard
creatures:
 - 1: Necromancer Wizard
```

## Environment

desert, urban