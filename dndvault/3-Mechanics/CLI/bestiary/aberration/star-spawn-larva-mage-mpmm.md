---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/environment/mountain
- monster/size/medium
- monster/type/aberration
aliases: ["Star Spawn Larva Mage"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 228, Mordenkainen's Tome of Foes p. 235
---
# [Star Spawn Larva Mage](3-Mechanics\CLI\bestiary\aberration/star-spawn-larva-mage-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 228, Mordenkainen's Tome of Foes p. 235*  

```statblock
"name": "Star Spawn Larva Mage (MPMM)"
"size": "Medium"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "168"
"hit_dice": "16d8 + 96"
"stats":
- !!int "17"
- !!int "12"
- !!int "23"
- !!int "18"
- !!int "12"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
"skillsaves":
  "Perception": !!int "6"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 16"
"languages": "Deep Speech"
"cr": "16"
"traits":
- "desc": "The mage casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 16):\n\nAt will:\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\n1/day: [dominate\
    \ monster](/3-Mechanics/CLI/spells/dominate-monster.md)"
  "name": "Spellcasting"
- "desc": "When the mage is reduced to 0 hit points, it breaks apart into a [swarm\
    \ of insects](/3-Mechanics/CLI/bestiary/beast/swarm-of-insects.md) in the same\
    \ space. Unless the swarm is destroyed, the mage reforms from it 24 hours later."
  "name": "Return to Worms"
"actions":
- "desc": "The mage makes three Slam or Eldritch Bolt attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) bludgeoning damage, and the target\
    \ must succeed on a DC 19 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the end of its next turn."
  "name": "Slam"
- "desc": "Ranged Spell Attack: dice: d20+8 (+8) to hit, range 60 ft., one target.\
    \ Hit: dice:3d10 + 3|text(19) (3d10 + 3) force damage."
  "name": "Eldritch Bolt"
- "desc": "Each creature other than a star spawn within 10 feet of the mage must succeed\
    \ on a DC 19 Dexterity saving throw or take dice:5d8|text(22) (5d8) necrotic\
    \ damage and be [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by masses of swarming worms. The affected creature takes dice:5d8|text(22)\
    \ (5d8) necrotic damage at the start of each of the mage's turns. The creature\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Plague of Worms (Recharge 6)"
"reactions":
- "desc": "When a creature within 20 feet of the mage fails a saving throw, the mage\
    \ gains 10 temporary hit points."
  "name": "Feed on Weakness"
"legendary_actions":
- "desc": "The mage makes one Slam attack."
  "name": "Slam"
- "desc": "The mage makes one Eldritch Bolt attack."
  "name": "Eldritch Bolt (Costs 2 Actions)"
- "desc": "Each creature [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by the mage's Plague of Worms takes dice:3d8|text(13) (3d8) necrotic damage,\
    \ and the mage gains 6 temporary hit points."
  "name": "Feed (Costs 3 Actions)"
"source":
- "MPMM"
- "MTF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Star%20Spawn%20Larva%20Mage.webp"
```
^statblock

```encounter-table
name: Star Spawn Larva Mage
creatures:
 - 1: Star Spawn Larva Mage
```

## Environment

mountain