---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/19
- monster/size/large
- monster/type/fiend/devil
aliases: ["Bael"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 54, Mordenkainen's Tome of Foes p. 170
---
# [Bael](3-Mechanics\CLI\bestiary\npc/bael-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 54, Mordenkainen's Tome of Foes p. 170*  

```statblock
"name": "Bael (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "24"
- !!int "17"
- !!int "20"
- !!int "21"
- !!int "24"
- !!int "24"
"speed": "30 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "9"
  "Intelligence": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Intimidation": !!int "13"
  "Perception": !!int "13"
  "Persuasion": !!int "13"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 23"
"languages": "all, telepathy 120 ft."
"cr": "19"
"traits":
- "desc": "Bael casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 21):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium), [charm\
    \ person](/3-Mechanics/CLI/spells/charm-person.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [major image](/3-Mechanics/CLI/spells/major-image.md)\n\
    \n1/day: [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md)\n\
    \n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Spellcasting"
- "desc": "Any creature, other than a devil, that starts its turn within 10 feet of\
    \ Bael must succeed on a DC 22 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of him until the start of its next turn. A creature succeeds on this saving\
    \ throw automatically if Bael wishes it or if he is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Dread"
- "desc": "If Bael fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Bael have advantage on saving throws against spells and other magical effects."
  "name": "Magic Resistance"
- "desc": "Bael regains 20 hit points at the start of his turn. If he takes cold or\
    \ radiant damage, this trait doesn't function at the start of his next turn. Bael\
    \ dies only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Bael makes two Hellish Morningstar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 20 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) force damage plus dice:2d8|text(9)\
    \ (2d8) necrotic damage."
  "name": "Hellish Morningstar"
- "desc": "Each of Bael's allies within 60 feet of him can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until the end\
    \ of his next turn."
  "name": "Infernal Command"
- "desc": "Bael teleports, along with any equipment he is wearing or carrying, up\
    \ to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Bael uses Spellcasting or Teleport."
  "name": "Fiendish Magic"
- "desc": "Bael uses Infernal Command."
  "name": "Infernal Command"
- "desc": "Bael makes one Hellish Morningstar attack."
  "name": "Attack (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Bael.webp"
```
^statblock

```encounter-table
name: Bael
creatures:
 - 1: Bael
```