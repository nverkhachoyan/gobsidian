---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-8
- monster/environment/forest
- monster/environment/hill
- monster/environment/underdark
- monster/environment/urban
- monster/size/small
- monster/type/fey
aliases: ["Boggle"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 128
---
# [Boggle](3-Mechanics\CLI\bestiary\fey/boggle-vgm.md)
*Source: Volo's Guide to Monsters p. 128*  

Boggles are the little bogeys of fairy tales. They lurk in the fringes of the Feywild and are also found on the Material Plane, where they hide under beds and in closets, waiting to frighten and bedevil folk with their mischief.

A boggle is born out of feelings of loneliness, materializing in a place where the Feywild touches the world in proximity to an intelligent being that feels isolated or abandoned. For example, a forsaken child might unintentionally conjure a boggle and see it as a sort of imaginary friend. A boggle might also appear in the attic of a lonely widower's house or in the caves of a hermit.

## Irksome Pests

Boggles engage in petty pranks to amuse themselves, passing the time at their hosts' expense. A boggle isn't above breaking dishes, hiding tools, making frightening sounds to startle cows and sour their milk, or hiding a baby in an attic. Although a boggle's antics might cause distress and unintentional harm, mischief-not mayhem-is usually its intent. If threatened, a boggle flees rather than stand and fight.

## Oily Excretions

A boggle excretes an oil from its pores and can make its oil slippery or sticky. The oil dries up and disappears an hour later.

## Twisting Space

A boggle can create magical openings to travel short distances or to pilfer items that would otherwise be beyond its reach. To create such a rift in space, a boggle must be adjacent to a space defined by a frame, such as an open window or a doorway, a gap between the bars of a cage, or the opening between the feet of a bed and the floor. The rift is invisible and disappears after a few seconds-enough time for the boggle to step, reach, or attack through it.

## Unreliable Allies

A boggle makes a decent servant for a strong-willed master, and wicked creatures such as fomorians and hags sometimes shelter boggles in their lairs. Warlocks who form pacts with archfey have also been known to command boggles, and charismatic individuals who make the right offers have enjoyed temporary alliances with these little tricksters. A bored boggle always finds some way to entertain itself.

```statblock
"name": "Boggle (VGM)"
"size": "Small"
"type": "fey"
"alignment": "typically  Chaotic Neutral"
"ac": !!int "14"
"hp": !!int "18"
"hit_dice": "4d6 + 4"
"stats":
- !!int "8"
- !!int "18"
- !!int "13"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Sleight of Hand": !!int "6"
  "Stealth": !!int "6"
  "Perception": !!int "3"
"damage_resistances": "fire"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Sylvan"
"cr": "1/8"
"traits":
- "desc": "The boggle excretes nonflammable oil from its pores. The boggle chooses\
    \ whether the oil is slippery or sticky and can change the oil on its skin from\
    \ one consistency to another as a bonus action.\n\nSlippery Oil: While coated\
    \ in slippery oil, the boggle gains advantage on Dexterity ([Acrobatics](/3-Mechanics/CLI/rules/skills.md#Acrobatics))\
    \ checks made to escape bonds, squeeze through narrow spaces, and end grapples.\n\
    \nSticky Oil: While coated in sticky oil, the boggle gains advantage on Strength\
    \ ([Athletics](/3-Mechanics/CLI/rules/skills.md#Athletics)) checks made to grapple\
    \ and any ability check made to maintain a hold on another creature, a surface,\
    \ or an object. The boggle can also climb difficult surfaces, including upside\
    \ down on ceilings, without needing to make an ability check."
  "name": "Boggle Oil"
- "desc": "As a bonus action, the boggle can create an [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ and immobile rift within an opening or frame it can see within 5 feet of it,\
    \ provided that the space is no bigger than 10 feet on any side. The dimensional\
    \ rift bridges the distance between that space and any point within 30 feet of\
    \ it that the boggle can see or specify by distance and direction (such as \"\
    30 feet straight up\"). While next to the rift, the boggle can see through it\
    \ and is considered to be next to the destination as well, and anything the boggle\
    \ puts through the rift (including a portion of its body) emerges at the destination.\
    \ Only the boggle can use the rift, and it lasts until the end of the boggle's\
    \ next turn."
  "name": "Dimensional Rift"
- "desc": "The boggle has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Uncanny Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage."
  "name": "Pummel"
- "desc": "The boggle creates a puddle of oil that is either slippery or sticky (boggle's\
    \ choice). The puddle is 1 inch deep and covers the ground in the boggle's space.\
    \ The puddle is difficult terrain For all creatures except boggles and lasts for\
    \ 1 hour.\n\nIf the oil is slippery, any creature that enters the puddle's area\
    \ or starts its turn there must succeed on a DC 11 Dexterity saving throw or fall\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\n\nIf the oil is sticky,\
    \ any creature that enters the puddle's area or starts its turn there must succeed\
    \ on a DC 11 Strength saving throw or be [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ On its turn. a creature can use an action to try to extricate itself from the\
    \ sticky puddle, ending the effect and moving into the nearest safe unoccupied\
    \ space with a successful DC 11 Strength check."
  "name": "Oil Puddle"
"source":
- "VGM"
- "IMR"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Boggle.webp"
```
^statblock

```encounter-table
name: Boggle
creatures:
 - 1: Boggle
```

## Environment

underdark, forest, hill, urban