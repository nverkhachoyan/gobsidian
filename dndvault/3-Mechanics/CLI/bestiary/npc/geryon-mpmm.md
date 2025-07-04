---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/22
- monster/size/huge
- monster/type/fiend/devil
aliases: ["Geryon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 136, Mordenkainen's Tome of Foes p. 173
---
# [Geryon](3-Mechanics\CLI\bestiary\npc/geryon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 136, Mordenkainen's Tome of Foes p. 173*  

```statblock
"name": "Geryon (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "300"
"hit_dice": "24d12 + 144"
"stats":
- !!int "29"
- !!int "17"
- !!int "22"
- !!int "19"
- !!int "16"
- !!int "23"
"speed": "30 ft., fly 50 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "10"
  "Wisdom": !!int "10"
  "Constitution": !!int "13"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Perception": !!int "10"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"damage_immunities": "cold, fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 20"
"languages": "all, telepathy 120 ft."
"cr": "22"
"traits":
- "desc": "Geryon casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 21):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium when\
    \ changing his appearance), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [ice storm](/3-Mechanics/CLI/spells/ice-storm.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [locate object](/3-Mechanics/CLI/spells/locate-object.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)\n\n1/day: [banishment](/3-Mechanics/CLI/spells/banishment.md)"
  "name": "Spellcasting"
- "desc": "If Geryon fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Geryon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Geryon regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Geryon dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Geryon makes one Claw attack and one Stinger attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:4d6 + 9|text(23) (4d6 + 9) cold damage. If the target\
    \ is Large or smaller, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (DC 24), and it is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. Geryon can grapple one creature at a time. If the target\
    \ is already [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) by Geryon,\
    \ the target takes an extra dice:6d8|text(27) (6d8) cold damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ creature. Hit: dice:2d4 + 9|text(14) (2d4 + 9) force damage, and the target\
    \ must succeed on a DC 21 Constitution saving throw or take dice:2d12|text(13)\
    \ (2d12) poison damage and become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until it finishes a short or long rest. The target's hit point maximum is reduced\
    \ by an amount equal to half the poison damage taken. This reduction lasts until\
    \ the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) condition is removed.\
    \ The target dies if its hit point maximum is reduced to 0."
  "name": "Stinger"
- "desc": "Geryon teleports, along with any equipment he is wearing and carrying,\
    \ up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Geryon targets one creature he can see within 60 feet of him. The target\
    \ must succeed on a DC 23 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of Geryon until the end of its next turn."
  "name": "Infernal Glare"
- "desc": "Geryon uses Teleport."
  "name": "Teleport"
- "desc": "Geryon makes one Stinger attack."
  "name": "Swift Sting (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Geryon.webp"
```
^statblock

```encounter-table
name: Geryon
creatures:
 - 1: Geryon
```