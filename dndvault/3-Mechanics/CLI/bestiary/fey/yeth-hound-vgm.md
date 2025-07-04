---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/large
- monster/type/fey
aliases: ["Yeth Hound"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 201
---
# [Yeth Hound](3-Mechanics\CLI\bestiary\fey/yeth-hound-vgm.md)
*Source: Volo's Guide to Monsters p. 201*  

Granted by powerful fey to individuals who please them, yeth hounds serve evil masters like hunting dogs. Yeth hounds fly in pursuit of their prey, often waiting until it is too exhausted to fight back. Only the threat of dawn drives the pack back into hiding.

## Minions of a Dark Master

A pack of yeth hounds can be created by powerful fey such as the Queen of Air and Darkness. Once it is brought into existence, a pack must have a master, who is often someone the creator wishes to reward. The master can telepathically communicate with its yeth hounds to give them commands from afar. If the master of a pack is killed, the hounds seek and choose a new master, typically an individual of great evil such as a vampire, a necromancer, or a hag.

A yeth hound stands about 5 feet tall at the shoulder and weighs around 400 pounds. Often all that can be seen of one in the darkness is the red glow of its eyes against its night-black fur. The head of a yeth hound has a human-like face, held up by a neck more flexible than a dog's. The creature gives off an odor like smoke.

Those that stand their ground and fight back discover that mundane weapons partially pass through the hound as if it was made of fog, but magic weapons and silvered weapons can strike true.

## Sound of Looming Death

Yeth hounds make a ghastly baying sound that can be heard all around. Creatures that can see a hound when it bays are filled with supernatural fear and usually flee in terror. When a victim tries to run away, a hound delights in chasing after it and tormenting it before bringing the hunt to a close.

## Foiled by Sunlight

Yeth hounds can't stand sunlight. A pack never willingly prolongs a hunt beyond the night hours and always seeks to return to its dark den before the first rays of dawn. No amount of coercion by a pack's master can deter this behavior. If a yeth hound is exposed to natural sunlight, it fades away, vanishing into the Ethereal Plane, from where its master can retrieve it only after the sun has set.

```statblock
"name": "Yeth Hound (VGM)"
"size": "Large"
"type": "fey"
"alignment": "Neutral Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "51"
"hit_dice": "6d10 + 18"
"stats":
- !!int "18"
- !!int "17"
- !!int "16"
- !!int "5"
- !!int "12"
- !!int "7"
"speed": "40 ft., fly 40 ft. (hover)"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks not\
  \ made with silvered weapons"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "understands Common, Elvish, and Sylvan but can't speak"
"cr": "4"
"traits":
- "desc": "The yeth hound has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
- "desc": "If the yeth hound starts its turn in sunlight, it is transported to the\
    \ Ethereal Plane. While sunlight shines on the spot from which it vanished, the\
    \ hound must remain in the Deep Ethereal. After sunset, it returns to the Border\
    \ Ethereal at the same spot, whereupon it typically sets out to find its pack\
    \ or its master. The hound is visible on the Material Plane while it is in the\
    \ Border Ethereal, and vice versa, but it can't affect or be affected by anything\
    \ on the other plane. Once it is adjacent to its master or a pack mate that is\
    \ on the Material Plane, a yeth hound in the Border Ethereal can return to the\
    \ Material Plane as an action."
  "name": "Sunlight Banishment"
- "desc": "While the yeth hound is on the same plane of existence as its master, it\
    \ can magically convey what it senses to its master, and the two can communicate\
    \ telepathically with each other."
  "name": "Telepathic Bond"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage, plus dice:4d6|text(14)\
    \ (4d6) psychic damage if the target is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Bite"
- "desc": "The yeth hound bays magically. Every enemy within 300 feet of the hound\
    \ that can hear it must succeed on a DC 13 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the hound's next turn or until the hound is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) target that\
    \ starts its turn within 30 feet of the hound must use all its movement on that\
    \ turn to get as far from the hound as possible, must finish the move before taking\
    \ an action, and must take the most direct route, even if hazards lie that way.\
    \ A target that successfully saves is immune to the baying of all yeth hounds\
    \ for the next 24 hours."
  "name": "Baleful Baying"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Yeth%20Hound.webp"
```
^statblock

```encounter-table
name: Yeth Hound
creatures:
 - 1: Yeth Hound
```

## Environment

grassland, forest, hill