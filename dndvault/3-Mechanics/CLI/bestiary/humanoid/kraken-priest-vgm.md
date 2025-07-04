---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Kraken Priest"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 215, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Kraken Priest](3-Mechanics\CLI\bestiary\humanoid/kraken-priest-vgm.md)
*Source: Volo's Guide to Monsters p. 215, Dragon of Icespire Peak, Storm Lord's Wrath*  

A kraken can seem godlike to folk who have witnessed its fury. Those who mistake its might for divine power and those who seek to appease the monster through veneration are sometimes rewarded with power, to serve thereafter as kraken priests.

The kraken can make itself dimly aware of a kraken priest's thoughts if the two are on the same plane of existence, and it can then push aside the priest's personality and control it. Kraken priests can thereby act as eyes and ears for their masters, and when the kraken has something to say, the priest becomes its mouthpiece.

Every kraken priest undergoes a change in appearance that reflects the kraken's influence, although each one differs in how its reverence is displayed. One kraken priest might have ink-black eyes and a suckered tentacle for a tongue, while another has a featureless face and a body covered in eyes and mouths that dribble seawater. These horrific manifestations intensify when the kraken possesses its minion to utter its dire pronouncements.

```statblock
"name": "Kraken Priest (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any Evil alignment"
"ac": !!int "10"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "12"
- !!int "10"
- !!int "16"
- !!int "10"
- !!int "15"
- !!int "14"
"speed": "30 ft., swim 30 ft."
"skillsaves":
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "passive Perception 15"
"languages": "any two languages"
"cr": "5"
"traits":
- "desc": "The priest's spellcasting ability is Wisdom (spell save DC 13, dice: d20+5\
    \ (+5) to hit with spell attacks). It can innately cast the following spells,\
    \ requiring no material components:\n\nAt will: [command](/3-Mechanics/CLI/spells/command.md),\
    \ [create or destroy water](/3-Mechanics/CLI/spells/create-or-destroy-water.md)\n\
    \n1/day each: [call lightning](/3-Mechanics/CLI/spells/call-lightning.md),\
    \ [Evard's black tentacles](/3-Mechanics/CLI/spells/evards-black-tentacles.md)\n\
    \n3/day each: [control water](/3-Mechanics/CLI/spells/control-water.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [water breathing](/3-Mechanics/CLI/spells/water-breathing.md), [water walk](/3-Mechanics/CLI/spells/water-walk.md)"
  "name": "Innate Spellcasting"
- "desc": "The priest can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "Melee Spell Attack: dice: d20+5 (+5) to hit, reach 5 ft., one creature.\
    \ Hit: dice:5d10|text(27) (5d10) thunder damage."
  "name": "Thunderous Touch"
- "desc": "A kraken speaks through the priest with a thunderous voice audible within\
    \ 300 feet. Creatures of the priest's choice that can hear the kraken's words\
    \ (which are spoken in Abyssal, Infernal, or Primordial) must succeed on a DC\
    \ 14 Charisma saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Voice of the Kraken (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "GoS"
- "DIP"
- "SLW"
- "LR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Kraken%20Priest.webp"
```
^statblock

```encounter-table
name: Kraken Priest
creatures:
 - 1: Kraken Priest
```

## Environment

coastal, underwater