---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/humanoid/kobold
aliases: ["Kobold Dragonshield"]
NoteIcon: monster
BestiaryType: humanoid (kobold)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 165, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Kobold Dragonshield](3-Mechanics\CLI\bestiary\humanoid/kobold-dragonshield-vgm.md)
*Source: Volo's Guide to Monsters p. 165, Dragon of Icespire Peak, Storm Lord's Wrath*  

A kobold dragonshield is a champion of its race. Almost all dragonshields begin life as normal kobolds, then are chosen by a dragon and invested with great powers for the purpose of protecting the dragon's eggs, but once every few years a kobold hatches with an innate version of the dragonshield's abilities. Accomplished at hand-to-hand combat, it bears many scars from desperate fights and carries a shield made out of cast-off dragon scales.

## Uncommon Courage

A dragonshield knows that it has a place of honor in the tribe, but-being kobolds at heart-most of them feel unworthy of their status and thus desperate to prove themselves deserving of it. A dragonshield's natural kobold cowardice is still present in its makeup, and thus it might still run away from a threat. But it also has the ability to rally in the face of certain death, inspiring other kobolds to follow it in a charge against the invaders of their warren.

```statblock
"name": "Kobold Dragonshield (VGM)"
"size": "Small"
"type": "humanoid"
"subtype": "kobold"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "44"
"hit_dice": "8d6 + 16"
"stats":
- !!int "12"
- !!int "15"
- !!int "14"
- !!int "8"
- !!int "9"
- !!int "10"
"speed": "20 ft."
"skillsaves":
  "Perception": !!int "1"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Draconic"
"cr": "1"
"traits":
- "desc": "The kobold has resistance to a type of damage based on the color of dragon\
    \ that invested it with power (choose or roll a dice: d10|avg|noform (d10)):\
    \ 1–2, acid (black); 3–4, cold (white); 5–6, fire (red); 7–8, lightning (blue);\
    \ 9–10, poison (green)."
  "name": "Dragon's Resistance"
- "desc": "If the kobold is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) by an effect\
    \ that allows a saving throw, it can repeat the save at the start of its turn\
    \ to end the effect on itself and all kobolds within 30 feet of it. Any kobold\
    \ that benefits from this trait (including the dragonshield) has advantage on\
    \ its next attack roll."
  "name": "Heart of the Dragon"
- "desc": "The kobold has advantage on an attack roll against a creature if at least\
    \ one of the kobold's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "While in sunlight, the kobold has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The kobold makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 1|text(4) (1d6 + 1) piercing\
    \ damage, or dice:1d8 + 1|text(5) (1d8 + 1) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Spear"
"source":
- "VGM"
- "DIP"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Kobold%20Dragonshield.webp"
```
^statblock

```encounter-table
name: Kobold Dragonshield
creatures:
 - 1: Kobold Dragonshield
```

## Environment

underdark, mountain, forest, hill