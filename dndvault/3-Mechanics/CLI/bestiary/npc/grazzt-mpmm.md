---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/24
- monster/size/large
- monster/type/fiend/demon
aliases: ["Graz'zt"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 148, Mordenkainen's Tome of Foes p. 149
---
# [Graz'zt](3-Mechanics\CLI\bestiary\npc/grazzt-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 148, Mordenkainen's Tome of Foes p. 149*  

```statblock
"name": "Graz'zt (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "346"
"hit_dice": "33d10 + 165"
"stats":
- !!int "22"
- !!int "15"
- !!int "21"
- !!int "23"
- !!int "21"
- !!int "26"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "12"
  "Constitution": !!int "12"
"skillsaves":
  "Deception": !!int "15"
  "Perception": !!int "12"
  "Persuasion": !!int "15"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "24"
"traits":
- "desc": "Graz'zt casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 23):\n\nAt will:\
    \ [charm person](/3-Mechanics/CLI/spells/charm-person.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md)\n\n1/day each: [dominate\
    \ monster](/3-Mechanics/CLI/spells/dominate-monster.md), [greater invisibility](/3-Mechanics/CLI/spells/greater-invisibility.md)\n\
    \n3/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [dominate person](/3-Mechanics/CLI/spells/dominate-person.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)"
  "name": "Spellcasting"
- "desc": "If Graz'zt fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Graz'zt has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Graz'zt makes two Wave of Sorrow attacks. He can replace one attack with\
    \ a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 10 ft., one\
    \ target. Hit: dice:4d6 + 6|text(20) (4d6 + 6) force damage plus dice:4d6|text(14)\
    \ (4d6) acid damage."
  "name": "Wave of Sorrow (Greatsword)"
- "desc": "Graz'zt teleports, along with any equipment he is wearing or carrying,\
    \ up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"bonus_actions":
- "desc": "Graz'zt transforms into a form that resembles a Medium Humanoid or back\
    \ into his true form. Aside from his size, his statistics are the same in each\
    \ form. Any equipment he is wearing or carrying isn't transformed."
  "name": "Change Shape"
"reactions":
- "desc": "Graz'zt tries to interrupt a spell he sees a creature casting within 60\
    \ feet of him. If the spell is 3rd level or lower, the spell fails and has no\
    \ effect. If the spell is 4th level or higher, Graz'zt makes a Charisma check\
    \ against a DC of 10 + the spell's level. On a success, the spell fails and has\
    \ no effect."
  "name": "Negate Spell (Recharge 5-6)"
"legendary_actions":
- "desc": "Graz'zt uses Spellcasting or Teleport."
  "name": "Abyssal Magic"
- "desc": "Graz'zt makes one Wave of Sorrow attack."
  "name": "Attack"
- "desc": "One creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by\
    \ Graz'zt that Graz'zt can see must use its reaction to move up to its speed as\
    \ Graz'zt directs."
  "name": "Dance, My Puppet!"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Graz%27zt.webp"
```
^statblock

```encounter-table
name: Graz'zt
creatures:
 - 1: Graz'zt
```