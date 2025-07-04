---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/arctic
- monster/environment/desert
- monster/environment/urban
- monster/size/large
- monster/type/undead
aliases: ["Boneclaw"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 121, Divine Contention, Dragon of Icespire Peak
---
# [Boneclaw](3-Mechanics\CLI\bestiary\undead/boneclaw-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 121, Divine Contention, Dragon of Icespire Peak*  

## Boneclaw

A wizard who tries to become a lich but fails might become a boneclaw instead. These hideous, cackling undead share a few of the lich's attributes—but where liches are immortal masters of the arcane, boneclaws are slaves to darkness, hatred, and pain.

The most important part of the transformation ritual occurs when the soul of the aspiring lich migrates to a prepared phylactery. If the spellcaster is too physically or magically weak to compel the soul into its prison, the soul instead seeks out a new master—a humanoid within a few miles who has an unusually hate-filled heart. The soul bonds itself to the foul essence it finds in that person, and the boneclaw becomes forever enslaved to its new master's wishes and subconscious whims. It forms near its master, sometimes appearing before that individual to receive orders and other times simply setting about the fulfillment of its master's desires.

### Limited Immortality

A boneclaw can't be destroyed while its master lives. No matter what happens to a boneclaw's body, it re-forms within hours and returns to whatever duty its master assigned.

The boneclaw can serve only evil. If its master finds redemption or sincerely turns away from the path of evil, the boneclaw is permanently destroyed.

## Cackling Slayers

Boneclaws delight in murder, and nothing pleases them more than causing horrific pain. They lurk like spiders in shadowy recesses, waiting for victims to approach within reach of their long, bony limbs. Once speared, a creature is pulled into the darkness to be sliced apart or teleported elsewhere to be tortured to death.

## Dark Reflections

A boneclaw's master might not want such a servant or even know it has one. Boneclaws bind to petty criminals, bullies, and even particularly cruel children. Even if the master is unaware of its new, horrid bodyguard, its local area will be plagued by disappearances and grisly murders, tied together by the common thread of the master's envy or hunger for revenge.

## Undead Nature

A boneclaw doesn't require air, food, drink, or sleep.

```statblock
"name": "Boneclaw (MTF)"
"size": "Large"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "17d10 + 34"
"stats":
- !!int "19"
- !!int "16"
- !!int "15"
- !!int "13"
- !!int "15"
- !!int "9"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "7"
  "Wisdom": !!int "6"
  "Constitution": !!int "6"
"skillsaves":
  "Stealth": !!int "7"
  "Perception": !!int "6"
"damage_resistances": "cold; necrotic; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 16"
"languages": "Common plus the main language of its master"
"cr": "12"
"traits":
- "desc": "While its master lives, a destroyed boneclaw gains a new body in dice:\
    \ 1d10|avg|noform (1d10) hours, with all its hit points. The new body appears\
    \ within 1 mile of the boneclaw's master."
  "name": "Rejuvenation"
- "desc": "While in dim light or darkness, the boneclaw can take the Hide action as\
    \ a bonus action."
  "name": "Shadow Stealth"
"actions":
- "desc": "The boneclaw makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one target.\
    \ Hit: dice:3d10 + 4|text(20) (3d10 + 4) piercing damage. If the target\
    \ is a creature, the boneclaw can pull the target up to 10 feet toward itself,\
    \ and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 14). The boneclaw has two claws. While a claw grapples a target,\
    \ the claw can attack only that target."
  "name": "Piercing Claw"
- "desc": "If the boneclaw is in dim light or darkness, each creature of the boneclaw's\
    \ choice within 5 feet of it must succeed on a DC 14 Constitution saving throw\
    \ or take dice:5d12 + 2|text(34) (5d12 + 2) necrotic damage.\n\nThe boneclaw\
    \ then magically teleports up to 60 feet to an unoccupied space it can see. It\
    \ can bring one creature it's grappling, teleporting that creature to an unoccupied\
    \ space it can see within 5 feet of its destination. The destination spaces of\
    \ this teleportation must be in dim light or darkness."
  "name": "Shadow Jump"
"reactions":
- "desc": "In response to a visible enemy moving into its reach, the boneclaw makes\
    \ one claw attack against that enemy. If the attack hits, the boneclaw can make\
    \ a second claw attack against the target."
  "name": "Deadly Reach"
"source":
- "MTF"
- "DC"
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Boneclaw.webp"
```
^statblock

```encounter-table
name: Boneclaw
creatures:
 - 1: Boneclaw
```

## Environment

arctic, desert, urban