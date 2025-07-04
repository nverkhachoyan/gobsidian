---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/beast
aliases: ["Swarm of Cranium Rats"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 133
---
# [Swarm of Cranium Rats](3-Mechanics\CLI\bestiary\beast/swarm-of-cranium-rats-vgm.md)
*Source: Volo's Guide to Monsters p. 133*  

Mind flayers create cranium rats by bombarding normal rats with psionic energy.

## Evil Collectives

Cranium rats are no smarter than ordinary rats and behave as such. However, if enough cranium rats come together to form a swarm, they merge their minds into a single intelligence with the accumulated memories of all the swarm's constituents. The rats become smarter as a result, and they retain their heightened intelligence for as long as the swarm persists. The swarm also awakens latent psionic abilities implanted within each cranium rat by its mind flayer creators, bestowing upon the swarm psionic powers similar to spells.

A rat separated from the swarm becomes an ordinary cranium rat with an Intelligence of 15. It loses 1 point of Intelligence each day that it remains separated from the swarm. Its Intelligence can't drop below 4 and becomes 15 again if it rejoins the swarm or another one.

## Telepathic Vermin

A single, low-intelligence cranium rat uses its natural telepathy to communicate hunger, fear, and other base emotions. A swarm of cranium rats communicating telepathically "speaks" as one creature, often referring to itself using the collective pronouns "we" and "us."

## Spies for an Elder Brain

Mind flayer colonies use cranium rats as spies. The rats invade surface communities and act as eyes and ears for the elder brain, transmitting their thoughts when they swarm and are within range of the elder brain's telepathy.

Cranium rats occasionally spread beyond the elder brain's range of influence. Whatever these rats do is of no concern to the elder brain, and the illithids can always make more if they so desire.

```statblock
"name": "Swarm of Cranium Rats (VGM)"
"size": "Medium"
"type": "beast"
"alignment": "Lawful Evil"
"ac": !!int "12"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "9"
- !!int "14"
- !!int "10"
- !!int "15"
- !!int "11"
- !!int "14"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 30 ft., passive Perception 10"
"languages": "telepathy 30 ft."
"cr": "5"
"traits":
- "desc": "The swarm's innate spellcasting ability is Intelligence (spell save DC\
    \ 13). As long as it has more than half of its hit points, it can innately cast\
    \ the following spells, requiring no components:\n\nAt will: [command](/3-Mechanics/CLI/spells/command.md),\
    \ [comprehend languages](/3-Mechanics/CLI/spells/comprehend-languages.md), [detect\
    \ thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md)\n\n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "As a bonus action, the swarm can shed dim light from its brains in a 5-foot\
    \ radius, increase the illumination to bright light in a 5 to 20-foot radius (and\
    \ dim light for an additional number of feet equal to the chosen radius), or extinguish\
    \ the light."
  "name": "Illumination"
- "desc": "The swarm can occupy another creature's space and vice versa, and the swarm\
    \ can move through any opening large enough for a Tiny rat. The swarm can't regain\
    \ hit points or gain temporary hit points."
  "name": "Swarm"
- "desc": "The swarm is immune to any effect that would sense its emotions or read\
    \ its thoughts. as well as to all divination spells."
  "name": "Telepathic Shroud"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 0 ft., one target\
    \ in the swarm's space. Hit: dice:4d6|text(14) (4d6) piercing damage, or\
    \ dice:2d6|text(7) (2d6) piercing damage if the swarm has half of its hit\
    \ points or fewer."
  "name": "Bites"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Swarm%20of%20Cranium%20Rats.webp"
```
^statblock

```encounter-table
name: Swarm of Cranium Rats
creatures:
 - 1: Swarm of Cranium Rats
```

## Environment

underdark, urban