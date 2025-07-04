---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/grassland
- monster/environment/hill
- monster/size/huge
- monster/type/giant/hill-giant
aliases: ["Mouth of Grolantor"]
NoteIcon: monster
BestiaryType: giant (hill giant)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 149
---
# [Mouth of Grolantor](3-Mechanics\CLI\bestiary\giant/mouth-of-grolantor-vgm.md)
*Source: Volo's Guide to Monsters p. 149*  

Hill giants know the kinds of foods that make them fatter, and they understand that exerting themselves too much tends to make them thinner. What the lazy brutes don't comprehend are the things that make them sick. They consume spoiled food and diseased carcasses with as much enthusiasm as children eating dessert. Fortunately for hill giants, they have a vulture's constitution and rarely suffer for such eating habits. This makes it all the more mysterious to them when one of their kind becomes ill and incapable of keeping down food. Vomiting hill giants are seen as vessels of a message from Grolantor.

The clan separates the sickened giant from the others, often trapping the giant in a cage or tying the giant to a post. A priest of Grolantor or chieftain visits the famished giant daily, trying to read portents in the puddles of bile the hill giant retched up. If the sickness soon passes, the hill giant can rejoin society. If not, the hill giant is instead starved to the point of madness so that Grolantor's hunger can be given a mouth in the world.

## Starved and Insane

A mouth of Grolantor is so disgraced that it ceases to be an individual and becomes an object. Paradoxically, that object is revered as a holy embodiment of Grolantor's eternal, aching hunger. Unlike a typical thick, sluggish, half-asleep hill giant, a mouth of Grolantor is thin as a whippet, alert like a bird, and constantly twitching around the edges. A mouth of Grolantor is kept perpetually imprisoned or shackled; if it breaks free, it's sure to kill a few hill giants before it's brought down or it sprints away on a killing spree. The only time a mouth of Grolantor is set loose is during a war, a raid against an enemy settlement, or in a last-ditch defense of the tribe's home. When the mouth of Grolantor has slaughtered and eaten its fill of the tribe's enemies, it passes out amid the gory remains of its victims, making it easy to recapture.

```statblock
"name": "Mouth of Grolantor (VGM)"
"size": "Huge"
"type": "giant"
"subtype": "hill giant"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "105"
"hit_dice": "10d12 + 40"
"stats":
- !!int "21"
- !!int "10"
- !!int "18"
- !!int "5"
- !!int "7"
- !!int "5"
"speed": "50 ft."
"skillsaves":
  "Perception": !!int "1"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "passive Perception 11"
"languages": "Giant"
"cr": "6"
"traits":
- "desc": "The giant is immune to [confusion](/3-Mechanics/CLI/spells/confusion.md)\
    \ spells and similar magic.\n\nOn each of its turns, the giant uses all its movement\
    \ to move toward the nearest creature or whatever else it might perceive as food.\
    \ Roll a dice: d10|avg|noform (d10) at the start of each of the giant's turns\
    \ to determine its action for that turn:\n\n1–3. The giant makes three attacks\
    \ with its fists against one random target within its reach. If no other creatures\
    \ are within its reach, the giant flies into a rage and gains advantage on all\
    \ attack rolls until the end of its next turn.\n\n4–5. The giant makes one attack\
    \ with its fist against every creature within its reach. If no other creatures\
    \ are within its reach, the giant makes one fist attack against itself.\n\n6–\
    7. The giant makes one attack with its bite against one random target within its\
    \ reach. If no other creatures are within its reach, its eyes glaze over and it\
    \ becomes [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until the start\
    \ of its next turn.\n\n8–10. The giant makes three attacks against one random\
    \ target within its reach: one attack with its bite and two with its fists. If\
    \ no other creatures are within its reach, the giant flies into a rage and gains\
    \ advantage on all attack rolls until the end of its next turn."
  "name": "Mouth of Madness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one creature.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage, and the giant magically\
    \ regains hit points equal to the damage dealt."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:3d8 + 5|text(18) (3d8 + 5) bludgeoning damage."
  "name": "Fist"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Mouth%20of%20Grolantor.webp"
```
^statblock

```encounter-table
name: Mouth of Grolantor
creatures:
 - 1: Mouth of Grolantor
```

## Environment

grassland, hill