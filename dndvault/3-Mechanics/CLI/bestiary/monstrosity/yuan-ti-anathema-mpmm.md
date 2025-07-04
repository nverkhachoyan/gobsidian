---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/desert
- monster/environment/forest
- monster/environment/underdark
- monster/size/huge
- monster/type/monstrosity
aliases: ["Yuan-ti Anathema"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 272, Volo's Guide to Monsters p. 202
---
# [Yuan-ti Anathema](3-Mechanics\CLI\bestiary\monstrosity/yuan-ti-anathema-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 272, Volo's Guide to Monsters p. 202*  

```statblock
"name": "Yuan-ti Anathema (MPMM)"
"size": "Huge"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d12 + 72"
"stats":
- !!int "23"
- !!int "13"
- !!int "19"
- !!int "19"
- !!int "17"
- !!int "20"
"speed": "40 ft., climb 40 ft., swim 40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "11"
"damage_resistances": "acid, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 21"
"languages": "Abyssal, Common, Draconic"
"cr": "12"
"traits":
- "desc": "The anathema casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 17):\n\nAt will:\
    \ [animal friendship](/3-Mechanics/CLI/spells/animal-friendship.md) (snakes only)\n\
    \n3/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [entangle](/3-Mechanics/CLI/spells/entangle.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [polymorph](/3-Mechanics/CLI/spells/polymorph.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Spellcasting (Anathema Form Only)"
- "desc": "The anathema has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Any creature of the anathema's choice, other than a snake or a yuan-ti,\
    \ that starts its turn within 30 feet of the anathema must succeed on a DC 17\
    \ Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of snakes and yuan-ti. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success. If a target's saving throw is successful or the\
    \ effect ends for it, the target is immune to this anathama's aura for the next\
    \ 24 hours."
  "name": "Ophidiophobia Aura"
- "desc": "The anathema has advantage on saves against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Six Heads"
"actions":
- "desc": "The anathema makes two Claw attacks and one Flurry of Bites attack."
  "name": "Multiattack (Anathema Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d6 + 6|text(13) (2d6 + 6) slashing damage."
  "name": "Claw (Anathema Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ creature. Hit: dice:6d6 + 6|text(27) (6d6 + 6) piercing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage."
  "name": "Flurry of Bites (Anathema Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 15 ft., one\
    \ Large or smaller creature. Hit: dice:3d6 + 6|text(16) (3d6 + 6) bludgeoning\
    \ damage plus dice:2d6|text(7) (2d6) acid damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 16). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and it takes dice:3d6 + 6|text(16) (3d6 + 6) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) acid damage at the start of each of its turns. The anathema can constrict\
    \ only one creature at a time."
  "name": "Constrict (Snake Form Only)"
"bonus_actions":
- "desc": "The anathema transforms into a Huge constrictor snake or back into its\
    \ true form. Its statistics are the same in each form. Any equipment it is wearing\
    \ or carrying isn't transformed."
  "name": "Change Shape"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Yuan-ti%20Anathema.webp"
```
^statblock

```encounter-table
name: Yuan-ti Anathema
creatures:
 - 1: Yuan-ti Anathema
```

## Environment

desert, forest, underdark