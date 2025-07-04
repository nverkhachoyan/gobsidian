---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-8
- monster/environment/hill
- monster/environment/underdark
- monster/size/tiny
- monster/type/aberration
aliases: ["Neogi Hatchling"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 179
---
# [Neogi Hatchling](3-Mechanics\CLI\bestiary\aberration/neogi-hatchling-vgm.md)
*Source: Volo's Guide to Monsters p. 179*  

Neogi are hateful slavers that consider most other creatures, even weaker neogi, to be servants and prey. A neogi looks like an outsize spider with an eel's neck and head. The creature can poison the body and the mind of its target, able to subjugate other beings that are otherwise physically superior.

## Alien Tyrants

Neogi usually dwell in far-flung locations on the Material Plane, as well as in the Feywild, the Shadowfell, and the Astral and Ethereal Planes. They invaded the world long ago from a remote location on the Material Plane, abandoning their home to conquer and devour creatures in other realms. To meet their need to navigate great distances, the neogi first dominated and assimilated the umber hulks of another lost world. Then, with these slaves providing the physical labor, the neogi designed and built sleek vessels, some capable of traversing the planes, to carry them to their new frontiers. Some neogi groups still create and use such vehicles, which have a distinct spidery aspect.

Some neogi use magic-the result of a pact between the neogi and aberrant entities they met during their journey from their home world. These entities look like stars and embody the essence of evil. They are known by such names as Acamar, Caiphon, Gibbeth, and Hadar.

Nothing about the neogi is more unfathomable than their mentality. Because they have the power to control minds, neogi consider doing so to be entirely appropriate. Their society makes no distinction between individuals, aside from the ability that a given creature has to control others, and they don't comprehend the emotional aspects of existence that humans and similar beings experience. To a neogi, hatred is as foreign a sensation as love, and showing loyalty in the absence of authority is foolishness.

Neogi mark themselves and their slaves through the use of dyes, transformational magic, and tattoos intended to signify rank, achievements, and ownership. By these signs, each neogi can identify its betters-and it must defer to those of higher station or risk harsh punishment.

Damn eel-spiders want to enslave us all! And no, they don't taste good.

-Volo

Outside the obligations of a servant to its master, neogi are willing to engage in any activity that profits them, and they are as devious as devils when doing so. Neogi buy and sell, but they pose a grave risk to potential patrons that might instead be easily enslaved, so their customers generally consist of desperate or evil individuals, or creatures that are formidable enough to treat with the neogi as equals. Neogi traders might set up shop in a planar bazaar, on the edge of a drow city, or near a mind flayer enclave. In other locations, the natives are more likely to join together to destroy a neogi caravan than to allow it safe conduct and trading privileges.

## Cycle of Death and Life

A neogi lives about as long as a human, and like a human it faces physical and mental infirmity as it ages. When an individual is rendered weak by advanced age, the other neogi in the group overpower it and inject it with a special poison. The toxin transforms the old neogi into a bloated, helpless mass of flesh called a great old master. Young neogi lay their eggs atop it, and when the hatchlings emerge, they devour the great old master and one another, until only a few of the strongest newborns are left.

## Hierarchy of Ownership

Surviving neogi hatchlings begin their lives under the control of adult neogi. They must learn about their society and earn a place in it, and each one starts its training by gaining mastery over a young umber hulk.

```statblock
"name": "Neogi Hatchling (VGM)"
"size": "Tiny"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "11"
"hp": !!int "7"
"hit_dice": "3d4"
"stats":
- !!int "3"
- !!int "13"
- !!int "10"
- !!int "6"
- !!int "10"
- !!int "9"
"speed": "20 ft., climb 20 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The hatchling has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and magic\
    \ can't put the hatchling to sleep."
  "name": "Mental Fortitude"
- "desc": "The hatchling can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) piercing damage plus dice:2d6|text(7)\
    \ (2d6) poison damage, and the target must succeed on a DC 10 Constitution saving\
    \ throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for\
    \ 1 minute. A target can repeat the saving throw at the end of each of its turns,\
    \ ending the effect on itself on a success."
  "name": "Bite"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Neogi%20Hatchling.webp"
```
^statblock

```encounter-table
name: Neogi Hatchling
creatures:
 - 1: Neogi Hatchling
```

## Environment

underdark, hill