---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/26
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Demogorgon"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 90, Mordenkainen's Tome of Foes p. 144
---
# [Demogorgon](3-Mechanics\CLI\bestiary\npc/demogorgon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 90, Mordenkainen's Tome of Foes p. 144*  

```statblock
"name": "Demogorgon (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "464"
"hit_dice": "32d12 + 256"
"stats":
- !!int "29"
- !!int "14"
- !!int "26"
- !!int "20"
- !!int "17"
- !!int "25"
"speed": "50 ft., swim 50 ft."
"saves":
  "Charisma": !!int "15"
  "Dexterity": !!int "10"
  "Wisdom": !!int "11"
  "Constitution": !!int "16"
"skillsaves":
  "Insight": !!int "11"
  "Perception": !!int "19"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 29"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Demogorgon casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 23):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [major image](/3-Mechanics/CLI/spells/major-image.md)\n\
    \n1/day each: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md), [project\
    \ image](/3-Mechanics/CLI/spells/project-image.md)\n\n3/day each: [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Spellcasting"
- "desc": "If Demogorgon fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Demogorgon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Demogorgon has advantage on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Two Heads"
"actions":
- "desc": "Demogorgon makes two Tentacle attacks. He can replace one attack with a\
    \ use of Gaze."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d12 + 9|text(28) (3d12 + 9) force damage. If the target\
    \ is a creature, it must succeed on a DC 23 Constitution saving throw, or its\
    \ hit point maximum is reduced by an amount equal to the damage taken. This reduction\
    \ lasts until the target finishes a long rest. The target dies if its hit point\
    \ maximum is reduced to 0."
  "name": "Tentacle"
- "desc": "Demogorgon turns his magical gaze toward one creature he can see within\
    \ 120 feet of him. The target must succeed on a DC 23 Wisdom saving throw or suffer\
    \ one of the following effects (choose one or roll a dice: d6|avg|noform (d6)):\n\
    \n- 1–2 Beguiling Gaze. The target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the start of Demogorgon's next turn or until Demogorgon is no longer within\
    \ line of sight.  \n- 3–4 Confusing Gaze. The target suffers the effect of\
    \ the [confusion](/3-Mechanics/CLI/spells/confusion.md) spell without making a\
    \ saving throw. The effect lasts until the start of Demogorgon's next turn. Demogorgon\
    \ doesn't need to concentrate on the spell.  \n- 5–6 Hypnotic Gaze. The target\
    \ is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by Demogorgon until\
    \ the start of Demogorgon's next turn. Demogorgon chooses how the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target uses its action, reaction, and movement.  "
  "name": "Gaze"
"legendary_actions":
- "desc": "Demogorgon uses Gaze and must use either Beguiling Gaze or Confusing Gaze."
  "name": "Gaze"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) bludgeoning damage plus\
    \ dice:2d10|text(11) (2d10) necrotic damage."
  "name": "Tail"
- "desc": "Demogorgon uses Spellcasting."
  "name": "Cast a Spell (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Demogorgon.webp"
```
^statblock

```encounter-table
name: Demogorgon
creatures:
 - 1: Demogorgon
```