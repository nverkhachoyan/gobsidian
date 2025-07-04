---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Bard"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 211, Sleeping Dragon's Wake
---
# [Bard](3-Mechanics\CLI\bestiary\humanoid/bard-vgm.md)
*Source: Volo's Guide to Monsters p. 211, Sleeping Dragon's Wake*  

Bards are gifted poets, storytellers, and entertainers who travel far and wide, but are commonly found in taverns or in the company of jolly bands of adventurers, rough-and-tumble mercenaries, and wealthy patrons.

```statblock
"name": "Bard (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "[chain shirt](/3-Mechanics/CLI/items/chain-shirt.md)"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "11"
- !!int "14"
- !!int "12"
- !!int "10"
- !!int "13"
- !!int "14"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "4"
  "Wisdom": !!int "3"
"skillsaves":
  "Perception": !!int "5"
  "Performance": !!int "6"
  "Acrobatics": !!int "4"
"senses": "passive Perception 15"
"languages": "any two languages"
"cr": "2"
"traits":
- "desc": "The bard is a 4th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). It has the\
    \ following bard spells prepared:\n\nCantrips (at will): [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [vicious mockery](/3-Mechanics/CLI/spells/vicious-mockery.md)\n\
    \n1st level (4 slots): [charm person](/3-Mechanics/CLI/spells/charm-person.md),\
    \ [healing word](/3-Mechanics/CLI/spells/healing-word.md), [heroism](/3-Mechanics/CLI/spells/heroism.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md), [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\
    \n2nd level (3 slots): [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [shatter](/3-Mechanics/CLI/spells/shatter.md)"
  "name": "Spellcasting"
- "desc": "The bard can perform a song while taking a short rest. Any ally who hears\
    \ the song regains an extra dice: 1d6|avg|noform (1d6) hit points if it spends\
    \ any Hit Dice to regain hit points at the end of that rest. The bard can confer\
    \ this benefit on itself as well."
  "name": "Song of Rest"
- "desc": "The bard can use a bonus action on its turn to target one creature within\
    \ 30 feet of it. If the target can hear the bard, the target must succeed on a\
    \ DC 12 Charisma saving throw or have disadvantage on ability checks, attack rolls,\
    \ and saving throws until the start of the bard's next turn."
  "name": "Taunt (2/Day)"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortbow"
"source":
- "VGM"
- "WDH"
- "GoS"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Bard.webp"
```
^statblock

```encounter-table
name: Bard
creatures:
 - 1: Bard
```

## Environment

urban