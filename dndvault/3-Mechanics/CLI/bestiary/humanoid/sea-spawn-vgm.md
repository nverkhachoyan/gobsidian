---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/humanoid
aliases: ["Sea Spawn"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 189
---
# [Sea Spawn](3-Mechanics\CLI\bestiary\humanoid/sea-spawn-vgm.md)
*Source: Volo's Guide to Monsters p. 189*  

Many of the stories sung as sea shanties and passed on as tales in dockside taverns tell of people lost to the sea-but not merely drowned and gone. These unfortunates are taken by the ocean and live on as sea spawn, haunting the waves like tortured reflections of their former selves. Coral encrusts them. Barnacles cling to their cold skin. Lungs that once filled with air can now breathe in water as well.

Tales provide myriad reasons for these strange transformations. "Be wary of falling in love with a sea elf or a merfolk," some say. "Return to port before a storm, no matter how tempting the catch." "Honor the sea gods as they demand, but never promise them your heart." Such cautionary tales disguise the deeper truth: things lurk beneath the waves that strive to claim the hearts and minds of land dwellers.

## Deep Thralls

Krakens, morkoths, sea hags, marids, storm giants, dragon turtles-all of these sea creatures and more can mark mortals as their own and claim them as minions. Such people might become beholden to their master through a bleak bargain, or they might find themselves cursed by such creatures. Once warped into a fishlike form, the person can't leave the sea for long without courting death.

## Anatomical Diversity

Sea spawn come in a wide variety of forms. An individual might have a tentacle for an arm, the jaws of a shark, a sea urchin's spines, a whale's fin, octopus eyes, seaweed hair, or any combination of such qualities. Some sea spawn have piscine body parts that provide them with special abilities beyond those of an ordinary humanoid.

## The Sea Spawn of Purple Rocks

Visitors to a string of islands called the Purple Rocks (in the Forgotten Realms setting) might notice one curious fact about the islands' human inhabitants: no infants or elderly are among them. This is because babies born to the Rocklanders are cast into the sea and claimed by a kraken named Slarkrethel. The experience transforms the children into fanatics dedicated to the kraken. They return from the sea as humans, but when they reach old age, they transform into sea spawn and rejoin their master in the dark depths. Some children return having suffered partial transformations, leaving them semi-bestial until their full transformation when they reach old age. These wretches are hidden until their final change, to keep the secret of the Purple Rocks.

Kraken priests (described in appendix B) are the tenders of the kraken's flock. Most of the priests are island natives, but some are merfolk, merrow, or sea elves that live in the water around the Purple Rocks.

```statblock
"name": "Sea Spawn (VGM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Neutral Evil"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "32"
"hit_dice": "5d8 + 10"
"stats":
- !!int "15"
- !!int "8"
- !!int "15"
- !!int "6"
- !!int "10"
- !!int "8"
"speed": "20 ft., swim 30 ft."
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "understands Aquan and Common but can't speak"
"cr": "1"
"traits":
- "desc": "The sea spawn can breathe air and water, but needs to be submerged in the\
    \ sea at least once a day for 1 minute to avoid suffocating."
  "name": "Limited Amphibiousness"
"actions":
- "desc": "The sea spawn makes three attacks: two unarmed strikes and one with its\
    \ Piscine Anatomy."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage."
  "name": "Unarmed Strike"
- "desc": "The sea spawn has one or more of the following attack options, provided\
    \ it has the appropriate anatomy:\n\n- Bite. Melee Weapon Attack: dice:\
    \ d20+5 (+5) to hit, reach 5 ft., one target. Hit: dice:1d4 + 2|text(4)\
    \ (1d4 + 2) piercing damage.  \n- Poison Quills. Melee Weapon Attack:\
    \ dice: d20+5 (+5) to hit, reach 5 ft., one creature. Hit: dice:1d6|text(3)\
    \ (1d6) poison damage, and the target must succeed on a DC 12 Constitution saving\
    \ throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1\
    \ minute. The target can repeat the saving throw at the end of each of its turns,\
    \ ending the effect on itself on a success.  \n- Tentacle. Melee Weapon Attack:\
    \ dice: d20+5 (+5) to hit, reach 10 ft., one target. Hit: dice:1d6 + 2|text(5)\
    \ (1d6 + 2) bludgeoning damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 12) if it is a Medium or smaller creature. Until this grapple ends,\
    \ the sea spawn can't use this tentacle on another target.  "
  "name": "Piscine Anatomy"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Sea%20Spawn.webp"
```
^statblock

```encounter-table
name: Sea Spawn
creatures:
 - 1: Sea Spawn
```

## Environment

underwater, coastal