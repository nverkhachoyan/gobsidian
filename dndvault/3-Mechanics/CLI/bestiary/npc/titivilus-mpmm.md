---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Titivilus"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 242, Mordenkainen's Tome of Foes p. 179
---
# [Titivilus](3-Mechanics\CLI\bestiary\npc/titivilus-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 242, Mordenkainen's Tome of Foes p. 179*  

```statblock
"name": "Titivilus (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "150"
"hit_dice": "20d8 + 60"
"stats":
- !!int "19"
- !!int "22"
- !!int "17"
- !!int "24"
- !!int "22"
- !!int "26"
"speed": "40 ft., fly 60 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "11"
  "Wisdom": !!int "11"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Insight": !!int "11"
  "Persuasion": !!int "13"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "all, telepathy 120 ft."
"cr": "16"
"traits":
- "desc": "Titivilus casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 21):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md), [major image](/3-Mechanics/CLI/spells/major-image.md),\
    \ [nondetection](/3-Mechanics/CLI/spells/nondetection.md), [sending](/3-Mechanics/CLI/spells/sending.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n3/day each: [mislead](/3-Mechanics/CLI/spells/mislead.md),\
    \ [modify memory](/3-Mechanics/CLI/spells/modify-memory.md)"
  "name": "Spellcasting"
- "desc": "If Titivilus fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Titivilus has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Titivilus regains 10 hit points at the start of his turn. If he takes cold\
    \ or radiant damage, this trait doesn't function at the start of his next turn.\
    \ Titivilus dies only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Whenever Titivilus speaks, he can choose a point within 60 feet of him;\
    \ his voice emanates from that point."
  "name": "Ventriloquism"
"actions":
- "desc": "Titivilus makes one Silver Sword attack, and he uses Frightful Word."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) force damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) force damage if used with two hands, plus dice:3d10|text(16)\
    \ (3d10) necrotic damage. If the target is a creature, its hit point maximum\
    \ is reduced by an amount equal to half the necrotic damage taken."
  "name": "Silver Sword"
- "desc": "Titivilus targets one creature he can see within 10 feet of him. The target\
    \ must succeed on a DC 21 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of him for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, the target must take the [Dash](/3-Mechanics/CLI/rules/actions.md#Dash)\
    \ action and move away from Titivilus by the safest available route on each of\
    \ its turns, unless there is nowhere to move, in which case it needn't take the\
    \ [Dash](/3-Mechanics/CLI/rules/actions.md#Dash) action. The target can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success."
  "name": "Frightful Word"
- "desc": "Titivilus teleports, along with any equipment he is wearing or carrying,\
    \ up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
- "desc": "Titivilus targets one creature he can see within 60 feet of him. The target\
    \ must succeed on a DC 21 Charisma saving throw or become [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by Titivilus for 1 minute. The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target can repeat the saving throw if Titivilus deals any damage to it. A creature\
    \ that succeeds on the saving throw is immune to Titivilus's Twisting Words for\
    \ 24 hours."
  "name": "Twisting Words"
"legendary_actions":
- "desc": "Titivilus uses Twisting Words. Alternatively, he targets one creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by him that is within 60 feet of him; that [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target must succeed on a DC 21 Charisma saving throw, or Titivilus decides how\
    \ the target acts during its next turn."
  "name": "Corrupting Guidance"
- "desc": "Titivilus uses Teleport."
  "name": "Teleport"
- "desc": "Titivilus makes one Silver Sword attack, or he uses Frightful Word."
  "name": "Assault (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Titivilus.webp"
```
^statblock

```encounter-table
name: Titivilus
creatures:
 - 1: Titivilus
```