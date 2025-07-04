---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Trapper"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 194
---
# [Trapper](3-Mechanics\CLI\bestiary\monstrosity/trapper-vgm.md)
*Source: Volo's Guide to Monsters p. 194*  

A trapper is a manta-like creature that lurks in subterranean and natural environments. It can change the color and texture of its tough, outward-facing side to help it blend in with its surroundings, while its soft, inward-facing side clings to the floor, wall, or ceiling in its hunting territory. It remains motionless as it waits for prey to come close. When a target is within its reach, it peels itself away from the surface and wraps around its prey, crushing, smothering, and then digesting it.

## Versatile Camouflage

A trapper can alter the color and texture of its outer side to match its surroundings. It can blend in with any surface made of stone, earth, or wood, masking its presence to any but the most rigorous scrutiny. It can't change its texture to that of a grassy or snow-covered surface, but it can change its color to match and then conceal itself under a thin layer of vegetation or actual snow.

Trappers know when prey draws near, so explore ruins and dungeons with equal wariness. For dumb beasts, they know very well what treasure is, what treasure chests are, and how these lure the likes of us.

-Volo

## Stationary Hunters

A trapper needs to eat about a halfling-sized meal once a week to remain sated. It is content to stay in one place, given a steady supply of food, and thus trappers are a threat along any well-traveled dungeon corridor and on routes through the wilderness that see a lot of traffic. When prey is scarce, a trapper enters a state of hibernation that can last for months, though it is still aware when prey comes near. A trapper on the verge of starvation might defy its instincts and begin creeping along, abandoning its old territory in search of better hunting.

## Beware of Leftovers

When its prey is dead, a trapper dissolves and absorbs the fleshy parts, leaving a scattering of bones, metal, treasure, and other indigestible bits in the place where the creature had been. A trapper that lurks on the floor of its hunting grounds can cover these remains with own body, making them look like irregularities in the surface. The creature might also attach itself to a wall or a ceiling close to a recent kill, effectively using the remnants as bait: a creature that stops to investigate the bones for valuables stands a good chance of becoming the trapper's next meal.

```statblock
"name": "Trapper (VGM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "85"
"hit_dice": "10d10 + 30"
"stats":
- !!int "17"
- !!int "10"
- !!int "17"
- !!int "2"
- !!int "13"
- !!int "4"
"speed": "10 ft., climb 10 ft."
"skillsaves":
  "Stealth": !!int "2"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"traits":
- "desc": "While the trapper is attached to a ceiling, floor, or wall and remains\
    \ motionless, it is almost indistinguishable from an ordinary section of ceiling,\
    \ floor, or wall. A creature that can see it and succeeds on a DC 20 Intelligence\
    \ ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation)) or Intelligence\
    \ ([Nature](/3-Mechanics/CLI/rules/skills.md#Nature)) check can discern its presence."
  "name": "False Appearance"
- "desc": "The trapper can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "One Large or smaller creature within 5 feet of the trapper must succeed\
    \ on a DC 14 Dexterity saving throw or be [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 14). Until the grapple ends, the target takes dice:4d6 + 3|text(17)\
    \ (4d6 + 3) bludgeoning damage plus dice:1d6|text(3) (1d6) acid damage at\
    \ the start of each of its turns. While [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ in this way, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded), and at risk of suffocating.\
    \ The trapper can smother only one creature at a time."
  "name": "Smother"
"source":
- "VGM"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Trapper.webp"
```
^statblock

```encounter-table
name: Trapper
creatures:
 - 1: Trapper
```

## Environment

underdark