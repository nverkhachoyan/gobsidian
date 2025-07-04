---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/26
- monster/size/large
- monster/type/fiend/devil
aliases: ["Zariel"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 280, Mordenkainen's Tome of Foes p. 180
---
# [Zariel](3-Mechanics\CLI\bestiary\npc/zariel-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 280, Mordenkainen's Tome of Foes p. 180*  

```statblock
"name": "Zariel (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "420"
"hit_dice": "29d10 + 261"
"stats":
- !!int "27"
- !!int "24"
- !!int "28"
- !!int "26"
- !!int "27"
- !!int "30"
"speed": "50 ft., fly 150 ft."
"saves":
  "Charisma": !!int "18"
  "Wisdom": !!int "16"
  "Intelligence": !!int "16"
"skillsaves":
  "Intimidation": !!int "18"
  "Perception": !!int "16"
"damage_resistances": "cold; fire; radiant; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 26"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Zariel casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 26):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium when\
    \ changing her appearance), [detect evil and good](/3-Mechanics/CLI/spells/detect-evil-and-good.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [major image](/3-Mechanics/CLI/spells/major-image.md), [wall of\
    \ fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n3/day each: [blade barrier](/3-Mechanics/CLI/spells/blade-barrier.md),\
    \ [dispel evil and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md), [finger\
    \ of death](/3-Mechanics/CLI/spells/finger-of-death.md)"
  "name": "Spellcasting"
- "desc": "Magical darkness doesn't impede Zariel's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "If Zariel fails a saving throw, she can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Zariel has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Zariel regains 20 hit points at the start of her turn. If she takes radiant\
    \ damage, this trait doesn't function at the start of her next turn. Zariel dies\
    \ only if she starts her turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Zariel makes three Flail or Longsword attacks. She can replace one attack\
    \ with a use of Horrid Touch, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) force damage plus dice:8d8|text(36)\
    \ (8d8) fire damage."
  "name": "Flail"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) radiant damage, or dice:2d10\
    \ + 8|text(19) (2d10 + 8) radiant damage when used with two hands, plus dice:8d8|text(36)\
    \ (8d8) fire damage."
  "name": "Longsword"
- "desc": "Zariel touches one creature within 10 feet of her. The target must succeed\
    \ on a DC 26 Constitution saving throw or take dice:8d10|text(44) (8d10) necrotic\
    \ damage and be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for\
    \ 1 minute. While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in\
    \ this way, the target is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened). The target can\
    \ repeat the saving throw at the end of each of its turns, ending the effect on\
    \ itself on a success."
  "name": "Horrid Touch (Recharge 5-6)"
- "desc": "Zariel teleports, along with any equipment she is wearing or carrying,\
    \ up to 120 feet to an unoccupied space she can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Zariel uses Teleport."
  "name": "Teleport"
- "desc": "Zariel turns her magical gaze toward one creature she can see within 120\
    \ feet of her and commands it to burn. The target must succeed on a DC 26 Wisdom\
    \ saving throw or take dice:4d10|text(22) (4d10) fire damage."
  "name": "Immolating Gaze (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Zariel.webp"
```
^statblock

```encounter-table
name: Zariel
creatures:
 - 1: Zariel
```