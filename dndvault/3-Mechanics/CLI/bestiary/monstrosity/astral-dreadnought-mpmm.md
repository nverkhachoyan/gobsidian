---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/21
- monster/size/gargantuan
- monster/type/monstrosity/titan
aliases: ["Astral Dreadnought"]
NoteIcon: monster
BestiaryType: monstrosity (titan)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 51, Mordenkainen's Tome of Foes p. 117
---
# [Astral Dreadnought](3-Mechanics\CLI\bestiary\monstrosity/astral-dreadnought-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 51, Mordenkainen's Tome of Foes p. 117*  

```statblock
"name": "Astral Dreadnought (MPMM)"
"size": "Gargantuan"
"type": "monstrosity"
"subtype": "titan"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "297"
"hit_dice": "17d20 + 119"
"stats":
- !!int "28"
- !!int "7"
- !!int "25"
- !!int "5"
- !!int "14"
- !!int "18"
"speed": "15 ft., fly 80 ft. (hover)"
"saves":
  "Dexterity": !!int "5"
  "Wisdom": !!int "9"
"skillsaves":
  "Perception": !!int "9"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 120 ft., passive Perception 19"
"languages": ""
"cr": "21"
"traits":
- "desc": "The dreadnought's eye creates an area of antimagic, as in the [antimagic\
    \ field](/3-Mechanics/CLI/spells/antimagic-field.md) spell, in a 150-foot cone.\
    \ At the start of each of its turns, it decides which way the cone faces. The\
    \ cone doesn't function while the eye is closed or while the dreadnought is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)."
  "name": "Antimagic Cone"
- "desc": "The dreadnought can't leave the Astral Plane, nor can it be banished or\
    \ otherwise transported out of that plane."
  "name": "Astral Entity"
- "desc": "Anything the dreadnought swallows is transported to a demiplane that can\
    \ be entered by no other means except a [wish](/3-Mechanics/CLI/spells/wish.md)\
    \ spell or the dreadnought's Bite and Donjon Visit. A creature can leave the demiplane\
    \ only by using magic that enables planar travel, such as the [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ spell. The demiplane resembles a stone cave roughly 1,000 feet in diameter with\
    \ a ceiling 100 feet high. Like a stomach, it contains the remains of past meals.\
    \ The dreadnought can't be harmed from within the demiplane. If the dreadnought\
    \ dies, the demiplane disappears, and everything inside it appears around the\
    \ dreadnought's corpse. The demiplane is otherwise indestructible."
  "name": "Demiplanar Donjon"
- "desc": "If the dreadnought fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "If the dreadnought scores a critical hit against a creature traveling by\
    \ means of the [astral projection](/3-Mechanics/CLI/spells/astral-projection.md)\
    \ spell, the dreadnought can cut the target's silver cord instead of dealing damage."
  "name": "Sever Silver Cord"
- "desc": "The dreadnought doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The dreadnought makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:5d10 + 9|text(36) (5d10 + 9) force damage. If the target\
    \ is a Huge or smaller creature and this damage reduces it to 0 hit points or\
    \ it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated), the\
    \ dreadnought swallows it. The swallowed target, along with everything it is wearing\
    \ and carrying, appears in an unoccupied space on the floor of the Demiplanar\
    \ Donjon."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ target. Hit: dice:3d6 + 9|text(19) (3d6 + 9) force damage."
  "name": "Claw"
"legendary_actions":
- "desc": "The dreadnought makes one Claw attack."
  "name": "Claw"
- "desc": "One Huge or smaller creature that the dreadnought can see within 60 feet\
    \ of it must succeed on a DC 19 Charisma saving throw or be teleported to an unoccupied\
    \ space on the floor of the Demiplanar Donjon. At the end of the target's next\
    \ turn, it reappears in the space it left or in the nearest unoccupied space if\
    \ that space is occupied."
  "name": "Donjon Visit (Costs 2 Actions)"
- "desc": "Each creature within 60 feet of the dreadnought must make a DC 19 Wisdom\
    \ saving throw, taking dice:4d10 + 4|text(26) (4d10 + 4) psychic damage on\
    \ a failed save, or half as much damage on a successful one."
  "name": "Psychic Projection (Costs 3 Actions)"
"source":
- "MPMM"
- "MTF"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Astral%20Dreadnought.webp"
```
^statblock

```encounter-table
name: Astral Dreadnought
creatures:
 - 1: Astral Dreadnought
```