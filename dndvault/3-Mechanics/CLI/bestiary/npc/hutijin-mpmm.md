---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/21
- monster/size/large
- monster/type/fiend/devil
aliases: ["Hutijin"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 157, Mordenkainen's Tome of Foes p. 175
---
# [Hutijin](3-Mechanics\CLI\bestiary\npc/hutijin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 157, Mordenkainen's Tome of Foes p. 175*  

```statblock
"name": "Hutijin (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "200"
"hit_dice": "16d10 + 112"
"stats":
- !!int "27"
- !!int "15"
- !!int "25"
- !!int "23"
- !!int "19"
- !!int "25"
"speed": "30 ft., fly 60 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "11"
  "Constitution": !!int "14"
"skillsaves":
  "Intimidation": !!int "14"
  "Perception": !!int "11"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "21"
"traits":
- "desc": "Hutijin casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 22):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium when\
    \ changing his appearance), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [hold monster](/3-Mechanics/CLI/spells/hold-monster.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n3/day: [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md)"
  "name": "Spellcasting"
- "desc": "Each creature within 30 feet of Hutijin that isn't a devil makes saving\
    \ throws with disadvantage."
  "name": "Infernal Despair"
- "desc": "If Hutijin fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Hutijin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Hutijin regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Hutijin dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Hutijin makes one Bite attack, one Claw attack, one Mace attack, and one\
    \ Tail attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) fire damage. The target must\
    \ succeed on a DC 22 Constitution saving throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way,\
    \ the target can't regain hit points, and it takes dice:3d6|text(10) (3d6)\
    \ poison damage at the start of each of its turns. The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) cold damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) force damage."
  "name": "Mace"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d10 + 8|text(19) (2d10 + 8) thunder damage."
  "name": "Tail"
- "desc": "Hutijin teleports, along with any equipment he is wearing and carrying,\
    \ up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"reactions":
- "desc": "In response to taking damage, Hutijin utters a dreadful word of power.\
    \ Each creature within 30 feet of him that isn't a devil must succeed on a DC\
    \ 22 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of him for 1 minute. A creature can repeat the saving throw at the end of each\
    \ of its turns, ending the effect on itself on a success. A creature that saves\
    \ against this effect is immune to his Fearful Voice for 24 hours."
  "name": "Fearful Voice (Recharge 5-6)"
"legendary_actions":
- "desc": "Hutijin makes one Claw, Mace, or Tail attack."
  "name": "Attack"
- "desc": "Hutijin uses Teleport."
  "name": "Teleport"
- "desc": "Hutijin releases lightning in a 30-foot radius, blocked only by total cover.\
    \ All other creatures in that area must each make a DC 22 Dexterity saving throw,\
    \ taking dice:4d8|text(18) (4d8) lightning damage on a failed save, or half\
    \ as much damage on a successful one."
  "name": "Lightning Storm (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Hutijin.webp"
```
^statblock

```encounter-table
name: Hutijin
creatures:
 - 1: Hutijin
```