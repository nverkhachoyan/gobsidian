---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/0
- monster/environment/underdark
- monster/environment/urban
- monster/size/tiny
- monster/type/beast
aliases: ["Cranium Rat"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 133
---
# [Cranium Rat](3-Mechanics\CLI\bestiary\beast/cranium-rat-vgm.md)
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
"name": "Cranium Rat (VGM)"
"size": "Tiny"
"type": "beast"
"alignment": "Lawful Evil"
"ac": !!int "12"
"hp": !!int "2"
"hit_dice": "1d4"
"stats":
- !!int "2"
- !!int "14"
- !!int "10"
- !!int "4"
- !!int "11"
- !!int "8"
"speed": "30 ft."
"senses": "darkvision 30 ft., passive Perception 10"
"languages": "telepathy 30 ft."
"cr": "0"
"traits":
- "desc": "As a bonus action, the cranium rat can shed dim light from its brain in\
    \ a 5-foot radius or extinguish the light."
  "name": "Illumination"
- "desc": "The cranium rat is immune to any effect that would sense its emotions or\
    \ read its thoughts, as well as to all divination spells."
  "name": "Telepathic Shroud"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: 1 piercing damage."
  "name": "Bite"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Cranium%20Rat.webp"
```
^statblock

```encounter-table
name: Cranium Rat
creatures:
 - 1: Cranium Rat
```

## Environment

underdark, urban