---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Cave Fisher"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 73, Volo's Guide to Monsters p. 130
---
# [Cave Fisher](3-Mechanics\CLI\bestiary\monstrosity/cave-fisher-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 73, Volo's Guide to Monsters p. 130*  

```statblock
"name": "Cave Fisher (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "16"
- !!int "13"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "3"
"speed": "20 ft., climb 20 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "2"
"senses": "blindsight 60 ft., passive Perception 12"
"languages": ""
"cr": "3"
"traits":
- "desc": "If the cave fisher drops to half its hit points or fewer, it gains vulnerability\
    \ to fire damage."
  "name": "Flammable Blood"
- "desc": "The cave fisher can climb difficult surfaces, including upside down on\
    \ ceilings, without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "The cave fisher makes two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) slashing damage."
  "name": "Claw"
- "desc": "One Large or smaller creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the cave fisher's Adhesive Filament must make a DC 13 Strength saving throw.\
    \ On a failed save, the target is pulled into an unoccupied space within 5 feet\
    \ of the cave fisher, and the cave fisher makes one Claw attack against it. Anyone\
    \ else who was attached to the filament is released. Until the grapple ends on\
    \ the target, the cave fisher can't use Adhesive Filament."
  "name": "Retract Filament"
"bonus_actions":
- "desc": "The cave fisher extends a sticky filament up to 60 feet, and the filament\
    \ adheres to anything that touches it. A creature the filament adheres to is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the cave fisher (escape DC 13), and ability checks made to escape this grapple\
    \ have disadvantage. The filament can be attacked (AC 15; 5 hit points; immunity\
    \ to poison and psychic damage). A weapon that fails to sever it becomes stuck\
    \ to it, requiring an action and a successful DC 13 Strength check to pull free.\
    \ Destroying the filament deals no damage to the cave fisher. The filament crumbles\
    \ away if the cave fisher takes this bonus action again."
  "name": "Adhesive Filament"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Cave%20Fisher.webp"
```
^statblock

```encounter-table
name: Cave Fisher
creatures:
 - 1: Cave Fisher
```

## Environment

underdark