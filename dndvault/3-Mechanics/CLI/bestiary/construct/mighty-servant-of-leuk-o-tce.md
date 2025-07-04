---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/huge
- monster/type/construct
aliases: ["Mighty Servant of Leuk-o"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 131
---
# [Mighty Servant of Leuk-o](3-Mechanics\CLI\bestiary\construct/mighty-servant-of-leuk-o-tce.md)
*Source: Tasha's Cauldron of Everything p. 131*  

```statblock
"name": "Mighty Servant of Leuk-o (TCE)"
"size": "Huge"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "310"
"hit_dice": "27d12 + 135"
"stats":
- !!int "30"
- !!int "14"
- !!int "20"
- !!int "1"
- !!int "14"
- !!int "10"
"speed": "60 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "9"
"skillsaves":
  "Perception": !!int "9"
"damage_resistances": "piercing, slashing"
"damage_immunities": "acid, bludgeoning, cold, fire, lightning, necrotic, poison,\
  \ psychic, radiant"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "blindsight 120 ft., passive Perception 19"
"languages": "understands the languages of creatures attuned to it but can't speak"
"traits":
- "desc": "The servant is immune to any spell or effect that would alter its form\
    \ or send it to another plane of existence."
  "name": "Immutable Existence"
- "desc": "The servant has advantage on saving throws against spells and other magical\
    \ effects, and spell attacks made against it have disadvantage."
  "name": "Magic Resistance"
- "desc": "The servant regains 10 hit points at the start of its turn. If it is reduced\
    \ to 0 hit points, this trait doesn't function until an attuned creature spends\
    \ 24 hours repairing the artifact or until the artifact is subjected to lightning\
    \ damage."
  "name": "Regeneration"
- "desc": "The servant's long jump is up to 50 feet and its high jump is up to 25\
    \ feet, with or without a running start."
  "name": "Standing Leap"
- "desc": "The servant doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:4d12 + 10|text(36) (4d12 + 10) force damage. Or Ranged\
    \ Weapon Attack: dice: d20+17 (+17) to hit, range 120 ft., one target. Hit:\
    \ dice:4d12 + 10|text(36) (4d12 + 10) force damage. If the target is an object,\
    \ it takes triple damage."
  "name": "Destructive Fist"
- "desc": "If the servant jumps at least 25 feet as part of its movement, it can then\
    \ use this action to land on its feet in a space that contains one or more other\
    \ creatures. Each of those creatures is pushed to an unoccupied space within 5\
    \ feet of the servant and must make a DC 25 Dexterity saving throw. On a failed\
    \ save, a creature takes dice:4d12|text(26) (4d12) bludgeoning damage and\
    \ is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone). On a successful\
    \ save, a creature takes half as much damage and isn't knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Crushing Leap"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Mighty%20Servant%20of%20Leuk-o.webp"
```
^statblock

```encounter-table
name: Mighty Servant of Leuk-o
creatures:
 - 1: Mighty Servant of Leuk-o
```