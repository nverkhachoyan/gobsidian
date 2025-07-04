---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/15
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/size/medium
- monster/type/undead
aliases: ["Skull Lord"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 230
---
# [Skull Lord](3-Mechanics\CLI\bestiary\undead/skull-lord-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 230*  

## Skull Lord

The skull lords have claimed vast regions of the Shadowfell as their dominion. From these blighted lands, they wage war against their rivals, commanding hordes of undead in a bid to establish dominance. Yet skull lords always prove to be their own worst enemies; as a combined being born from three hateful individuals, they constantly plot against themselves.

## Creatures of Betrayal

Infighting and treachery brought the skull lords into existence. The first of them appeared in the aftermath of Vecna's bid to conquer the world of Greyhawk, after the vampire Kas betrayed Vecna and took his eye and hand. In the confusion resulting from this turn of events, Vecna's warlords turned against each other, and the dark one's plans were dashed. In a rage, Vecna gathered up his generals and captains and bound them in groups of three, fusing them into undead abominations cursed to fight among themselves for all time. Since the first skull lords were exiled into shadow, others have joined them, typically after being created from other leaders who betrayed their masters.

## Undead Nature

A skull lord doesn't require air, food, drink, or sleep.

```statblock
"name": "Skull Lord (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "105"
"hit_dice": "14d8 + 42"
"stats":
- !!int "14"
- !!int "16"
- !!int "17"
- !!int "16"
- !!int "15"
- !!int "21"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "7"
  "Stealth": !!int "8"
  "Perception": !!int "12"
  "History": !!int "8"
"damage_resistances": "cold; necrotic; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "poison"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
  \ [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 22"
"languages": "all the languages it knew in life"
"cr": "15"
"traits":
- "desc": "The skull lord is a 13th-level spellcaster. Its spellcasting ability is\
    \ Charisma (spell save DC 18, dice: d20+10 (+10) to hit with spell attacks).\
    \ The skull lord knows the following sorcerer spells:\n\nCantrips (at will):\
    \ [chill touch](/3-Mechanics/CLI/spells/chill-touch.md), [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md), [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\
    \n1st level (4 slots): [magic missile](/3-Mechanics/CLI/spells/magic-missile.md),\
    \ [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md), [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\
    \n2nd level (3 slots): [mirror image](/3-Mechanics/CLI/spells/mirror-image.md),\
    \ [scorching ray](/3-Mechanics/CLI/spells/scorching-ray.md)\n\n3rd level (3\
    \ slots): [fear](/3-Mechanics/CLI/spells/fear.md), [haste](/3-Mechanics/CLI/spells/haste.md)\n\
    \n4th level (3 slots): [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [ice storm](/3-Mechanics/CLI/spells/ice-storm.md)\n\n5th level (2 slots):\
    \ [cloudkill](/3-Mechanics/CLI/spells/cloudkill.md), [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md)\n\
    \n6th level (1 slots): [eyebite](/3-Mechanics/CLI/spells/eyebite.md)\n\n7th\
    \ level (1 slots): [finger of death](/3-Mechanics/CLI/spells/finger-of-death.md)"
  "name": "Spellcasting"
- "desc": "If the skull lord fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "While within 30 feet of the skull lord, any undead ally of the skull lord\
    \ makes saving throws with advantage, and that ally regains dice: 1d6|avg|noform\
    \ (1d6) hit points whenever it starts its turn there."
  "name": "Master of the Grave"
- "desc": "If the skull lord is subjected to an effect that allows it to make a Dexterity\
    \ saving throw to take only half the damage, the skull lord instead takes no damage\
    \ if it succeeds on the saving throw, and only half damage if it fails."
  "name": "Evasion"
"actions":
- "desc": "The skull lord makes three bone staff attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) bludgeoning damage plus dice:4d6|text(14)\
    \ (4d6) necrotic damage."
  "name": "Bone Staff"
"legendary_actions":
- "desc": "The skull lord makes a bone staff attack."
  "name": "Bone Staff (Costs 2 Actions)"
- "desc": "The skull lord casts a cantrip."
  "name": "Cantrip"
- "desc": "The skull lord moves up to its speed without provoking opportunity attacks."
  "name": "Move"
- "desc": "Up to five [skeletons](/3-Mechanics/CLI/bestiary/undead/skeleton.md) or\
    \ [zombies](/3-Mechanics/CLI/bestiary/undead/zombie.md) appear in unoccupied spaces\
    \ within 30 feet of the skull lord and remain until destroyed. Undead summoned\
    \ in this way roll initiative and act in the next available turn. The skull lord\
    \ can have up to five undead summoned by this ability at a time."
  "name": "Summon Undead (Costs 3 Actions)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Skull%20Lord.webp"
```
^statblock

```encounter-table
name: Skull Lord
creatures:
 - 1: Skull Lord
```

## Environment

desert, swamp, underdark