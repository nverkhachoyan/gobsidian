---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/elemental
aliases: ["Gargoyle"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 140. Available in the SRD and the Basic Rules.
---
# [Gargoyle](3-Mechanics\CLI\bestiary\elemental/gargoyle.md)
*Source: Monster Manual p. 140. Available in the SRD and the Basic Rules.*  

The inanimate gargoyles that perch atop great buildings are inspired by these malevolent creatures of elemental earth that resemble grotesque, fiendish statues. A gargoyle lurks among masonry and ruins, as still as any stone sculpture, and delights in the terror it creates when it breaks from its suspended pose, as well as the pain it inflicts on its victims.

## Animate Stone

Gargoyles cling to rocky cliffs and mountains, or roost on ledges in underground caves. They haunt city rooftops, perching vulture-like among the high stone arches and buttresses of castles and cathedrals, and they can hold themselves so still that they appear inanimate. Able to maintain this state for years, a gargoyle makes an ideal sentry.

## Deadly Reputation

Gargoyles have a reputation for cruelty. Statues carved into the likenesses of gargoyles appear in the architecture of countless cultures to frighten away trespassers. Although such sculptures are only decorative, real gargoyles can hide among them to ambush unsuspecting victims. A gargoyle might alleviate the tedium of its watch by catching and tormenting birds or rodents, but its long wait only increases its craving for harming sentient creatures.

## Cruel Servants

Gargoyles are easily inspired by the cunning of an intelligent master. They enjoy simple tasks such as guarding a master's home, torturing and killing interlopers, and anything else that involves minimum effort and maximum pain and carnage.

Gargoyles sometimes serve demons for their propensity for wanton chaos and destruction. Powerful spellcasters can also easily enlist gargoyle guardians to keep watch over their gates and walls. Gargoyles have the patience and fortitude of stone, and will serve even the cruelest master for years without complaint.

## Elemental Nature

A gargoyle doesn't require air, food, drink, or sleep.

> [!note] Shards of Elemental Evil
> 
> As Ogrémoch, the evil Prince of Elemental Earth, treads his stony realm, it leaves shards of broken rock in his wake. Imbued with slivers of sentience, these shards thrum with the essence of the elemental prince, growing over long years into vaguely humanoid rock formations that resolve at last into the hard, cruel shapes of gargoyles.
> 
> Ogrémoch doesn't create gargoyles deliberately, but they are a physical manifestation of his evil. Gargoyles are mockeries of the elemental air that Ogrémoch despises. They are heavy creatures of living stone, yet capable of flight. Like their creator, they possess a fundamental hatred for beings of elemental air, aarakocra in particular, and relish every opportunity to destroy such creatures.
> 
> On their home plane, gargoyles carve out earth motes that Ogrémoch hurtles into Aaqa, the domain of the aarakocra and the benevolent Wind Dukes the bird folk serve in the Elemental Plane of Air.
^shards-of-elemental-evil

```statblock
"name": "Gargoyle"
"size": "Medium"
"type": "elemental"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "7d8 + 21"
"stats":
- !!int "15"
- !!int "11"
- !!int "16"
- !!int "6"
- !!int "11"
- !!int "7"
"speed": "30 ft., fly 60 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't adamantine"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Terran"
"cr": "2"
"traits":
- "desc": "While the gargoyle remains motionless, it is indistinguishable from an\
    \ inanimate statue."
  "name": "False Appearance"
"actions":
- "desc": "The gargoyle makes two attacks: one with its bite and one with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) slashing damage."
  "name": "Claws"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "PotA"
- "SKT"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "BGDIA"
- "ERLW"
- "IDRotF"
- "CM"
- "JttRC"
- "BMT"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Gargoyle.webp"
```
^statblock

```encounter-table
name: Gargoyle
creatures:
 - 1: Gargoyle
```

## Environment

underdark, urban