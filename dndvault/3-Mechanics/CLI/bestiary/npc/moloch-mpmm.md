---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/21
- monster/size/large
- monster/type/fiend/devil
aliases: ["Moloch"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 183, Mordenkainen's Tome of Foes p. 177
---
# [Moloch](3-Mechanics\CLI\bestiary\npc/moloch-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 183, Mordenkainen's Tome of Foes p. 177*  

```statblock
"name": "Moloch (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "253"
"hit_dice": "22d10 + 132"
"stats":
- !!int "26"
- !!int "19"
- !!int "22"
- !!int "21"
- !!int "18"
- !!int "23"
"speed": "30 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "11"
  "Wisdom": !!int "11"
  "Constitution": !!int "13"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Perception": !!int "11"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "21"
"traits":
- "desc": "Moloch casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 21):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium when\
    \ changing his appearance), [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Spellcasting"
- "desc": "If Moloch fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Moloch has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Moloch regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Moloch dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Moloch makes one Bite attack, one Claw attack, and one Many-Tailed Whip\
    \ attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d8 + 8|text(26) (4d8 + 8) fire damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) force damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 30 ft., one\
    \ target. Hit: dice:2d4 + 8|text(13) (2d4 + 8) lightning damage plus dice:2d10|text(11)\
    \ (2d10) thunder damage. If the target is a creature, it must succeed on a DC\
    \ 24 Strength saving throw or be pulled up to 30 feet in a straight line toward\
    \ Moloch."
  "name": "Many-Tailed Whip"
- "desc": "Moloch exhales in a 30-foot cube. Each creature in that area must succeed\
    \ on a DC 21 Wisdom saving throw or take dice:5d10|text(27) (5d10) psychic\
    \ damage, drop whatever it is holding, and become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of Moloch for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a creature must take the [Dash](/3-Mechanics/CLI/rules/actions.md#Dash)\
    \ action and move away from Moloch by the safest available route on each of its\
    \ turns, unless there is nowhere to move, in which case it needn't take the [Dash](/3-Mechanics/CLI/rules/actions.md#Dash)\
    \ action. If the creature ends its turn in a location where it doesn't have line\
    \ of sight to Moloch, the creature can repeat the saving throw, ending the effect\
    \ on itself on a success."
  "name": "Breath of Despair (Recharge 5-6)"
- "desc": "Moloch teleports, along with any equipment he is wearing or carrying, up\
    \ to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Moloch makes one Bite, Claw, or Many-Tailed Whip attack."
  "name": "Attack"
- "desc": "Moloch uses Teleport."
  "name": "Teleport"
- "desc": "Moloch uses Spellcasting."
  "name": "Cast a Spell (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Moloch.webp"
```
^statblock

```encounter-table
name: Moloch
creatures:
 - 1: Moloch
```