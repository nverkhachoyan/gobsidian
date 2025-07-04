---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/urban
- monster/size/large
- monster/type/monstrosity
aliases: ["Banderhobb"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 122
---
# [Banderhobb](3-Mechanics\CLI\bestiary\monstrosity/banderhobb-vgm.md)
*Source: Volo's Guide to Monsters p. 122*  

A banderhobb is a hybrid of shadow and flesh. Through dark magic, these components take on an enormous and vile humanoid shape, resembling a bipedal toad. In this form, a banderhobb temporarily serves its creator as a thug, a thief, and a kidnapper.

## Birthed by Hags

In the earliest days of the world, a coven of night hags devised a ritual that led to the creation of the first banderhobb. A hag that knows the ritual might be willing to teach it for the right price. Some other dark fey and powerful fiends also know of the process, as do a few mortal mages. Instructions might also be found in a tome devoted to debased wizardry.

A banderhobb fulfills its duties until its existence ends. When it expires, usually several days after its birth, it leaves behind only tarry goo and wisps of shadow. Legends tell of a dark tower in the Shadowfell where the shadows sometimes reform, and banderhobbs roam.

## Silent and Deadly

When the ritual to create a banderhobb is complete, flesh, spirit, and shadow combine to produce a creature as big as an ogre. The newly formed monstrosity has spindly limbs that belie great strength. Its broad maw holds a long tongue and rows of fangs, both of which it uses to grab and swallow a creature or perhaps an object the banderhobb intends to steal. Despite its size, a banderhobb makes little noise, moving as silently as the shadows that infuse it. A banderhobb isn't capable of speech, but it can understand orders given to it by its creator and communicates with nearby banderhobbs in a psychic manner.

## Agents of Evil

During its brief existence, a banderhobb attempts to carry out the bidding of the one who birthed it. It accomplishes its mission with no concern for the harm it suffers or creates. Its only desire is to serve and succeed. A banderhobb that is assigned to track down a target is particularly dangerous when it is provided with a lock of hair, a personal belonging, or other object connected to the target. Possession of such an item allows it to sense the creature's location from as far as a mile away.

```statblock
"name": "Banderhobb (VGM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "20"
- !!int "12"
- !!int "20"
- !!int "11"
- !!int "14"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "8"
  "Stealth": !!int "7"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "Common and the languages of its creator but can't speak"
"cr": "5"
"traits":
- "desc": "If the banderhobb has even a tiny piece of a creature or an object in its\
    \ possession, such as a lock of hair or a splinter of wood, it knows the most\
    \ direct route to that creature or object if it is within 1 mile of the banderhobb."
  "name": "Resonant Connection"
- "desc": "While in dim light or darkness, the banderhobb can take the Hide action\
    \ as a bonus action."
  "name": "Shadow Stealth"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:5d6 + 5|text(22) (5d6 + 5) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15) if\
    \ it is a Large or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the banderhobb can't use its bite attack or tongue attack on another target."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one target.\
    \ Hit: dice:3d6|text(10) (3d6) necrotic damage, and the target must make\
    \ a DC 15 Strength saving throw. On a failed save, the target is pulled to a space\
    \ within 5 feet of the banderhobb, which can use a bonus action to make a bite\
    \ attack against the target."
  "name": "Tongue"
- "desc": "The banderhobb makes a bite attack against a Medium or smaller creature\
    \ it is grappling. If the attack hits, the target is swallowed, and the grapple\
    \ ends. The swallowed creature is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), it has total\
    \ cover against attacks and other effects outside the banderhobb and it takes\
    \ dice:3d6|text(10) (3d6) necrotic damage at the start of each of the banderhobb's\
    \ turns. A creature reduced to 0 hit points in this way stops taking necrotic\
    \ damage and becomes stable.\n\nThe banderhobb can have only one target swallowed\
    \ at a time. While the banderhobb isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ it can regurgitate the creature at any time (no action required) in a space\
    \ within 5 feet of it. The creature exits [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ If the banderhobb dies, it likewise regurgitates a swallowed creature."
  "name": "Swallow"
- "desc": "The banderhobb magically teleports up to 30 feet to an unoccupied space\
    \ of dim light or darkness that it can see. Before or after teleporting, it can\
    \ make a bite or tongue attack."
  "name": "Shadow Step"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Banderhobb.webp"
```
^statblock

```encounter-table
name: Banderhobb
creatures:
 - 1: Banderhobb
```

## Environment

urban