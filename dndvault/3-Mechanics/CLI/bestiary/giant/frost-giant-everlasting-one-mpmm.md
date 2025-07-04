---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/coastal
- monster/size/huge
- monster/type/giant
aliases: ["Frost Giant Everlasting One"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 131, Volo's Guide to Monsters p. 148
---
# [Frost Giant Everlasting One](3-Mechanics\CLI\bestiary\giant/frost-giant-everlasting-one-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 131, Volo's Guide to Monsters p. 148*  

```statblock
"name": "Frost Giant Everlasting One (MPMM)"
"size": "Huge"
"type": "giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "patchwork armor"
"hp": !!int "189"
"hit_dice": "14d12 + 98"
"stats":
- !!int "25"
- !!int "9"
- !!int "24"
- !!int "9"
- !!int "10"
- !!int "12"
"speed": "40 ft."
"saves":
  "Wisdom": !!int "4"
  "Strength": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Athletics": !!int "11"
  "Perception": !!int "4"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Giant"
"cr": "12"
"traits":
- "desc": "The giant has a 25% chance chance of having more than one head. If it has\
    \ more than one, it has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks and on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Extra Heads"
- "desc": "The giant regains 10 hit points at the start of its turn. If the giant\
    \ takes acid or fire damage, this trait doesn't function at the start of its next\
    \ turn. The giant dies only if it starts its turn with 0 hit points and doesn't\
    \ regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The giant makes two Greataxe or Rock attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d12 + 7|text(26) (3d12 + 7) slashing damage, or dice:3d12\
    \ + 11|text(30) (3d12 + 11) slashing damage while raging."
  "name": "Greataxe"
- "desc": "Ranged Weapon Attack: dice: d20+11 (+11) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 7|text(29) (4d10 + 7) bludgeoning damage."
  "name": "Rock"
"bonus_actions":
- "desc": "The giant enters a rage. The rage lasts for 1 minute or until the giant\
    \ is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated). While\
    \ raging, the giant gains the following benefits:\n\n- The giant has advantage\
    \ on Strength checks and Strength saving throws.  \n- When it makes a melee weapon\
    \ attack, the giant gains a +4 bonus to the damage roll.  \n- The giant has resistance\
    \ to bludgeoning, piercing, and slashing damage.  "
  "name": "Vaprak's Rage (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Frost%20Giant%20Everlasting%20One.webp"
```
^statblock

```encounter-table
name: Frost Giant Everlasting One
creatures:
 - 1: Frost Giant Everlasting One
```

## Environment

coastal