---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underdark
- monster/size/large
- monster/type/plant
aliases: ["Myconid Sovereign"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Monster Manual p. 232
---
# [Myconid Sovereign](3-Mechanics\CLI\bestiary\plant/myconid-sovereign.md)
*Source: Monster Manual p. 232*  

Myconids are intelligent, ambulatory fungi that live in the Underdark, seek enlightenment, and deplore violence. If approached peacefully, myconids gladly provide shelter or allow safe passage through their colonies.

## Circles and Melds

The largest myconid in a colony is its sovereign, which presides over one or more social groups called circles. A circle consists of twenty or more myconids that work, live, and meld together.

A meld is a form of communal meditation that allows myconids to transcend their dull subterranean existence. The myconids' rapport spores bind the participants into a group consciousness. Hallucination spores then induce a shared dream that provides entertainment and social interaction. Myconids consider melding to be the purpose of their existence. They use it in the pursuit of higher consciousness, collective union, and spiritual apotheosis. Myconids also use their rapport spores to communicate telepathically with other sentient creatures.

## Myconid Reproduction

Like other fungi, myconids reproduce by mundane sporing. They carefully control their spores' release to avoid overpopulation.

```statblock
"name": "Myconid Sovereign"
"size": "Large"
"type": "plant"
"alignment": "Lawful Neutral"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "60"
"hit_dice": "8d10 + 16"
"stats":
- !!int "12"
- !!int "10"
- !!int "14"
- !!int "13"
- !!int "15"
- !!int "10"
"speed": "30 ft."
"senses": "darkvision 120 ft., passive Perception 12"
"languages": ""
"cr": "2"
"traits":
- "desc": "When the myconid takes damage, all other myconids within 240 feet of it\
    \ can sense its pain."
  "name": "Distress Spores"
- "desc": "While in sunlight, the myconid has disadvantage on ability checks, attack\
    \ rolls, and saving throws. The myconid dies if it spends more than 1 hour in\
    \ direct sunlight."
  "name": "Sun Sickness"
"actions":
- "desc": "The myconid uses either its Hallucination Spores or its Pacifying Spores,\
    \ then makes a fist attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d4 + 1|text(8) (3d4 + 1) bludgeoning damage plus dice:3d4|text(7)\
    \ (3d4) poison damage."
  "name": "Fist"
- "desc": "The myconid targets one corpse of a humanoid or a Large or smaller beast\
    \ within 5 feet of it and releases spores at the corpse. In 24 hours, the corpse\
    \ rises as a spore servant. The corpse stays animated for dice: 1d4 + 1|avg|noform\
    \ (1d4 + 1) weeks or until destroyed, and it can't be animated again in this\
    \ way."
  "name": "Animating Spores (3/Day)"
- "desc": "The myconid ejects spores at one creature it can see within 5 feet of it.\
    \ The target must succeed on a DC 12 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 minute. The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ target is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ while it hallucinates. The target can repeat the saving throw at the end of\
    \ each of its turns, ending the effect on itself on a success."
  "name": "Hallucination Spores"
- "desc": "The myconid ejects spores at one creature it can see within 5 feet of it.\
    \ The target must succeed on a DC 12 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Pacifying Spores"
- "desc": "A 30-foot radius of spores extends from the myconid. These spores can go\
    \ around corners and affect only creatures with an Intelligence of 2 or higher\
    \ that aren't undead, constructs, or elementals. Affected creatures can communicate\
    \ telepathically with one another while they are within 30 feet of each other.\
    \ The effect lasts for 1 hour."
  "name": "Rapport Spores"
"source":
- "MM"
- "WDMM"
- "GoS"
- "IMR"
- "IDRotF"
- "KftGV"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Myconid%20Sovereign.webp"
```
^statblock

```encounter-table
name: Myconid Sovereign
creatures:
 - 1: Myconid Sovereign
```

## Environment

underdark