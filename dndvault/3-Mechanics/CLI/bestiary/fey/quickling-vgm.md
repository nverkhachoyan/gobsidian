---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/forest
- monster/size/tiny
- monster/type/fey
aliases: ["Quickling"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 187
---
# [Quickling](3-Mechanics\CLI\bestiary\fey/quickling-vgm.md)
*Source: Volo's Guide to Monsters p. 187*  

Quicklings rocket through haunting, twisted forests where the unseelie fey hold sway, both in the Feywild and in the world. Racing faster than the eye can track, each appears as little more than a blurry wavering in the air.

A quickling is a small, slender fey, similar to a miniature elf with sharp, feral features. Its cold, cruel eyes gleam like jewels.

## Live Fast, Die Young

Quicklings owe their existence-and their plight-to the Queen of Air and Darkness, the dread ruler of the Gloaming Court. Once a race of lazy and egotistical fey, the creatures that would become the quicklings were late in answering the queen's summons one time too many. To hasten their pace and teach them to mind her will, the queen shrank their stature and sped up their internal clocks. The queen's curse gave quicklings their amazing speed but also accelerated their passage through life-no quickling lives longer than fifteen years.

To other creatures, a quickling seems blindingly fast, vanishing into an indistinct blur as it moves. Its cruel laughter is a burst of rapid staccato sounds, its speech a shrill squeal. Only when a quickling deliberately slows down, which it prefers not to do, can other beings properly see, hear, and comprehend it. Never truly at rest, a "stationary" quickling constantly paces and shifts in place, as though it can't wait to be off again.

Tricks of that sort are hardly the limit of their artful malice, however. They don't commit outright murder, but quicklings can ruin lives in plenty of other ways: stealing an important letter, swiping coins collected for the poor, planting a stolen item in someone's bag. Quicklings enjoy causing suffering that transcends mere mis chief, especially when the blame for their actions falls on other creatures and creates discord.

## Too Fast for Words

The mortal realm is a ponderous place to a quickling's eye: a hurricane creeps gradually across the sky, a torrent of rain drifts earthward like lazy snowflakes, lightning crawls in a meandering path from cloud to cloud. The slow and boring world seems to be populated by torpid creatures whose deep, mooing speech lacks meaning.

## Mischief, Not Murder

Quicklings have a capricious nature that goes well with their energy level: they think as fast as they run, and they are always up to something. A quickling spends most of its time perpetrating acts of mischief on slower creatures. One rarely passes up an opportunity to tie a person's bootlaces together, move the stool a creature is about to sit on, or unbuckle a saddle while no one's looking.

```statblock
"name": "Quickling (VGM)"
"size": "Tiny"
"type": "fey"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"hp": !!int "10"
"hit_dice": "3d4 + 3"
"stats":
- !!int "4"
- !!int "23"
- !!int "13"
- !!int "10"
- !!int "12"
- !!int "7"
"speed": "120 ft."
"skillsaves":
  "Sleight of Hand": !!int "8"
  "Stealth": !!int "8"
  "Perception": !!int "5"
  "Acrobatics": !!int "8"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Common, Sylvan"
"cr": "1"
"traits":
- "desc": "Attack rolls against the quickling have disadvantage unless the quickling\
    \ is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) or [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)."
  "name": "Blurred Movement"
- "desc": "If the quickling is subjected to an effect that allows it to make a Dexterity\
    \ saving throw to take only half damage, it instead takes no damage if it succeeds\
    \ on the saving throw, and only half damage if it fails."
  "name": "Evasion"
"actions":
- "desc": "The quickling makes three dagger attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 6|text(8) (1d4 + 6) piercing\
    \ damage."
  "name": "Dagger"
"source":
- "VGM"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Quickling.webp"
```
^statblock

```encounter-table
name: Quickling
creatures:
 - 1: Quickling
```

## Environment

forest