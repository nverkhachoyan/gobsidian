---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Dybbuk"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 113, Mordenkainen's Tome of Foes p. 132
---
# [Dybbuk](3-Mechanics\CLI\bestiary\fiend/dybbuk-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 113, Mordenkainen's Tome of Foes p. 132*  

```statblock
"name": "Dybbuk (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"hp": !!int "37"
"hit_dice": "5d8 + 15"
"stats":
- !!int "6"
- !!int "19"
- !!int "16"
- !!int "16"
- !!int "15"
- !!int "14"
"speed": "40 ft. (hover)"
"skillsaves":
  "Intimidation": !!int "4"
  "Deception": !!int "6"
  "Perception": !!int "4"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "Abyssal, Common, telepathy 120 ft."
"cr": "4"
"traits":
- "desc": "The dybbuk casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 12):\n\nAt will:\
    \ [dimension door](/3-Mechanics/CLI/spells/dimension-door.md)\n\n3/day: [phantasmal\
    \ force](/3-Mechanics/CLI/spells/phantasmal-force.md)"
  "name": "Spellcasting"
- "desc": "The dybbuk can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "The dybbuk has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) necrotic damage. If the target is\
    \ a creature, its hit point maximum is also reduced by dice:1d6|text(3) (1d6).\
    \ This reduction lasts until the target finishes a short or long rest. The target\
    \ dies if its hit point maximum is reduced to 0."
  "name": "Tentacle"
- "desc": "The dybbuk disappears into an intact corpse within 5 feet of it that belonged\
    \ to a Large or smaller Beast or Humanoid. The dybbuk gains 20 temporary hit points.\
    \ While possessing the corpse, the dybbuk adopts the corpse's size and can't use\
    \ Incorporeal Movement. Its game statistics otherwise remain the same. The possession\
    \ lasts until the temporary hit points are lost or the dybbuk ends it as a bonus\
    \ action. When the possession ends, the dybbuk appears in an unoccupied space\
    \ within 5 feet of the corpse."
  "name": "Possess Corpse (Recharge 6)"
"bonus_actions":
- "desc": "While Possess Corpse is active, the dybbuk makes the corpse do something\
    \ unnatural, such as vomit blood, twist its head all the way around, or cause\
    \ a quadruped to move as a biped. Any Beast or Humanoid that sees this behavior\
    \ must succeed on a DC 12 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the dybbuk for 1 minute. The [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. A creature that succeeds on a saving throw\
    \ against this ability is immune to Control Corpse for 24 hours."
  "name": "Control Corpse"
"source":
- "MPMM"
- "MTF"
- "AATM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Dybbuk.webp"
```
^statblock

```encounter-table
name: Dybbuk
creatures:
 - 1: Dybbuk
```

## Environment

desert, urban