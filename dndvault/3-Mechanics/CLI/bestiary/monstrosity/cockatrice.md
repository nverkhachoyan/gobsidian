---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/grassland
- monster/size/small
- monster/type/monstrosity
aliases: ["Cockatrice"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 42. Available in the SRD and the Basic Rules.
---
# [Cockatrice](3-Mechanics\CLI\bestiary\monstrosity/cockatrice.md)
*Source: Monster Manual p. 42. Available in the SRD and the Basic Rules.*  

The cockatrice looks like a hideous hybrid of lizard, bird, and bat, and it is infamous for its ability to turn flesh to stone. These omnivores have a diet that consists of berries, nuts, flowers, and small animals such as insects, mice, and frogs-things they can swallow whole. They would be no threat to anything else if not for their fierce and frenzied response to even a hint of danger. A cockatrice flies into the face of any threat, squawking and madly beating its wings as its head darts out to peck. The smallest scratch from a cockatrice's beak can spell doom as its victim slowly turns to stone from the injury.

```statblock
"name": "Cockatrice"
"size": "Small"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "11"
"hp": !!int "27"
"hit_dice": "6d6 + 6"
"stats":
- !!int "6"
- !!int "12"
- !!int "12"
- !!int "2"
- !!int "13"
- !!int "5"
"speed": "20 ft., fly 40 ft."
"senses": "darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "1/2"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) piercing damage, and the target must\
    \ succeed on a DC 11 Constitution saving throw against being magically [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified).\
    \ On a failed save, the creature begins to turn to stone and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ It must repeat the saving throw at the end of its next turn. On a success, the\
    \ effect ends. On a failure, the creature is [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ for 24 hours."
  "name": "Bite"
"source":
- "MM"
- "TftYP"
- "BGDIA"
- "ERLW"
- "MOT"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Cockatrice.webp"
```
^statblock

```encounter-table
name: Cockatrice
creatures:
 - 1: Cockatrice
```

## Environment

grassland