---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githzerai Anarch"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 142, Mordenkainen's Tome of Foes p. 207
---
# [Githzerai Anarch](3-Mechanics\CLI\bestiary\humanoid/githzerai-anarch-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 142, Mordenkainen's Tome of Foes p. 207*  

```statblock
"name": "Githzerai Anarch (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Any alignment"
"ac": !!int "20"
"ac_class": "psychic defense"
"hp": !!int "144"
"hit_dice": "17d8 + 68"
"stats":
- !!int "16"
- !!int "21"
- !!int "18"
- !!int "18"
- !!int "20"
- !!int "14"
"speed": "30 ft., fly 40 ft. (hover)"
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "10"
  "Intelligence": !!int "9"
  "Strength": !!int "8"
"skillsaves":
  "Insight": !!int "10"
  "Perception": !!int "10"
  "Arcana": !!int "9"
"senses": "passive Perception 20"
"languages": "Gith"
"cr": "16"
"traits":
- "desc": "The githzerai casts one of the following spells, requiring no spell components\
    \ and using Wisdom as the spellcasting ability (spell save DC 18):\n\nAt will:\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md) (the hand is invisible)\n\n\
    1/day each: [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md), [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\
    \n3/day each: [see invisibility](/3-Mechanics/CLI/spells/see-invisibility.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "If the githzerai fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "While the githzerai is wearing no armor and wielding no shield, its AC\
    \ includes its Wisdom modifier."
  "name": "Psychic Defense"
"actions":
- "desc": "The githzerai makes three Unarmed Strike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d8 + 5|text(9) (1d8 + 5) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) psychic damage."
  "name": "Unarmed Strike"
"legendary_actions":
- "desc": "The githzerai makes one Unarmed Strike attack."
  "name": "Strike"
- "desc": "The githzerai teleports, along with any equipment it is wearing or carrying,\
    \ to an unoccupied space it can see within 30 feet of it."
  "name": "Teleport"
- "desc": "The githzerai casts the [reverse gravity](/3-Mechanics/CLI/spells/reverse-gravity.md)\
    \ spell, using Wisdom as the spellcasting ability. The spell has the normal effect,\
    \ except that the githzerai can orient the area in any direction and creatures\
    \ and objects fall toward the end of the area."
  "name": "Change Gravity (Costs 3 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Githzerai%20Anarch.webp"
```
^statblock

```encounter-table
name: Githzerai Anarch
creatures:
 - 1: Githzerai Anarch
```