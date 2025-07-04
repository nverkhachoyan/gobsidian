---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/7
- monster/environment/forest
- monster/size/small
- monster/type/fey
aliases: ["Korred"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 168
---
# [Korred](3-Mechanics\CLI\bestiary\fey/korred-vgm.md)
*Source: Volo's Guide to Monsters p. 168*  

Korreds are unpredictable, secretive fey with strong ties to earth and stone. Because of their magical hair and their mystical understanding of minerals, they are sought after by treasure-hunters, dwarves, and others that desire wealth beneath the earth.

## Earthy Fey

Korreds prefer to keep their own company and occasionally consort with creatures of elemental earth such as galeb duhr. A tribe of korreds gathers weekly to perform ceremonial dances, beating out rhythms on stone with their hooves and clubs. In the depths of the Material Plane, korreds typically flee from other creatures but become aggressive when they feel insulted or are annoyed by the sounds of mining.

Korreds can hurl boulders far larger than it seems they should be able to, shape stone as though it were clay, swim through rock, and summon earth elementals and other creatures. They also gain supernatural strength just from standing on the ground.

There's a legend about a merchant who tried to cut a korred's hair with golden shears. The korred fed him those shears, from his swallow to his sitter.

-Volo

## Stone Sympathy

No one knows the ways of stone and earth better than a korred. Korreds can seemingly smell veins of metal or gems. A korred on the surface can feel the rise and fall of bedrock under the earth and where caves lie, and underground it knows the pathways through the stone for miles. Secret doors that lead through stone are as obvious as windows to a korred.

## Enchanted Hair

Korreds have hair all over their bodies, but the hair that grows from their heads is magical. When cut, it transforms into whatever material was used to cut it. Korreds use iron shears to cut lengths of their hair, then weave the strands together to create iron ropes that they can manipulate, animating them to bind or snake around creatures and objects. Korreds take great pride in their hair, and equally great offense at any one who attempts to cut it without permission.

```statblock
"name": "Korred (VGM)"
"size": "Small"
"type": "fey"
"alignment": "Chaotic Neutral"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "102"
"hit_dice": "12d6 + 60"
"stats":
- !!int "23"
- !!int "14"
- !!int "20"
- !!int "10"
- !!int "15"
- !!int "9"
"speed": "30 ft., burrow 30 ft."
"skillsaves":
  "Athletics": !!int "9"
  "Stealth": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 120 ft., tremorsense 120 ft., passive Perception 15"
"languages": "Dwarvish, Gnomish, Sylvan, Terran, Undercommon"
"cr": "7"
"traits":
- "desc": "The korred's innate spellcasting ability is Wisdom (save DC 13). It can\
    \ innately cast the following spells, requiring no components:\n\nAt will:\
    \ [commune with nature](/3-Mechanics/CLI/spells/commune-with-nature.md), [meld\
    \ into stone](/3-Mechanics/CLI/spells/meld-into-stone.md), [stone shape](/3-Mechanics/CLI/spells/stone-shape.md)\n\
    \n1/day each: [conjure elemental](/3-Mechanics/CLI/spells/conjure-elemental.md)\
    \ (as 6th-level spell; [galeb duhr](/3-Mechanics/CLI/bestiary/elemental/galeb-duhr.md),\
    \ [gargoyle](/3-Mechanics/CLI/bestiary/elemental/gargoyle.md), [earth elemental](/3-Mechanics/CLI/bestiary/elemental/earth-elemental.md),\
    \ or [xorn](/3-Mechanics/CLI/bestiary/elemental/xorn.md) only), [Otto's irresistible\
    \ dance](/3-Mechanics/CLI/spells/ottos-irresistible-dance.md)"
  "name": "Innate Spellcasting"
- "desc": "The korred has at least one 50-foot-long rope woven out of its hair. As\
    \ a bonus action, the korred commands one such rope within 30 feet of it to move\
    \ up to 20 feet and entangle a Large or smaller creature that the korred can see.\
    \ The target must succeed on a DC 13 Dexterity saving throw or become [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the rope (escape DC 13). Until this grapple ends. the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ The korred can use a bonus action to release the target, which is also freed\
    \ if the korred dies or becomes [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\n\
    \nA rope of korred hair has AC 20 and 20 hit points. It regains 1 hit point at\
    \ the start of each of the korred's turns while it has at least 1 hit point and\
    \ the korred is alive. If the rope drops to 0 hit points, it is destroyed."
  "name": "Command Hair"
- "desc": "The korred has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide in rocky terrain."
  "name": "Stone Camouflage"
- "desc": "While on the ground, the korred deals 2 extra dice of damage with any weapon\
    \ attack (included in its attacks)."
  "name": "Stone's Strength"
"actions":
- "desc": "The korred makes two attacks with its greatclub or hurls two rocks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 6|text(10) (1d8 + 6) bludgeoning damage, or dice:3d8\
    \ + 6|text(19) (3d8 + 6) bludgeoning damage if the korred is on the ground."
  "name": "Greatclub"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 60/120 ft.,\
    \ one target. Hit: dice:2d8 + 6|text(15) (2d8 + 6) bludgeoning damage, or\
    \ dice:4d8 + 6|text(24) (4d8 + 6) bludgeoning damage if the korred is on the\
    \ ground."
  "name": "Rock"
"source":
- "VGM"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Korred.webp"
```
^statblock

```encounter-table
name: Korred
creatures:
 - 1: Korred
```

## Environment

forest