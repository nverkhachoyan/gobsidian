---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/5
- monster/environment/swamp
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Allip"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 116, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Allip](3-Mechanics\CLI\bestiary\undead/allip-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 116, Dragon of Icespire Peak, Storm Lord's Wrath*  

## Allip

When a mind uncovers a secret that a powerful being has protected with a mighty curse, the result is often the emergence of an allip. Secrets protected in this manner range in scope from a demon lord's true name to the hidden truths of the cosmic order. The allip acquires the secret, but the curse annihilates its body and leaves behind a spectral creature composed of fragments from the victim's psyche and overwhelming psychic agony.

## Blasphemous Secrets

Every allip is wracked with a horrifying insight that torments what remains of its mind. In the presence of other creatures, an allip seeks to relieve this burden by sharing its secret. The creature can impart only a shard of the knowledge that doomed it, but that piece is enough to wrack the recipient with temporary madness.

The survivors of an allip's attack are sometimes left with a compulsion to learn more about what spawned this monstrosity. Strange phrases echo through their minds, and weird visions occupy their dreams. The sense that some colossal truth sits just outside their recall plagues them for days, months, and sometimes years after their fateful encounter.

## Undead Nature

An allip doesn't require air, food, drink, or sleep.

## Insidious Lore

An allip might attempt to share its lore to escape its curse and enter the afterlife. It can transfer knowledge from its mind by guiding another creature to write down what it knows. This process takes days or possibly weeks. An allip can accomplish this task by lurking in the study or workplace of a scholar. If the allip remains hidden, its victim is gradually overcome by manic energy. A scholar, driven by sudden insights to work night and day, produces reams of text with little memory of exactly what the documents contain. If the allip succeeds, it passes from the world—and its terrible secret hides somewhere in the scholar's text, waiting to be discovered by its next victim.

```statblock
"name": "Allip (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "13"
"hp": !!int "40"
"hit_dice": "9d8"
"stats":
- !!int "6"
- !!int "17"
- !!int "10"
- !!int "17"
- !!int "15"
- !!int "16"
"speed": "0 ft., fly 40 ft. (hover)"
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "5"
"damage_resistances": "acid; fire; lightning; thunder; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "the languages it knew in life"
"cr": "5"
"traits":
- "desc": "The allip can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
"actions":
- "desc": "Melee Spell Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d6 + 3|text(17) (4d6 + 3) psychic damage."
  "name": "Maddening Touch"
- "desc": "The allip chooses up to three creatures it can see within 60 feet of it.\
    \ Each target must succeed on a DC 14 Wisdom saving throw, or it takes dice:1d8\
    \ + 3|text(7) (1d8 + 3) psychic damage and must use its reaction to make a\
    \ melee weapon attack against one creature of the allip's choice that the allip\
    \ can see. Constructs and undead are immune to this effect."
  "name": "Whispers of Madness"
- "desc": "Each creature within 30 feet of the allip that can hear it must make a\
    \ DC 14 Wisdom saving throw. On a failed save, a target takes dice:2d8 + 3|text(12)\
    \ (2d8 + 3) psychic damage, and it is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of its next turn. On a successful save, it takes half as much\
    \ damage and isn't [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned). Constructs\
    \ and undead are immune to this effect."
  "name": "Howling Babble (Recharge 6)"
"source":
- "MTF"
- "DIP"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Allip.webp"
```
^statblock

```encounter-table
name: Allip
creatures:
 - 1: Allip
```

## Environment

swamp, urban