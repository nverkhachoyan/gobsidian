---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/hill
- monster/environment/underdark
- monster/size/small
- monster/type/fey/goblinoid
aliases: ["Nilbog"]
NoteIcon: monster
BestiaryType: fey (goblinoid)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 195, Volo's Guide to Monsters p. 182
---
# [Nilbog](3-Mechanics\CLI\bestiary\fey/nilbog-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 195, Volo's Guide to Monsters p. 182*  

```statblock
"name": "Nilbog (MPMM)"
"size": "Small"
"type": "fey"
"subtype": "goblinoid"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "7"
"hit_dice": "2d6"
"stats":
- !!int "8"
- !!int "14"
- !!int "10"
- !!int "10"
- !!int "8"
- !!int "15"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "Common, Goblin"
"cr": "1"
"traits":
- "desc": "The nilbog casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [Tasha's hideous laughter](/3-Mechanics/CLI/spells/tashas-hideous-laughter.md)"
  "name": "Spellcasting"
- "desc": "Any creature that attempts to damage the nilbog must first succeed on a\
    \ DC 12 Charisma saving throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ until the end of the creature's next turn. A creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way must use its action praising the nilbog.\n\nThe nilbog can't regain\
    \ hit points, including through magical healing, except through its Reversal of\
    \ Fortune reaction."
  "name": "Nilbogism"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage."
  "name": "Fool's Scepter"
- "desc": "The nilbog targets one creature it can see within 60 feet of it. The target\
    \ must succeed on a DC 12 Wisdom saving throw or take dice:2d4|text(5) (2d4)\
    \ psychic damage and have disadvantage on its next attack roll before the end\
    \ of its next turn."
  "name": "Mocking Word"
"bonus_actions":
- "desc": "The nilbog takes the [Disengage](/3-Mechanics/CLI/rules/actions.md#Disengage)\
    \ or [Hide](/3-Mechanics/CLI/rules/actions.md#Hide) action."
  "name": "Nimble Escape"
"reactions":
- "desc": "In response to another creature dealing damage to the nilbog, the nilbog\
    \ reduces the damage to 0 and regains dice:1d6|text(3) (1d6) hit points."
  "name": "Reversal of Fortune"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Nilbog.webp"
```
^statblock

```encounter-table
name: Nilbog
creatures:
 - 1: Nilbog
```

## Environment

forest, hill, underdark